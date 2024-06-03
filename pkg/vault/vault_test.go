package vault_test

import (
	"context"
	"testing"

	"github.com/di-dao/sonr/pkg/vault"
	"github.com/stretchr/testify/assert"
)

func TestNewVault(t *testing.T) {
	vt, err := vault.Generate(context.Background())
	assert.NotNil(t, vt)
	assert.NoError(t, err)
}
