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

package pseudsys

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/sonr-io/core/crypto/common"
	"github.com/sonr-io/core/crypto/ec"
	"github.com/sonr-io/core/crypto/schnorr"
)

type CA struct {
	verifier   *schnorr.Verifier
	a          *big.Int
	b          *big.Int
	privateKey *ecdsa.PrivateKey
}

type CACert struct {
	BlindedA *big.Int
	BlindedB *big.Int
	R        *big.Int
	S        *big.Int
}

func NewCACert(blindedA, blindedB, r, s *big.Int) *CACert {
	return &CACert{
		BlindedA: blindedA,
		BlindedB: blindedB,
		R:        r,
		S:        s,
	}
}

func NewCA(group *schnorr.Group, d *big.Int, caPubKey *PubKey) *CA {
	c := ec.GetCurve(ec.P256)
	pubKey := ecdsa.PublicKey{Curve: c, X: caPubKey.H1, Y: caPubKey.H2}
	privateKey := ecdsa.PrivateKey{PublicKey: pubKey, D: d}

	schnorrVerifier := schnorr.NewVerifier(group)
	ca := CA{
		verifier:   schnorrVerifier,
		privateKey: &privateKey,
	}

	return &ca
}

func (ca *CA) GetChallenge(a, b, x *big.Int) *big.Int {
	// TODO: check if b is really a valuable external user's public master key; if not, close the session

	ca.a = a
	ca.b = b
	base := []*big.Int{a} // only one base
	ca.verifier.SetProofRandomData(x, base, b)
	challenge := ca.verifier.GetChallenge()
	return challenge
}

func (ca *CA) Verify(z *big.Int) (*CACert, error) {
	verified := ca.verifier.Verify([]*big.Int{z})
	if verified {
		r := common.GetRandomInt(ca.verifier.Group.Q)
		blindedA := ca.verifier.Group.Exp(ca.a, r)
		blindedB := ca.verifier.Group.Exp(ca.b, r)
		// blindedA, blindedB must be used only once (never use the same pair for two
		// different organizations)

		hashed := common.HashIntoBytes(blindedA, blindedB)
		r, s, err := ecdsa.Sign(rand.Reader, ca.privateKey, hashed)
		if err != nil {
			return nil, err
		} else {
			return NewCACert(blindedA, blindedB, r, s), nil
		}
	} else {
		return nil, fmt.Errorf("knowledge of secret was not verified")
	}
}
