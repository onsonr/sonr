package qr

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestIsQRInvalid checks that IsQR returns an
// error when p is not a prime.
func TestIsQRInvalid(t *testing.T) {
	_, err := isQR(big.NewInt(0), big.NewInt(10))
	assert.Error(t, err)
}

func TestIsQR(t *testing.T) {
	x := big.NewInt(2)
	p := big.NewInt(11)

	a := new(big.Int).Exp(x, big.NewInt(2), p)

	isqr, err := isQR(a, p)
	assert.NoError(t, err)
	assert.True(t, isqr)
}
