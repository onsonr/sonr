/*
 * Copyright 2017 XLAB d.o.o.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package df

import (
	"fmt"
	"math/big"

	"github.com/sonr-io/core/crypto/common"
)

// Common values used for comparisons and big integer arithmetic, defined here for convenience.
var (
	zero = big.NewInt(0)
	one  = big.NewInt(1)
	two  = big.NewInt(2)
	four = big.NewInt(4)
)

// lagrange is an array that will hold the roots of the decomposed integer.
// It has exactly 4 elements - if decomposition has less than four roots, the remaining
// ones are set to zero.
type lagrange [4]*big.Int

// NewLagrange initializes lagrange array with zero values
func NewLagrange() *lagrange {
	w := &lagrange{}
	for i := 0; i < 4; i++ {
		w[i] = new(big.Int).Set(zero)
	}
	return w
}

func (l *lagrange) Set(w0, w1, w2, w3 *big.Int) {
	l[0].Set(w0)
	l[1].Set(w1)
	l[2].Set(w2)
	l[3].Set(w3)
}

// lipmaaDecomposition calls lipmaaDecompose and converts its result from a
// fixed-size 4-element array to a slice containing only non-zero roots.
func lipmaaDecomposition(n *big.Int) ([]*big.Int, error) {
	roots, err := lipmaaDecompose(n)
	if err != nil {
		return nil, err
	}

	rootsSlice := make([]*big.Int, 0, 4)
	for _, root := range roots {
		if root.Cmp(zero) != 0 {
			rootsSlice = append(rootsSlice, root)
		}
	}

	return rootsSlice, nil
}

// lipmaaDecompose takes a positive integer and computes its lagrange representation as the
// sum of (at most) four squares.
// Returns the roots of the decomposed integer in an array of size 4 - when squared,
// they sum up to exactly n.
func lipmaaDecompose(n *big.Int) (*lagrange, error) {
	// roots of the decomposed integer
	var w *lagrange
	var err error

	isSpecial, w, err := getSpecialDecomposition(n)
	if err != nil {
		return w, err
	}
	if isSpecial {
		return w, nil
	}

	// Write n in the form 2^t(2k+1), where t,k >= 0
	// t is the index (starting at LSB) of the first set bit
	t := 0
	for n.Bit(t) == 0 {
		t++
	}
	k := new(big.Int).Rsh(n, uint(t+1))

	if t == 1 {
		p, w1, w2 := findPrimeAndTwoRoots(n)
		sqrtOfMinus1 := sqrtOfMinus1(p)
		w3, w4 := decomposePrimeToTwoSquares(sqrtOfMinus1, p)
		w.Set(w1, w2, w3, w4)

	} else if t%2 == 1 {
		// if t is odd but not 1, find a representation (w1,w2,w3,w4) of integer m
		// and return representation of integer n = (sw1,sw2,sw3,sw4), where s = 2^((t-1)/2).
		e := uint(t-1) / 2 // exponent of s
		m := new(big.Int).Rsh(n, uint(t-1))

		// recurse to find representation (w1,w2,w3,w4) of integer m
		w, err = lipmaaDecompose(m)
		if err != nil {
			return w, err
		}
		// because s is a power of two (e.g. s = 2^((t-1)/2) = 2^e), multiplication with s
		// is equivalent to performing e = (t-1)/2 left shifts
		w[0].Lsh(w[0], e)
		w[1].Lsh(w[1], e)
		w[2].Lsh(w[2], e)
		w[3].Lsh(w[3], e)
	} else {
		// if t is even, first find a representation (w1,w2,w3,w4) of integer m = 2(2k+1)
		m := new(big.Int).Mul(k, four)
		m.Add(m, two)

		// recurse to find representation (w1,w2,w3,w4) of integer m
		wM, err := lipmaaDecompose(m)
		if err != nil {
			return w, err
		}

		// group the roots of decomposed integer m into a representation (w1,w2,w3,w4)
		// so that w1 is congruent w2 mod 2 and w3 is congruent w4 mod 2 (e.g. w1-w2 and w3-w4 are
		// both divisible by 2). This way we get the representation of 2k+1 instead of
		// representation for 2(2k+1)
		if wM[0].Bit(0) == wM[1].Bit(0) { // w1 and w2 are both odd, OK
			w = wM
		}
		if wM[0].Bit(0) == wM[2].Bit(0) { //
			w = wM
			tmp := new(big.Int).Set(wM[2])
			wM[2].Set(wM[1])
			wM[1].Set(tmp)
		}
		if wM[0].Bit(0) == wM[3].Bit(0) {
			w = wM
			tmp := new(big.Int).Set(wM[3])
			wM[3].Set(wM[1])
			wM[1].Set(tmp)
		}

		// return representation of n = (s(w1+w2), s(w1-w2), s(w3+w4), s(w3-w4))
		// where s = 2^(t/2-1)
		w1Plusw2 := new(big.Int).Add(w[0], w[1])
		w1Minusw2 := new(big.Int).Sub(w[0], w[1])
		w3Plusw4 := new(big.Int).Add(w[2], w[3])
		w3Minusw4 := new(big.Int).Sub(w[2], w[3])

		if t/2-1 >= 0 {
			e := big.NewInt(int64(t/2 - 1))
			s := new(big.Int).Exp(two, e, nil)
			w.Set(new(big.Int).Mul(w1Plusw2, s),
				new(big.Int).Mul(w1Minusw2, s),
				new(big.Int).Mul(w3Plusw4, s),
				new(big.Int).Mul(w3Minusw4, s))

		} else { // t/2-1 < 0
			w.Set(new(big.Int).Div(w1Plusw2, two),
				new(big.Int).Div(w1Minusw2, two),
				new(big.Int).Div(w3Plusw4, two),
				new(big.Int).Div(w3Minusw4, two))
		}
	}

	for i := 0; i < 4; i++ {
		w[i].Abs(w[i])
	}

	return w, nil
}

// getSpecialDecomposition checks for validity of integral argument and returns decomposition of a
// special case (0, 1 and 2) if aporopriate. Otherwise it returns a zero-filled decomposition.
func getSpecialDecomposition(n *big.Int) (bool, *lagrange, error) {
	w := NewLagrange()

	if n.Cmp(zero) < 0 {
		return true, w, fmt.Errorf("n cannot be negative")
	}

	if n.Cmp(zero) == 0 {
		return true, w, nil
	}
	if n.Cmp(one) == 0 {
		w[0].Set(one)
		return true, w, nil
	}
	if n.Cmp(two) == 0 {
		w[0].Set(one)
		w[1].Set(one)
		return true, w, nil
	}

	return false, w, nil
}

// findPrimeAndTwoRoots finds such a prime p = n-w1²-w2² that p is congruent to 1 modulo 4
// (e.g. p-1 is divisible by 4), given randomly selected w1 <= sqrt(n) and w2 <= sqrt(n - w1²),
// where exactly one of w1, w2 is even.
// w1 and w2 are also returned as two of the roots of n's decomposition.
func findPrimeAndTwoRoots(n *big.Int) (*big.Int, *big.Int, *big.Int) {
	// we will need to get random w1, w2 within appropriate upper boundaries
	w1, w2 := new(big.Int), new(big.Int)
	w1Upper, w2Upper := new(big.Int), new(big.Int)

	w1Upper.Sqrt(n) // upper bound for w1 is sqrt(n)

	p := new(big.Int)
	// repeat the procedure until we're fairly confident that p really is a prime
	for !p.ProbablyPrime(20) {
		// choose random w1 <= sqrt(n)
		w1 = common.GetRandomInt(w1Upper)

		// choose random w2 <= sqrt(n - w1²)
		w2Upper.Sqrt(new(big.Int).Sub(n, new(big.Int).Mul(w1, w1)))
		w2 = common.GetRandomInt(w2Upper)

		// we need to ensure that exactly one of w1, w2 is even
		if w1.Bit(0) == w2.Bit(0) {
			w1.Sub(w1, one) // decrease w1 by one to obtain an odd number
			if w1.Cmp(zero) <= 0 {
				continue
			}
		}

		// p = n - w1² - w2²
		p.Sub(n, new(big.Int).Mul(w1, w1))
		p.Sub(p, new(big.Int).Mul(w2, w2))
	}

	return p, w1, w2
}

// decomposePrimeToTwoSquares aplies partial Euclidean algorithm over integers in
// order to obtain such x and y that produce p = x²+y² for a prime p.
// x and y are obtained as the two largest remainders that are less than sqrt(p).
func decomposePrimeToTwoSquares(u, p *big.Int) (*big.Int, *big.Int) {
	a := new(big.Int).Set(u)
	b := new(big.Int).Set(p)
	r := new(big.Int).Mod(a, b)

	x := new(big.Int).Set(a)
	y := new(big.Int).Set(r)
	sqrtP := new(big.Int).Sqrt(p)

	for x.Cmp(sqrtP) > 0 || x.Cmp(y) == 0 {
		x.Set(a)
		y.Set(r)
		r.Mod(a, b)

		a.Set(b)
		b.Set(r)
	}

	return x, y
}

// sqrtOfMinus1 efficiently computes the square root of -1 modulo p,
// where p is a prime.
func sqrtOfMinus1(p *big.Int) *big.Int {
	if p.Cmp(one) == 0 {
		return zero
	}
	if p.Cmp(two) == 0 {
		return one
	}

	u := new(big.Int).Sub(p, one) // u = p-1
	k := new(big.Int).Set(u)

	var s uint = 0
	for k.Bit(int(s)) == 0 {
		s++
	}

	k.Rsh(k, uint(s))
	k.Sub(k, one)
	k.Rsh(k, 1)

	r := new(big.Int).Exp(u, k, p)

	n := new(big.Int).Exp(r, two, nil)
	n.Rem(n, p)
	n.Mul(n, u)
	n.Rem(n, p)

	r.Mul(r, u)
	r.Rem(r, p)

	if n.Cmp(one) == 0 {
		return r
	}

	z := new(big.Int).Set(two)
	for new(big.Int).Exp(z, new(big.Int).Div(new(big.Int).Sub(p, one), two), p).Cmp(one) == 0 {
		z.Add(z, one)
	}

	v := new(big.Int).Set(k)
	v.Lsh(v, 1)
	v.Add(v, one)

	c := new(big.Int).Exp(z, v, p)

	t := uint(0)
	for n.Cmp(one) > 0 {
		k.Set(n)
		t = s
		s = 0

		for k.Cmp(one) != 0 {
			k.Mod(new(big.Int).Exp(k, two, nil), p)
			s++
		}
		t -= s

		v.Set(one)
		v.Lsh(v, t-1)
		c.Exp(c, v, p)
		r.Rem(new(big.Int).Mul(r, c), p)
		c.Rem(new(big.Int).Mul(c, c), p)
		n.Mod(new(big.Int).Mul(n, c), p)
	}

	return r
}
