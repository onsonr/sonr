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
	"math/big"
)

// Polynomial with coefficients in Z_prime. Coefficients are given as [a_0, a_1, ..., a_degree] where
// polynomial is p(x) = a_0 + a_1 * x + ... + a_degree * x^degree
type Polynomial struct {
	coefficients []*big.Int
	degree       int
	prime        *big.Int // coefficients are in Z_prime
}

func NewRandomPolynomial(degree int, prime *big.Int) (*Polynomial, error) {
	var coefficients []*big.Int
	for i := 0; i <= degree; i++ {
		coef := GetRandomInt(prime) // coeff has to be < prime
		coefficients = append(coefficients, coef)
	}
	polynomial := Polynomial{
		coefficients: coefficients,
		degree:       degree,
		prime:        prime,
	}

	return &polynomial, nil
}

func (polynomial *Polynomial) SetCoefficient(coeff_ind int, coefficient *big.Int) {
	polynomial.coefficients[coeff_ind] = coefficient
}

// Computes polynomial values at given points.
func (polynomial *Polynomial) GetValues(points []*big.Int) map[*big.Int]*big.Int {
	m := make(map[*big.Int]*big.Int)
	for _, value := range points {
		m[value] = polynomial.GetValue(value)
	}
	return m
}

// Computes polynomial values at given point.
func (polynomial *Polynomial) GetValue(point *big.Int) *big.Int {
	value := big.NewInt(0)
	for i, coeff := range polynomial.coefficients {
		// a_i * point^i
		tmp := new(big.Int).Exp(point, big.NewInt(int64(i)), polynomial.prime) // point^i % prime
		tmp.Mul(coeff, tmp)
		value.Add(value, tmp)
	}
	value.Mod(value, polynomial.prime)
	return value
}

// Given degree+1 points which are on the polynomial, LagrangeInterpolation computes p(a).
func LagrangeInterpolation(a *big.Int, points map[*big.Int]*big.Int, prime *big.Int) *big.Int {
	value := big.NewInt(0)
	for key, val := range points {
		numerator := big.NewInt(1)
		denominator := big.NewInt(1)
		t := new(big.Int)
		for key1 := range points {
			if key == key1 {
				continue
			}

			t.Sub(a, key1)
			numerator.Mul(numerator, t)
			numerator.Mod(numerator, prime)

			t.Sub(key, key1)
			denominator.Mul(denominator, t)
			denominator.Mod(denominator, prime)
		}
		t1 := new(big.Int)
		denominator_inv := new(big.Int)
		denominator_inv.ModInverse(denominator, prime)
		t1.Mul(numerator, denominator_inv)
		t1.Mod(t1, prime)

		t2 := new(big.Int)
		// (prime + value + t1 * val) % prime
		t2.Mul(val, t1)
		t2.Add(t2, prime)
		t2.Add(t2, prime)

		value.Add(value, t2)
		value.Add(value, prime)
	}
	value.Mod(value, prime)
	return value
}
