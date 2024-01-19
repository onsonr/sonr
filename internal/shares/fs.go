package shares

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/ipfs/boxo/files"

	modulev1 "github.com/sonrhq/sonr/api/identity/module/v1"
	"github.com/sonrhq/sonr/crypto/core/protocol"
	"github.com/sonrhq/sonr/pkg/did"
)

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
