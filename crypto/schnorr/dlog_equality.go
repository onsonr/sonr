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

package schnorr

import (
	"math/big"

	"github.com/sonr-io/core/crypto/common"
)

// ProveEquality demonstrates how prover can prove the knowledge of log_g1(t1), log_g2(t2) and
// that log_g1(t1) = log_g2(t2).
func ProveEquality(secret, g1, g2, t1, t2 *big.Int, group *Group) bool {
	eProver := NewEqualityProver(group)
	eVerifier := NewEqualityVerifier(group)

	x1, x2 := eProver.GetProofRandomData(secret, g1, g2)

	challenge := eVerifier.GetChallenge(g1, g2, t1, t2, x1, x2)
	z := eProver.GetProofData(challenge)
	verified := eVerifier.Verify(z)
	return verified
}

type EqualityProver struct {
	Group  *Group
	r      *big.Int
	secret *big.Int
	g1     *big.Int
	g2     *big.Int
}

func NewEqualityProver(group *Group) *EqualityProver {
	prover := EqualityProver{
		Group: group,
	}

	return &prover
}

func (p *EqualityProver) GetProofRandomData(secret, g1, g2 *big.Int) (*big.Int, *big.Int) {
	// Sets the values that are needed before the protocol can be run.
	// The protocol proves the knowledge of log_g1(t1), log_g2(t2) and
	// that log_g1(t1) = log_g2(t2).
	p.secret = secret
	p.g1 = g1
	p.g2 = g2

	r := common.GetRandomInt(p.Group.Q)
	p.r = r
	x1 := p.Group.Exp(p.g1, r)
	x2 := p.Group.Exp(p.g2, r)
	return x1, x2
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
	Group     *Group
	challenge *big.Int
	g1        *big.Int
	g2        *big.Int
	x1        *big.Int
	x2        *big.Int
	t1        *big.Int
	t2        *big.Int
}

func NewEqualityVerifier(group *Group) *EqualityVerifier {
	verifier := EqualityVerifier{
		Group: group,
	}

	return &verifier
}

func (v *EqualityVerifier) GetChallenge(g1, g2, t1, t2, x1, x2 *big.Int) *big.Int {
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

	r11 := v.Group.Exp(v.t1, v.challenge)
	r12 := v.Group.Exp(v.t2, v.challenge)
	right1 := v.Group.Mul(r11, v.x1)
	right2 := v.Group.Mul(r12, v.x2)

	return left1.Cmp(right1) == 0 && left2.Cmp(right2) == 0
}
