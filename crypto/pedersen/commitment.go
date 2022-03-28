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

package pedersen

import (
	"math/big"

	"fmt"

	"github.com/sonr-io/core/crypto/common"
	"github.com/sonr-io/core/crypto/schnorr"
)

// TODO: might be better to have only one method (like GetCommitment) instead of
// GetCommitMsg and GetDecommit msg, which would return c, r. Having two methods and storing r into
// committer might be awkward when having more commitments (like in RSABasedCommitment when
// proving multiplication property, see commitments_test.go).

// Committer first needs to know H (it gets it from the receiver).
// Then committer can commit to some value x - it sends to receiver c = g^x * h^r.
// When decommitting, committer sends to receiver r, x; receiver checks whether c = g^x * h^r.

type Params struct {
	Group *schnorr.Group
	H     *big.Int
	a     *big.Int
	// trapdoor a can be nil (doesn't need to be known), it is rarely needed -
	// for example in one of techniques to turn sigma to ZKP
}

func NewParams(group *schnorr.Group, H, a *big.Int) *Params {
	return &Params{
		Group: group,
		H:     H, // H = group.G^a
		a:     a,
	}
}

func GenerateParams(bitLengthGroupOrder int) (*Params, error) {
	group, err := schnorr.NewGroup(bitLengthGroupOrder)
	if err != nil {
		return nil, fmt.Errorf("error when creating SchnorrGroup: %s", err)
	}
	a := common.GetRandomInt(group.Q)
	return NewParams(group, group.Exp(group.G, a), a), nil
}

type Committer struct {
	Params         *Params
	Commitment     *big.Int
	committedValue *big.Int
	r              *big.Int
}

func NewCommitter(pedersenParams *Params) *Committer {
	committer := Committer{
		Params: pedersenParams,
	}
	return &committer
}

// It receives a value x (to this value a commitment is made), chooses a random x, outputs c = g^x * g^r.
func (c *Committer) GetCommitMsg(val *big.Int) (*big.Int, error) {
	if val.Cmp(c.Params.Group.Q) == 1 || val.Cmp(big.NewInt(0)) == -1 {
		err := fmt.Errorf("committed value needs to be in Z_q (order of a base point)")
		return nil, err
	}

	// c = g^x * h^r
	r := common.GetRandomInt(c.Params.Group.Q)

	c.r = r
	c.committedValue = val
	t1 := c.Params.Group.Exp(c.Params.Group.G, val)
	t2 := c.Params.Group.Exp(c.Params.H, r)
	comm := c.Params.Group.Mul(t1, t2)
	c.Commitment = comm

	return comm, nil
}

// It returns values x and r (commitment was c = g^x * g^r).
func (c *Committer) GetDecommitMsg() (*big.Int, *big.Int) {
	val := c.committedValue
	r := c.r
	return val, r
}

func (c *Committer) VerifyTrapdoor(trapdoor *big.Int) bool {
	h := c.Params.Group.Exp(c.Params.Group.G, trapdoor)
	return h.Cmp(c.Params.H) == 0
}

type Receiver struct {
	Params     *Params
	commitment *big.Int
}

func NewReceiver(bitLengthGroupOrder int) (*Receiver, error) {
	params, err := GenerateParams(bitLengthGroupOrder)
	if err != nil {
		return nil, err
	}
	return NewReceiverFromParams(params), nil
}

func NewReceiverFromParams(params *Params) *Receiver {
	return &Receiver{
		Params: params,
	}
}

func (r *Receiver) GetTrapdoor() *big.Int {
	return r.Params.a
}

// When receiver receives a commitment, it stores the value using SetCommitment method.
func (r *Receiver) SetCommitment(el *big.Int) {
	r.commitment = el
}

// When receiver receives a decommitment, CheckDecommitment verifies it against the stored value
// (stored by SetCommitment).
func (r *Receiver) CheckDecommitment(R, val *big.Int) bool {
	t1 := r.Params.Group.Exp(r.Params.Group.G, val) // g^x
	t2 := r.Params.Group.Exp(r.Params.H, R)         // h^r
	c := r.Params.Group.Mul(t1, t2)                 // g^x * h^r

	return c.Cmp(r.commitment) == 0
}
