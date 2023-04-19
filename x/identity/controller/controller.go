package controller

import (
	"context"
	"fmt"
	"strings"

	// "github.com/sonrhq/core/internal/vault"
	// "github.com/sonrhq/core/internal/vault"
	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/x/identity/keeper"
	"github.com/sonrhq/core/x/identity/types"
	"github.com/sonrhq/core/x/identity/types/models"
)

var PrimaryAccountaddress string = "primary"

type Controller interface {
	// Address returns the controller's address
	Address() string

	// Get the controller's DID
	Did() string

	// PrimaryIdentity returns the controller's DID document
	PrimaryIdentity() *types.DidDocument

	// PrimaryTxHash returns the controller's primary identity transaction hash
	PrimaryTxHash() string

	// BlockchainIdentities returns the controller's blockchain identities
	BlockchainIdentities() []*types.DidDocument

	// Createmodels.Account creates a new models.Account for the controller
	CreateAccount(name string, coinType crypto.CoinType) (models.Account, error)

	// GetAccount returns an account by Address or DID
	GetAccount(id string) (models.Account, error)

	// Listmodels.Accounts returns the controller's models.Accounts
	ListAccounts() ([]models.Account, error)

	// SendMail sends a message between two Controllers
	SendMail(address string, to string, body string) error

	// ReadMail reads the controller's inbox
	ReadMail(address string) ([]*models.InboxMessage, error)

	// Sign signs a message with the controller's models.Account
	Sign(address string, msg []byte) ([]byte, error)

	// Verify verifies a signature with the controller's models.Account
	Verify(address string, msg []byte, sig []byte) (bool, error)
}

type didController struct {
	primary    models.Account
	primaryDoc *types.DidDocument
	blockchain []models.Account

	currCredential *crypto.WebauthnCredential
	disableIPFS    bool
	aka            string
	txHash         string
}

func NewController(options ...Option) (Controller, error) {
	opts := defaultOptions()
	for _, option := range options {
		option(opts)
	}

	doneCh := make(chan models.Account)
	errCh := make(chan error)
	go generateInitialAccount(context.Background(), opts.WebauthnCredential, doneCh, errCh, opts)

	select {
	case acc := <-doneCh:
		cn, err := setupController(context.Background(), acc, opts)
		if err != nil {
			return nil, err
		}
		return cn, nil
	case err := <-errCh:
		return nil, err
	}
}

func LoadController(doc *types.DidDocument) (Controller, error) {
	acc, err := keeper.GetAccount(doc.Id)
	if err != nil {
		return nil, err
	}
	blockAccDids := doc.ListBlockchainIdentities()
	var blockAccs []models.Account
	for _, did := range blockAccDids {
		acc, err := keeper.GetAccount(did)
		if err != nil {
			return nil, err
		}
		blockAccs = append(blockAccs, acc)
	}
	cn := &didController{
		primary:    acc,
		primaryDoc: doc,
		blockchain: blockAccs,
	}
	return cn, nil
}

func (dc *didController) Address() string {
	return dc.primary.Address()
}

func (dc *didController) Did() string {
	return dc.primaryDoc.Id
}

func (dc *didController) PrimaryIdentity() *types.DidDocument {
	return dc.primaryDoc
}

func (dc *didController) BlockchainIdentities() []*types.DidDocument {
	var docs []*types.DidDocument
	for _, acc := range dc.blockchain {
		docs = append(docs, acc.DidDocument())
	}
	return docs
}

func (dc *didController) ListAccounts() ([]models.Account, error) {
	if dc.primary == nil {
		return nil, fmt.Errorf("no Primary Account found")
	}
	return append([]models.Account{dc.primary}, dc.blockchain...), nil
}

func (dc *didController) CreateAccount(name string, coinType crypto.CoinType) (models.Account, error) {
	ctCount := 0
	for _, acc := range dc.blockchain {
		if acc.CoinType() == coinType {
			ctCount++
		}
	}
	newAcc, err := dc.primary.DeriveAccount(coinType, ctCount, name)
	if err != nil {
		return nil, err
	}

	// Add account to the vault
	if !dc.disableIPFS {
		err = keeper.InsertAccount(newAcc)
		if err != nil {
			return nil, err
		}
	}

	// Add the new models.Account to the controller
	dc.blockchain = append(dc.blockchain, newAcc)
	dc.primaryDoc.AddBlockchainIdentity(newAcc.DidDocument())
	dc.UpdatePrimaryIdentity(newAcc.DidDocument())
	return newAcc, nil
}

// Getmodels.Account returns the controller's models.Account from the Address
func (dc *didController) GetAccount(address string) (models.Account, error) {
	if strings.Contains(address, "did:") {
		return dc.GetAccountByDid(address)
	}
	for _, acc := range dc.blockchain {
		if acc.Address() == address {
			return acc, nil
		}
	}
	return nil, fmt.Errorf("models.Account not found")
}

// GetAccountByDid returns the controller's models.Account from the DID
func (dc *didController) GetAccountByDid(did string) (models.Account, error) {
	if dc.primaryDoc.Id == did {
		return dc.primary, nil
	}
	for _, acc := range dc.blockchain {
		if acc.DidDocument().Id == did {
			return acc, nil
		}
	}
	return nil, fmt.Errorf("models.Account not found")
}

// Sign signs a message with the controller's selected models.Account
func (dc *didController) Sign(address string, msg []byte) ([]byte, error) {
	acc, err := dc.GetAccount(address)
	if err != nil {
		return nil, err
	}
	return acc.Sign(msg)
}

// Verify verifies a signature with the controller's selected models.Account
func (dc *didController) Verify(address string, msg []byte, sig []byte) (bool, error) {
	acc, err := dc.GetAccount(address)
	if err != nil {

		return false, err
	}
	return acc.Verify(msg, sig)
}

// SendMail sends a mail from the controller's selected models.Account
func (dc *didController) SendMail(address string, to string, body string) error {
	acc, err := dc.GetAccount(address)
	if err != nil {
		return err
	}
	msg, err := acc.CreateInboxMessage(to, body)
	if err != nil {
		return err
	}
	err = keeper.WriteInbox(to, msg)
	if err != nil {
		return err
	}
	return nil
}

// ReadMail reads a mail from the controller's selected models.Account
func (dc *didController) ReadMail(address string) ([]*models.InboxMessage, error) {
	acc, err := dc.GetAccount(address)
	if err != nil {
		return nil, err
	}
	return keeper.ReadInbox(acc.Address())
}

// PrimaryTxHash returns the transaction hash of the primary models.Account
func (dc *didController) PrimaryTxHash() string {
	return dc.txHash
}
