package node

import (
	"context"

	// "log"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"

	// "github.com/sonr-io/sonr/core/device"
	// "github.com/sonr-io/sonr/pkg/crypto"

	"github.com/tendermint/starport/starport/pkg/cosmosaccount"
	"github.com/tendermint/starport/starport/pkg/cosmosclient"
)

var (
	sonrDeviceKeyFile = "sonr-provision.priv"
)

// Client is a client for the Sonr network
type Client struct {
	cosmosclient.Client
	//highwayv1.HighwayServiceClient
	Host    HostImpl
	ctx     context.Context
	Account cosmosaccount.Account
}

// NewCosmos creates a new client for the given name
func NewCosmos(ctx context.Context, addr string, sname string, passphrase string) (*Client, error) {
	// create an instance of cosmosclient
	cosmos, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix("snr"))
	if err != nil {
		return nil, err
	}

	// Create a new highway client
	return &Client{
		Client: cosmos,
		ctx:    ctx,
		// Host:                 h,
	}, nil
}

// Keyring returns the keyring for the given name
func (c *Client) Keyring() keyring.Keyring {
	return c.Client.Context.Keyring
}
