package wallet

import (
	"errors"

	"github.com/kataras/golog"
)

const (
	device_pub_key  = "did:sonr:device-public-key"
	device_priv_key = "did:sonr:device-private-key"
	mk_file_name    = "sonr_master_key_file"
)

// Error Definitions
var (
	logger             = golog.Default.Child("core/wallet")
	ErrInvalidKeyType  = errors.New("Invalid KeyPair Type provided")
	ErrKeychainUnready = errors.New("Keychain has not been loaded")
	ErrNoPrivateKey    = errors.New("No private key in KeyPair")
	ErrNoPublicKey     = errors.New("No public key in KeyPair")
)

// Option is a function that modifies the wallet options.
type Option func(*options)

// WithPassphrase sets the passphrase of the wallet.
func WithPassphrase(passphrase string) Option {
	return func(o *options) {
		o.passphrase = passphrase
	}
}

// WithSName sets the name of the wallet.
func WithSName(sname string) Option {
	return func(o *options) {
		o.sname = sname
	}
}

// options is a collection of options for the wallet.
type options struct {
	passphrase string
	sname      string
}

// defaultOptions returns the default wallet options.
func defaultOptions() *options {
	return &options{
		passphrase: "wagmi",
		sname:      "test",
	}
}

// createDefaultKeys creates the default keys
func createDefaultKeys(sname string) error {
	// res, err := Instance.GetAll(token, wallet.Credential)
	// if err != nil {
	// 	return err
	// }

	// if len(res) == 0 {
	// 	doc, err := newDeviceDID()
	// 	if err != nil {
	// 		return err
	// 	}

	// 	raw, err := doc.MarshalJSON()
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// convert doc to raw json message
	// 	logger.Infof("Created Device DID: %v", string(raw))
	// 	err = Instance.Add(token, wallet.Credential, raw)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	logger.Infof("Created new Key")
	// }
	return nil
}

// // newDeviceDID returns the device DID
// func newDeviceDID() (*did.Doc, error) {
// 	privKey, pubKey, err := crypto.GenerateEd25519Key(rand.Reader)
// 	if err != nil {
// 		return nil, err
// 	}

// 	pubBuf, err := crypto.MarshalPublicKey(pubKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	privBuf, err := crypto.MarshalPrivateKey(privKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	devid, err := device.ID()
// 	if err != nil {
// 		return nil, err
// 	}

// 	devicePubVerify := did.NewVerificationMethodFromBytes(device_pub_key, pubKey.Type().String(), devid, pubBuf)
// 	devicePrivVerify := did.NewVerificationMethodFromBytes(device_priv_key, privKey.Type().String(), devid, privBuf)
// 	verificationMethod := []did.VerificationMethod{*devicePubVerify, *devicePrivVerify}
// 	didDoc := did.BuildDoc(did.WithVerificationMethod(verificationMethod))
// 	didDoc.ID = fmt.Sprintf("did:sonr:%s", devid)
// 	return didDoc, nil
// }

// // setupWallet sets up the wallet
// func setupWallet(sname, passphrase string, masterLock secretlock.Service) error {
// 	// create a master key reader from this file
// 	// Note: Now that the protected master key file is created, future calls to aries.New()
// 	//       start from the below lines + the masterLock creation above, ie the above code is for master key prep only.
// 	mkReader, err := local.MasterKeyFromPath(device.Wallet.JoinPath(mk_file_name))
// 	if err != nil {
// 		logger.Errorf("Failed to create master key reader: %s", err)
// 		return err
// 	}

// 	// finally create a new instance of local secret lock service using masterLock (key wrapper) and mkReader
// 	secLock, err := local.NewService(mkReader, masterLock)
// 	if err != nil {
// 		logger.Errorf("Failed to create local secret lock service: %s", err.Error())
// 		return err
// 	}

// 	// finally create the framework with custom secret lock service created above
// 	Framework, err = aries.New(aries.WithSecretLock(secLock))
// 	if err != nil {
// 		logger.Errorf("Failed to create framework: %s", err.Error())
// 		return err
// 	}

// 	if Instance == nil {
// 		provider, err := Framework.Context()
// 		if err != nil {
// 			logger.Errorf("Failed to create context: %s", err.Error())
// 			return err
// 		}

// 		if err := wallet.ProfileExists(sname, provider); err != nil {
// 			logger.Warnf("Profile does not exist, creating new...")
// 			// Check if wallet exists
// 			err = wallet.CreateProfile(sname, provider, wallet.WithPassphrase(passphrase))
// 			if err != nil {
// 				logger.Errorf("Failed to create profile: %s", err.Error())
// 				return err
// 			}
// 		} else {
// 			logger.Infof("Profile exists, updating...")
// 			err = wallet.UpdateProfile(sname, provider, wallet.WithPassphrase(passphrase))
// 			if err != nil {
// 				return err
// 			}
// 		}

// 		Instance, err = wallet.New(sname, provider)
// 		if err != nil {
// 			return err
// 		}
// 		token, err = Instance.Open(wallet.WithUnlockByPassphrase(passphrase))
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return createDefaultKeys(sname)
// }
