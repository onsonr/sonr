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

package qoneway

import (
	"testing"

	"github.com/sonr-io/core/crypto/common"
	"github.com/stretchr/testify/assert"
)

func TestRSABasedCommitment(t *testing.T) {
	receiver, err := NewReceiver(1024)
	if err != nil {
		t.Errorf("Error when initializing Receiver: %v", err)
	}

	committer, err := NewCommitter(receiver.RSABased, receiver.Y)
	if err != nil {
		t.Errorf("Error when initializing Committer: %v", err)
	}

	a := common.GetRandomInt(committer.Q)
	c, _ := committer.GetCommitMsg(a)

	receiver.SetCommitment(c)
	committedVal, r := committer.GetDecommitMsg()
	success := receiver.CheckDecommitment(r, committedVal)

	assert.Equal(t, true, success, "RSABasedCommitment does not work correctly")
}
