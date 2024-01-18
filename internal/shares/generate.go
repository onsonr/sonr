package shares

import (
	"encoding/json"
	"fmt"
	"path"

	"github.com/ipfs/boxo/files"

	modulev1 "github.com/sonrhq/sonr/api/identity/module/v1"
	"github.com/sonrhq/sonr/crypto/core/curves"
	"github.com/sonrhq/sonr/crypto/core/protocol"
	"github.com/sonrhq/sonr/crypto/tecdsa/dklsv1"
	"github.com/sonrhq/sonr/pkg/did"
)

// K_DEFAULT_MPC_CURVE is the default curve for the controller.
var K_DEFAULT_MPC_CURVE = curves.P256()

func Generate(coinType modulev1.CoinType) (files.Directory, []byte, string, error) {
	bob := dklsv1.NewBobDkg(K_DEFAULT_MPC_CURVE, protocol.Version1)
	alice := dklsv1.NewAliceDkg(K_DEFAULT_MPC_CURVE, protocol.Version1)
	err := runIteratedProtocol(bob, alice)
	if err != nil {
		return nil, nil, "", err
	}
	aliceRes, err := alice.Result(protocol.Version1)
	if err != nil {
		return nil, nil, "", err
	}
	bobRes, err := bob.Result(protocol.Version1)
	if err != nil {
		return nil, nil, "", err
	}
	aliceOut, err := dklsv1.DecodeAliceDkgResult(aliceRes)
	if err != nil {
		return nil, nil, "", err
	}
	pub := aliceOut.PublicKey.ToAffineCompressed()
	addr, err := did.GetAddressByPublicKey(pub, coinType)
	if err != nil {
		return nil, nil, "", err
	}
	dir, err := writeSharesToDisk(coinType, addr, bobRes, aliceRes)
	if err != nil {
		return nil, nil, "", err
	}
	return dir, pub, addr, nil
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
