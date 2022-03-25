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

package encryption

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/sonr-io/core/crypto/common"
)

// https://pirk.incubator.apache.org/papers/1999_asiacrypt_paillier_paper.pdf
type Paillier struct {
	primeLength int
	lambda      *big.Int
	pubKey      *PaillierPubKey
}

type PaillierPubKey struct {
	n  *big.Int
	n2 *big.Int
	g  *big.Int
}

func NewPaillier(primeLength int) *Paillier {
	paillier := Paillier{
		primeLength: primeLength,
	}
	paillier.generateKey()

	return &paillier
}

func NewPubPaillier(pubKey *PaillierPubKey) *Paillier {
	return &Paillier{
		pubKey: pubKey,
	}
}

func (paillier *Paillier) Encrypt(m *big.Int) (*big.Int, error) {
	if m.Cmp(paillier.pubKey.n) >= 0 {
		err := fmt.Errorf("msg is too big")
		return nil, err
	}

	// c = g^m * r^n mod n^2
	// r should be from Z_n*, but as it is very unlikely that we get an element which is not
	// invertible, we don't check
	r := common.GetRandomInt(paillier.pubKey.n)
	t1 := new(big.Int).Exp(paillier.pubKey.g, m, paillier.pubKey.n2) // g^m
	t2 := new(big.Int).Exp(r, paillier.pubKey.n, paillier.pubKey.n2) // r^n
	c := new(big.Int).Mul(t1, t2)
	c.Mod(c, paillier.pubKey.n2)

	return c, nil
}

func (paillier *Paillier) Decrypt(c *big.Int) (*big.Int, error) {
	if c.Cmp(paillier.pubKey.n2) >= 0 {
		err := fmt.Errorf("cipertext is too big")
		return nil, err
	}

	// p = (c^lambda - 1) / (g^lambda - 1) mod n
	c1 := c.Exp(c, paillier.lambda, paillier.pubKey.n2)
	c1.Sub(c1, big.NewInt(1))
	c1.Div(c1, paillier.pubKey.n)

	g1 := new(big.Int).Exp(paillier.pubKey.g, paillier.lambda, paillier.pubKey.n2)
	g1.Sub(g1, big.NewInt(1))
	g1.Div(g1, paillier.pubKey.n)

	g1_inv := new(big.Int).ModInverse(g1, paillier.pubKey.n2)

	p := new(big.Int).Mul(c1, g1_inv)
	p.Mod(p, paillier.pubKey.n)
	return p, nil
}

func (paillier *Paillier) GetPubKey() *PaillierPubKey {
	return paillier.pubKey
}

func (paillier *Paillier) generateKey() {
	p, _ := rand.Prime(rand.Reader, paillier.primeLength)
	q, _ := rand.Prime(rand.Reader, paillier.primeLength)
	p_min := new(big.Int).Sub(p, big.NewInt(1)) // p-1
	q_min := new(big.Int).Sub(q, big.NewInt(1)) // q-1

	paillier.lambda = common.LCM(p_min, q_min)
	n := new(big.Int).Mul(p, q)
	n2 := new(big.Int).Mul(n, n)

	pubKey := PaillierPubKey{
		n:  n,
		n2: n2,
	}

	for {
		g := common.GetRandomInt(n2)
		// check whether it is of order k * n
		// g = (1+n)^x * y^n mod n^2
		// g^lambda = (1+n)^(lambda * x) * y^(n * lambda) mod n^2
		// due to Carmichael:
		// g^lambda = (1+n)^(lambda * x) mod n^2
		// due to binomial theorem:
		// g^lambda = (1 + lambda * x * n) mod n^2
		t := new(big.Int).Exp(g, paillier.lambda, n2)

		x := new(big.Int).Sub(t, big.NewInt(1)) // (g^lambda - 1)
		x.Div(x, paillier.lambda)               // (g^lambda - 1) / lambda
		x.Div(x, n)                             // (g^lambda - 1) / (lambda * n)

		gcd := new(big.Int).GCD(nil, nil, x, n)
		if gcd.Cmp(big.NewInt(1)) == 0 {
			pubKey.g = g
			paillier.pubKey = &pubKey
			break
		}
	}

}
