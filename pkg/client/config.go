package client

import (
	"github.com/duo-labs/webauthn/webauthn"
	"github.com/sonr-io/sonr/pkg/config"
	"github.com/tendermint/starport/starport/pkg/cosmosaccount"
	"github.com/tendermint/starport/starport/pkg/cosmosclient"
)

// CosmosOptions returns the cosmos options for the highway node
func toCosmosOptions(c *config.Config) []cosmosclient.Option {
	// Create the options
	opts := make([]cosmosclient.Option, 0)
	if c.CosmosUseFaucet {
		opts = append(opts, cosmosclient.WithUseFaucet(c.CosmosFaucetAddress, c.CosmosFaucetDenom, c.CosmosFaucetMinAmount))
	}

	// Convert the keyring backend string
	var keyring cosmosaccount.KeyringBackend
	if c.CosmosKeyringBackend == "os" {
		keyring = cosmosaccount.KeyringOS
	} else {
		keyring = cosmosaccount.KeyringTest
	}

	// Add remaining cosmos options
	opts = append(opts, cosmosclient.WithNodeAddress(c.CosmosNodeAddress),
		cosmosclient.WithAddressPrefix(c.CosmosAddressPrefix),
		cosmosclient.WithHome(c.CosmosHomePath),
		cosmosclient.WithKeyringBackend(keyring),
		cosmosclient.WithKeyringServiceName(c.CosmosKeyringServiceName))

	return opts
}

// WebauthnConfig returns the configuration for the WebAuthn module
func toWebauthnConfig(c *config.Config) *webauthn.Config {
	return &webauthn.Config{
		RPDisplayName: c.WebAuthNRPDisplayName,
		RPID:          c.WebAuthNRPID,
		RPOrigin:      c.WebAuthNRPOrigin,
		RPIcon:        c.WebAuthNRPIcon,
		Debug:         c.WebAuthNDebug,
	}

}
