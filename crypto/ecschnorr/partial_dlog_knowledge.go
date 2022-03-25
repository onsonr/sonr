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

type ECTriple struct {
	A *ec.GroupElement
	B *ec.GroupElement
	C *ec.GroupElement
}

func NewECTriple(a, b, c *ec.GroupElement) *ECTriple {
	triple := ECTriple{A: a, B: b, C: c}
	return &triple
}

// ProvePartialDLogKnowledge demonstrates how prover can prove that he knows dlog_a2(b2) and
// the verifier does not know whether knowledge of dlog_a1(b1) or knowledge of dlog_a2(b2) was proved.
func ProvePartialDLogKnowledge(group *ec.Group, secret1 *big.Int,
	a1, a2, b2 *ec.GroupElement) bool {
	prover := NewPartialProver(group)
	verifier := NewPartialVerifier(group)

	b1 := prover.Group.Exp(a1, secret1)
	triple1, triple2 := prover.GetProofRandomData(secret1, a1, b1, a2, b2)

	verifier.SetProofRandomData(triple1, triple2)
	challenge := verifier.GetChallenge()

	c1, z1, c2, z2 := prover.GetProofData(challenge)
	verified := verifier.Verify(c1, z1, c2, z2)

	return verified
}

// Proving that it knows either secret1 such that a1^secret1 = b1 or
//  secret2 such that a2^secret2 = b2.
type PartialProver struct {
	Group   *ec.Group
	secret1 *big.Int
	a1      *ec.GroupElement
	a2      *ec.GroupElement
	r1      *big.Int
	c2      *big.Int
	z2      *big.Int
	ord     int
}

func NewPartialProver(group *ec.Group) *PartialProver {
	return &PartialProver{
		Group: group,
	}
}

func (p *PartialProver) GetProofRandomData(secret1 *big.Int, a1, b1, a2,
	b2 *ec.GroupElement) (*ECTriple, *ECTriple) {
	p.a1 = a1
	p.a2 = a2
	p.secret1 = secret1
	r1 := common.GetRandomInt(p.Group.Q)
	c2 := common.GetRandomInt(p.Group.Q)
	z2 := common.GetRandomInt(p.Group.Q)
	p.r1 = r1
	p.c2 = c2
	p.z2 = z2
	x1 := p.Group.Exp(a1, r1)
	x2 := p.Group.Exp(a2, z2)
	b2ToC2 := p.Group.Exp(b2, c2)
	b2ToC2Inv := p.Group.Inv(b2ToC2)
	x2 = p.Group.Mul(x2, b2ToC2Inv)

	// we need to make sure that the order does not reveal which secret we do know:
	ord := common.GetRandomInt(big.NewInt(2))
	triple1 := NewECTriple(x1, a1, b1)
	triple2 := NewECTriple(x2, a2, b2)

	if ord.Cmp(big.NewInt(0)) == 0 {
		p.ord = 0
		return triple1, triple2
	} else {
		p.ord = 1
		return triple2, triple1
	}
}

func (p *PartialProver) GetProofData(challenge *big.Int) (*big.Int, *big.Int,
	*big.Int, *big.Int) {
	c1 := new(big.Int).Xor(p.c2, challenge)

	z1 := new(big.Int)
	z1.Mul(c1, p.secret1)
	z1.Add(z1, p.r1)
	z1.Mod(z1, p.Group.Q)

	if p.ord == 0 {
		return c1, z1, p.c2, p.z2
	} else {
		return p.c2, p.z2, c1, z1
	}
}

type PartialVerifier struct {
	Group     *ec.Group
	triple1   *ECTriple // contains x1, a1, b1
	triple2   *ECTriple // contains x2, a2, b2
	challenge *big.Int
}

func NewPartialVerifier(group *ec.Group) *PartialVerifier {
	return &PartialVerifier{
		Group: group,
	}
}

func (v *PartialVerifier) SetProofRandomData(triple1, triple2 *ECTriple) {
	v.triple1 = triple1
	v.triple2 = triple2
}

func (v *PartialVerifier) GetChallenge() *big.Int {
	challenge := common.GetRandomInt(v.Group.Q)
	v.challenge = challenge
	return challenge
}

func (v *PartialVerifier) verifyTriple(triple *ECTriple,
	challenge, z *big.Int) bool {
	left := v.Group.Exp(triple.B, z)      // a.X, a.Y, z
	r := v.Group.Exp(triple.C, challenge) // b.X, b.Y, challenge
	right := v.Group.Mul(r, triple.A)     // r1, r2, x.X, x.Y

	return left.Equals(right)
}

func (v *PartialVerifier) Verify(c1, z1, c2, z2 *big.Int) bool {
	c := new(big.Int).Xor(c1, c2)
	if c.Cmp(v.challenge) != 0 {
		return false
	}

	verified1 := v.verifyTriple(v.triple1, c1, z1)
	verified2 := v.verifyTriple(v.triple2, c2, z2)
	return verified1 && verified2
}
