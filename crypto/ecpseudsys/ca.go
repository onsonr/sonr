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

package ecpseudsys

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/sonr-io/core/crypto/common"
	"github.com/sonr-io/core/crypto/ec"
	"github.com/sonr-io/core/crypto/ecschnorr"
	"github.com/sonr-io/core/crypto/pseudsys"
)

type CA struct {
	group      *ec.Group
	verifier   *ecschnorr.Verifier
	a          *ec.GroupElement
	b          *ec.GroupElement
	privateKey *ecdsa.PrivateKey
}

type CACert struct {
	BlindedA *ec.GroupElement
	BlindedB *ec.GroupElement
	R        *big.Int
	S        *big.Int
}

func NewCACert(blindedA, blindedB *ec.GroupElement, r, s *big.Int) *CACert {
	return &CACert{
		BlindedA: blindedA,
		BlindedB: blindedB,
		R:        r,
		S:        s,
	}
}

func NewCA(d *big.Int, caPubKey *pseudsys.PubKey, curve ec.Curve) *CA {
	c := ec.GetCurve(curve)
	pubKey := ecdsa.PublicKey{Curve: c, X: caPubKey.H1, Y: caPubKey.H2}
	privateKey := ecdsa.PrivateKey{PublicKey: pubKey, D: d}

	return &CA{
		verifier:   ecschnorr.NewVerifier(curve),
		privateKey: &privateKey,
	}
}

func (ca *CA) GetChallenge(a, b, x *ec.GroupElement) *big.Int {
	// TODO: check if b is really a valuable external user's public master key; if not, close the session

	ca.a = a
	ca.b = b
	ca.verifier.SetProofRandomData(x, a, b)

	return ca.verifier.GetChallenge()
}

func (ca *CA) Verify(z *big.Int) (*CACert, error) {
	if verified := ca.verifier.Verify(z); !verified {
		return nil, fmt.Errorf("knowledge of secret was not verified")
	}

	r := common.GetRandomInt(ca.verifier.Group.Q)
	blindedA := ca.verifier.Group.Exp(ca.a, r)
	blindedB := ca.verifier.Group.Exp(ca.b, r)
	// blindedA, blindedB must be used only once (never use the same pair for two
	// different organizations)

	hashed := common.HashIntoBytes(blindedA.X, blindedA.Y, blindedB.X, blindedB.Y)
	r, s, err := ecdsa.Sign(rand.Reader, ca.privateKey, hashed)
	if err != nil {
		return nil, err
	}

	return NewCACert(blindedA, blindedB, r, s), nil
}
