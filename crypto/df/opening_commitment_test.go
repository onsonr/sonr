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

package df

import (
	"math/big"
	"testing"

	"github.com/sonr-io/core/crypto/common"
	"github.com/stretchr/testify/assert"
)

// TestDFCommitmentOpening demonstrates how to prove that you can open DamgardFujisaki commitment.
func TestDFCommitmentOpening(t *testing.T) {
	receiver, err := NewReceiver(128, 80)
	if err != nil {
		t.Errorf("Error in NewReceiver: %v", err)
	}

	// n^2 is used for T - but any other value can be used as well
	T := new(big.Int).Mul(receiver.QRSpecialRSA.N, receiver.QRSpecialRSA.N)
	committer := NewCommitter(receiver.QRSpecialRSA.N,
		receiver.G, receiver.H, T, receiver.K)

	x := common.GetRandomInt(committer.T)
	c, err := committer.GetCommitMsg(x)
	if err != nil {
		t.Errorf("Error in computing commit msg: %v", err)
	}
	receiver.SetCommitment(c)

	challengeSpaceSize := 80
	prover := NewOpeningProver(committer, challengeSpaceSize)
	verifier := NewOpeningVerifier(receiver, challengeSpaceSize)

	proofRandomData := prover.GetProofRandomData()
	verifier.SetProofRandomData(proofRandomData)

	challenge := verifier.GetChallenge()
	s1, s2 := prover.GetProofData(challenge)
	proved := verifier.Verify(s1, s2)

	assert.Equal(t, true, proved, "DamgardFujisaki opening proof failed.")
}
