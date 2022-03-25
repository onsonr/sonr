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

package ecpedersen

import (
	"fmt"
	"math/big"

	"github.com/sonr-io/core/crypto/common"
	"github.com/sonr-io/core/crypto/ec"
)

type Params struct {
	Group *ec.Group
	H     *ec.GroupElement
	a     *big.Int
	// trapdoor a can be nil (doesn't need to be known), it is rarely needed -
	// for example in one of techniques to turn sigma to ZKP
}

func NewParams(group *ec.Group, H *ec.GroupElement, a *big.Int) *Params {
	return &Params{
		Group: group,
		H:     H, // H = g^a
		a:     a,
	}
}

func GenerateParams(curveType ec.Curve) *Params {
	group := ec.NewGroup(curveType)
	a := common.GetRandomInt(group.Q)
	return NewParams(group, group.ExpBaseG(a), a)
}

// Committer can commit to some value x - it sends to receiver c = g^x * h^r.
// When decommitting, committer sends to receiver r, x; receiver checks whether c = g^x * h^r.
type Committer struct {
	Params         *Params
	Commitment     *ec.GroupElement
	committedValue *big.Int
	r              *big.Int
}

func NewCommitter(params *Params) *Committer {
	committer := Committer{
		Params: params,
	}
	return &committer
}

// It receives a value x (to this value a commitment is made), chooses a random x, outputs c = g^x * g^r.
func (c *Committer) GetCommitMsg(val *big.Int) (*ec.GroupElement, error) {
	if val.Cmp(c.Params.Group.Q) == 1 || val.Cmp(big.NewInt(0)) == -1 {
		err := fmt.Errorf("the committed value needs to be in Z_q (order of a base point)")
		return nil, err
	}

	// c = g^x * h^r
	r := common.GetRandomInt(c.Params.Group.Q)

	c.r = r
	c.committedValue = val
	x1 := c.Params.Group.ExpBaseG(val)
	x2 := c.Params.Group.Exp(c.Params.H, r)
	comm := c.Params.Group.Mul(x1, x2)
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
	h := c.Params.Group.ExpBaseG(trapdoor)
	return h.Equals(c.Params.H)
}

type Receiver struct {
	Params     *Params
	commitment *ec.GroupElement
}

func NewReceiver(curve ec.Curve) *Receiver {
	return &Receiver{
		Params: GenerateParams(curve),
	}
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
func (r *Receiver) SetCommitment(el *ec.GroupElement) {
	r.commitment = el
}

// When receiver receives a decommitment, CheckDecommitment verifies it against the stored value
// (stored by SetCommitment).
func (r *Receiver) CheckDecommitment(R, val *big.Int) bool {
	a := r.Params.Group.ExpBaseG(val)      // g^x
	b := r.Params.Group.Exp(r.Params.H, R) // h^r
	c := r.Params.Group.Mul(a, b)          // g^x * h^r

	return c.Equals(r.commitment)
}
