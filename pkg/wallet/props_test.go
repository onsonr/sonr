package wallet_test

import (
	"testing"

	"github.com/di-dao/sonr/crypto/mpc"
	props "github.com/di-dao/sonr/pkg/wallet"
	"github.com/stretchr/testify/require"
)

func TestLinkUnlinkProperty(t *testing.T) {
	// Create properties
	props := props.NewProperties()
	require.NotNil(t, props)

	// Create KSS and get public key
	kss, err := mpc.GenerateKss()
	require.NoError(t, err)
	pk := kss.PublicKey()

	// Property key
	propertyKey := "email"

	// Link a property
	propertyValue := "user@example.com"
	witness, err := props.Set(pk, propertyKey, propertyValue)
	require.NoError(t, err)
	require.NotEmpty(t, witness)

	// Validate the linked property
	valid := props.Check(pk, propertyKey, witness)
	require.True(t, valid)

	// Unlink the property
	err = props.Remove(pk, propertyKey, propertyValue)
	require.NoError(t, err)

	// Validate the unlinked property
	valid = props.Check(pk, propertyKey, witness)
	require.False(t, valid)
}
