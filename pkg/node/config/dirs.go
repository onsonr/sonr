package config

import (
	"crypto/rand"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/client"
	"golang.org/x/crypto/nacl/box"
)

func kEncPrivKeyPath(cctx client.Context) string {
	return filepath.Join(cctx.HomeDir, ".sonr", "highway", "encryption_key")
}

func kEncPubKeyPath(cctx client.Context) string {
	return filepath.Join(cctx.HomeDir, ".sonr", "highway", "encryption_key.pub")
}

func hasEncryptionKey(cctx client.Context) bool {
	_, err := os.Stat(kEncPrivKeyPath(cctx))
	return err == nil
}

func hasEncryptionPubKey(cctx client.Context) bool {
	_, err := os.Stat(kEncPubKeyPath(cctx))
	return err == nil
}

func hasKeys(cctx client.Context) bool {
	return hasEncryptionKey(cctx) && hasEncryptionPubKey(cctx)
}

func generateBoxKeys(cctx client.Context) error {
	pub, priv, err := box.GenerateKey(rand.Reader)
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Dir(kEncPrivKeyPath(cctx)), 0755)
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Dir(kEncPubKeyPath(cctx)), 0755)
	if err != nil {
		return err
	}
	err = os.WriteFile(kEncPrivKeyPath(cctx), priv[:], 0600)
	if err != nil {
		return err
	}
	err = os.WriteFile(kEncPubKeyPath(cctx), pub[:], 0600)
	if err != nil {
		return err
	}
	return nil
}

func loadBoxKeys(cctx client.Context) (*[32]byte, *[32]byte, error) {
	if !hasKeys(cctx) {
		return nil, nil, fmt.Errorf("no keys found")
	}
	priv, err := os.ReadFile(kEncPrivKeyPath(cctx))
	if err != nil {
		return nil, nil, err
	}
	pub, err := os.ReadFile(kEncPubKeyPath(cctx))
	if err != nil {
		return nil, nil, err
	}
	var privKey [32]byte
	var pubKey [32]byte
	copy(privKey[:], priv)
	copy(pubKey[:], pub)
	return &privKey, &pubKey, nil
}
