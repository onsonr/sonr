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

	"github.com/sonr-io/core/crypto/schnorr"
)

type CredVerifier struct {
	group  *schnorr.Group
	secKey *SecKey

	verifier *schnorr.EqualityVerifier
	a        *big.Int
	b        *big.Int
}

func NewCredVerifier(group *schnorr.Group, secKey *SecKey) *CredVerifier {
	return &CredVerifier{
		group:    group,
		secKey:   secKey,
		verifier: schnorr.NewEqualityVerifier(group),
	}
}

func (v *CredVerifier) GetChallenge(a, b, a1, b1, x1, x2 *big.Int) *big.Int {
	// TODO: check if (a, b) is registered; if not, close the session

	v.a = a
	v.b = b
	challenge := v.verifier.GetChallenge(a, a1, b, b1, x1, x2)
	return challenge
}

func (v *CredVerifier) Verify(z *big.Int, cred *Cred, orgPubKeys *PubKey) bool {
	if !v.verifier.Verify(z) {
		return false
	}

	valid1 := cred.T1.Verify(v.group, v.group.G, orgPubKeys.H2,
		cred.SmallBToGamma, cred.AToGamma)

	aAToGamma := v.group.Mul(cred.SmallAToGamma, cred.AToGamma)
	valid2 := cred.T2.Verify(v.group, v.group.G, orgPubKeys.H1,
		aAToGamma, cred.BToGamma)

	return valid1 && valid2
}
