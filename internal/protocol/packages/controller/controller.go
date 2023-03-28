package controller

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/derekparker/trie"
	"github.com/getsentry/sentry-go"
	"github.com/sonrhq/core/internal/protocol/packages/resolver"
	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/pkg/crypto/mpc"
	"github.com/sonrhq/core/x/identity/types"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

var PrimaryAccountName string = "primary"

type Controller interface {
	// Address returns the controller's address
	Address() string

	// Get the controller's DID
	Did() string

	// Get the controller's DID document
	PrimaryIdentity() *types.DidDocument

	// Blockchain returns the controller's blockchain identities
	BlockchainIdentities() []*types.DidDocument

	// Authorize the client to access the controller's account
	Authorize(cred *crypto.WebauthnCredential) error

	// CreateAccount creates a new account for the controller
	CreateAccount(name string, coinType crypto.CoinType) (Account, error)

	// GetAccount returns the controller's account
	GetAccount(name string, coinType crypto.CoinType) (Account, error)

	// ListAccounts returns the controller's accounts
	ListAccounts(ct crypto.CoinType) ([]Account, error)

	// ListLocalAccounts returns the controller's local accounts (WARNING: this is not secure)
	ListLocalAccounts() []Account

	// Sign signs a message with the controller's account
	Sign(name string, coinType crypto.CoinType, msg []byte) ([]byte, error)

	// Verify verifies a signature with the controller's account
	Verify(name string, coinType crypto.CoinType, msg []byte, sig []byte) (bool, error)
}

type didController struct {
	primary    Account
	primaryDoc *types.DidDocument
	blockchain []Account
}

type Options struct {
	// The controller's DID document
	GenerateInitialAccounts []string

	// The controller's on config generated handler
	OnConfigGenerated []mpc.OnConfigGenerated

	// Credential to authorize the controller
	WebauthnCredential *crypto.WebauthnCredential
}

type Option func(*Options)

func WithInitialAccounts(accounts ...string) Option {
	return func(o *Options) {
		o.GenerateInitialAccounts = accounts
	}
}

func WithConfigHandlers(handlers ...mpc.OnConfigGenerated) Option {
	return func(o *Options) {
		o.OnConfigGenerated = handlers
	}
}

func WithWebauthnCredential(cred *crypto.WebauthnCredential) Option {
	return func(o *Options) {
		o.WebauthnCredential = cred
	}
}

func NewController(ctx context.Context, options ...Option) (Controller, Account, error) {
	opts := &Options{}
	for _, option := range options {
		option(opts)
	}
	cred, err := ValidateWebauthnCredential(opts.WebauthnCredential)
	if err != nil {
		fmt.Println("Warning - Error validating webauthn credential: ", err)
	}
	doneCh := make(chan Account)
	errCh := make(chan error)

	go generateInitialAccount(ctx, cred, doneCh, errCh, opts)

	select {
	case acc := <-doneCh:
		cn, err := setupController(ctx, acc, opts)
		if err != nil {
			return nil, nil, err
		}
		return cn, acc, nil
	case err := <-errCh:
		sentry.CaptureException(err)
		return nil, nil, err
	}
}

// AuthorizeController loads a controller from the given DID document using the underlying IPFS store
func AuthorizeController(ctx context.Context, didDoc *types.DidDocument) (Controller, error) {
	// Get the IPFS store service
	ksm, err := resolver.ListKeyShares(didDoc.Id)
	if err != nil {
		return nil, err
	}

	// Get the primary account
	kss := make([]KeyShare, 0)
	for did, bz := range ksm {
		ks, err := LoadKeyshareFromStore(did, bz)
		if err != nil {
			return nil, err
		}
		kss = append(kss, ks)
	}
	primary := NewAccount(kss, crypto.SONRCoinType)
	return &didController{
		primary: primary,
		primaryDoc: didDoc,
		blockchain: []Account{},
	}, nil
}

func (dc *didController) Address() string {
	return dc.primary.Address()
}

