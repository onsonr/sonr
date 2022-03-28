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

// TestDFCommitmentRange demonstrates how to prove that the commitment
// hides a number x such that a <= x <= b. Given c, prove that c = g^x * h^r (mod n) where a<= x <= b.
func TestDFCommitmentRange(t *testing.T) {
	receiver, err := NewReceiver(128, 80)
	if err != nil {
		t.Errorf("error in NewReceiver: %v", err)
	}

	// n^2 is used for T - but any other value can be used as well
	T := new(big.Int).Mul(receiver.QRSpecialRSA.N, receiver.QRSpecialRSA.N)
	committer := NewCommitter(receiver.QRSpecialRSA.N,
		receiver.G, receiver.H, T, receiver.K)

	x := common.GetRandomInt(committer.QRSpecialRSA.N)
	a := new(big.Int).Sub(x, big.NewInt(10))
	b := new(big.Int).Add(x, big.NewInt(10))
	c, err := committer.GetCommitMsg(x)
	if err != nil {
		t.Errorf("error in computing commit msg: %v", err)
	}
	receiver.SetCommitment(c)

	challengeSpaceSize := 80
	prover, err := NewRangeProver(committer, x, a, b, challengeSpaceSize)
	if err != nil {
		t.Errorf("error in instantiating RangeProver: %v", err)
	}

	smallCommitments1, bigCommitments1, smallCommitments2, bigCommitments2 :=
		prover.GetVerifierInitializationData()
	verifier, err := NewRangeVerifier(receiver, a, b, smallCommitments1,
		bigCommitments1, smallCommitments2, bigCommitments2, challengeSpaceSize)
	if err != nil {
		t.Errorf("error in instantiating RangeVerifier: %v", err)
	}

	proofRandomData1, proofRandomData2 := prover.GetProofRandomData()
	challenges1, challenges2 := verifier.GetChallenges()
	err = verifier.SetProofRandomData(proofRandomData1, proofRandomData2)
	if err != nil {
		t.Errorf("error when calling SetProofRandomData: %v", err)
	}

	proofData1, proofData2, err := prover.GetProofData(challenges1, challenges2)
	if err != nil {
		t.Errorf("error when calling GetProofData: %v", err)
	}

	proved, err := verifier.Verify(proofData1, proofData2)
	if err != nil {
		t.Errorf("error when calling Verify: %v", err)
	}
	assert.Equal(t, true, proved, "DamgardFujisaki range proof failed.")
}
