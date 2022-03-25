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

func TestPartialECDLogKnowledge(t *testing.T) {
	group := ec.NewGroup(ec.P256)

	exp1 := common.GetRandomInt(group.Q)
	exp2 := common.GetRandomInt(group.Q)
	a1 := group.ExpBaseG(exp1)
	a2 := group.ExpBaseG(exp2)

	secret1 := common.GetRandomInt(group.Q)
	x := common.GetRandomInt(group.Q)

	//b1X, b1Y := dlog.ExponentiateBaseG(secret1)
	// we pretend that we don't know x:
	b2 := group.ExpBaseG(x)

	proved := ProvePartialDLogKnowledge(group, secret1, a1, a2, b2)

	assert.Equal(t, proved, true, "partial dlog knowledge proof does not work")
}
