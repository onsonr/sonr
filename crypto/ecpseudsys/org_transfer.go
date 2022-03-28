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

	"github.com/sonr-io/core/crypto/ec"
	"github.com/sonr-io/core/crypto/ecschnorr"
	"github.com/sonr-io/core/crypto/pseudsys"
)

type CredVerifier struct {
	secKey *pseudsys.SecKey

	verifier *ecschnorr.EqualityVerifier
	a        *ec.GroupElement
	b        *ec.GroupElement
	curve    ec.Curve
}

func NewCredVerifier(secKey *pseudsys.SecKey, c ec.Curve) *CredVerifier {
	return &CredVerifier{
		secKey:   secKey,
		verifier: ecschnorr.NewEqualityVerifier(c),
		curve:    c,
	}
}

// TODO GetChallenge?
func (v *CredVerifier) GetChallenge(a, b, a1, b1,
	x1, x2 *ec.GroupElement) *big.Int {
	// TODO: check if (a, b) is registered; if not, close the session

	v.a = a
	v.b = b

	return v.verifier.GetChallenge(a, a1, b, b1, x1, x2)
}

func (v *CredVerifier) Verify(z *big.Int,
	credential *Cred, orgPubKeys *PubKey) bool {
	verified := v.verifier.Verify(z)
	if !verified {
		return false
	}

	g := ec.NewGroupElement(v.verifier.Group.Curve.Params().Gx,
		v.verifier.Group.Curve.Params().Gy)

	valid1 := credential.T1.Verify(ec.P256, g, orgPubKeys.H2,
		credential.SmallBToGamma, credential.AToGamma)

	aAToGamma := v.verifier.Group.Mul(credential.SmallAToGamma, credential.AToGamma)
	valid2 := credential.T2.Verify(ec.P256, g, orgPubKeys.H1,
		aAToGamma, credential.BToGamma)

	return valid1 && valid2
}
