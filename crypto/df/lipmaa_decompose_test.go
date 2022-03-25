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

package df

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestLipmaaDecompositionNeg tests whether decomposition function returns an error
// when an invaid argument (negative integer) is provided.
func TestLipmaaDecompositionNeg(t *testing.T) {
	integer := big.NewInt(-1)
	_, err := lipmaaDecompose(integer)
	if err == nil {
		t.Errorf("decomposition of negative integer should produce an error")
	}
}

// TestLipmaaDecompositionSmall iteratively tests whether the roots of retrieved
// decompositions sum up exactly to a given integer when they are squared.
func TestLipmaaDecompositionSmall(t *testing.T) {
	var bigInt = new(big.Int)
	for i := 0; i <= 10000; i++ {
		bigInt.SetInt64(int64(i))
		roots, _ := lipmaaDecompose(bigInt)
		squaredRootsSum := squareRootsAndSum(roots)
		if bigInt.Cmp(squaredRootsSum) != 0 {
			t.Errorf("decomposition does not work correctly for small integers ("+
				"expected root to sum up to %d, got %d)", bigInt.Int64(), squaredRootsSum.Int64())
		}
	}
}

// TestLipmaaDecompositionLarge tests whether decomposition works for a large integer.
func TestLipmaaDecompositionLarge(t *testing.T) {
	bigInt, _ := new(big.Int).SetString("16714772973240639959372252262788596420406994288943442724185217359247384753656472309049760952976644136858333233015922583099687128195321947212684779063190875332970679291085543110146729439665070418750765330192961290161474133279960593149307037455272278582955789954847238104228800942225108143276152223829168166008095539967222363070565697796008563529948374781419181195126018918350805639881625937503224895840081959848677868603567824611344898153185576740445411565094067875133968946677861528581074542082733743513314354002186235230287355796577107626422168586230066573268163712626444511811717579062108697723640288393001520781671", 10)
	roots, _ := lipmaaDecompose(bigInt)
	if squareRootsAndSum(roots).Cmp(bigInt) != 0 {
		t.Errorf("decomposition does not work correctly for a large integer")
	}

}

// TestLipmaaDecomposition checks whether filtering of Lagrangian works - e.g.
// whether the function returns a slice containing only non-zero roots.
func TestLipmaaDecomposition(t *testing.T) {
	tests := []struct {
		n           *big.Int
		expectedLen int
	}{
		{big.NewInt(0), 0},
		{big.NewInt(1), 1},
		{big.NewInt(5), 2},
		{big.NewInt(24), 3},
		{big.NewInt(7), 4},
	}

	for _, test := range tests {
		res, _ := lipmaaDecomposition(test.n)
		assert.Equal(t, len(res), test.expectedLen, "filtering lagrangian does not work")
	}
}

// squareRootsAndSum calculates the sum of squares for a given lagrangian.
func squareRootsAndSum(w *lagrange) *big.Int {
	sum := new(big.Int)
	for _, root := range w {
		sum.Add(sum, new(big.Int).Mul(root, root))
	}

	return sum
}
