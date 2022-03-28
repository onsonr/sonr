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

package qr

// Zero-knowledge proof	of quadratic residousity (implemented for historical reasons)

import (
	"math/big"

	"fmt"

	"github.com/sonr-io/core/crypto/common"
	"github.com/sonr-io/core/crypto/schnorr"
)

// ProveQR demonstrates how the prover can prove that y1^2 is QR.
func ProveQR(y1 *big.Int, group *schnorr.Group) bool {
	y := group.Mul(y1, y1)
	prover := NewProver(group, y1)
	verifier := NewVerifier(y, group)
	m := group.P.BitLen()

	for i := 0; i < m; i++ {
		x := prover.GetProofRandomData()
		c := verifier.GetChallenge(x)

		z, _ := prover.GetProofData(c)

		proved := verifier.Verify(z)
		if !proved {
			return false
		}
	}
	return true
}

type Prover struct {
	Group *schnorr.Group
	Y     *big.Int
	y1    *big.Int
	r     *big.Int
}

func NewProver(group *schnorr.Group, y1 *big.Int) *Prover {
	y := group.Mul(y1, y1)
	return &Prover{
		Group: group,
		Y:     y,
		y1:    y1,
	}
}

func (p *Prover) GetProofRandomData() *big.Int {
	r := common.GetRandomInt(p.Group.P)
	p.r = r
	x := p.Group.Exp(r, big.NewInt(2))
	return x
}

func (p *Prover) GetProofData(challenge *big.Int) (*big.Int, error) {
	if challenge.Cmp(big.NewInt(0)) == 0 {
		return p.r, nil
	} else if challenge.Cmp(big.NewInt(1)) == 0 {
		z := new(big.Int).Mul(p.r, p.y1)
		z.Mod(z, p.Group.P)
		return z, nil
	} else {
		err := fmt.Errorf("challenge is not valid")
		return nil, err
	}
}

type Verifier struct {
	Group     *schnorr.Group
	x         *big.Int
	y         *big.Int
	challenge *big.Int
}

func NewVerifier(y *big.Int, group *schnorr.Group) *Verifier {
	return &Verifier{
		Group: group,
		y:     y,
	}
}

func (v *Verifier) GetChallenge(x *big.Int) *big.Int {
	v.x = x
	c := common.GetRandomInt(big.NewInt(2)) // 0 or 1
	v.challenge = c
	return c
}

func (v *Verifier) Verify(z *big.Int) bool {
	z2 := new(big.Int).Mul(z, z)
	z2.Mod(z2, v.Group.P)
	if v.challenge.Cmp(big.NewInt(0)) == 0 {
		return z2.Cmp(v.x) == 0
	} else {
		s := new(big.Int).Mul(v.x, v.y)
		s.Mod(s, v.Group.P)
		return z2.Cmp(s) == 0
	}
}

// isQR accepts integer a and prime p, and returns true if
// a is quadratic residue in Z_p group, false otherwise.
// If p is not a prime, error is returned.
func isQR(a *big.Int, p *big.Int) (bool, error) {
	if !p.ProbablyPrime(20) {
		return false, fmt.Errorf("p is not a prime")
	}

	one := big.NewInt(1)

	// check whether a^((p-1)/2) is 1 or -1 (Euler's criterion)
	pMin1 := new(big.Int).Sub(p, one)
	exp := new(big.Int).Div(pMin1, big.NewInt(2)) // exponent
	cr := new(big.Int).Exp(a, exp, p)

	return cr.Cmp(one) == 0, nil
}
