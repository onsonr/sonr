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

// ProvePreimageKnowledge demonstrates how given Homomorphism f:H->G and element u from G
// prover can prove the knowledge of v such that f(v) = u.
func ProvePreimageKnowledge(homomorphism func(*big.Int) *big.Int, H crypto.Group,
	u, v *big.Int, iterations int) bool {
	prover := NewProver(homomorphism, H, v)
	verifier := NewVerifier(homomorphism, H, u)

	// The proof needs to be repeated sequentially because one-bit challenges are used. Note
	// that when one-bit challenges are used, the prover has in one iteration 50% chances
	// that guesses the challenge. Thus, sufficient number of iterations is needed (like 80).
	// One-bit challenges are required - otherwise proof of knowledge extractor might
	// not work (algorithm to extract preimage when prover is used as a black-box and
	// rewinded to use the same first message in both executions).
	for j := 0; j < iterations; j++ {
		proofRandomData := prover.GetProofRandomData()
		verifier.SetProofRandomData(proofRandomData)
		challenge := verifier.GetChallenge()
		z := prover.GetProofData(challenge)
		if !verifier.Verify(z) {
			return false
		}
	}

	return true
}

// Prover proves that it knows v such that f(v) = u, given homomorphism f: H -> G
// and u from group G.
// This is a generalized Schnorr prover, but one-bit challenges need to be used
// to enable extractor (more to be added in docs).
type Prover struct {
	Homomorphism func(*big.Int) *big.Int
	H            crypto.Group
	v            *big.Int
	r            *big.Int
}

func NewProver(homomorphism func(*big.Int) *big.Int, H crypto.Group,
	v *big.Int) *Prover {
	return &Prover{
		Homomorphism: homomorphism,
		H:            H,
		v:            v,
	}
}

// Chooses random r from H and returns QOneWayHomomorpism(r).
func (p *Prover) GetProofRandomData() *big.Int {
	// TODO: see SchnorrProver comment, note that here setting of the required parameters (v) is
	// done in the constructor.

	// x = Homomorphism(r), where r is random
	r := p.H.GetRandomElement()
	p.r = r
	x := p.Homomorphism(r)
	return x
}

// GetProofData receives challenge defined by a verifier, and returns z = r * v^challenge.
func (p *Prover) GetProofData(challenge *big.Int) *big.Int {
	// z = r * v^challenge
	z := p.H.Exp(p.v, challenge)
	z = p.H.Mul(p.r, z)
	return z
}

type Verifier struct {
	Homomorphism func(*big.Int) *big.Int
	H            crypto.Group
	challenge    *big.Int
	u            *big.Int
	x            *big.Int
}

func NewVerifier(homomorphism func(*big.Int) *big.Int, H crypto.Group,
	u *big.Int) *Verifier {
	return &Verifier{
		Homomorphism: homomorphism,
		H:            H,
		u:            u,
	}
}

func (v *Verifier) SetProofRandomData(x *big.Int) {
	v.x = x
}

func (v *Verifier) GetChallenge() *big.Int {
	challenge := common.GetRandomInt(big.NewInt(2)) // challenges need to be binary
	v.challenge = challenge
	return challenge
}

// It receives z = r * v^challenge. It returns true if Homomorphism(z) = x * u^challenge, otherwise false.
func (v *Verifier) Verify(z *big.Int) bool {
	left := v.Homomorphism(z)
	right := v.H.Exp(v.u, v.challenge)
	right = v.H.Mul(v.x, right)
	return left.Cmp(right) == 0
}
