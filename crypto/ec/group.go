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

package ec

import (
	"crypto/elliptic"
	"math/big"

	"github.com/sonr-io/core/crypto/common"
)

// TODO Insert appropriate comment with description of this struct
type GroupElement struct {
	X *big.Int
	Y *big.Int
}

func NewGroupElement(x, y *big.Int) *GroupElement {
	return &GroupElement{
		X: x,
		Y: y,
	}
}

func (e *GroupElement) Equals(b *GroupElement) bool {
	return e.X.Cmp(b.X) == 0 && e.Y.Cmp(b.Y) == 0
}

// Group is a wrapper around elliptic.Curve. It is a cyclic group with generator
// (c.Params().Gx, c.Params().Gy) and order c.Params().N (which is exposed as Q in a wrapper).
type Group struct {
	Curve elliptic.Curve
	Q     *big.Int
}

func NewGroup(curveType Curve) *Group {
	c := GetCurve(curveType)
	group := Group{
		Curve: c,
		Q:     c.Params().N, // order of generator G
	}
	return &group
}

// GetRandomElement returns a random element from this group.
func (g *Group) GetRandomElement() *GroupElement {
	r := common.GetRandomInt(g.Q)
	el := g.ExpBaseG(r)
	return el
}

// Mul computes a * b in Group. This actually means a + b as this is additive group.
func (g *Group) Mul(a, b *GroupElement) *GroupElement {
	// computes (x1, y1) + (x2, y2) as this is g on elliptic curves
	x, y := g.Curve.Add(a.X, a.Y, b.X, b.Y)
	return NewGroupElement(x, y)
}

// Exp computes base^exponent in Group. This actually means exponent * base as this is
// additive group.
func (g *Group) Exp(base *GroupElement, exponent *big.Int) *GroupElement {
	// computes (x, y) * exponent
	hx, hy := g.Curve.ScalarMult(base.X, base.Y, exponent.Bytes())
	return NewGroupElement(hx, hy)
}

// Exp computes base^exponent in Group where base is the generator.
// This actually means exponent * G as this is additive group.
func (g *Group) ExpBaseG(exponent *big.Int) *GroupElement {
	// computes g ^^ exponent or better to say g * exponent as this is elliptic ((gx, gy) * exponent)
	hx, hy := g.Curve.ScalarBaseMult(exponent.Bytes())
	return NewGroupElement(hx, hy)
}

// Inv computes inverse of x in Group. This is done by computing x^(order-1) as:
// x * x^(order-1) = x^order = 1. Note that this actually means x * (order-1) as this is
// additive group.
func (g *Group) Inv(x *GroupElement) *GroupElement {
	orderMinOne := new(big.Int).Sub(g.Q, big.NewInt(1))
	inv := g.Exp(x, orderMinOne)
	return inv
}
