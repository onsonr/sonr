package sfs

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"

	crypto_rand "crypto/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/pkg/mpc"
	servicetypes "github.com/sonrhq/core/x/service/types"
	"github.com/sonrhq/core/x/vault/types"
	"golang.org/x/crypto/nacl/box"
)

// The function generates a new DID and associated keyshares for a given coin type.
// It returns the DID and the associated account information.
func GetAccountInfo(ctx context.Context, did string) (*types.AccountInfo, error) {
	vks, err := GetPublicKeyshare(did)
	if err != nil {
		return nil, err
	}
	ct := crypto.CoinTypeFromBipPath(vks.CoinType)
	newDid, newAddr := ct.FormatDID(vks.PubKey())
	// Return the account info
	return &types.AccountInfo{
		Address:   newAddr,
		CoinType:  ct.Name(),
		Did:       newDid,
		Type:      vks.PubKey().KeyType,
		PublicKey: vks.PubKey().Multibase(),
	}, nil
}

// The function retrieves the public keyshare and the associated encrypted keyshare from a vault based on a given accounts DID, then derives a new set of keyshares for the given coin type.
func DeriveWithKeyshares(ctx context.Context, did string, cred *servicetypes.WebauthnCredential, coinType crypto.CoinType) (*types.AccountInfo, error) {
	vks, err := GetPublicKeyshare(did)
	if err != nil {
		return nil, err
	}
	eks, err := GetEncryptedKeyshare(did, cred)
	if err != nil {
		return nil, err
	}
	newVks, err := vks.DeriveBip44(coinType.BipPath())
	if err != nil {
		return nil, err
	}
	newEks, err := eks.DeriveBip44(coinType.BipPath())
	if err != nil {
		return nil, err
	}
	go InsertPublicKeyshare(newVks, coinType)
	go InsertEncryptedKeyshare(newEks, cred, coinType)

	newDid, newAddr := coinType.FormatDID(newVks.PubKey())
	// Return the account info
	return &types.AccountInfo{
		Address:   newAddr,
		CoinType:  coinType.Name(),
		Did:       newDid,
		Type:      newVks.PubKey().KeyType,
		PublicKey: newVks.PubKey().Multibase(),
	}, nil
}

// The function retrieves the public keyshare and the associated encrypted keyshare from a vault based on a given accounts DID, then signs a given message with all keyshares.
func SignWithKeyshares(ctx context.Context, did string, cred *servicetypes.WebauthnCredential, msg []byte) ([]byte, error) {
	vks, err := GetPublicKeyshare(did)
	if err != nil {
		return nil, fmt.Errorf("failed to get public keyshare: %w", err)
	}
	eks, err := GetEncryptedKeyshare(did, cred)
	if err != nil {
		return nil, fmt.Errorf("failed to get encrypted keyshare: %w", err)
	}
	sig, err := mpc.SignCMP([]*crypto.MPCCmpConfig{vks.CMPConfig(), eks.CMPConfig()}, msg)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

// The function retrieves the public keyshare and the associated encrypted keyshare from a vault based on a given accounts DID, then signs a given message with all keyshares.
func SignCosmosTxWithKeyshares(ctx context.Context, did string, cred *servicetypes.WebauthnCredential, msgs ...sdk.Msg) ([]byte, error) {
	vks, err := GetPublicKeyshare(did)
	if err != nil {
		return nil, fmt.Errorf("failed to get public keyshare: %w", err)
	}
	eks, err := GetEncryptedKeyshare(did, cred)
	if err != nil {
		return nil, fmt.Errorf("failed to get encrypted keyshare: %w", err)
	}
	ksc := types.NewKSS(eks, vks)
	sig, err := types.SignAnyTransactions(ksc, msgs...)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transactions: %w", err)
	}
	return sig, nil
}

