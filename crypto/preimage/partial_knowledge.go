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

package preimage

import (
	"math/big"

	"github.com/sonr-io/core/crypto"
	"github.com/sonr-io/core/crypto/common"
)

// ProvePartialPreimageKnowledge demonstrates how prover can prove that he knows f^(-1)(u1) and
// the verifier does not know whether knowledge of f^(-1)(u1) or f^(-1)(u2) was proved.
// Note that PartialDLogKnowledge is a special case of PartialPreimageKnowledge.
func ProvePartialPreimageKnowledge(homomorphism func(*big.Int) *big.Int, H crypto.Group,
	v1, u1, u2 *big.Int, iterations int) bool {
	prover := NewPartialProver(homomorphism, H, v1, u1, u2)
	verifier := NewPartialVerifier(homomorphism, H)

	// The proof needs to be repeated sequentially because one-bit challenges are used. Note
	// that when one-bit challenges are used, the prover has in one iteration 50% chances
	// that guesses the challenge. Thus, sufficient number of iterations is needed (like 80).
	// One-bit challenges are required - otherwise proof of knowledge extractor might
	// not work (algorithm to extract preimage when prover is used as a black-box and
	// rewinded to use the same first message in both executions).
	for j := 0; j < iterations; j++ {
		pair1, pair2 := prover.GetProofRandomData()
		verifier.SetProofRandomData(pair1, pair2)
		challenge := verifier.GetChallenge()
		c1, z1, c2, z2 := prover.GetProofData(challenge)
		if !verifier.Verify(c1, z1, c2, z2) {
			return false
		}
	}

	return true
}

type PartialProver struct {
	Homomorphism func(*big.Int) *big.Int
	H            crypto.Group
	v1           *big.Int
	u1           *big.Int
	u2           *big.Int
	r1           *big.Int
	c2           *big.Int
	z2           *big.Int
	ord          int
}

func NewPartialProver(homomorphism func(*big.Int) *big.Int, H crypto.Group,
	v1, u1, u2 *big.Int) *PartialProver {
	return &PartialProver{
		Homomorphism: homomorphism,
		H:            H,
		v1:           v1,
		u1:           u1,
		u2:           u2,
	}
}

// GetProofRandomData returns Homomorphism(r1) and Homomorphism(z2)/(u2^c2)
// in random order and where r1, z2, c2 are random from H.
func (p *PartialProver) GetProofRandomData() (*common.Pair, *common.Pair) {
	r1 := p.H.GetRandomElement()
	c2 := common.GetRandomInt(big.NewInt(2)) // challenges need to be binary
	z2 := p.H.GetRandomElement()
	p.r1 = r1
	p.c2 = c2
	p.z2 = z2
	x1 := p.Homomorphism(r1)
	x2 := p.Homomorphism(z2)
	u2ToC2 := p.H.Exp(p.u2, c2)
	u2ToC2Inv := p.H.Inv(u2ToC2)
	x2 = p.H.Mul(x2, u2ToC2Inv)

	// we need to make sure that the order does not reveal which secret we do know:
	ord := common.GetRandomInt(big.NewInt(2))
	pair1 := common.NewPair(x1, p.u1)
	pair2 := common.NewPair(x2, p.u2)

	if ord.Cmp(big.NewInt(0)) == 0 {
		p.ord = 0
		return pair1, pair2
	} else {
		p.ord = 1
		return pair2, pair1
	}
}

func (p *PartialProver) GetProofData(challenge *big.Int) (*big.Int, *big.Int,
	*big.Int, *big.Int) {
	c1 := new(big.Int).Xor(p.c2, challenge)
	// z1 = r*v^e
	z1 := p.H.Exp(p.v1, c1)
	z1 = p.H.Mul(p.r1, z1)

	if p.ord == 0 {
		return c1, z1, p.c2, p.z2
	} else {
		return p.c2, p.z2, c1, z1
	}
}

type PartialVerifier struct {
	Homomorphism func(*big.Int) *big.Int
	H            crypto.Group
	pair1        *common.Pair
	pair2        *common.Pair
	challenge    *big.Int
}

func NewPartialVerifier(homomorphism func(*big.Int) *big.Int,
	H crypto.Group) *PartialVerifier {
	return &PartialVerifier{
		Homomorphism: homomorphism,
		H:            H,
	}
}

func (v *PartialVerifier) SetProofRandomData(pair1, pair2 *common.Pair) {
	v.pair1 = pair1
	v.pair2 = pair2
}

func (v *PartialVerifier) GetChallenge() *big.Int {
	challenge := common.GetRandomInt(big.NewInt(2)) // challenges need to be binary
	v.challenge = challenge
	return challenge
}

func (v *PartialVerifier) verifyPair(pair *common.Pair,
	challenge, z *big.Int) bool {
	left := v.Homomorphism(z)
	r1 := v.H.Exp(pair.B, challenge)
	right := v.H.Mul(r1, pair.A)
	return left.Cmp(right) == 0
}

func (v *PartialVerifier) Verify(c1, z1, c2, z2 *big.Int) bool {
	c := new(big.Int).Xor(c1, c2)
	if c.Cmp(v.challenge) != 0 {
		return false
	}

	verified1 := v.verifyPair(v.pair1, c1, z1)
	verified2 := v.verifyPair(v.pair2, c2, z2)
	return verified1 && verified2
}
