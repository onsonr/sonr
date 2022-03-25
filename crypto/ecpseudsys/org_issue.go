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
	"math/big"

	"fmt"

	"github.com/sonr-io/core/crypto/ec"
	"github.com/sonr-io/core/crypto/ecschnorr"
	"github.com/sonr-io/core/crypto/pseudsys"
)

type Cred struct {
	SmallAToGamma *ec.GroupElement
	SmallBToGamma *ec.GroupElement
	AToGamma      *ec.GroupElement
	BToGamma      *ec.GroupElement
	T1            *ecschnorr.BlindedTrans
	T2            *ecschnorr.BlindedTrans
}

func NewCred(aToGamma, bToGamma, AToGamma, BToGamma *ec.GroupElement,
	t1, t2 *ecschnorr.BlindedTrans) *Cred {
	return &Cred{
		SmallAToGamma: aToGamma,
		SmallBToGamma: bToGamma,
		AToGamma:      AToGamma,
		BToGamma:      BToGamma,
		T1:            t1,
		T2:            t2,
	}
}

type CredIssuer struct {
	secKey *pseudsys.SecKey

	// the following fields are needed for issuing a credential
	verifier *ecschnorr.Verifier
	prover1  *ecschnorr.BTEqualityProver
	prover2  *ecschnorr.BTEqualityProver
	a        *ec.GroupElement
	b        *ec.GroupElement
}

func NewCredIssuer(secKey *pseudsys.SecKey, curveType ec.Curve) *CredIssuer {
	// g1 = a_tilde, t1 = b_tilde,
	// g2 = a, t2 = b
	return &CredIssuer{
		secKey:   secKey,
		verifier: ecschnorr.NewVerifier(curveType),
		prover1:  ecschnorr.NewBTEqualityProver(curveType),
		prover2:  ecschnorr.NewBTEqualityProver(curveType),
	}
}

// TODO GetChallenge?
func (i *CredIssuer) GetChallenge(a, b, x *ec.GroupElement) *big.Int {
	// TODO: check if (a, b) is registered; if not, close the session

	i.a = a
	i.b = b
	i.verifier.SetProofRandomData(x, a, b)

	return i.verifier.GetChallenge()
}

// Verifies that user knows log_a(b). Sends back proof random data (g1^r, g2^r) for both equality proofs.
func (i *CredIssuer) Verify(z *big.Int) (
	*ec.GroupElement, *ec.GroupElement, *ec.GroupElement,
	*ec.GroupElement, *ec.GroupElement, *ec.GroupElement, error) {
	if verified := i.verifier.Verify(z); !verified {
		err := fmt.Errorf("authentication with organization failed")
		return nil, nil, nil, nil, nil, nil, err
	}

	A := i.verifier.Group.Exp(i.b, i.secKey.S2)
	aA := i.verifier.Group.Mul(i.a, A)
	B := i.verifier.Group.Exp(aA, i.secKey.S1)

	g1 := ec.NewGroupElement(i.verifier.Group.Curve.Params().Gx,
		i.verifier.Group.Curve.Params().Gy)
	g2 := ec.NewGroupElement(i.b.X, i.b.Y)

	x11, x12 := i.prover1.GetProofRandomData(i.secKey.S2, g1, g2)
	x21, x22 := i.prover2.GetProofRandomData(i.secKey.S1, g1, aA)

	return x11, x12, x21, x22, A, B, nil
}

func (i *CredIssuer) GetProofData(challenge1, challenge2 *big.Int) (*big.Int, *big.Int) {
	z1 := i.prover1.GetProofData(challenge1)
	z2 := i.prover2.GetProofData(challenge2)
	return z1, z2
}
