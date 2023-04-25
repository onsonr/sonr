package controller

import (
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/x/identity/internal/vault"
	"github.com/sonrhq/core/x/identity/types"
	"github.com/sonrhq/core/x/identity/types/models"
	srvtypes "github.com/sonrhq/core/x/service/types"
)

type WalletClaims interface {
	GetClaimableWallet() *types.ClaimableWallet
	IssueChallenge() (protocol.URLEncodedBase64, error)
	Assign(cred *srvtypes.WebauthnCredential, alias string) (Controller, error)
}

type walletClaims struct {
	Claims  *types.ClaimableWallet `json:"claims" yaml:"claims"`
	Creator string                 `json:"creator" yaml:"creator"`
}

// The function creates a new wallet claim with a given creator and key shares.
func NewWalletClaims(creator string, kss []models.KeyShare) (WalletClaims, error) {
	pub := kss[0].PubKey()
	keyIds := make([]string, 0)
	for _, ks := range kss {
		keyIds = append(keyIds, ks.Did())
	}

	cw := &types.ClaimableWallet{
		Creator:   creator,
		PublicKey: pub.Base64(),
		Keyshares: keyIds,
		Count:     int32(len(kss)),
		Claimed:   false,
	}

	return &walletClaims{
		Claims:  cw,
		Creator: creator,
	}, nil
}

// The function returns a WalletClaims interface that contains the claimable wallet and its creator.
func LoadClaimableWallet(cw *types.ClaimableWallet) WalletClaims {
	return &walletClaims{
		Claims:  cw,
		Creator: cw.Creator,
	}
}

// The `GetClaimableWallet()` function is a method of the `walletClaims` struct that returns a pointer
// to the `ClaimableWallet` object stored in the struct. This allows other parts of the code to access
// the `ClaimableWallet` object and its properties.
func (wc *walletClaims) GetClaimableWallet() *types.ClaimableWallet {
	return wc.Claims
}

// This function is used to issue a challenge for the claimable wallet. It returns the public key of
// the claimable wallet as a URL-encoded base64 string, which can be used as a challenge for WebAuthn
// authentication. If the public key is empty, it returns an error.
func (wc *walletClaims) IssueChallenge() (protocol.URLEncodedBase64, error) {
	if wc.Claims.PublicKey == "" {
		return nil, fmt.Errorf("public key is empty")
	}
	return protocol.URLEncodedBase64(wc.Claims.PublicKey), nil
}

// This function assigns a WebAuthn credential to a claimable wallet by creating a new DID document and
// adding the credential as an additional authentication method. It then creates a new `didController`
// instance with the new DID document and returns it as a `Controller` interface. The `alias` parameter
// is used to set the `AlsoKnownAs` field in the DID document.
func (wc *walletClaims) Assign(cred *srvtypes.WebauthnCredential, alias string) (Controller, error) {
	kss := make([]models.KeyShare, 0)
	for _, ks := range wc.Claims.Keyshares {
		ks, err := vault.GetKeyshare(ks)
		if err != nil {
			return nil, err
		}
		kss = append(kss, ks)
	}

	acc := models.NewAccount(kss, crypto.SONRCoinType)
	doc := acc.DidDocument()
	credential := srvtypes.NewCredential(cred, doc.Id)
	vm := credential.ToVerificationMethod()
	_, err := doc.LinkAdditionalAuthenticationMethod(vm)
	if err != nil {
		return nil, err
	}

	doc.AlsoKnownAs = []string{alias}

	cn := &didController{
		primary:    acc,
		primaryDoc: doc,
		blockchain: []models.Account{},
		txHash: "",
		disableIPFS: false,
		currCredential: cred,
	}
	cn.CreatePrimaryIdentity(doc, acc, alias, uint32(wc.Claims.Id))
	return cn, nil
}
