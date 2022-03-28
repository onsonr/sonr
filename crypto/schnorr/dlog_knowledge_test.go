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
	"testing"

	"github.com/sonr-io/core/crypto/common"
	"github.com/stretchr/testify/assert"
)

// TestDLogKnowledge demonstrates how the prover proves that it knows (x_1,...,x_k)
// such that y = g_1^x_1 * ... * g_k^x_k where g_i are given generators of cyclic group G.
func TestDLogKnowledge(t *testing.T) {
	group, err := NewGroup(256)
	if err != nil {
		t.Errorf("error when creating Schnorr group: %v", err)
	}

	var bases [3]*big.Int
	for i := 0; i < len(bases); i++ {
		r := common.GetRandomInt(group.Q)
		bases[i] = group.Exp(group.G, r)
	}

	var secrets [3]*big.Int
	for i := 0; i < 3; i++ {
		secrets[i] = common.GetRandomInt(group.Q)
	}

	// y = g_1^x_1 * ... * g_k^x_k where g_i are bases and x_i are secrets
	y := big.NewInt(1)
	for i := 0; i < 3; i++ {
		f := group.Exp(bases[i], secrets[i])
		y = group.Mul(y, f)
	}

	prover, err := NewProver(group, secrets[:], bases[:], y)
	if err != nil {
		t.Errorf("error when creating Prover: %v", err)
	}
	verifier := NewVerifier(group)

	proofRandomData := prover.GetProofRandomData()
	verifier.SetProofRandomData(proofRandomData, bases[:], y)

	challenge := verifier.GetChallenge()
	proofData := prover.GetProofData(challenge)
	verified := verifier.Verify(proofData)

	assert.Equal(t, verified, true, "dlog knowledge proof does not work")
}
