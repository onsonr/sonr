package wallet

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"path"

	"github.com/ipfs/boxo/files"
	"github.com/tink-crypto/tink-go/v2/keyset"

	modulev1 "github.com/sonrhq/sonr/api/sonr/identity/module/v1"
	"github.com/sonrhq/sonr/crypto/core/protocol"
	"github.com/sonrhq/sonr/crypto/daead"
	"github.com/sonrhq/sonr/internal/wallet/kss"
	"github.com/sonrhq/sonr/pkg/did"
)

var keyHandle *keyset.Handle

func init() {
	kh, err := daead.NewKeyHandle()
	if err != nil {
		panic(err)
	}
	keyHandle = kh
}

// Wallet is a local temp file system which spawns shares as proto actors
type Wallet struct {
	Accounts  []*modulev1.Account
	Address   string
	PublicKey []byte
}

// New takes request context and root directory and returns a new Keychain
// 1. It requires an initial credential id to be passed as a value within the accumulator object
func New(ctx context.Context) (files.Directory, *Wallet, error) {
	pub, aliceOut, bobOut, err := kss.Generate()
	if err != nil {
		return nil, nil, err
	}
	ct := modulev1.CoinType_COIN_TYPE_SONR
	addr, err := did.GetAddressByPublicKey(pub, ct)
	if err != nil {
		return nil, nil, err
	}
	dir, err := writeSharesToDisk(ct, addr, bobOut, aliceOut)
	if err != nil {
		return nil, nil, err
	}
	kc := &Wallet{
		Address:   addr,
		PublicKey: pub,
	}
	return dir, kc, nil
}

// Encrypt takes a keyset handle and uses DAED to encrypt a message
func (kc *Wallet) Encrypt(dir files.Directory, associatedData []byte) (files.Directory, error) {
	it := dir.Entries()

	mapDir := make(map[string]files.Node)
	addFile := func(name string, file files.File) {
		mapDir[name] = file
	}

	for ok := it.Next(); ok; ok = it.Next() {
		// Encrypt each file
		file := files.FileFromEntry(it)
		oldBytes, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}
		newBytes, err := daead.Encrypt(keyHandle, oldBytes, associatedData)
		if err != nil {
			return nil, err
		}
		newFile := files.NewBytesFile(newBytes)
		addFile(it.Name(), newFile)
	}
	return files.NewMapDirectory(mapDir), nil
}

// Decrypt takes a keyset handle and uses DAED to decrypt a message
func (kc *Wallet) Decrypt(dir files.Directory, associatedData []byte) (files.Directory, error) {
	it := dir.Entries()

	mapDir := make(map[string]files.Node)
	addFile := func(name string, file files.File) {
		mapDir[name] = file
	}

	for ok := it.Next(); ok; ok = it.Next() {
		// Decrypt each file
		file := files.FileFromEntry(it)
		oldBytes, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}
		newBytes, err := daead.Decrypt(keyHandle, oldBytes, associatedData)
		if err != nil {
			return nil, err
		}
		newFile := files.NewBytesFile(newBytes)
		addFile(it.Name(), newFile)
	}
	return files.NewMapDirectory(mapDir), nil
}

func writeSharesToDisk(coinType modulev1.CoinType, address string, bobOut *protocol.Message, aliceOut *protocol.Message) (files.Directory, error) {
	pathPrefix := fmt.Sprintf("%s%s", did.GetCoinTypeDIDMethod(coinType), address)
	outBz, err := json.Marshal(aliceOut)
	if err != nil {
		return nil, err
	}
	aliceFile := files.NewBytesFile(outBz)
	alicePath := path.Join(".keyshares", fmt.Sprintf("%s.privshare", pathPrefix))

	outBz, err = json.Marshal(bobOut)
	if err != nil {
		return nil, err
	}
	bobFile := files.NewBytesFile(outBz)
	bobPath := path.Join(".keyshares", fmt.Sprintf("%s.pubshare", pathPrefix))
	dir := files.NewMapDirectory(map[string]files.Node{
		bobPath:   bobFile,
		alicePath: aliceFile,
	})
	return dir, nil
}
