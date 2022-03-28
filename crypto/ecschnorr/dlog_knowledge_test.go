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
	"testing"

	"github.com/sonr-io/core/crypto/common"
	"github.com/sonr-io/core/crypto/ec"
	"github.com/stretchr/testify/assert"
)

// TestECDLogKnowledge demonstrates how prover can prove the knowledge of log_g1(t1) - that
// means g1^secret = t1 in EC group.
func TestECDLogKnowledge(t *testing.T) {
	group := ec.NewGroup(ec.P256)
	exp1 := common.GetRandomInt(group.Q)
	a1 := group.ExpBaseG(exp1)
	secret := common.GetRandomInt(group.Q)
	b1 := group.Exp(a1, secret)

	prover := NewProver(ec.P256)
	verifier := NewVerifier(ec.P256)

	x := prover.GetProofRandomData(secret, a1)
	verifier.SetProofRandomData(x, a1, b1)

	challenge := verifier.GetChallenge()
	z := prover.GetProofData(challenge)
	verified := verifier.Verify(z)

	assert.Equal(t, verified, true, "dlog equality proof does not work")
}
