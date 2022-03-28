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
	"testing"

	"github.com/sonr-io/core/crypto/common"
	"github.com/sonr-io/core/crypto/ec"
	"github.com/stretchr/testify/assert"
)

// TestPedersenEC demonstrates how a value can be committed and later opened (decommitted) using Pedersen EC committer.
func TestPedersenEC(t *testing.T) {
	receiver := NewReceiver(ec.P256)
	committer := NewCommitter(receiver.Params)

	a := common.GetRandomInt(committer.Params.Group.Q)
	c, err := committer.GetCommitMsg(a)
	if err != nil {
		t.Errorf("Error in GetCommitMsg: %v", err)
	}

	receiver.SetCommitment(c)
	committedVal, r := committer.GetDecommitMsg()
	success := receiver.CheckDecommitment(r, committedVal)

	assert.Equal(t, true, success, "Pedersen EC commitment failed.")
}
