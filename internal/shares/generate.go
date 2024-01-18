package shares

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	modulev1 "github.com/sonrhq/sonr/api/identity/module/v1"
	"github.com/sonrhq/sonr/crypto/core/curves"
	"github.com/sonrhq/sonr/crypto/core/protocol"
	"github.com/sonrhq/sonr/crypto/tecdsa/dklsv1"
	"github.com/sonrhq/sonr/pkg/did"
)

// K_DEFAULT_MPC_CURVE is the default curve for the controller.
var K_DEFAULT_MPC_CURVE = curves.P256()

func Generate(rootDir string, coinType modulev1.CoinType) ([]byte, string, error) {
	bob := dklsv1.NewBobDkg(K_DEFAULT_MPC_CURVE, protocol.Version1)
	alice := dklsv1.NewAliceDkg(K_DEFAULT_MPC_CURVE, protocol.Version1)
	err := runIteratedProtocol(bob, alice)
	if err != nil {
		return nil, "", err
	}
	aliceRes, err := alice.Result(protocol.Version1)
	if err != nil {
		return nil, "", err
	}
	bobRes, err := bob.Result(protocol.Version1)
	if err != nil {
		return nil, "", err
	}
	aliceOut, err := dklsv1.DecodeAliceDkgResult(aliceRes)
	if err != nil {
		return nil, "", err
	}
	pub := aliceOut.PublicKey.ToAffineCompressed()
	addr, err := did.GetAddressByPublicKey(pub, coinType)
	if err != nil {
		return nil, "", err
	}
	err = os.MkdirAll(path.Join(rootDir, ".keyshares"), os.ModePerm)
	if err != nil {
		return nil, "", err
	}
	err = writeSharesToDisk(rootDir, coinType, addr, bobRes, aliceRes)
	if err != nil {
		return nil, "", err
	}
	return pub, addr, nil
}

func writeSharesToDisk(rootDir string, coinType modulev1.CoinType, address string, bobOut *protocol.Message, aliceOut *protocol.Message) error {
	pathPrefix := fmt.Sprintf("%s%s", did.GetCoinTypeDIDMethod(coinType), address)
	outBz, err := json.Marshal(aliceOut)
	if err != nil {
		return err
	}
	alicePath := path.Join(rootDir, ".keyshares", fmt.Sprintf("%s.privshare", pathPrefix))
	err = os.WriteFile(alicePath, outBz, os.ModePerm)
	if err != nil {
		return err
	}
	fmt.Printf("Keyshare written to disk: %s", alicePath)
	outBz, err = json.Marshal(bobOut)
	if err != nil {
		return err
	}
	bobPath := path.Join(rootDir, ".keyshares", fmt.Sprintf("%s.pubshare", pathPrefix))
	err = os.WriteFile(bobPath, outBz, os.ModePerm)
	if err != nil {
		return err
	}
	fmt.Printf("Keyshare written to disk: %s", bobPath)
	return nil
}
