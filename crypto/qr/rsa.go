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
	"math/big"

	"fmt"
)

// RSA presents QR_N - group of quadratic residues modulo N where N is a product
// of two primes. This group is in general NOT cyclic (it is only when (P-1)/2 and (Q-1)/2 are primes,
// see RSASpecial). The group QR_N is isomorphic to QR_P x QR_Q.
type RSA struct {
	N     *big.Int // N = P * Q
	P     *big.Int
	Q     *big.Int
	Order *big.Int // Order = (P-1)/2 * (Q-1)/2
}

func NewRSA(P, Q *big.Int) (*RSA, error) {
	if !P.ProbablyPrime(20) || !Q.ProbablyPrime(20) {
		return nil, fmt.Errorf("P and Q must be primes")
	}
	pMin := new(big.Int).Sub(P, big.NewInt(1))
	pMinHalf := new(big.Int).Div(pMin, big.NewInt(2))
	qMin := new(big.Int).Sub(Q, big.NewInt(1))
	qMinHalf := new(big.Int).Div(qMin, big.NewInt(2))
	order := new(big.Int).Mul(pMinHalf, qMinHalf)
	return &RSA{
		N:     new(big.Int).Mul(P, Q),
		P:     P,
		Q:     Q,
		Order: order,
	}, nil
}

func NewRSAPublic(N *big.Int) *RSA {
	return &RSA{
		N: N,
	}
}

// Add computes x + y (mod N)
func (g *RSA) Add(x, y *big.Int) *big.Int {
	r := new(big.Int)
	r.Add(x, y)
	return r.Mod(r, g.N)
}

// Mul computes x * y in QR_N. This means x * y mod N.
func (g *RSA) Mul(x, y *big.Int) *big.Int {
	r := new(big.Int)
	r.Mul(x, y)
	return r.Mod(r, g.N)
}

// Inv computes inverse of x in QR_N. This means xInv such that x * xInv = 1 mod N.
func (g *RSA) Inv(x *big.Int) *big.Int {
	return new(big.Int).ModInverse(x, g.N)
}

// Exp computes base^exponent in QR_N. This means base^exponent mod rsa.N.
func (g *RSA) Exp(base, exponent *big.Int) *big.Int {
	expAbs := new(big.Int).Abs(exponent)
	if expAbs.Cmp(exponent) == 0 {
		return new(big.Int).Exp(base, exponent, g.N)
	} else {
		t := new(big.Int).Exp(base, expAbs, g.N)
		return g.Inv(t)
	}
}

// IsElementInGroup returns true if a is in QR_N and false otherwise.
func (g *RSA) IsElementInGroup(a *big.Int) (bool, error) {
	if g.P == nil {
		return false,
			fmt.Errorf("IsElementInGroup not available for RSA with only public parameters")
	}

	factors := []*big.Int{g.P, g.Q}
	for _, p := range factors {
		factorIsQR, err := isQR(a, p)
		if err != nil {
			return false, err
		}
		if !factorIsQR {
			return false, nil
		}
	}
	return true, nil
}
