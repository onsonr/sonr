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

package qr_test

import (
	"math/big"
	"testing"

	"github.com/sonr-io/core/crypto/qr"
	"github.com/stretchr/testify/assert"
)

func TestGeneratorOfCompositeQR(t *testing.T) {
	rsa, err := qr.NewRSASpecial(512)
	if err != nil {
		t.Errorf("Error when instantiating RSASpecial: %v", err)
	}

	g, err := rsa.GetRandomGenerator()
	if err != nil {
		t.Errorf("Error when searching for RSASpecial generator: %v", err)
	}

	// order of g should be p1*q1
	order := new(big.Int).Mul(rsa.P1, rsa.Q1)
	tmp := new(big.Int).Exp(g, order, rsa.N)

	assert.Equal(t, tmp, big.NewInt(1), "g is not a generator")
	// other possible orders in this group are: p1, q1
	tmp = new(big.Int).Exp(g, rsa.P1, rsa.N)
	assert.NotEqual(t, tmp, big.NewInt(1), "g is not a generator")

	tmp = new(big.Int).Exp(g, rsa.Q1, rsa.N)
	assert.NotEqual(t, tmp, big.NewInt(1), "g is not a generator")
}
