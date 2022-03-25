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

package ecpseudsys

import (
	"github.com/sonr-io/core/crypto/common"
	"github.com/sonr-io/core/crypto/ec"
	"github.com/sonr-io/core/crypto/pseudsys"
)

type PubKey struct {
	H1, H2 *ec.GroupElement
}

func NewPubKey(h1, h2 *ec.GroupElement) *PubKey {
	return &PubKey{h1, h2}
}

// GenerateKeyPair takes EC group and constructs a public key for pseudonym system scheme in EC
// arithmetic.
func GenerateKeyPair(group *ec.Group) (*pseudsys.SecKey, *PubKey) {
	s1 := common.GetRandomInt(group.Q)
	s2 := common.GetRandomInt(group.Q)
	h1 := group.ExpBaseG(s1)
	h2 := group.ExpBaseG(s2)

	return pseudsys.NewSecKey(s1, s2), NewPubKey(h1, h2)
}
