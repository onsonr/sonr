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

package secretsharing

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/sonr-io/core/crypto/common"
)

// Shamir's secret sharing scheme
type Dealer struct {
}

func NewDealer() (*Dealer, error) {
	dealer := Dealer{}

	return &dealer, nil
}

func (dealer *Dealer) SplitSecret(secret string, threshold int,
	numberOfShares int) (map[*big.Int]*big.Int, *big.Int, error) {
	if threshold < 2 {
		err := fmt.Errorf("the threshold should be at least 2")
		return nil, nil, err
	}
	if threshold > numberOfShares {
		err := fmt.Errorf("the threshold should be smaller than the number of shares")
		return nil, nil, err
	}

	b := []byte(secret)
	secretNum := new(big.Int).SetBytes(b)

	// generate some random prime which will be bigger than numberOfShares and secretNum
	if secretNum.Cmp(big.NewInt(int64(numberOfShares))) < 0 {
		err := fmt.Errorf("the number of shares (participants) is too high")
		return nil, nil, err
	}

	prime, err := rand.Prime(rand.Reader, secretNum.BitLen()+1)
	if err != nil {
		return nil, nil, err
	}

	if prime.Cmp(big.NewInt(int64(numberOfShares))) < 0 {
		err := fmt.Errorf("the number of shares (participants) is too high")
		return nil, nil, err
	}

	polynomial, _ := common.NewRandomPolynomial(threshold-1, prime)
	polynomial.SetCoefficient(0, secretNum)

	var ps []*big.Int
	for i := 0; i < numberOfShares; i++ {
		ps = append(ps, big.NewInt(int64(i)))
	}
	points := polynomial.GetValues(ps)

	return points, prime, nil
}

// It takes threshold
func (dealer *Dealer) RecoverSecret(points map[*big.Int]*big.Int, prime *big.Int) string {
	secretNum := common.LagrangeInterpolation(big.NewInt(0), points, prime)

	b := secretNum.Bytes()
	return string(b)
}
