package sfs

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/kataras/jwt"
	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/internal/stores"
	servicetypes "github.com/sonrhq/core/x/service/types"
	"github.com/sonrhq/core/x/vault/types"
)

const (
	kCategoryUnclaimed = "unclaimed-wallets"
	kCategoryPending   = "pending-wallets"
	kCategoryVault     = "vault-wallets"
)

var (
	kTempSonrJWTSigningKey = []byte("sercrethatmaycontainch@r$32chars")
)

// The function inserts a keyshare into a table and returns an error if there is one.
func InsertUnclaimedKeyshare(ks *types.VaultKeyshare, coinType crypto.CoinType, idx int) error {
	frg := fmt.Sprintf("ucw-%d", idx)
	did, _ := coinType.FormatDID(ks.PubKey())
	dat, err := json.Marshal(ks)
	if err != nil {
		return err
	}
	err = store.AddSetItem(stores.CategorySetKeyName(kCategoryUnclaimed), did)
	if err != nil {
		return err
	}

	err = store.AddMapItem(stores.KeyshareMapKeyName(did), frg, crypto.Base64Encode(dat))
	if err != nil {
		return err
	}
	return nil
}

// The function inserts the public vault keyshare into a table and returns an error if there is one.
func InsertPublicKeyshare(ks *types.VaultKeyshare, coinType crypto.CoinType) error {
	did, _ := coinType.FormatDID(ks.PubKey())
	dat, err := json.Marshal(ks)
	if err != nil {
		return err
	}
	err = store.AddMapItem(stores.KeyshareMapKeyName(did), "vault", crypto.Base64Encode(dat))
	if err != nil {
		return err
	}
	return nil
}

// The function inserts a keyshare into a table and returns an error if there is one.
func InsertEncryptedKeyshare(ks *types.VaultKeyshare, cred *servicetypes.WebauthnCredential, coinType crypto.CoinType) error {
	did, _ := coinType.FormatDID(ks.PubKey())
	frg := cred.ShortID()
	dat, err := json.Marshal(ks)
	if err != nil {
		return err
	}
	datCh := make(chan []byte)
	errCh := make(chan error)
	go func() {
		encDat, err := cred.Encrypt(dat)
		if err != nil {
			errCh <- err
			return
		}
		datCh <- encDat
	}()
	encDat := <-datCh
	err = <-errCh
	if err != nil {
		return err
	}
	err = store.AddMapItem(stores.KeyshareMapKeyName(did), frg, crypto.Base64Encode(encDat))
	if err != nil {
		return err
	}
	return nil
}

// The function retrieves all unclaimed keyshares from a vault based on a given accounts DID.
func GetUnclaimedKeyshares(did string) (types.KeyShareCollection, error) {
	ok, err := store.ExistsInSet(stores.CategorySetKeyName(kCategoryPending), did)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("This DID is not in the pending set. Please verify that the DID is correct")
	}
	ksDid1 := "ucw-1"
	ksDid2 := "ucw-2"
	ksMap, err := store.GetMap(stores.KeyshareMapKeyName(did))
	if err != nil {
		return nil, err
	}
	ks1Bz, err := crypto.Base64Decode(ksMap[ksDid1])
	ks2Bz,err  := crypto.Base64Decode(ksMap[ksDid2])
	ks1, err := types.LoadKeyshare(ks1Bz)
	if err != nil {
		return nil, err
	}
	ks2, err := types.LoadKeyshare(ks2Bz)
	if err != nil {
		return nil, err
	}
	_, err = store.DelMapItem(stores.KeyshareMapKeyName(did), ksDid1)
	if err != nil {
		return nil, err
	}
	_, err = store.DelMapItem(stores.KeyshareMapKeyName(did), ksDid2)
	if err != nil {
		return nil, err
	}
	_, err = store.DelSetItem(stores.CategorySetKeyName(kCategoryPending), did)
	if err != nil {
		return nil, err
	}
	return []*types.VaultKeyshare{ks1, ks2}, nil
}

