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

package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"math/big"

	"github.com/sonr-io/core/crypto/zn"
)

// Group presents Z_n* - group of all integers smaller than n and coprime with n,
// where n is a product of two distinct large primes. Note that this group
// is NOT cyclic (as opposed for example to RSASpecial which is a subgroup of Group group).
type Group struct {
	zn.Group          // Group group is a special case of Group where n = P * Q)
	N        *big.Int // N = P * Q
	P        *big.Int
	Q        *big.Int
	E        *big.Int
}

func NewGroup(nBitLength int) (*Group, error) {
	priv, err := rsa.GenerateKey(rand.Reader, nBitLength)
	if err != nil {
		return nil, err
	}
	p := priv.Primes[0]
	q := priv.Primes[1]
	n := new(big.Int).Mul(p, q)

	return &Group{
		P:     p,
		Q:     q,
		N:     n,
		Group: *zn.NewGroup(n),
	}, nil
}

// Homomorphism returns x^E mod N (it is not called Encrypt, because there is no padding).
func (group *Group) Homomorphism(x *big.Int) *big.Int {
	return new(big.Int).Exp(x, group.E, group.N)
}
