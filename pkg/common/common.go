package common

import (
	"io/ioutil"

	crypto "github.com/libp2p/go-libp2p/core/crypto"
	tm_crypto "github.com/tendermint/tendermint/crypto"
	tm_json "github.com/tendermint/tendermint/libs/json"
)

func LoadPrivKeyFromJsonPath(path string) (crypto.PrivKey, error) {
	// Load the key from the given path.
	key, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// Create new private key interface
	var vnPk tm_crypto.PrivKey

	// Unmarshal the key into the interface.
	err = tm_json.Unmarshal(key, &vnPk)
	if err != nil {
		return nil, err
	}
	priv, err := crypto.UnmarshalPrivateKey(vnPk.Bytes())
	if err != nil {
		return nil, err
	}
	return priv, nil
}
