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

package crypto

import (
	"math/big"
)

// Group interface is used to enable the usage of different groups in some schemes.
// For example when we have a homomorphism f between two groups and
// we are proving that we know an f-preimage of an element - meaning that for a given v we
// know u such that f(u) = v.
// Note that this is an interface for modular arithmetic groups. For elliptic curve
// groups at the moment there is no need for an interface.
type Group interface {
	GetRandomElement() *big.Int
	Mul(*big.Int, *big.Int) *big.Int
	Exp(*big.Int, *big.Int) *big.Int
	Inv(*big.Int) *big.Int
}
