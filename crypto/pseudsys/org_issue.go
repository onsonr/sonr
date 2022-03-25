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
	"math/big"

	"fmt"

	"github.com/sonr-io/core/crypto/schnorr"
)

type Cred struct {
	SmallAToGamma *big.Int
	SmallBToGamma *big.Int
	AToGamma      *big.Int
	BToGamma      *big.Int
	T1            *schnorr.BlindedTrans
	T2            *schnorr.BlindedTrans
}

func NewCred(aToGamma, bToGamma, AToGamma, BToGamma *big.Int,
	t1, t2 *schnorr.BlindedTrans) *Cred {
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
	group  *schnorr.Group
	secKey *SecKey

	// the following fields are needed for issuing a credential
	verifier *schnorr.Verifier
	prover1  *schnorr.BTEqualityProver
	prover2  *schnorr.BTEqualityProver
	a        *big.Int
	b        *big.Int
}

func NewCredIssuer(group *schnorr.Group, secKey *SecKey) *CredIssuer {
	// g1 = a_tilde, t1 = b_tilde,
	// g2 = a, t2 = b
	return &CredIssuer{
		group:    group,
		secKey:   secKey,
		verifier: schnorr.NewVerifier(group),
		prover1:  schnorr.NewBTEqualityProver(group),
		prover2:  schnorr.NewBTEqualityProver(group),
	}
}

func (i *CredIssuer) GetChallenge(a, b, x *big.Int) *big.Int {
	// TODO: check if (a, b) is registered; if not, close the session

	i.a = a
	i.b = b
	base := []*big.Int{a} // only one base
	i.verifier.SetProofRandomData(x, base, b)

	return i.verifier.GetChallenge()
}

// Verifies that user knows log_a(b). Sends back proof random data (g1^r, g2^r) for both equality proofs.
func (i *CredIssuer) Verify(z *big.Int) (
	*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	if verified := i.verifier.Verify([]*big.Int{z}); !verified {
		err := fmt.Errorf("authentication with organization failed")
		return nil, nil, nil, nil, nil, nil, err
	}

	A := i.group.Exp(i.b, i.secKey.S2)
	aA := i.group.Mul(i.a, A)
	B := i.group.Exp(aA, i.secKey.S1)

	x11, x12 := i.prover1.GetProofRandomData(i.secKey.S2, i.group.G, i.b)
	x21, x22 := i.prover2.GetProofRandomData(i.secKey.S1, i.group.G, aA)

	return x11, x12, x21, x22, A, B, nil

}

func (i *CredIssuer) GetProofData(challenge1,
	challenge2 *big.Int) (*big.Int, *big.Int) {
	z1 := i.prover1.GetProofData(challenge1)
	z2 := i.prover2.GetProofData(challenge2)
	return z1, z2
}
