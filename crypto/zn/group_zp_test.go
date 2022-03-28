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

package zn

import (
	"math/big"
	"testing"

	"github.com/sonr-io/core/crypto/common"
	"github.com/stretchr/testify/assert"
)

func TestGetGeneratorOfZnSubgroup(t *testing.T) {
	p, err := common.GetSafePrime(512)
	if err != nil {
		t.Errorf("Error in GetSafePrime: %v", err)
	}

	zp, _ := NewGroupZp(p)
	p1 := new(big.Int).Div(zp.Order, big.NewInt(2))
	g, err := zp.GetGeneratorOfSubgroup(p1)
	if err != nil {
		t.Errorf("Error in GetGeneratorOfSubgroup: %v", err)
	}
	g.Exp(g, big.NewInt(0).Sub(p1, big.NewInt(1)), p1) // g^(p1-1) % p1 should be 1

	assert.Equal(t, g, big.NewInt(1), "not a generator")
}
