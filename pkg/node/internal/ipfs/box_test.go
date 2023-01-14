package ipfs

import (
	"context"
	"crypto/rand"
	"os"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/sonr-hq/sonr/pkg/node/config"
)

// TestBoxer is a test function
func TestBoxer(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cnfg, err := TempClientConfig()
	if err != nil {
		t.Fatal(err)
	}
	node, err := Initialize(ctx, cnfg)
	if err != nil {
		t.Fatal(err)
	}
	pubKey, err := TempPubKey()
	if err != nil {
		t.Fatal(err)
	}

	msg := []byte("hello world")
	cid, err := node.AddEncrypted(msg, pubKey)
	if err != nil {
		t.Fatal(err)
	}
	decryptedMsg, err := node.GetDecrypted(cid, pubKey)
	if err != nil {
		t.Fatal(err)
	}
	if string(decryptedMsg) != string(msg) {
		t.Fatal("decrypted message does not match original message")
	}
	t.Logf("decrypted message: %s", decryptedMsg)
}

// TempClientContext returns a temporary client context
func TempClientConfig() (*config.Config, error) {
	path, err := os.MkdirTemp("", "sonr")
	if err != nil {
		return nil, err
	}
	cctx := client.Context{
		HomeDir: path,
	}
	cnfg := config.DefaultConfig()
	err = cnfg.Apply(config.WithClientContext(cctx, true))
	if err != nil {
		return nil, err
	}
	return cnfg, nil
}

// TempPubKey returns a temporary public key
func TempPubKey() ([]byte, error) {
	_, pub, err := crypto.GenerateSecp256k1Key(rand.Reader)
	if err != nil {
		return nil, err
	}
	return pub.Raw()
}