func (dc *didController) Did() string {
	return dc.primary.Did()
}

func (dc *didController) PrimaryIdentity() *types.DidDocument {
	return dc.primaryDoc
}

func (dc *didController) BlockchainIdentities() []*types.DidDocument {
	var docs []*types.DidDocument
	for _, acc := range dc.blockchain {
		docs = append(docs, acc.DidDocument(dc.Did()))
	}
	return docs
}

func (dc *didController) Authorize(cred *crypto.WebauthnCredential) error {
	return nil
}

func (dc *didController) ListLocalAccounts() []Account {
	return dc.blockchain
}

func (dc *didController) CreateAccount(name string, coinType crypto.CoinType) (Account, error) {
	kss, err := dc.primary.ListKeyshares()
	if err != nil {
		return nil, err
	}
	var cmpcnfs []*cmp.Config
	for _, ks := range kss {
		cmpcnfs = append(cmpcnfs, ks.Config())
	}

	newAccCh := make(chan Account)
	errCh := make(chan error)
	go func() {
		var newKss []KeyShare
		for _, conf := range cmpcnfs {
			newConf, err := conf.DeriveBIP32(uint32(coinType.BipPath()))
			if err != nil {
				errCh <- err
				return
			}
			ksb, err := newConf.MarshalBinary()
			if err != nil {
				errCh <- err
				return
			}
			ks, err := NewKeyshare(string(newConf.ID), ksb, coinType, name)
			if err != nil {
				errCh <- err
				return
			}
			newKss = append(newKss, ks)
		}
		newAccCh <- NewAccount(newKss, coinType)
	}()

	// Create the new account and map the keyshares to the resolver
	select {
	case newAcc := <-newAccCh:
		err = resolver.InsertAccountInfo(newAcc.ToStore())
		if err != nil {
			return nil, err
		}
		dc.blockchain = append(dc.blockchain, newAcc)
		dc.primaryDoc.AddBlockchainIdentity(newAcc.DidDocument(dc.Did()))
		return newAcc, nil
	case err := <-errCh:
		sentry.CaptureException(err)
		return nil, err
	}
}

// GetAccount returns the controller's account from the Address
func (dc *didController) GetAccount(name string, coinType crypto.CoinType) (Account, error) {
	return nil, fmt.Errorf("account not found")
}

// ListAccounts returns the controller's accounts
func (dc *didController) ListAccounts(ct crypto.CoinType) ([]Account, error) {
	// Get the IPFS store service
	// mapKv, err := resolver.ListKeyShares()
	// if err != nil {
	// 	return nil, err
	// }
	var accs []Account
	// filtered := fuzzySearch(mapKv, dc.Address(), FilterOptions{})
	// for _, k := range filtered {
	// 	acc, err := dc.GetAccount(k.AccountName(), ct)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	accs = append(accs, acc)
	// }
	return accs, nil
}

// Sign signs a message with the controller's selected account
func (dc *didController) Sign(name string, coinType crypto.CoinType, msg []byte) ([]byte, error) {
	acc, err := dc.GetAccount(name, coinType)
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}
	return acc.Sign(msg)
}

