package bulletproof

import (
	crand "crypto/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/onsonr/sonr/crypto/core/curves"
)

func TestInnerProductHappyPath(t *testing.T) {
	curve := curves.ED25519()
	a := randScalarVec(3, *curve)
	b := randScalarVec(3, *curve)
	_, err := innerProduct(a, b)
	require.NoError(t, err)
}

func TestInnerProductMismatchedLengths(t *testing.T) {
	curve := curves.ED25519()
	a := randScalarVec(3, *curve)
	b := randScalarVec(4, *curve)
	_, err := innerProduct(a, b)
	require.Error(t, err)
}

func TestInnerProductEmptyVector(t *testing.T) {
	curve := curves.ED25519()
	a := randScalarVec(0, *curve)
	b := randScalarVec(0, *curve)
	_, err := innerProduct(a, b)
	require.Error(t, err)
}

func TestInnerProductOut(t *testing.T) {
	curve := curves.ED25519()
	a := randScalarVec(2, *curve)
	b := randScalarVec(2, *curve)
	c, err := innerProduct(a, b)
	require.NoError(t, err)

	// Calculate manually a0*b0 + a1*b1
	cPrime := a[0].Mul(b[0]).Add(a[1].Mul(b[1]))
	require.Equal(t, c, cPrime)
}

func TestSplitListofPointsHappyPath(t *testing.T) {
	curve := curves.ED25519()
	points := randPointVec(10, *curve)
	firstHalf, secondHalf, err := splitPointVector(points)
	require.NoError(t, err)
	require.Len(t, firstHalf, 5)
	require.Len(t, secondHalf, 5)
}

func TestSplitListofPointsOddLength(t *testing.T) {
	curve := curves.ED25519()
	points := randPointVec(11, *curve)
	_, _, err := splitPointVector(points)
	require.Error(t, err)
}

func TestSplitListofPointsZeroLength(t *testing.T) {
	curve := curves.ED25519()
	points := randPointVec(0, *curve)
	_, _, err := splitPointVector(points)
	require.Error(t, err)
}

func randScalarVec(length int, curve curves.Curve) []curves.Scalar {
	out := make([]curves.Scalar, length)
	for i := 0; i < length; i++ {
		out[i] = curve.Scalar.Random(crand.Reader)
	}
	return out
}

func randPointVec(length int, curve curves.Curve) []curves.Point {
	out := make([]curves.Point, length)
	for i := 0; i < length; i++ {
		out[i] = curve.Point.Random(crand.Reader)
	}
	return out
}
