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

	"github.com/stretchr/testify/assert"
)

func TestBitCommitmentProof(t *testing.T) {
	verified, err := ProveBitCommitment()
	if err != nil {
		t.Errorf("Error in bit commitment proof: %v", err)
	}

	assert.Equal(t, true, verified, "Bit commitment does not work correctly")
}

func TestCommitmentMultiplicationProof(t *testing.T) {
	proved, err := ProveMultiplicationCommitment()
	if err != nil {
		t.Errorf("Error in multiplication proof: %v", err)
	}

	assert.Equal(t, true, proved, "Commitments multiplication proof failed.")
}
