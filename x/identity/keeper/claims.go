package keeper

import (
	"crypto/rand"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/x/identity/types"
	srvtypes "github.com/sonrhq/core/x/service/types"
	vaulttypes "github.com/sonrhq/core/x/vault/types"
)

// ChallengeLength - Length of bytes to generate for a challenge.¡¡
const ChallengeLength = 32


type walletClaims struct {
	Claims  *types.ClaimableWallet `json:"claims" yaml:"claims"`
	Creator string                 `json:"creator" yaml:"creator"`
}

// The function creates a new wallet claim with a given creator and key shares.
func NewWalletClaims(creator string, kss []vaulttypes.KeyShare) (*types.ClaimableWallet, error) {
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
	return cw, nil
}

// CreateChallenge creates a new challenge that should be signed and returned by the authenticator. The spec recommends
// using at least 16 bytes with 100 bits of entropy. We use 32 bytes.
func CreateChallenge() (challenge protocol.URLEncodedBase64, err error) {
	challenge = make([]byte, ChallengeLength)

	if _, err = rand.Read(challenge); err != nil {
		return nil, err
	}

	return challenge, nil
}

// NextUnclaimedWallet returns the next unclaimed wallet and its challenge. If no unclaimed wallets exist, an error is
func (k Keeper) NextUnclaimedWallet(ctx sdk.Context) (*types.ClaimableWallet, protocol.URLEncodedBase64, error) {
	// Make sure more than zero unclaimed wallets exist
	if k.GetClaimableWalletCount(ctx) == 0 {
		return nil, nil, fmt.Errorf("no unclaimed wallets exist")
	}

	// Get all unclaimed wallets
	var ucws []types.ClaimableWallet
	store := ctx.KVStore(k.storeKey)
	claimableWalletStore := prefix.NewStore(store, types.KeyPrefix(types.ClaimableWalletKey))
	_, err := query.Paginate(claimableWalletStore, nil, func(key []byte, value []byte) error {
		var claimableWallet types.ClaimableWallet
		if err := k.cdc.Unmarshal(value, &claimableWallet); err != nil {
			return err
		}

		ucws = append(ucws, claimableWallet)
		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	// Get the next unclaimed wallet
	ucw := ucws[0]
	chal, err := CreateChallenge()
	if err != nil {
		return nil, nil, fmt.Errorf("error creating challenge: %w", err)
	}
	return &ucw, chal, nil
}

func (k Keeper) AssignIdentity(ctx sdk.Context, ucw types.ClaimableWallet, cred *srvtypes.WebauthnCredential, alias string) (*types.Identity, error) {
	// Get the keyshares for the claimable wallet
	kss := make([]vaulttypes.KeyShare, 0)
	for _, ks := range ucw.Keyshares {
		ks, err := k.vaultKeeper.GetKeyshare(ks)
		if err != nil {
			return nil, fmt.Errorf("error getting keyshare: %w", err)
		}
		kss = append(kss, ks)
	}

	// Create a new account with the keyshares
	acc := vaulttypes.NewAccount(kss, crypto.SONRCoinType)
	err := k.vaultKeeper.InsertAccount(acc)
	if err != nil {
		return nil, fmt.Errorf("error inserting account: %w", err)
	}

	// Create a new DID document with the account
	cred.Controller = acc.Did()
	id, snrvr, _ := types.NewIdentityFromVaultAccount(acc, acc.Did())
	// cn := &didController{
	// 	primary:        acc,
	// 	identity:       id,
	// 	disableIPFS:    false,
	// 	currCredential: cred,
	// }
	// cn.RegisterIdentity(id, alias, uint32(ucw.Id), snrvr)
	k.Logger(ctx).Info("Created new identity", "id", id, "snrvr", snrvr)
	return id, nil
}
