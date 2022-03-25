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

// TestDFCommitmentPositive demonstrates how to prove that the commitment
// hides a positive number. Given c, prove that c = g^x * h^r (mod n) where x >= 0.
func TestDFCommitmentPositive(t *testing.T) {
	receiver, err := NewReceiver(128, 80)
	if err != nil {
		t.Errorf("error in NewReceiver: %v", err)
	}

	// n^2 is used for T - but any other value can be used as well
	T := new(big.Int).Mul(receiver.QRSpecialRSA.N, receiver.QRSpecialRSA.N)
	committer := NewCommitter(receiver.QRSpecialRSA.N,
		receiver.G, receiver.H, T, receiver.K)

	x := common.GetRandomInt(committer.QRSpecialRSA.N)
	c, err := committer.GetCommitMsg(x)
	if err != nil {
		t.Errorf("error in computing commit msg: %v", err)
	}
	receiver.SetCommitment(c)
	_, r := committer.GetDecommitMsg()

	challengeSpaceSize := 80
	prover, err := NewPositiveProver(committer, x, r,
		challengeSpaceSize)
	if err != nil {
		t.Errorf("error in instantiating PositiveProver: %v", err)
	}

	smallCommitments, bigCommitments := prover.GetVerifierInitializationData()
	verifier, err := NewPositiveVerifier(receiver, receiver.Commitment,
		smallCommitments, bigCommitments, challengeSpaceSize)
	if err != nil {
		t.Errorf("error in instantiating PositiveVerifier: %v", err)
	}

	proofRandomData := prover.GetProofRandomData()
	challenges := verifier.GetChallenges()
	err = verifier.SetProofRandomData(proofRandomData)
	if err != nil {
		t.Errorf("error when calling SetProofRandomData: %v", err)
	}
	proofData := prover.GetProofData(challenges)
	proved := verifier.Verify(proofData)
	assert.Equal(t, true, proved, "DamgardFujisaki positive proof failed.")
}
