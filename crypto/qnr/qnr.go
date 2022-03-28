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

package qnr

// Statistical zero-knowledge proof of quadratic non-residousity.

import (
	"fmt"
	"math/big"

	"github.com/sonr-io/core/crypto/common"
	"github.com/sonr-io/core/crypto/qr"
)

// ProveQNR demonstrates how the prover can prove that y is not quadratic residue (there does
// not exist element y1 such that y1^2 = y in group RSA.
func ProveQNR(y *big.Int, qr *qr.RSA) (bool, error) {
	prover := NewProver(qr, y)
	verifier := NewQNRVerifier(qr, y)
	m := qr.N.BitLen()

	for i := 0; i < m; i++ {
		w, pairs := verifier.GetChallenge()
		prover.SetProofRandomData(w)
		// get challenge from prover for proving that verifier is not cheating
		randVector := prover.GetChallenge()

		verProof := verifier.GetProofData(randVector)

		verifierIsHonest := prover.Verify(pairs, verProof)
		if !verifierIsHonest {
			err := fmt.Errorf("verifier is not honest")
			return false, err
		}

		typ, err := prover.GetProofData(w)
		if err != nil {
			return false, nil
		}

		proved := verifier.Verify(typ)
		if !proved {
			return false, nil
		}
	}
	return true, nil
}

type Prover struct {
	QR *qr.RSA
	Y  *big.Int
	w  *big.Int
}

func NewProver(qr *qr.RSA, y *big.Int) *Prover {
	return &Prover{
		QR: qr,
		Y:  y,
	}
}

func (p *Prover) GetChallenge() []int {
	m := p.QR.N.BitLen()
	var randVector []int
	for i := 0; i < m; i++ {
		// todo: remove big.Int
		b := common.GetRandomInt(big.NewInt(2)) // 0 or 1
		var r int
		if b.Cmp(big.NewInt(0)) == 0 {
			r = 0
		} else {
			r = 1
		}
		randVector = append(randVector, r)
	}
	return randVector
}

func (p *Prover) SetProofRandomData(w *big.Int) {
	p.w = w
}

func (p *Prover) GetProofData(challenge *big.Int) (int, error) {
	isQR, err := p.QR.IsElementInGroup(challenge)
	if err != nil {
		return 0, err
	}
	var typ int
	if isQR {
		typ = 1 // challenge is of type r^2
	} else {
		typ = 2 // challenge is of type r^2 * y
	}
	return typ, nil
}

func (p *Prover) Verify(pairs, verProof []*common.Pair) bool {
	for ind, proofPair := range verProof {
		if proofPair.B.Cmp(big.NewInt(0)) == 0 {
			t := p.QR.Mul(proofPair.A, proofPair.A)
			Aw := p.QR.Mul(pairs[ind].A, p.w)
			Bw := p.QR.Mul(pairs[ind].B, p.w)

			if (t.Cmp(Aw) != 0) && (t.Cmp(Bw) != 0) {
				return false
			}
		} else {
			r1Squared := p.QR.Mul(proofPair.A, proofPair.A)
			r2Squared := p.QR.Mul(proofPair.B, proofPair.B)
			r2Squaredy := p.QR.Mul(r2Squared, p.Y)
			pair := pairs[ind]

			if !((r1Squared.Cmp(pair.A) == 0 && r2Squaredy.Cmp(pair.B) == 0) ||
				(r1Squared.Cmp(pair.B) == 0 && r2Squaredy.Cmp(pair.A) == 0)) {
				return false
			}
		}
	}
	return true
}

type Verifier struct {
	QR    *qr.RSA
	x     *big.Int
	y     *big.Int
	typ   int
	r     *big.Int
	pairs []*common.Pair
}

func NewQNRVerifier(qr *qr.RSA, y *big.Int) *Verifier {
	return &Verifier{
		QR: qr,
		y:  y,
	}
}

func (v *Verifier) GetChallenge() (*big.Int, []*common.Pair) {
	r := common.GetRandomInt(v.QR.N)
	// checking that gcd(r, N) = 1 is not needed as the probability is low
	v.r = r
	v.pairs = v.pairs[:0] // clear v.pairs
	r2 := v.QR.Mul(r, r)

	b := common.GetRandomInt(big.NewInt(2)) // 0 or 1
	var w *big.Int

	if b.Cmp(big.NewInt(0)) == 0 {
		w = r2
		v.typ = 1
	} else {
		w = v.QR.Mul(r2, v.y)
		v.typ = 2
	}

	m := v.QR.N.BitLen()
	var pairs []*common.Pair
	for i := 0; i < m; i++ {
		r1 := common.GetRandomInt(v.QR.N)
		r2 := common.GetRandomInt(v.QR.N)
		aj := v.QR.Mul(r1, r1) // r1^2
		bj := v.QR.Mul(r2, r2)
		bj = v.QR.Mul(bj, v.y) // r2^2 * y

		bitj := common.GetRandomInt(big.NewInt(2)) // 0 or 1

		v.pairs = append(v.pairs, &common.Pair{A: r1, B: r2})

		var pair *common.Pair
		if bitj.Cmp(big.NewInt(1)) == 0 {
			pair = &common.Pair{
				A: aj,
				B: bj,
			}
		} else {
			pair = &common.Pair{
				A: bj,
				B: aj,
			}
		}
		pairs = append(pairs, pair)
	}

	return w, pairs
}

func (v *Verifier) GetProofData(randVector []int) []*common.Pair {
	var pairs []*common.Pair
	for ind, i := range randVector {
		if i == 0 {
			pair := &common.Pair{
				A: v.pairs[ind].A,
				B: v.pairs[ind].B,
			}
			pairs = append(pairs, pair)
		} else {
			if v.typ == 1 { // w = r^2
				r1 := v.pairs[ind].A
				t := v.QR.Mul(v.r, r1)
				// t = r * r1
				pairs = append(pairs, &common.Pair{A: t, B: big.NewInt(0)})
			} else { // w = r^2 * y
				r2 := v.pairs[ind].B
				t := v.QR.Mul(v.r, r2)
				t = v.QR.Mul(t, v.y)
				// t = r * r2 * y
				pairs = append(pairs, &common.Pair{A: t, B: big.NewInt(0)})
			}
		}
	}
	return pairs
}

func (v *Verifier) Verify(typ int) bool {
	return v.typ == typ
}
