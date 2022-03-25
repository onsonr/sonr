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

	"fmt"
)

// SquareProver proves that the commitment hides the square. Given c,
// prove that c = g^(x^2) * h^r (mod n).
type SquareProver struct {
	*EqualityProver
	// We have two commitments with the same value: SmallCommitment = g^x * h^r1 and
	// c = SmallCommitment^x * h^r2. Also c = g^(x^2) * h^r.
	SmallCommitment *big.Int
}

func NewSquareProver(committer *Committer,
	x *big.Int, challengeSpaceSize int) (*SquareProver, error) {

	// Input committer contains c = g^(x^2) * h^r (mod n).
	// We now create two committers - committer1 will contain SmallCommitment = g^x * h^r1 (mod n),
	// committer2 will contain the same c as committer, but using a different
	// base c = SmallCommitment^x * h^r2.
	// Note that c = SmallCommitment^x * h^r2 = g^(x^2) * h^(r1*x) * h^r2, so we choose r2 = r - r1*x.
	// SquareProver proves that committer1 and committer2 hide the same value (x) -
	// using EqualityProver.

	committer1 := NewCommitter(committer.QRSpecialRSA.N,
		committer.G, committer.H, committer.T, committer.K)
	smallCommitment, err := committer1.GetCommitMsg(x)
	if err != nil {
		return nil, fmt.Errorf("error when creating commit msg")
	}

	committer2 := NewCommitter(committer.QRSpecialRSA.N,
		smallCommitment, committer.H, committer.T, committer.K)
	_, r := committer.GetDecommitMsg()
	_, r1 := committer1.GetDecommitMsg()
	r1x := new(big.Int).Mul(r1, x)
	r2 := new(big.Int).Sub(r, r1x)

	// we already know the commitment (it is c), so we ignore the first variable -
	// just need to set the committer2 committedValue and r:
	_, err = committer2.GetCommitMsgWithGivenR(x, r2)
	if err != nil {
		return nil, fmt.Errorf("error when creating commit msg with given r")
	}

	prover := NewEqualityProver(committer1, committer2, challengeSpaceSize)

	return &SquareProver{
		EqualityProver:  prover,
		SmallCommitment: smallCommitment,
	}, nil
}

type SquareVerifier struct {
	*EqualityVerifier
}

func NewSquareVerifier(receiver *Receiver,
	c1 *big.Int, challengeSpaceSize int) (*SquareVerifier, error) {

	receiver1, err := NewReceiverFromParams(receiver.QRSpecialRSA.GetPrimes(),
		receiver.G, receiver.H, receiver.K)
	if err != nil {
		return nil, fmt.Errorf("error when calling NewReceiverFromParams")
	}
	receiver1.SetCommitment(c1)

	receiver2, err := NewReceiverFromParams(receiver.QRSpecialRSA.GetPrimes(),
		c1, receiver.H, receiver.K)
	if err != nil {
		return nil, fmt.Errorf("error when calling NewReceiverFromParams")
	}
	receiver2.SetCommitment(receiver.Commitment)

	verifier := NewEqualityVerifier(receiver1, receiver2, challengeSpaceSize)

	return &SquareVerifier{
		verifier,
	}, nil
}
