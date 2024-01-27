package wallet

import (
	"context"
	"encoding/json"
	"fmt"
	"path"

	"github.com/ipfs/boxo/files"

	modulev1 "github.com/sonrhq/sonr/api/sonr/identity/module/v1"
	"github.com/sonrhq/sonr/internal/kss"
	"github.com/sonrhq/sonr/pkg/crypto/core/protocol"
	"github.com/sonrhq/sonr/pkg/did"
)

// Wallet is a local temp file system which spawns shares as proto actors
type Wallet struct {
	Accounts  []*modulev1.Account
	Address   string
	PublicKey []byte
	Directory files.Directory
}

// New takes request context and root directory and returns a new Keychain
// 1. It requires an initial credential id to be passed as a value within the accumulator object

func New(ctx context.Context) (*Wallet, error) {
	pub, aliceOut, bobOut, err := kss.Generate()
	if err != nil {
		return nil, err
	}
	ct := modulev1.CoinType_COIN_TYPE_SONR
	addr, err := did.GetAddressByPublicKey(pub, ct)
	if err != nil {
		return nil, err
	}
	dir, err := writeSharesToDisk(ct, addr, bobOut, aliceOut)
	if err != nil {
		return nil, err
	}
	kc := &Wallet{
		Address:   addr,
		PublicKey: pub,
		Directory: dir,
	}
	return kc, nil
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
