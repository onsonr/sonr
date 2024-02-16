package wallet_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sonrhq/sonr/internal/wallet"
)

func TestNewKeychain(t *testing.T) {
	_, kc, err := wallet.New(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, kc.Address)
}

func BenchmarkNewKeychain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, kc, err := wallet.New(context.Background())
		if err != nil {
			b.Fail()
		}
		if len(kc.Address) == 0 {
			b.Fail()
		}
	}
}

func TestNewWalletEncryptDecrypt(t *testing.T) {
	associatedData := []byte("associated data")
	dir, kc, err := wallet.New(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, kc.Address)

	encrypted, err := kc.Encrypt(dir, associatedData)
	require.NoError(t, err)
	require.NotEmpty(t, encrypted)

	decrypted, err := kc.Decrypt(encrypted, associatedData)
	require.NoError(t, err)
	require.NotEmpty(t, decrypted)
}
