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

// ProveDLogEquality demonstrates how prover can prove the knowledge of log_g1(t1), log_g2(t2) and
// that log_g1(t1) = log_g2(t2) in EC group.
func ProveDLogEquality(secret *big.Int, g1, g2, t1, t2 *ec.GroupElement,
	curve ec.Curve) bool {
	eProver := NewEqualityProver(curve)
	eVerifier := NewEqualityVerifier(curve)

	x1, x2 := eProver.GetProofRandomData(secret, g1, g2)

	challenge := eVerifier.GetChallenge(g1, g2, t1, t2, x1, x2)
	z := eProver.GetProofData(challenge)
	verified := eVerifier.Verify(z)
	return verified
}

type EqualityProver struct {
	Group  *ec.Group
	r      *big.Int
	secret *big.Int
	g1     *ec.GroupElement
	g2     *ec.GroupElement
}

func NewEqualityProver(curve ec.Curve) *EqualityProver {
	group := ec.NewGroup(curve)
	prover := EqualityProver{
		Group: group,
	}

	return &prover
}

func (p *EqualityProver) GetProofRandomData(secret *big.Int,
	g1, g2 *ec.GroupElement) (*ec.GroupElement, *ec.GroupElement) {
	// Sets the values that are needed before the protocol can be run.
	// The protocol proves the knowledge of log_g1(t1), log_g2(t2) and
	// that log_g1(t1) = log_g2(t2).
	p.secret = secret
	p.g1 = g1
	p.g2 = g2

	r := common.GetRandomInt(p.Group.Q)
	p.r = r
	a := p.Group.Exp(p.g1, r)
	b := p.Group.Exp(p.g2, r)
	return a, b
}

func (p *EqualityProver) GetProofData(challenge *big.Int) *big.Int {
	// z = r + challenge * secret
	z := new(big.Int)
	z.Mul(challenge, p.secret)
	z.Add(z, p.r)
	z.Mod(z, p.Group.Q)
	return z
}

type EqualityVerifier struct {
	Group     *ec.Group
	challenge *big.Int
	g1        *ec.GroupElement
	g2        *ec.GroupElement
	x1        *ec.GroupElement
	x2        *ec.GroupElement
	t1        *ec.GroupElement
	t2        *ec.GroupElement
}

func NewEqualityVerifier(curve ec.Curve) *EqualityVerifier {
	group := ec.NewGroup(curve)
	return &EqualityVerifier{
		Group: group,
	}
}

func (v *EqualityVerifier) GetChallenge(g1, g2, t1, t2, x1,
	x2 *ec.GroupElement) *big.Int {
	// Set the values that are needed before the protocol can be run.
	// The protocol proves the knowledge of log_g1(t1), log_g2(t2) and
	// that log_g1(t1) = log_g2(t2).
	v.g1 = g1
	v.g2 = g2
	v.t1 = t1
	v.t2 = t2

	// Sets the values g1^r1 and g2^r2.
	v.x1 = x1
	v.x2 = x2

	challenge := common.GetRandomInt(v.Group.Q)
	v.challenge = challenge
	return challenge
}

// It receives z = r + secret * challenge.
//It returns true if g1^z = g1^r * (g1^secret) ^ challenge and g2^z = g2^r * (g2^secret) ^ challenge.
func (v *EqualityVerifier) Verify(z *big.Int) bool {
	left1 := v.Group.Exp(v.g1, z)
	left2 := v.Group.Exp(v.g2, z)

	r1 := v.Group.Exp(v.t1, v.challenge)
	r2 := v.Group.Exp(v.t2, v.challenge)
	right1 := v.Group.Mul(r1, v.x1)
	right2 := v.Group.Mul(r2, v.x2)

	return left1.Equals(right1) && left2.Equals(right2)
}
