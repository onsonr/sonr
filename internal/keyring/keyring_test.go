package keyring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DSC(t *testing.T) {
	newDsc, err := CreateDSC()
	assert.NoError(t, err, "create DSC")

	dsc, err := GetDSC()
	assert.NoError(t, err, "get DSC")

	assert.Equal(t, newDsc, dsc, "keys match")
}

func Test_PSK(t *testing.T) {
	newPsk, err := CreatePSK()
	assert.NoError(t, err, "create PSK")

	psk, err := GetPSK()
	assert.NoError(t, err, "get PSK")

	assert.Equal(t, newPsk, psk, "keys match")
}