// Verify verifies a signature with the controller's selected account
func (dc *didController) Verify(name string, coinType crypto.CoinType, msg []byte, sig []byte) (bool, error) {
	acc, err := dc.GetAccount(name, coinType)
	if err != nil {
		sentry.CaptureException(err)
		return false, err
	}
	return acc.Verify(msg, sig)
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                          Helper Methods for Controller                         ||
// ! ||--------------------------------------------------------------------------------||

func generateInitialAccount(ctx context.Context, credential *crypto.WebauthnCredential, doneCh chan Account, errChan chan error, opts *Options) {
	shardName := crypto.PartyID(base64.RawStdEncoding.EncodeToString(credential.Id))
	// Call Handler for keygen
	confs, err := mpc.Keygen(shardName, 1, []crypto.PartyID{"vault"}, opts.OnConfigGenerated...)
	if err != nil {
		errChan <- err
	}

	pubKey, err := crypto.NewPubKeyFromCmpConfig(confs[0])
	if err != nil {
		errChan <- err
	}

	rootDid, _ := crypto.SONRCoinType.FormatDID(pubKey)
	var kss []KeyShare
	for _, conf := range confs {
		ksb, err := conf.MarshalBinary()
		if err != nil {
			errChan <- err
		}
		ksDid := fmt.Sprintf("%s#%s", rootDid, conf.ID)
		err = resolver.InsertKSItem(ksDid, ksb)
		if err != nil {
			errChan <- err
		}
		ks, err := NewKeyshare(ksDid, ksb, crypto.SONRCoinType, "Primary")
		if err != nil {
			errChan <- err
		}
		kss = append(kss, ks)
	}
	doneCh <- NewAccount(kss, crypto.SONRCoinType)
}

func setupController(ctx context.Context, primary Account, opts *Options) (Controller, error) {
	err := resolver.InsertAccountInfo(primary.ToStore())
	if err != nil {
		sentry.CaptureException(err)
		return nil, err
	}

	var doc *types.DidDocument
	if opts.WebauthnCredential != nil {
		cvm, err := resolver.EncodeCredentialVerificationMethod(opts.WebauthnCredential, primary.Did())
		if err != nil {
			sentry.CaptureException(err)
			return nil, err
		}
		doc = types.NewPrimaryIdentity(primary.Did(), primary.PubKey(), cvm)
	} else {
		doc = types.NewPrimaryIdentity(primary.Did(), primary.PubKey(), nil)
	}


	cont := &didController{
		primary:    primary,
		blockchain: []Account{},
		primaryDoc: doc,
	}
	cts := []crypto.CoinType{}
	for _, ct := range opts.GenerateInitialAccounts {
		cts = append(cts, crypto.CoinTypeFromName(ct))
	}

	for _, ct := range cts {
		acc, err := cont.CreateAccount("Account 1", ct)
		if err != nil {
			sentry.CaptureException(err)
			return nil, err
		}
		cont.blockchain = append(cont.blockchain, acc)
	}
	return cont, nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                              Map & Slice Filtering                             ||
// ! ||--------------------------------------------------------------------------------||
type FilterOptions struct {
	CoinType    crypto.CoinType
	AccountName *string
	Index       *int
}

func fuzzySearch(m []resolver.KVStoreItem, query string, options FilterOptions) []KeyShare {
	mapKv := make(map[string][]byte)
	// Create a trie and insert keys
	t := trie.New()
	for _, i := range m {
		t.Add(i.Did(), i.Bytes())
		mapKv[i.Did()] = i.Bytes()
	}

	// Perform fuzzy search with a query
	matches := t.FuzzySearch(query)

	// Filter results based on the provided options
	results := make([]KeyShare, 0)
	for _, match := range matches {
		ksr, err := ParseKeyShareDid(match)
		if err != nil {
			continue
		}
		if ksr.CoinType != options.CoinType {
			continue
		}
		if options.AccountName != nil && ksr.AccountName != *options.AccountName {
			continue
		}
		ks, err := NewKeyshare(ksr.DID, mapKv[match], ksr.CoinType, ksr.AccountName)
		if err != nil {
			continue
		}
		results = append(results, ks)
	}

	return results
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                       WebauthnCredential utility methods                       ||
// ! ||--------------------------------------------------------------------------------||

func ValidateWebauthnCredential(credential *crypto.WebauthnCredential) (*crypto.WebauthnCredential, error) {
	// Check for nil credential
	if credential == nil {
		return &crypto.WebauthnCredential{
			Id:        []byte("user1"),
			PublicKey: []byte{0x00},
		}, errors.New("credential is nil")
	}

	// Check for nil credential id
	if credential.Id == nil {
		credential.Id = []byte("user1")
	}

	// Check for nil credential public key
	if credential.PublicKey == nil {
		credential.PublicKey = []byte{0x00}
	}
	return credential, nil
}
