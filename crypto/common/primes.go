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
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
)

// GetSafePrime returns a safe prime p (p = 2*p1 + 2 where p1 is prime too).
func GetSafePrime(bits int) (p *big.Int, err error) {
	p1 := GetGermainPrime(bits - 1)
	p = big.NewInt(0)
	p.Mul(p1, big.NewInt(2))
	p.Add(p, big.NewInt(1))

	if p.BitLen() == bits {
		return p, nil
	} else {
		err := fmt.Errorf("bit length not correct")
		return nil, err
	}
}

// GetGermainPrime returns a prime number p for which 2*p + 1 is also prime. Note that conversely p
// is called safe prime.
func GetGermainPrime(bits int) (p *big.Int) {
	// multiple germainPrime goroutines are called and we assume at least one will compute a
	// safe prime and send it to the channel, thus we do not handle errors in germainPrime
	var c chan *big.Int = make(chan *big.Int)
	var quit chan int = make(chan int)
	for j := int(0); j < 8; j++ {
		go germainPrime(bits, c, quit)
	}
	msg := <-c
	// for small values for parameter bits (which should be small only for testing) it sometimes
	// happen "send on closed channel" - so we leave the channel c to a garbage collector
	//close(c)
	close(quit)
	return msg
}

var smallPrimes = []uint8{
	3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53,
}

// smallPrimesProduct is the product of the values in smallPrimes and allows us
// to reduce a candidate prime by this number and then determine whether it's
// coprime to all the elements of smallPrimes without further big.Int
// operations.
var smallPrimesProduct = new(big.Int).SetUint64(16294579238595022365)

// germainPrime is slightly modified Prime function from:
// https://github.com/golang/go/blob/master/src/crypto/rand/util.go
// germainPrime returns a number, p, of the given size, such that p and 2*p+1 are primes
// with high probability.
// germainPrime will return error for any error returned by rand.Read or if bits < 2.
func germainPrime(bits int, c chan *big.Int, quit chan int) (p *big.Int, err error) {
	rand := rand.Reader

	if bits < 2 {
		err = fmt.Errorf("crypto/rand: prime size must be at least 2-bit")
		return
	}

	b := uint(bits % 8)
	if b == 0 {
		b = 8
	}

	bytes := make([]byte, (bits+7)/8)
	p = new(big.Int)
	p1 := new(big.Int)

	bigMod := new(big.Int)

	for {
		select {
		case <-quit:
			return
		default:
			// this is to make it non-blocking
		}

		_, err = io.ReadFull(rand, bytes)
		if err != nil {
			return nil, err
		}

		// Clear bits in the first byte to make sure the candidate has a size <= bits.
		bytes[0] &= uint8(int(1<<b) - 1)
		// Don't let the value be too small, i.e, set the most significant two bits.
		// Setting the top two bits, rather than just the top bit,
		// means that when two of these values are multiplied together,
		// the result isn't ever one bit short.
		if b >= 2 {
			bytes[0] |= 3 << (b - 2)
		} else {
			// Here b==1, because b cannot be zero.
			bytes[0] |= 1
			if len(bytes) > 1 {
				bytes[1] |= 0x80
			}
		}
		// Make the value odd since an even number this large certainly isn't prime.
		bytes[len(bytes)-1] |= 1

		p.SetBytes(bytes)

		// Calculate the value mod the product of smallPrimes. If it's
		// a multiple of any of these primes we add two until it isn't.
		// The probability of overflowing is minimal and can be ignored
		// because we still perform Miller-Rabin tests on the result.
		bigMod.Mod(p, smallPrimesProduct)
		mod := bigMod.Uint64()

	NextDelta:
		for delta := uint64(0); delta < 1<<20; delta += 2 {
			m := mod + delta
			for _, prime := range smallPrimes {
				if m%uint64(prime) == 0 && (bits > 6 || m != uint64(prime)) {
					continue NextDelta
				}

				// 2*mod + 2*delta + 1	should not be divisible by smallPrimes as well
				m1 := (2*m + 1) % smallPrimesProduct.Uint64()

				if m1%uint64(prime) == 0 && (bits > 6 || m1 != uint64(prime)) {
					continue NextDelta
				}
			}

			if delta > 0 {
				bigMod.SetUint64(delta)
				p.Add(p, bigMod)
			}

			p1.Add(p, p)
			p1.Add(p1, big.NewInt(1))
			break
		}

		// There is a tiny possibility that, by adding delta, we caused
		// the number to be one bit too long. Thus we check BitLen
		// here.
		if p.ProbablyPrime(20) && p.BitLen() == bits {
			if p1.ProbablyPrime(20) {
				// waiting for a message about channel being closed is repeated here,
				// because it might happen that channel is closed after waiting at the
				// beginning of for loop above (but we want to have it there also,
				// otherwise it this goroutine might be searching for a germain
				// prime for some time after one was found by another goroutine
				select {
				case <-quit:
					return
				default:
					// this is to make it non-blocking
				}

				c <- p
				return
			}
		}
	}
}
