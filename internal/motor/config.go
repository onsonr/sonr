package motor

import (
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/crypto"
)

func setupDefault() (*MotorNode, error) {
	// Create Client instance
	c := client.NewClient(client.ConnEndpointType_LOCAL)

	// Generate wallet
	w, err := crypto.GenerateWallet()
	if err != nil {
		return nil, err
	}
	bechAddr, err := w.Address()
	if err != nil {
		return nil, err
	}
	err = c.RequestFaucet(bechAddr)
	if err != nil {
		return nil, err
	}

	pk, err := w.PublicKeyProto()
	if err != nil {
		return nil, err
	}

	return &MotorNode{
		Cosmos:  c,
		Wallet:  w,
		Address: bechAddr,
		PubKey:  pk,
		DIDDoc:  w.DIDDocument,
	}, nil
}
