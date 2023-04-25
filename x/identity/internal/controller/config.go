package controller

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/internal/crypto/mpc"
	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/x/identity/internal/vault"
	"github.com/sonrhq/core/x/identity/types"
	"github.com/sonrhq/core/x/identity/types/models"
	servicetypes "github.com/sonrhq/core/x/service/types"
)

// ! ||--------------------------------------------------------------------------------||
// ! ||                              Client Facing Options                             ||
// ! ||--------------------------------------------------------------------------------||

type Options struct {
	// The controller's on config generated handler
	OnConfigGenerated []mpc.OnConfigGenerated

	// Credential to authorize the controller
	WebauthnCredential *servicetypes.WebauthnCredential

	// Disable IPFS
	DisableIPFS bool

	// Broadcast the transaction
	BroadcastTx bool

	// Username for the controller
	Username string

	errChan chan error
	broadcastChan chan *local.BroadcastTxResponse
}

func defaultOptions() *Options {
	return &Options{
		OnConfigGenerated: []mpc.OnConfigGenerated{},
		DisableIPFS:       false,
		BroadcastTx:       false,
		Username:          "",
		errChan:           make(chan error),
		broadcastChan: make(chan *local.BroadcastTxResponse),
	}
}

type Option func(*Options)

func WithUsername(username string) Option {
	return func(o *Options) {
		o.Username = username
	}
}

func WithConfigHandlers(handlers ...mpc.OnConfigGenerated) Option {
	return func(o *Options) {
		o.OnConfigGenerated = handlers
	}
}

func WithWebauthnCredential(cred *servicetypes.WebauthnCredential) Option {
	return func(o *Options) {
		o.WebauthnCredential = cred
	}
}

func WithIPFSDisabled() Option {
	return func(o *Options) {
		o.DisableIPFS = true
	}
}

func WithBroadcastTx(brdcastChan chan *local.BroadcastTxResponse) Option {
	return func(o *Options) {
		o.BroadcastTx = true
		o.broadcastChan = brdcastChan
	}
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                          Helper Methods for Controller                         ||
// ! ||--------------------------------------------------------------------------------||

func generateInitialAccount(ctx context.Context, credential *servicetypes.WebauthnCredential, doneCh chan models.Account, errChan chan error, opts *Options) {
	shardName := crypto.PartyID(base64.RawStdEncoding.EncodeToString(credential.Id))
	// Call Handler for keygen
	confs, err := mpc.Keygen(shardName, mpc.WithHandlers(opts.OnConfigGenerated...))
	if err != nil {
		errChan <- err
	}

	pubKey, err := crypto.NewPubKeyFromCmpConfig(confs[0])
	if err != nil {
		errChan <- err
	}

	rootDid, _ := crypto.SONRCoinType.FormatDID(pubKey)
	var kss []models.KeyShare
	for _, conf := range confs {
		ksb, err := conf.MarshalBinary()
		if err != nil {
			errChan <- err
		}
		ksDid := fmt.Sprintf("%s#%s", rootDid, conf.ID)
		ks, err := models.NewKeyshare(ksDid, ksb, crypto.SONRCoinType)
		if err != nil {
			errChan <- err
		}
		kss = append(kss, ks)
	}
	doneCh <- models.NewAccount(kss, crypto.SONRCoinType)
}

func setupController(ctx context.Context, primary models.Account, opts *Options) (Controller, error) {
	if !opts.DisableIPFS {
		err := vault.InsertAccount(primary)
		if err != nil {
			return nil, err
		}
	}

	doc := types.NewPrimaryIdentity(primary.Did(), primary.PubKey(), nil)
	if opts.WebauthnCredential != nil {
		cred, err := servicetypes.ValidateWebauthnCredential(opts.WebauthnCredential, primary.Did())
		if err != nil {
			return nil, err
		}
		doc = types.NewPrimaryIdentity(primary.Did(), primary.PubKey(), cred.ToVerificationMethod())
		if !opts.DisableIPFS {
			err = vault.StoreCredential(cred)
			if err != nil {
				return nil, err
			}
		}
	}

	if opts.Username != "" {
		doc.AlsoKnownAs = []string{opts.Username}
	}

	cont := &didController{
		primary:     primary,
		blockchain:  []models.Account{},
		primaryDoc:  doc,
		disableIPFS: opts.DisableIPFS,
		txHash:      "",
		aka:         doc.FindUsername(),
	}
	return cont, nil
}
