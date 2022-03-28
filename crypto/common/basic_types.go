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

package common

import "math/big"

// Pair is the same as ECGroupElement, but to be used in non EC schemes when a pair of
// *big.Int is needed.
type Pair struct {
	A *big.Int
	B *big.Int
}

func NewPair(a, b *big.Int) *Pair {
	pair := Pair{A: a, B: b}
	return &pair
}

type Triple struct {
	A *big.Int
	B *big.Int
	C *big.Int
}

func NewTriple(a, b, c *big.Int) *Triple {
	triple := Triple{A: a, B: b, C: c}
	return &triple
}
