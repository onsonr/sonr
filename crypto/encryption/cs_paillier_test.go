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
	"math/big"
	"testing"

	"github.com/sonr-io/core/crypto/common"
	"github.com/stretchr/testify/assert"
)

func TestCSPaillier(t *testing.T) {
	csp := NewCSPaillier(
		&CSPaillierSecParams{
			L:        512,
			RoLength: 160,
			K:        158,
			K1:       158,
		})

	cspSec, _ := NewCSPaillierFromSecKey(csp.SecKey)
	cspPub := NewCSPaillierFromPubKey(csp.PubKey)

	m := common.GetRandomInt(big.NewInt(8685849))
	label := common.GetRandomInt(big.NewInt(340002223232))

	u, e, v, _ := cspPub.Encrypt(m, label)
	p, _ := cspSec.Decrypt(u, e, v, label)
	assert.Equal(t, m, p, "Camenisch-Shoup modified Paillier encryption/decryption does not work correctly")

	l, delta := cspPub.GetOpeningMsg(p)
	u1, e1, v1, delta1, l1, _ := cspPub.GetProofRandomData(u, e, label)

	cspVer := NewCSPaillierFromPubKey(csp.PubKey)

	cspVer.SetVerifierEncData(u, e, v, delta, label, l)
	challenge := cspVer.GetChallenge()
	cspVer.SetProofRandomData(u1, e1, v1, delta1, l1, challenge)

	rTilde, sTilde, mTilde := cspPub.GetProofData(challenge)

	assert.True(t, cspVer.Verify(rTilde, sTilde, mTilde), "Camenisch-Shoup modified Paillier verifiable encryption proof does not work correctly")
	assert.True(t, new(big.Int).Abs(v).Cmp(v) == 0, "Camenisch-Shoup modified Paillier verifiable encryption proof does not work correctly")
}
