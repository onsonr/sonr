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

package common

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGermainPrime(t *testing.T) {
	p := GetGermainPrime(512)
	p1 := new(big.Int).Add(p, p)
	p1.Add(p1, big.NewInt(1))

	assert.Equal(t, p.ProbablyPrime(20), true, "p should be prime")
	assert.Equal(t, p1.ProbablyPrime(20), true, "p1 should be prime")
}

func TestGetSafePrime(t *testing.T) {
	p, err := GetSafePrime(512)
	if err != nil {
		t.Errorf("Error in GetSafePrime: %v", err)
	}
	p1 := new(big.Int)
	p1.Sub(p, big.NewInt(1))
	p1.Div(p1, big.NewInt(2))

	assert.Equal(t, p.ProbablyPrime(20), true, "p should be prime")
	assert.Equal(t, p1.ProbablyPrime(20), true, "p1 should be prime")
}
