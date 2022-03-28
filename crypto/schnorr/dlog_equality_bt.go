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

// BlindedTrans represents a blinded transcript.
type BlindedTrans struct {
	A      *big.Int
	B      *big.Int
	Hash   *big.Int
	ZAlpha *big.Int
}

func NewBlindedTrans(a, b, hash, zAlpha *big.Int) *BlindedTrans {
	return &BlindedTrans{
		A:      a,
		B:      b,
		Hash:   hash,
		ZAlpha: zAlpha,
	}
}

// Verifies that the blinded transcript is valid. That means the knowledge of log_g1(t1), log_G2(T2)
// and log_g1(t1) = log_G2(T2). Note that G2 = g2^gamma, T2 = t2^gamma where gamma was chosen
// by verifier.
func (t *BlindedTrans) Verify(group *Group, g1, t1, G2, T2 *big.Int) bool {
	// BlindedTrans should be in the following form: [alpha1, beta1, hash(alpha1, beta1), z+alpha]

	// check hash:
	hashNum := common.Hash(t.A, t.B)
	if hashNum.Cmp(t.Hash) != 0 {
		return false
	}

	// We need to verify (note that c-beta = hash(alpha1, beta1))
	// g1^(z+alpha) = alpha1 * t1^(c-beta)
	// G2^(z+alpha) = beta1 * T2^(c-beta)
	left1 := group.Exp(g1, t.ZAlpha)
	right1 := group.Exp(t1, t.Hash)
	right1 = group.Mul(t.A, right1)

	left2 := group.Exp(G2, t.ZAlpha)
	right2 := group.Exp(T2, t.Hash)
	right2 = group.Mul(t.B, right2)

	if left1.Cmp(right1) == 0 && left2.Cmp(right2) == 0 {
		return true
	} else {
		return false
	}
}

type BTEqualityProver struct {
	Group  *Group
	r      *big.Int
	secret *big.Int
	g1     *big.Int
	g2     *big.Int
}

func NewBTEqualityProver(group *Group) *BTEqualityProver {
	prover := BTEqualityProver{
		Group: group,
	}
	return &prover
}

// Prove that you know dlog_g1(h1), dlog_g2(h2) and that dlog_g1(h1) = dlog_g2(h2).
func (p *BTEqualityProver) GetProofRandomData(secret, g1, g2 *big.Int) (*big.Int,
	*big.Int) {
	// Set the values that are needed before the protocol can be run.
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

func (p *BTEqualityProver) GetProofData(challenge *big.Int) *big.Int {
	// z = r + challenge * secret
	z := new(big.Int)
	z.Mul(challenge, p.secret)
	z.Add(z, p.r)
	z.Mod(z, p.Group.Q)
	return z
}

type BTEqualityVerifier struct {
	Group      *Group
	gamma      *big.Int
	challenge  *big.Int
	g1         *big.Int
	g2         *big.Int
	x1         *big.Int
	x2         *big.Int
	t1         *big.Int
	t2         *big.Int
	alpha      *big.Int
	transcript *BlindedTrans
}

func NewBTEqualityVerifier(group *Group,
	gamma *big.Int) *BTEqualityVerifier {
	if gamma == nil {
		gamma = common.GetRandomInt(group.Q)
	}
	verifier := BTEqualityVerifier{
		Group: group,
		gamma: gamma,
	}

	return &verifier
}

func (v *BTEqualityVerifier) GetChallenge(g1, g2, t1, t2, x1, x2 *big.Int) *big.Int {
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
	hashNum := common.Hash(alpha1, beta1)
	challenge := new(big.Int).Add(hashNum, beta)
	challenge.Mod(challenge, v.Group.Q)

	v.challenge = challenge
	v.transcript = NewBlindedTrans(alpha1, beta1, hashNum, nil)
	v.alpha = alpha

	return challenge
}

// It receives z = r + secret * challenge.
//It returns true if g1^z = g1^r * (g1^secret) ^ challenge and g2^z = g2^r * (g2^secret) ^ challenge.
func (v *BTEqualityVerifier) Verify(z *big.Int) (bool, *BlindedTrans,
	*big.Int, *big.Int) {
	left1 := v.Group.Exp(v.g1, z)
	left2 := v.Group.Exp(v.g2, z)

	r11 := v.Group.Exp(v.t1, v.challenge)
	r12 := v.Group.Exp(v.t2, v.challenge)
	right1 := v.Group.Mul(r11, v.x1)
	right2 := v.Group.Mul(r12, v.x2)

	// transcript [(alpha1, beta1), hash(alpha1, beta1), z+alpha]
	// however, we are actually returning [alpha1, beta1, hash(alpha1, beta1), z+alpha]
	z1 := new(big.Int).Add(z, v.alpha)
	v.transcript.ZAlpha = z1

	G2 := v.Group.Exp(v.g2, v.gamma)
	T2 := v.Group.Exp(v.t2, v.gamma)

	if left1.Cmp(right1) == 0 && left2.Cmp(right2) == 0 {
		return true, v.transcript, G2, T2
	} else {
		return false, nil, nil, nil
	}
}
