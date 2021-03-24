package user

import (
	"crypto/rand"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/pkg/errors"
)

// ^ Get Key: Returns Private key from disk if found ^ //
func (u *User) GetPrivateKey() (crypto.PrivKey, error) {
	// @ Set Path
	privKeyFileName := filepath.Join(u.FS.Root, "snr-peer.privkey")
	var generate bool

	// @ Find Key
	privKeyBytes, err := ioutil.ReadFile(privKeyFileName)
	if os.IsNotExist(err) {
		generate = true
	} else if err != nil {
		return nil, err
	}

	// @ Check for Generate
	if generate {
		// Create New Key
		privKey, _, err := crypto.GenerateEd25519Key(rand.Reader)
		if err != nil {
			return nil, errors.Wrap(err, "generating identity private key")
		}

		// Marshal Data
		privKeyBytes, err := crypto.MarshalPrivateKey(privKey)
		if err != nil {
			return nil, errors.Wrap(err, "marshalling identity private key")
		}

		// Create File
		f, err := os.Create(privKeyFileName)
		if err != nil {
			return nil, errors.Wrap(err, "creating identity private key file")
		}
		defer f.Close()

		// Write Key to file
		if _, err := f.Write(privKeyBytes); err != nil {
			return nil, errors.Wrap(err, "writing identity private key to file")
		}
		return privKey, nil
	}

	privKey, err := crypto.UnmarshalPrivateKey(privKeyBytes)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshalling identity private key")
	}
	return privKey, nil
}
