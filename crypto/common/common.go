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

import (
	"crypto/sha512"
	"math/big"
)

// It takes big.Int numbers, transform them to bytes, and concatenate the bytes.
func ConcatenateNumbers(numbers ...*big.Int) []byte {
	var bs []byte
	for _, n := range numbers {
		bs = append(bs, n.Bytes()...)
	}
	return bs
}

// It concatenates numbers (their bytes), computes a hash and outputs a hash as []byte.
func HashIntoBytes(numbers ...*big.Int) []byte {
	toBeHashed := ConcatenateNumbers(numbers...)
	sha512 := sha512.New()
	sha512.Write(toBeHashed)
	hashBytes := sha512.Sum(nil)
	return hashBytes
}

// It concatenates numbers (their bytes), computes a hash and outputs a hash as *big.Int.
func Hash(numbers ...*big.Int) *big.Int {
	hashBytes := HashIntoBytes(numbers...)
	hashNum := new(big.Int).SetBytes(hashBytes)
	return hashNum
}

// It computes x^y mod m. Negative y are supported.
func Exponentiate(x, y, m *big.Int) *big.Int {
	var r *big.Int
	if y.Cmp(big.NewInt(0)) >= 0 {
		r = new(big.Int).Exp(x, y, m)
	} else {
		r = new(big.Int).Exp(x, new(big.Int).Abs(y), m)
		r.ModInverse(r, m)
	}
	return r
}

// Computes least common multiple.
func LCM(x, y *big.Int) *big.Int {
	n := new(big.Int)
	n.Mul(x, y)
	d := new(big.Int)
	d.GCD(nil, nil, x, y)
	t := new(big.Int)
	t.Div(n, d)
	return t
}

// Contains returns true if array contains a given element, otherwise false.
func Contains(arr []int, el int) bool {
	for _, i := range arr {
		if el == i {
			return true
		}
	}

	return false
}
