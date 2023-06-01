package sfs

import (
	"strings"

	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/x/vault/types"
)

// Resolve account takes a list of key shares and a coin type and returns an account.
func ClaimAccount( ksDidList []string, coinType crypto.CoinType) (types.Account, error) {
	kss := make([]types.KeyShare, 0)
	for _, ks := range ksDidList {
		ks, err := GetKeyshare( ks)
		if err != nil {
			return nil, err
		}
		kss = append(kss, ks)
	}

	acc := types.NewAccount(kss, coinType)
	go InsertAccount(acc)
	return acc, nil
}

// The function inserts an account and its associated key shares into a vault.
func InsertAccount(acc types.Account) {
	ksAccListVal := strings.Join(acc.ListKeyShares(), ",")
	_, err := ksTable.Put(ctx, types.AccountPrefix(acc.Did()), []byte(ksAccListVal))
	if err != nil {
		return
	}
	secKey, err := acc.GenerateSecretKey(types.DefaultParams().KeyshareSeedFragment)
	if err != nil {
		return
	}
	acc.MapKeyShare(func(ks types.KeyShare) types.KeyShare {
		go insertEncKeyshare(ks, secKey)
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
	acc := types.NewAccount(ksList, ksr.CoinType)
	return acc, nil
}

// The function retrieves a keyshare from a vault based on a given key DID.
func GetKeyshare( keyDid string) (types.KeyShare, error) {
	ksr, err := types.ParseKeyShareDID(keyDid)
	if err != nil {
		return nil, err
	}
	vBiz, err := ksTable.Get(ctx, types.KeysharePrefix(keyDid))
	if err != nil {
		return nil, err
	}
	ks, err := types.NewKeyshare(keyDid, vBiz, ksr.CoinType)
	if err != nil {
		return nil, err
	}
	return ks, nil
}

