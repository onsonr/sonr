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

package zn

import (
	"math/big"

	"github.com/sonr-io/core/crypto/common"
)

// Group represents Z_n* - group of all integers smaller than n and coprime with n.
// Note that this group is not cyclic in the general case (as opposed to, for
// example, Schnorr group).
//
// When n is a prime, use a special case of this group, the GroupZp struct instead.
type Group struct {
	N *big.Int
}

func NewGroup(n *big.Int) *Group {
	return &Group{
		N: n,
	}
}

// GetRandomElement returns a random element from this group. Elements of this group
// are integers that are coprime with N.
func (g *Group) GetRandomElement() *big.Int {
	return common.GetRandomZnInvertibleElement(g.N)
}

// Add computes x + y mod group.N.
func (g *Group) Add(x, y *big.Int) *big.Int {
	r := new(big.Int)
	r.Add(x, y)
	r.Mod(r, g.N)
	return r
}

// Mul computes x * y mod group.N.
func (g *Group) Mul(x, y *big.Int) *big.Int {
	r := new(big.Int)
	r.Mul(x, y)
	r.Mod(r, g.N)
	return r
}

// Exp computes x^exponent mod group.N.
func (g *Group) Exp(x, exponent *big.Int) *big.Int {
	r := new(big.Int)
	r.Exp(x, exponent, g.N)
	return r
}

// Inv computes inverse of x, that means xInv such that x * xInv = 1 mod group.N.
func (g *Group) Inv(x *big.Int) *big.Int {
	return new(big.Int).ModInverse(x, g.N)
}

// IsElementInGroup returns true if x is in the group and false otherwise. An element x is
// in Group when it is coprime with group.N, that means gcd(x, group.N) = 1.
func (g *Group) IsElementInGroup(x *big.Int) bool {
	c := new(big.Int).GCD(nil, nil, x, g.N)
	return x.Cmp(g.N) < 0 && c.Cmp(big.NewInt(1)) == 0
}
