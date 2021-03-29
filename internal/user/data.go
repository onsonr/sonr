package user

import (
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/pkg/errors"
	md "github.com/sonr-io/core/internal/models"
)

// ^ Get Peer returns Users Contact ^ //
func (u *User) Contact() *md.Contact {
	return u.contact
}

// ^ Get Peer returns Users Current device ^ //
func (u *User) Device() *md.Device {
	return u.device
}

// ^ Get Key: Returns Private key from disk if found ^ //
func (u *User) PrivateKey() (crypto.PrivKey, error) {
	// @ Get Private Key
	if ok := u.FS.IsFile(K_SONR_PRIV_KEY); ok {
		// Get Key File
		buf, err := u.FS.ReadFile(K_SONR_PRIV_KEY)
		if err != nil {
			return nil, err
		}

		// Get Key from Buffer
		key, err := crypto.UnmarshalPrivateKey(buf)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshalling identity private key")
		}

		// Set Key Ref
		return key, nil
	} else {
		// Create New Key
		privKey, _, err := crypto.GenerateKeyPair(crypto.Ed25519, -1)
		if err != nil {
			return nil, errors.Wrap(err, "generating identity private key")
		}

		// Marshal Data
		buf, err := crypto.MarshalPrivateKey(privKey)
		if err != nil {
			return nil, errors.Wrap(err, "marshalling identity private key")
		}

		// Write Key to File
		_, err = u.FS.WriteFile(K_SONR_PRIV_KEY, buf)
		if err != nil {
			return nil, err
		}

		// Set Key Ref
		return privKey, nil
	}

}

// ^ Updates Current Contact Card ^
func (u *User) SetContact(newContact *md.Contact) error {
	// Set Node Contact
	u.contact = newContact

	// Load User
	user, err := u.LoadUser()
	if err != nil {
		return err
	}

	// Save User
	if err := u.SaveUser(user); err != nil {
		return err
	}
	return nil
}
