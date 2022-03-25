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

package ecschnorr

import (
	"math/big"

	"github.com/sonr-io/core/crypto/common"
	"github.com/sonr-io/core/crypto/ec"
)

// Prover proves knowledge of a discrete logarithm.
type Prover struct {
	Group  *ec.Group
	a      *ec.GroupElement
	secret *big.Int
	r      *big.Int // ProofRandomData
}

func NewProver(curveType ec.Curve) *Prover {
	return &Prover{
		Group: ec.NewGroup(curveType),
	}
}

// It contains also value b = a^secret.
func (p *Prover) GetProofRandomData(secret *big.Int,
	a *ec.GroupElement) *ec.GroupElement {
	r := common.GetRandomInt(p.Group.Q)
	p.r = r
	p.a = a
	p.secret = secret
	x := p.Group.Exp(a, r)
	return x
}

// It receives challenge defined by a verifier, and returns z = r + challenge * w.
func (p *Prover) GetProofData(challenge *big.Int) *big.Int {
	// z = r + challenge * secret
	z := new(big.Int)
	z.Mul(challenge, p.secret)
	z.Add(z, p.r)
	z.Mod(z, p.Group.Q)
	return z
}

type Verifier struct {
	Group     *ec.Group
	x         *ec.GroupElement
	a         *ec.GroupElement
	b         *ec.GroupElement
	challenge *big.Int
}

func NewVerifier(curveType ec.Curve) *Verifier {
	return &Verifier{
		Group: ec.NewGroup(curveType),
	}
}

// TODO: t transferred at some other stage?
func (v *Verifier) SetProofRandomData(x, a, b *ec.GroupElement) {
	v.x = x
	v.a = a
	v.b = b
}

func (v *Verifier) GetChallenge() *big.Int {
	challenge := common.GetRandomInt(v.Group.Q)
	v.challenge = challenge
	return challenge
}

// SetChallenge is used when Fiat-Shamir is used - when challenge is generated using hash by the prover.
func (v *Verifier) SetChallenge(challenge *big.Int) {
	v.challenge = challenge
}

func (v *Verifier) Verify(z *big.Int) bool {
	left := v.Group.Exp(v.a, z)
	r := v.Group.Exp(v.b, v.challenge)
	right := v.Group.Mul(r, v.x)
	return left.Equals(right)
}