// The function retrieves the public keyshare from a vault based on a given accounts DID and verifies a given signature.
func VerifyWithPublicKeyshare(ctx context.Context, did string, sig []byte, msg []byte) (bool, error) {
	ks, err := GetPublicKeyshare(did)
	if err != nil {
		return false, fmt.Errorf("failed to get public keyshare: %w", err)
	}
	return mpc.VerifyCMP(ks.CMPConfig(), msg, sig)
}
// CreateInbox sets up a default inbox for the account
func CreateInbox(accDid string) error {
	_, err := types.CreateDefaultInboxMap(accDid)
	if err != nil {
		return nil
	}
	// _, err = store.MailTable().Put(context.Background(), inbox)
	// if err != nil {
	// 	return err
	// }
	return nil
}

// HasInbox checks if the account has an inbox
func HasInbox(accDid string) (bool, error) {
	// inboxRaw, err := store.MailTable().Get(context.Background(), accDid, &iface.DocumentStoreGetOptions{})
	// if err != nil {
	// 	return false, err
	// }
	// if len(inboxRaw) == 0 {
	// 	return false, nil
	// }
	return true, nil
}

// LoadInbox loads the inbox for the account
func LoadInbox(accDid string) (*types.Inbox, error) {
	// // Check if the inbox exists
	// hasInbox, err := HasInbox(accDid)
	// if err != nil {
	// 	return nil, err
	// }
	// if !hasInbox {
	// 	err := CreateInbox(accDid)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	// // Load the inbox
	// inboxRaw, err := store.MailTable().Get(context.Background(), accDid, &iface.DocumentStoreGetOptions{})
	// inboxMap, ok := inboxRaw[0].(map[string]interface{})
	// if !ok {
	// 	return nil, errors.New("invalid inbox")
	// }
	// inbox, err := types.NewInboxFromMap(inboxMap)
	// if err != nil {
	// 	return nil, err
	// }
	// return inbox, nil
	return nil, errors.New("not implemented")
}

// ReadInbox reads the inbox for the account
func ReadInbox(accDid string) ([]*types.WalletMail, error) {
	// inbox, err := LoadInbox(accDid)
	// if err != nil {
	// 	return nil, err
	// }
	// return inbox.Messages, nil
	return nil, errors.New("not implemented")
}

// WriteInbox writes the inbox to the database
func WriteInbox(toDid string, msg *types.WalletMail) error {
	// // Get the inbox
	// inbox, err := LoadInbox(toDid)
	// if err != nil {
	// 	return err
	// }
	// // Add the message to the inbox
	// inboxMap, err := inbox.AddMessageToMap(msg)
	// if err != nil {
	// 	return err
	// }
	// // Update the inbox
	// _, err = store.MailTable().Put(context.Background(), inboxMap)
	// if err != nil {
	// 	return err
	// }
	// return nil
	return errors.New("not implemented")
}

type mbKey struct {
	sk   []byte
	pub  *[32]byte
	priv *[32]byte
}

func NewMailboxKey(secretKey []byte) (*mbKey, error) {
	if len(secretKey) != 32 {
		return nil, fmt.Errorf("invalid secret key length: %d, needs to be 32", len(secretKey))
	}
	pub, priv, err := box.GenerateKey(bytes.NewReader(secretKey))
	if err != nil {
		return nil, err
	}
	return &mbKey{
		sk:   secretKey,
		pub:  pub,
		priv: priv,
	}, nil
}

// Seal message with the public key of the recipient.
func (k *mbKey) SealMessage(message []byte, recipient []byte) ([]byte, error) {
	peersPublicKey := bytesToPointer(recipient)
	var nonce [24]byte
	if _, err := io.ReadFull(crypto_rand.Reader, nonce[:]); err != nil {
		return nil, err
	}
	// This encrypts msg and appends the result to the nonce.
	encrypted := box.Seal(nonce[:], message, &nonce, peersPublicKey, k.priv)
	return encrypted, nil
}

func (k *mbKey) PublicKeyBase64() string {
	return base64.StdEncoding.EncodeToString(k.NormalizePublicKey())
}

func (k *mbKey) Type() string {
	return "mailbox"
}

func (k *mbKey) NormalizePrivateKey() []byte {
	return k.priv[:]
}

func (k *mbKey) NormalizePublicKey() []byte {
	return k.pub[:]
}

func bytesToPointer(b []byte) *[32]byte {
	var a [32]byte
	copy(a[:], b)
	return &a
}
