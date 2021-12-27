package wallet

import (
	"errors"

	"github.com/sonr-io/core/device"
)

var (
	token     string
)

// Exists checks if the wallet exists
func Exists() bool {
	return device.Wallet.Exists(mk_file_name)
}

// New creates a new wallet
func New(pp, sname string) error {
	// if err := device.Wallet.Delete(mk_file_name); err != nil {
	// 	logger.Warnf("Couldnt delete master key file: %s", err.Error())
	// }

	// // keySize to be used to create master key
	// keySize := sha256.Size

	// // create a master lock to protect the master key
	// // (salt is optional, if using one, ensure it is stored and passed in for future uses)
	// masterLock, err := hkdf.NewMasterLock(pp, sha256.New, nil)
	// if err != nil {
	// 	logger.Errorf("Failed to create master lock: %s", err.Error())
	// 	return err
	// }

	// // generate a random master key
	// masterKeyContent := make([]byte, keySize)
	// _, err = rand.Read(masterKeyContent)
	// if err != nil {
	// 	logger.Errorf("Failed to create master key: %s", err.Error())
	// 	return err
	// }

	// // encrypt it
	// masterKeyEnc, err := masterLock.Encrypt("", &secretlock.EncryptRequest{
	// 	Plaintext: string(masterKeyContent)})
	// if err != nil {
	// 	logger.Errorf("Failed to encrypt master key: %s", err.Error())
	// 	return err
	// }

	// err = device.Wallet.WriteFile(mk_file_name, []byte(masterKeyEnc.Ciphertext))
	// if err != nil {
	// 	logger.Errorf("Failed to write master key file: %s", err.Error())
	// 	return err
	// }
	// return setupWallet(sname, pp, masterLock)
	return nil
}

// Open opens the Sonr Wallet to interact with the blockchain
func Open(options ...Option) error {
	opts := defaultOptions()
	for _, opt := range options {
		opt(opts)
	}
	if device.Wallet.Exists(mk_file_name) {
		// create a master lock to protect the master key
		// (salt is optional, if using one, ensure it is stored and passed in for future uses)
		// masterLock, err := hkdf.NewMasterLock(opts.passphrase, sha256.New, nil)
		// if err != nil {
		// 	logger.Errorf("Failed to create master lock: %s", err.Error())
		// 	return err
		// }
		// return setupWallet(opts.sname, opts.passphrase, masterLock)
	} else {

	}
	return errors.New("Wallet Master Key File does not exist")
}

// SignWith signs a message with the specified keypair
func Sign(msg []byte) ([]byte, error) {
	privKey, err := DevicePrivKey()
	if err != nil {
		logger.Errorf("%s - Failed to get local host's private key", err)
		return nil, err
	}
	return privKey.Sign(msg)
}

// VerifyWith verifies a signature with specified pair
func Verify(msg []byte, sig []byte) error {
	privKey, err := DevicePrivKey()
	if err != nil {
		logger.Errorf("%s - Failed to get local host's private key", err)
		return err
	}
	ok, err := privKey.GetPublic().Verify(msg, sig)
	if ok {
		return nil
	}
	return err
}
