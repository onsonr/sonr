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

// TestDFCommitmentSquare demonstrates how to prove that the commitment
// hides the square. Given c, prove that c = g^(x^2) * h^r (mod n).
func TestDFCommitmentSquare(t *testing.T) {
	receiver, err := NewReceiver(128, 80)
	if err != nil {
		t.Errorf("Error in NewReceiver: %v", err)
	}

	// n^2 is used for T - but any other value can be used as well
	T := new(big.Int).Mul(receiver.QRSpecialRSA.N, receiver.QRSpecialRSA.N)
	committer := NewCommitter(receiver.QRSpecialRSA.N,
		receiver.G, receiver.H, T, receiver.K)

	x := common.GetRandomInt(committer.QRSpecialRSA.N)
	x2 := new(big.Int).Mul(x, x)
	c, err := committer.GetCommitMsg(x2)
	if err != nil {
		t.Errorf("Error in computing commit msg: %v", err)
	}
	receiver.SetCommitment(c)

	challengeSpaceSize := 80
	prover, err := NewSquareProver(committer, x, challengeSpaceSize)
	if err != nil {
		t.Errorf("Error in instantiating SquareProver: %v", err)
	}

	verifier, err := NewSquareVerifier(receiver, prover.SmallCommitment, challengeSpaceSize)
	if err != nil {
		t.Errorf("Error in instantiating SquareVerifier: %v", err)
	}

	proofRandomData1, proofRandomData2 := prover.GetProofRandomData()
	verifier.SetProofRandomData(proofRandomData1, proofRandomData2)

	challenge := verifier.GetChallenge()
	s1, s21, s22 := prover.GetProofData(challenge)
	proved := verifier.Verify(s1, s21, s22)

	assert.Equal(t, true, proved, "DamgardFujisaki square proof failed.")
}
