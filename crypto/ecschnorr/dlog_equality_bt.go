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

// BlindedTrans represents a blinded transcript.
type BlindedTrans struct {
	Alpha_1 *big.Int
	Alpha_2 *big.Int
	Beta_1  *big.Int
	Beta_2  *big.Int
	Hash    *big.Int
	ZAlpha  *big.Int
}

func NewBlindedTrans(alpha_1, alpha_2, beta_1, beta_2, hash, zAlpha *big.Int) *BlindedTrans {
	return &BlindedTrans{
		Alpha_1: alpha_1,
		Alpha_2: alpha_2,
		Beta_1:  beta_1,
		Beta_2:  beta_2,
		Hash:    hash,
		ZAlpha:  zAlpha,
	}
}

// Verifies that the blinded transcript is valid. That means the knowledge of log_g1(t1), log_G2(T2)
// and log_g1(t1) = log_G2(T2). Note that G2 = g2^gamma, T2 = t2^gamma where gamma was chosen
// by verifier.
func (t *BlindedTrans) Verify(curve ec.Curve, g1, t1, G2, T2 *ec.GroupElement) bool {
	group := ec.NewGroup(curve)

	// check hash:
	hashNum := common.Hash(t.Alpha_1, t.Alpha_2,
		t.Beta_1, t.Beta_2)
	if hashNum.Cmp(t.Hash) != 0 {
		return false
	}

	// We need to verify (note that c-beta = hash(alpha11, alpha12, beta11, beta12))
	// g1^(z+alpha) = (alpha11, alpha12) * t1^(c-beta)
	// G2^(z+alpha) = (beta11, beta12) * T2^(c-beta)
	left1 := group.Exp(g1, t.ZAlpha)
	right1 := group.Exp(t1, t.Hash)
	Alpha := ec.NewGroupElement(t.Alpha_1, t.Alpha_2)
	right1 = group.Mul(Alpha, right1)

	left2 := group.Exp(G2, t.ZAlpha)
	right2 := group.Exp(T2, t.Hash)
	Beta := ec.NewGroupElement(t.Beta_1, t.Beta_2)
	right2 = group.Mul(Beta, right2)

	return left1.Equals(right1) && left2.Equals(right2)
}

type BTEqualityProver struct {
	Group  *ec.Group
	r      *big.Int
	secret *big.Int
	g1     *ec.GroupElement
	g2     *ec.GroupElement
}

func NewBTEqualityProver(curve ec.Curve) *BTEqualityProver {
	group := ec.NewGroup(curve)
	prover := BTEqualityProver{
		Group: group,
	}
	return &prover
}

// Prove that you know dlog_g1(h1), dlog_g2(h2) and that dlog_g1(h1) = dlog_g2(h2).
func (p *BTEqualityProver) GetProofRandomData(secret *big.Int,
	g1, g2 *ec.GroupElement) (*ec.GroupElement, *ec.GroupElement) {
	// Set the values that are needed before the protocol can be run.
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

func (p *BTEqualityProver) GetProofData(challenge *big.Int) *big.Int {
	// z = r + challenge * secret
	z := new(big.Int)
	z.Mul(challenge, p.secret)
	z.Add(z, p.r)
	z.Mod(z, p.Group.Q)
	return z
}

type BTEqualityVerifier struct {
	Group      *ec.Group
	gamma      *big.Int
	challenge  *big.Int
	g1         *ec.GroupElement
	g2         *ec.GroupElement
	x1         *ec.GroupElement
	x2         *ec.GroupElement
	t1         *ec.GroupElement
	t2         *ec.GroupElement
	alpha      *big.Int
	transcript *BlindedTrans
}

func NewBTEqualityVerifier(curve ec.Curve,
	gamma *big.Int) *BTEqualityVerifier {
	group := ec.NewGroup(curve)
	if gamma == nil {
		gamma = common.GetRandomInt(group.Q)
	}
	verifier := BTEqualityVerifier{
		Group: group,
		gamma: gamma,
	}

	return &verifier
}

func (v *BTEqualityVerifier) GetChallenge(g1, g2, t1, t2, x1,
	x2 *ec.GroupElement) *big.Int {
	// Set the values that are needed before the protocol can be run.
	// The protocol proves the knowledge of log_g1(t1), log_g2(t2) and
	// that log_g1(t1) = log_g2(t2).
	v.g1 = g1
	v.g2 = g2
	v.t1 = t1
	v.t2 = t2

	// Set the values g1^r1 and g2^r2.
	v.x1 = x1
	v.x2 = x2

	alpha := common.GetRandomInt(v.Group.Q)
	beta := common.GetRandomInt(v.Group.Q)

	// alpha1 = g1^r * g1^alpha * t1^beta
	// beta1 = (g2^r * g2^alpha * t2^beta)^gamma
	alpha1 := v.Group.Exp(v.g1, alpha)
	alpha1 = v.Group.Mul(v.x1, alpha1)
	tmp := v.Group.Exp(v.t1, beta)
	alpha1 = v.Group.Mul(alpha1, tmp)

	beta1 := v.Group.Exp(v.g2, alpha)
	beta1 = v.Group.Mul(v.x2, beta1)
	tmp = v.Group.Exp(v.t2, beta)
	beta1 = v.Group.Mul(beta1, tmp)
	beta1 = v.Group.Exp(beta1, v.gamma)

	// c = hash(alpha1, beta) + beta mod q
	hashNum := common.Hash(alpha1.X, alpha1.Y, beta1.X, beta1.Y)
	challenge := new(big.Int).Add(hashNum, beta)
	challenge.Mod(challenge, v.Group.Q)

	v.challenge = challenge
	v.transcript = NewBlindedTrans(alpha1.X, alpha1.Y, beta1.X, beta1.Y, hashNum, nil)
	v.alpha = alpha

	return challenge
}

// It receives z = r + secret * challenge.
//It returns true if g1^z = g1^r * (g1^secret) ^ challenge and g2^z = g2^r * (g2^secret) ^ challenge.
func (v *BTEqualityVerifier) Verify(z *big.Int) (bool, *BlindedTrans,
	*ec.GroupElement, *ec.GroupElement) {
	left1 := v.Group.Exp(v.g1, z)
	left2 := v.Group.Exp(v.g2, z)

	r1 := v.Group.Exp(v.t1, v.challenge)
	r2 := v.Group.Exp(v.t2, v.challenge)
	right1 := v.Group.Mul(r1, v.x1)
	right2 := v.Group.Mul(r2, v.x2)

	// transcript [(alpha11, alpha12, beta11, beta12), hash(alpha11, alpha12, beta11, beta12), z+alpha]
	// however, we are actually returning:
	// [alpha11, alpha12, beta11, beta12, hash(alpha11, alpha12, beta11, beta12), z+alpha]
	z1 := new(big.Int).Add(z, v.alpha)
	v.transcript.ZAlpha = z1

	G2 := v.Group.Exp(v.g2, v.gamma)
	T2 := v.Group.Exp(v.t2, v.gamma)

	if left1.Equals(right1) && left2.Equals(right2) {
		return true, v.transcript, G2, T2
	} else {
		return false, nil, nil, nil
	}
}
