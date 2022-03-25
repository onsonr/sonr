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

// TestDFCommitmentEquality demonstrates how to prove that the two commitments c1, c2
// hide the same number. Given c1, c2, prove that c1 = g1^x * h1^r1 (mod n1) and
// c2 = g2^x * h2^r2 (mod n2) for some x, r1, r2 (note that both commitments hide x).
func TestDFCommitmentEquality(t *testing.T) {
	receiver1, err := NewReceiver(128, 80)
	if err != nil {
		t.Errorf("Error in NewReceiver: %v", err)
	}

	// n^2 is used for T - but any other value can be used as well
	T := new(big.Int).Mul(receiver1.QRSpecialRSA.N, receiver1.QRSpecialRSA.N)
	committer1 := NewCommitter(receiver1.QRSpecialRSA.N,
		receiver1.G, receiver1.H, T, receiver1.K)

	receiver2, err := NewReceiver(128, 80)
	if err != nil {
		t.Errorf("Error in NewReceiver: %v", err)
	}
	committer2 := NewCommitter(receiver2.QRSpecialRSA.N,
		receiver2.G, receiver2.H, T, receiver2.K)

	x := common.GetRandomInt(committer1.T)
	c1, err := committer1.GetCommitMsg(x)
	if err != nil {
		t.Errorf("Error in computing commit msg: %v", err)
	}
	receiver1.SetCommitment(c1)

	c2, err := committer2.GetCommitMsg(x)
	if err != nil {
		t.Errorf("Error in computing commit msg: %v", err)
	}
	receiver2.SetCommitment(c2)

	challengeSpaceSize := 80
	prover := NewEqualityProver(committer1, committer2, challengeSpaceSize)
	verifier := NewEqualityVerifier(receiver1, receiver2, challengeSpaceSize)

	proofRandomData1, proofRandomData2 := prover.GetProofRandomData()
	verifier.SetProofRandomData(proofRandomData1, proofRandomData2)

	challenge := verifier.GetChallenge()
	s1, s21, s22 := prover.GetProofData(challenge)
	proved := verifier.Verify(s1, s21, s22)

	assert.Equal(t, true, proved, "DamgardFujisaki equality proof failed.")
}
