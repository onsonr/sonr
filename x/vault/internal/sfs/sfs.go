package sfs

import (
	"context"
	"fmt"
	"strings"

	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/pkg/node"
	servicetypes "github.com/sonrhq/core/x/service/types"
	"github.com/sonrhq/core/x/vault/types"
)

var (
	ksTable   node.IPFSKVStore
	mailTable node.IPFSDocsStore

	ctx = context.Background()
)

// The function "Init" is declared with the parameter "error" and its purpose is not specified as the code is incomplete.
func Init() error {
	err := node.StartLocalIPFS()
	if err != nil {
		return err
	}
	params := types.NewParams()

	kv, err := node.OpenKeyValueStore(ctx, params.KeyshareSeedFragment)
	if err != nil {
		return err
	}

	docs, err := node.OpenDocumentStore(ctx, params.InboxSeedFragment, nil)
	if err != nil {
		return err
	}
	ksTable = kv
	mailTable = docs
	return nil
}

// Resolve account takes a list of key shares and a coin type and returns an account.
func ClaimAccount(ucwDid string, coinType crypto.CoinType, cred *servicetypes.WebauthnCredential) (types.Account, *types.VaultKeyshare, error) {
	// Configure unclaimed wallet dids
	ks1Did := fmt.Sprintf("%s#ucw-1", ucwDid)
	ks2Did := fmt.Sprintf("%s#ucw-2", ucwDid)

	// Fetch first keyshare
	ks1, err := GetKeyshare(ks1Did)
	if err != nil {
		return nil, nil, err
	}

	// Fetch second keyshare
	ks2, err := GetKeyshare(ks2Did)
	if err != nil {
		return nil, nil, err
	}

	// Rename keyshares
	vaultKs := ks1.Rename(types.SetClaimed("vault"))
	authKs := ks2.Rename(types.SetClaimed(cred.ShortID()))

	// Insert keyshares
	go InsertKeyshare(vaultKs)
	go InsertEncryptedKeyshare(authKs, cred)

	// Return account interface
	acc := types.NewAccount(coinType, vaultKs, authKs)
	return acc, authKs.ToProto(), nil
}

// The function inserts an account and its associated key shares into a vault.
func InsertAccount(acc types.Account) {
	ksAccListVal := strings.Join(acc.ListKeyShares(), ",")
	_, err := ksTable.Put(ctx, types.AccountPrefix(acc.Did()), []byte(ksAccListVal))
	if err != nil {
		return
	}
	acc.MapKeyShare(func(ks types.KeyShare) types.KeyShare {
		go InsertKeyshare(ks)
		return ks
	})
	return
}

// The function inserts a keyshare into a table and returns an error if there is one.
func InsertKeyshare(ks types.KeyShare) {
	_, err := ksTable.Put(ctx, types.KeysharePrefix(ks.Did()), ks.Bytes())
	if err != nil {
		return
	}
	return
}

// The function inserts a keyshare into a table and returns an error if there is one.
func InsertEncryptedKeyshare(ks types.KeyShare, cred *servicetypes.WebauthnCredential) {
	dat := ks.Bytes()
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
	err := <-errCh
	_, err = ksTable.Put(ctx, types.KeysharePrefix(ks.Did()), encDat)
	if err != nil {
		return
	}

	// Check that the encrypted keyshare can be retrieved
	_, err = GetEncryptedKeyshare(ks.Did(), cred)
	if err != nil {
		panic(err)
	}
	return
}

// The function retrieves an account from a key store table using the account's DID and returns it as a
// model.
func GetAccount(accDid string) (types.Account, error) {
	ksr, err := types.ParseAccountDID(accDid)
	if err != nil {
		return nil, err
	}

	vBiz, err := ksTable.Get(ctx, types.AccountPrefix(accDid))
	if err != nil {
		return nil, err
	}

	ksAccListVal := strings.Split(string(vBiz), ",")
	var ksList []types.KeyShare
	for _, ksDid := range ksAccListVal {
		ks, err := GetKeyshare(ksDid)
		if err != nil {
			return nil, err
		}
		ksList = append(ksList, ks)
	}
	acc := types.NewAccount(ksr.CoinType, ksList...)
	return acc, nil
}

// The function retrieves a keyshare from a vault based on a given key DID.
func GetKeyshare(ksDid string) (types.KeyShare, error) {
	ksr, err := types.ParseKeyShareDID(ksDid)
	if err != nil {
		return nil, err
	}
	vBiz, err := ksTable.Get(ctx, types.KeysharePrefix(ksDid))
	if err != nil {
		return nil, err
	}
	ks, err := types.NewKeyshare(vBiz, ksr.CoinType)
	if err != nil {
		return nil, err
	}
	return ks, nil
}

// The function retrieves a keyshare from a vault based on a given key DID.
func GetEncryptedKeyshare(ksDid string, cred *servicetypes.WebauthnCredential) (types.KeyShare, error) {
	ksr, err := types.ParseKeyShareDID(ksDid)
	if err != nil {
		return nil, err
	}
	vBizch := make(chan []byte)
	errCh := make(chan error)
	go func() {
		vEnc, err := ksTable.Get(ctx, types.KeysharePrefix(ksDid))
		if err != nil {
			errCh <- err
			return
		}
		vBiz, err := cred.Decrypt(vEnc)
		if err != nil {
			errCh <- err
			return
		}
		vBizch <- vBiz
	}()
	vBiz := <-vBizch
	err = <-errCh
	ks, err := types.NewKeyshare(vBiz, ksr.CoinType)
	if err != nil {
		return nil, err
	}
	return ks, nil
}
