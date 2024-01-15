package keychain_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sonrhq/sonr/internal/keychain"
)

func TestNewKeychain(t *testing.T) {
	kc, err := keychain.New()
	require.NoError(t, err)
	fmt.Println(kc.RootDir)
	kc.Burn()
}
