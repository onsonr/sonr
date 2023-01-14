package controller

import (
	"fmt"

	"github.com/sonr-hq/sonr/pkg/common/crypto"
	"github.com/sonr-hq/sonr/pkg/vault/internal/fs"
	"github.com/sonr-hq/sonr/pkg/vault/internal/network"
	"github.com/sonr-hq/sonr/x/identity/types"
)

type DIDDocumentController interface {
	// Address returns the address of the DID document.
	Address() string

	// Authenticate authenticates a WebAuthnCredential request.
	Authenticate(credentialRequestJson string) error

	// Verify verifies the given data with the DID document.
	Verify(data, signature []byte) error
}

type didDocumentController struct {
	didDocument *types.DidDocument
	vault       fs.VaultFS
	wallet      network.OfflineWallet
	accounts    map[string]crypto.Wallet
}

func NewDIDDocumentController(didDocument *types.DidDocument, vault fs.VaultFS, wallet network.OfflineWallet) didDocumentController {
	didController := didDocumentController{
		didDocument: didDocument,
		vault:       vault,
		wallet:      wallet,
	}

	return didController
}

func (c *didDocumentController) initAccounts() error {
	accounts := c.didDocument.ListBlockchainAccounts()
	for _, account := range accounts {
		//	wallet, err := c.wallet.GetWallet(account.Prefix, account.Index)
		fmt.Println(account.Address())
	}
	return nil
}
