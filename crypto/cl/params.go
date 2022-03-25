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

package cl

// Params presents parameters that organization (which is issuing credentials) needs to set.
type Params struct {
	// There are only a few possibilities for RhoBitLen. 256 implies that the modulus
	// bit length is 2048
	RhoBitLen         int // bit length of order of the commitment group
	NLength           int // bit length of RSA modulus
	KnownAttrsNum     int // number of attributes known to both - credential issuer and receiver
	CommittedAttrsNum int // number of attributes for which the issuer knows only commitments
	HiddenAttrsNum    int // number of attributes known only to the receiver
	AttrBitLen        int // bit length of attribute
	HashBitLen        int // bit length of hash output used for Fiat-Shamir
	SecParam          int // security parameter
	EBitLen           int // size of e values of certificates
	E1BitLen          int // size of the interval the e values are taken from
	VBitLen           int // size of the v values of the certificates
	ChallengeSpace    int // bit length of challenges for DF commitment proofs
}

// TODO: add method to load params from file or blockchain or wherever they will be stored.
func GetDefaultParamSizes() *Params {
	return &Params{
		RhoBitLen:         256,
		NLength:           256, // should be at least 2048 when not testing
		KnownAttrsNum:     5,
		CommittedAttrsNum: 1,
		HiddenAttrsNum:    0,
		AttrBitLen:        256,
		HashBitLen:        512,
		SecParam:          80,
		EBitLen:           597,
		E1BitLen:          120,
		VBitLen:           2724,
		ChallengeSpace:    80,
	}
}
