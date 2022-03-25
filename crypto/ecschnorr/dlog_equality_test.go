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

package ecschnorr

import (
	"testing"

	"github.com/sonr-io/core/crypto/common"
	"github.com/sonr-io/core/crypto/ec"
	"github.com/stretchr/testify/assert"
)

func TestECDLogEquality(t *testing.T) {
	group := ec.NewGroup(ec.P256)
	secret := common.GetRandomInt(group.Q)

	r1 := common.GetRandomInt(group.Q)
	r2 := common.GetRandomInt(group.Q)

	g1 := group.ExpBaseG(r1)
	g2 := group.ExpBaseG(r2)

	t1 := group.Exp(g1, secret)
	t2 := group.Exp(g2, secret)

	proved := ProveDLogEquality(secret, g1, g2, t1, t2, ec.P256)
	assert.Equal(t, proved, true, "dlog equality proof does not work")
}
