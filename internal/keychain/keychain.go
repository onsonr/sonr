package keychain

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/ipfs/boxo/files"
	"github.com/ipfs/kubo/client/rpc"

	modulev1 "github.com/sonrhq/sonr/api/identity/module/v1"
	"github.com/sonrhq/sonr/crypto/core/protocol"
	"github.com/sonrhq/sonr/internal/shares"
	"github.com/sonrhq/sonr/pkg/did"
)

// Keychain is a local temp file system which spawns shares as proto actors
type Keychain struct {
	Wallets   []*modulev1.Account
	Address   string
	PublicKey []byte
	Directory files.Directory
}

// New takes request context and root directory and returns a new Keychain
// 1. It requires an initial credential id to be passed as a value within the accumulator object

func New(ctx context.Context) (*Keychain, error) {
	pub, aliceOut, bobOut, err := shares.Generate()
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
	kc := &Keychain{
		Address:   addr,
		PublicKey: pub,
		Directory: dir,
	}
	return kc, nil
}

func getIpfsClient() *rpc.HttpApi {
	// The `IPFSClient()` function is a method of the `context` struct that returns an instance of the `rpc.HttpApi` type.
	ipfsC, err := rpc.NewLocalApi()
	if err != nil {
		panic(err)
	}

	return ipfsC
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

func readSharesFromDisk(coinType modulev1.CoinType, address string) (*protocol.Message, *protocol.Message, error) {
	pathPrefix := fmt.Sprintf("%s%s", did.GetCoinTypeDIDMethod(coinType), address)
	bobPath := path.Join(".keyshares", fmt.Sprintf("%s.pubshare", pathPrefix))
	alicePath := path.Join(".keyshares", fmt.Sprintf("%s.privshare", pathPrefix))
	bobBz, err := os.ReadFile(bobPath)
	if err != nil {
		return nil, nil, err
	}
	bobOut := &protocol.Message{}
	err = json.Unmarshal(bobBz, bobOut)
	if err != nil {
		return nil, nil, err
	}
	aliceBz, err := os.ReadFile(alicePath)
	if err != nil {
		return nil, nil, err
	}
	aliceOut := &protocol.Message{}
	err = json.Unmarshal(aliceBz, aliceOut)
	if err != nil {
		return nil, nil, err
	}
	return bobOut, aliceOut, nil
}
