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

package qr

import (
	"fmt"
	"math/big"

	"github.com/sonr-io/core/crypto/common"
)

// RSASpecial presents QR_N - group of quadratic residues modulo N where N is a product
// of two SAFE primes. This group is cyclic and a generator is easy to find.
// The group QR_N is isomorphic to QR_P x QR_Q. The order of QR_P and QR_Q are
// P1 and Q1 respectively. Because gcd(P1, Q1) = 1, QR_P x QR_Q is cyclic as well.
// The order of RSASpecial is P1 * Q1.
type RSASpecial struct {
	RSA
	P1 *big.Int
	Q1 *big.Int
}

func NewRSASpecial(safePrimeBitLength int) (*RSASpecial, error) {
	specialRSAPrimes, err := GetRSASpecialPrimes(safePrimeBitLength)
	if err != nil {
		return nil, err
	}
	return NewRSASpecialFromParams(specialRSAPrimes)
}

func NewRSASpecialFromParams(specialRSAPrimes *RSASpecialPrimes) (*RSASpecial, error) {
	qrRSA, err := NewRSA(specialRSAPrimes.P, specialRSAPrimes.Q)
	if err != nil {
		return nil, err
	}
	return &RSASpecial{
		RSA: *qrRSA,
		P1:  specialRSAPrimes.P1,
		Q1:  specialRSAPrimes.Q1}, nil
}

func NewRSApecialPublic(N *big.Int) *RSASpecial {
	return &RSASpecial{
		RSA: *NewRSAPublic(N),
	}
}

func (rs *RSASpecial) GetPrimes() *RSASpecialPrimes {
	return NewRSASpecialPrimes(rs.P, rs.Q, rs.P1, rs.Q1)
}

// GetRandomGenerator returns a random generator of a group of quadratic residues QR_N.
func (rs *RSASpecial) GetRandomGenerator() (*big.Int, error) {
	// We know Z_n* and Z_p* x Z_q* are isomorphic (Chinese Remainder Theorem).
	// Let's take x from Z_n* and its counterpart from (x mod p, x mod q) from Z_p* x Z_q*.
	// Because of the isomorphism, if we compute x^2 mod n, the counterpart of this
	// element in Z_p* x Z_q* is (x^2 mod p, x^2 mod q).
	// Thus QR_n = QR_p x QR_q.
	// The order of QR_p is (p-1)/2 = p1 and the order of QR_q is (q-1)/2 = q1.
	// Because p1 and q1 are primes, QR_p and QR_q are cyclic. Thus, also QR_n is cyclic
	// (because the product of two cyclic groups is cyclic iff the two orders are coprime)
	// and of order p1 * q1.
	// Thus the possible orders of elements in QR_n are: p1, q1, p1 * q1.
	// We need to find an element of order p1 * q1 (we rule out elements of order p1 and q1).

	if rs.P == nil {
		return nil,
			fmt.Errorf("GetRandomGenerator not available for RSASpecial with only public parameters")
	}

	for {
		a := common.GetRandomZnInvertibleElement(rs.N)
		a.Exp(a, big.NewInt(2), rs.N) // make it quadratic residue

		// check if the order is p1
		check1 := rs.Exp(a, rs.P1)
		if check1.Cmp(big.NewInt(1)) == 0 {
			continue
		}

		// check if the order is q1
		check2 := rs.Exp(a, rs.Q1)
		if check2.Cmp(big.NewInt(1)) == 0 {
			continue
		}

		return a, nil
	}
}

// GetRandomElement returns a random element from this group. First a random generator
// is chosen and then it is exponentiated to the random int between 0 and order
// of QR_N (P1 * Q1).
func (rs *RSASpecial) GetRandomElement() (*big.Int, error) {
	if rs.P == nil {
		return nil,
			fmt.Errorf("GetRandomElement not available for RSASpecial with only public parameters")
	}
	g, err := rs.GetRandomGenerator()
	if err != nil {
		return nil, err
	}
	r := common.GetRandomInt(rs.Order)
	el := rs.Exp(g, r)
	return el, nil
}

type RSASpecialPrimes struct {
	P  *big.Int
	Q  *big.Int
	P1 *big.Int
	Q1 *big.Int
}

func NewRSASpecialPrimes(P, Q, p, q *big.Int) *RSASpecialPrimes {
	return &RSASpecialPrimes{
		P:  P,
		Q:  Q,
		P1: p,
		Q1: q,
	}
}

// GetRSASpecialPrimes returns primes P, Q, p, q such that P = 2*p + 1 and Q = 2*q + 1.
func GetRSASpecialPrimes(bits int) (*RSASpecialPrimes, error) {
	p1 := common.GetGermainPrime(bits - 1)
	p := big.NewInt(0)
	p.Mul(p1, big.NewInt(2))
	p.Add(p, big.NewInt(1))

	q1 := common.GetGermainPrime(bits - 1)
	q := big.NewInt(0)
	q.Mul(q1, big.NewInt(2))
	q.Add(q, big.NewInt(1))

	if p.BitLen() == bits && q.BitLen() == bits {
		return NewRSASpecialPrimes(p, q, p1, q1), nil
	} else {
		err := fmt.Errorf("bit length not correct")
		return NewRSASpecialPrimes(nil, nil, nil, nil), err
	}
}