// The function retrieves the public keyshare from a vault based on a given accounts DID.
func GetPublicKeyshare(did string) (*types.VaultKeyshare, error) {
	ok, err := store.ExistsInSet(stores.CategorySetKeyName(kCategoryVault), did)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("This DID is not in the controlled vault set. Please verify that the DID is correct")
	}
	ksDid := "vault"
	ksMap, err := store.GetMap(stores.KeyshareMapKeyName(did))
	if err != nil {
		return nil, err
	}
	vBiz, err := crypto.Base64Decode(ksMap[ksDid])
	if err != nil {
		return nil, err
	}
	ks, err := types.LoadKeyshare(vBiz)
	if err != nil {
		return nil, err
	}
	return ks, nil
}

// The function retrieves a keyshare from a vault based on a given key DID.
func GetEncryptedKeyshare(did string, cred *servicetypes.WebauthnCredential) (*types.VaultKeyshare, error) {
	ok, err := store.ExistsInSet(stores.CategorySetKeyName(kCategoryVault), did)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("This DID is not in the controlled vault set. Please verify that the DID is correct")
	}
	ksDid := cred.ShortID()
	ksMap, err := store.GetMap(stores.KeyshareMapKeyName(did))
	if err != nil {
		return nil, err
	}
	vEnc, err := crypto.Base64Decode(ksMap[ksDid])
	if err != nil {
		return nil, err
	}
	vBiz, err := cred.Decrypt(vEnc)
	if err != nil {
		return nil, err
	}
	ks, err := types.LoadKeyshare(vBiz)
	if err != nil {
		return nil, err
	}
	return ks, nil
}

// RandomUnclaimedWallet returns a random unclaimed wallet from the store
func RandomUnclaimedWallet() (string, error) {
	allUnclaimed, err := store.GetSet(stores.CategorySetKeyName(kCategoryUnclaimed))
	if err != nil {
		return "", err
	}
	if len(allUnclaimed) == 0 {
		return "", fmt.Errorf("No unclaimed wallets found")
	}
	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn(len(allUnclaimed))
	ucwDid := allUnclaimed[randIndex]

	// Add the wallet to the pending set
	err = store.AddSetItem(stores.CategorySetKeyName(kCategoryPending), ucwDid)
	if err != nil {
		return "", err
	}

	// Remove the wallet from the unclaimed set
	_, err = store.DelSetItem(stores.CategorySetKeyName(kCategoryUnclaimed), ucwDid)
	if err != nil {
		return "", err
	}
	return ucwDid, nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                 JWT credential                                 ||
// ! ||--------------------------------------------------------------------------------||

// RetreiveCredential returns the *servicetypes.WebauthnCredential from the JWT token string
func RetreiveCredential(tokenString string) (*servicetypes.WebauthnCredential, error) {
	// Verify and extract claims from a token:
	verifiedToken, err := jwt.Verify(jwt.HS256, kTempSonrJWTSigningKey, []byte(tokenString))
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	// Extract the claims
	var claims JWTClaims
	err = verifiedToken.Claims(&claims)
	if err != nil {
		panic(err)
	}
	var credential servicetypes.WebauthnCredential
	err = credential.Unmarshal(claims.Credential)
	if err != nil {
		return nil, err
	}
	return &credential, nil
}

// StoreCredential stores the credential in the icefire redis store with a 1 hour expiry. The client is returned the JWT token string.
func StoreCredential(did string, cred *servicetypes.WebauthnCredential) (string, error) {
	bz, err := cred.Marshal()
	if err != nil {
		return "", err
	}
	claims := JWTClaims{
		Did: did,
		Credential: bz,
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
	}

	token, err := jwt.Sign(jwt.HS256, kTempSonrJWTSigningKey, claims)
	if err != nil {
		return "", err
	}
	return string(token), nil
}
