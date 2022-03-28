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
	"fmt"
	"math/big"

	"github.com/sonr-io/core/crypto/common"
	"github.com/sonr-io/core/crypto/qr"
)

// Based on:
// I. Damgard and E. Fujisaki. An integer commitment scheme based on groups with hidden order. http://eprint.iacr.org/2001, 2001.
//
// Damgard-Fujisaki is statistically-hiding integer commitment scheme which works in groups
// with hidden order, like RSASpecial.
// This scheme can be used to commit to any integer (that is not generally true for commitment
// schemes, usually there is some boundary), however a boundary (denoted by T) is needed
// for the associated proofs.

// df represents what is common in Committer and Receiver.
type df struct {
	QRSpecialRSA *qr.RSASpecial
	H            *big.Int
	G            *big.Int // G = H^alpha % RSASpecial.N, where alpha is chosen when Receiver is created (Committer does not know alpha)
	K            int      // security parameter
}

// ComputeCommit returns G^a * H^r % group.N for a given a and r. Note that this is exactly
// the commitment, but with a given a and r. It serves as a helper function for
// associated proofs where g^x * h^y % group.N needs to be computed several times.
func (df *df) ComputeCommit(a, r *big.Int) *big.Int {
	tmp1 := df.QRSpecialRSA.Exp(df.G, a)
	tmp2 := df.QRSpecialRSA.Exp(df.H, r)
	c := df.QRSpecialRSA.Mul(tmp1, tmp2)
	return c
}

type Committer struct {
	df
	B              int      // 2^B is upper bound estimation for group order, it can be len(RSASpecial.N) - 2
	T              *big.Int // we can commit to values between -T and T
	committedValue *big.Int
	r              *big.Int
}

// TODO: switch h and g
func NewCommitter(n, g, h, t *big.Int, k int) *Committer {
	// n.BitLen() - 2 is used as B
	return &Committer{df: df{
		QRSpecialRSA: qr.NewRSApecialPublic(n),
		G:            g,
		H:            h,
		K:            k},
		B: n.BitLen() - 2,
		T: t}
}

// TODO: the naming is not OK because it also sets committer.committedValue and committer.r
func (c *Committer) GetCommitMsg(a *big.Int) (*big.Int, error) {
	abs := new(big.Int).Abs(a)
	if abs.Cmp(c.T) != -1 {
		return nil, fmt.Errorf("committed value needs to be in (-T, T)")
	}
	// c = G^a * H^r % group.N
	// choose r from 2^(B + k)
	exp := big.NewInt(int64(c.B + c.K))
	boundary := new(big.Int).Exp(big.NewInt(2), exp, nil)
	r := common.GetRandomInt(boundary)
	commitment := c.ComputeCommit(a, r)

	c.committedValue = a
	c.r = r
	return commitment, nil
}

func (c *Committer) GetCommitMsgWithGivenR(a, r *big.Int) (*big.Int, error) {
	abs := new(big.Int).Abs(a)
	if abs.Cmp(c.T) != -1 {
		return nil, fmt.Errorf("committed value needs to be in (-T, T)")
	}
	// c = G^a * H^r % group.N
	commitment := c.ComputeCommit(a, r)
	c.committedValue = a
	c.r = r
	return commitment, nil
}

func (c *Committer) GetDecommitMsg() (*big.Int, *big.Int) {
	return c.committedValue, c.r
}

type Receiver struct {
	df
	Commitment *big.Int
}

// NewReceiver receives two parameters: safePrimeBitLength tells the length of the
// primes in RSASpecial group and should be at least 1024, k is security parameter on which it depends
// the hiding property (commitment c = G^a * H^r where r is chosen randomly from (0, 2^(B+k)) - the distribution of
// c is statistically close to uniform, 2^B is upper bound estimation for group order).
func NewReceiver(safePrimeBitLength, k int) (*Receiver, error) {
	qr, err := qr.NewRSASpecial(safePrimeBitLength)
	if err != nil {
		return nil, err
	}

	h, err := qr.GetRandomGenerator()
	if err != nil {
		return nil, err
	}

	alpha := common.GetRandomInt(qr.Order)
	g := qr.Exp(h, alpha)
	if err != nil {
		return nil, err
	}

	return &Receiver{df: df{
			QRSpecialRSA: qr,
			H:            h,
			G:            g,
			K:            k}},
		nil
}

// NewReceiverFromParams returns an instance of a receiver with the
// parameters as given by input. Different instances are needed because
// each sets its own Commitment value.
func NewReceiverFromParams(specialRSAPrimes *qr.RSASpecialPrimes, g, h *big.Int,
	k int) (
	*Receiver, error) {
	group, err := qr.NewRSASpecialFromParams(specialRSAPrimes)
	if err != nil {
		return nil, err
	}

	return &Receiver{df: df{
		QRSpecialRSA: group,
		G:            g,
		H:            h,
		K:            k},
	}, nil
}

// When receiver receives a commitment, it stores the value using SetCommitment method.
func (r *Receiver) SetCommitment(c *big.Int) {
	r.Commitment = c
}

func (r *Receiver) CheckDecommitment(R, a *big.Int) bool {
	tmp1 := r.QRSpecialRSA.Exp(r.G, a)
	tmp2 := r.QRSpecialRSA.Exp(r.H, R)
	c := r.QRSpecialRSA.Mul(tmp1, tmp2)

	return c.Cmp(r.Commitment) == 0
}
