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

// TestDFCommitmentMultiplication demonstrates how to prove that for given commitments
// c1 = g^x1 * h^r1, c2 = g^x2 * h^r2, c3 = g^x3 * h^r3, it holds x3 = x1 * x2
func TestDFCommitmentMultiplication(t *testing.T) {
	receiver1, err := NewReceiver(128, 80)
	if err != nil {
		t.Errorf("Error in NewReceiver: %v", err)
	}

	// n^2 is used for T - but any other value can be used as well
	T := new(big.Int).Mul(receiver1.QRSpecialRSA.N, receiver1.QRSpecialRSA.N)

	committer1 := NewCommitter(receiver1.QRSpecialRSA.N,
		receiver1.G, receiver1.H, T, receiver1.K)

	receiver2, err := NewReceiverFromParams(receiver1.QRSpecialRSA.GetPrimes(),
		receiver1.G, receiver1.H, receiver1.K)
	if err != nil {
		t.Errorf("Error in NewReceiver: %v", err)
	}
	committer2 := NewCommitter(receiver2.QRSpecialRSA.N,
		receiver2.G, receiver2.H, T, receiver2.K)

	receiver3, err := NewReceiverFromParams(receiver1.QRSpecialRSA.GetPrimes(),
		receiver1.G, receiver1.H, receiver1.K)
	if err != nil {
		t.Errorf("Error in NewReceiver: %v", err)
	}
	committer3 := NewCommitter(receiver3.QRSpecialRSA.N,
		receiver3.G, receiver3.H, T, receiver3.K)

	x1 := common.GetRandomInt(committer1.QRSpecialRSA.N)
	x1.Neg(x1) // test with negative
	x2 := common.GetRandomInt(committer2.QRSpecialRSA.N)
	x3 := new(big.Int).Mul(x1, x2)
	c1, err := committer1.GetCommitMsg(x1)
	if err != nil {
		t.Errorf("Error in computing commit msg: %v", err)
	}

	c2, err := committer2.GetCommitMsg(x2)
	if err != nil {
		t.Errorf("Error in computing commit msg: %v", err)
	}

	c3, err := committer3.GetCommitMsg(x3)
	if err != nil {
		t.Errorf("Error in computing commit msg: %v", err)
	}

	receiver1.SetCommitment(c1)
	receiver2.SetCommitment(c2)
	receiver3.SetCommitment(c3)

	challengeSpaceSize := 80
	prover := NewMultiplicationProver(committer1, committer2, committer3, challengeSpaceSize)
	verifier := NewMultiplicationVerifier(receiver1, receiver2, receiver3, challengeSpaceSize)

	d1, d2, d3 := prover.GetProofRandomData()
	verifier.SetProofRandomData(d1, d2, d3)

	challenge := verifier.GetChallenge()
	u1, u, v1, v2, v3 := prover.GetProofData(challenge)
	proved := verifier.Verify(u1, u, v1, v2, v3)

	assert.Equal(t, true, proved, "DamgardFujisaki multiplication proof failed.")
}
