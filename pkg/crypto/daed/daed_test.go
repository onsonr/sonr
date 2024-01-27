package daed_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sonrhq/sonr/pkg/crypto/daed"
)

func TestNewKeyset(t *testing.T) {
	err := daed.NewKeyset()
	assert.NoError(t, err)
}
