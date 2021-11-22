package wallet

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/hyperledger/aries-framework-go/pkg/framework/aries"
	"github.com/hyperledger/aries-framework-go/pkg/secretlock"
	"github.com/hyperledger/aries-framework-go/pkg/secretlock/local"
	"github.com/hyperledger/aries-framework-go/pkg/secretlock/local/masterlock/hkdf"
	"github.com/hyperledger/aries-framework-go/pkg/wallet"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/device"
)

const (
	MASTER_KEY_FILE_NAME  = "sonr_master_key_file"
	WALLET_INFO_FILE_NAME = "sonr_wallet_info_file"
)

var (
	Framework *aries.Aries
	Instance  *wallet.Wallet
	token     string
)

func New(pp, sname string) error {
	if err := device.Wallet.Delete(MASTER_KEY_FILE_NAME); err != nil {
		logger.Warnf("Couldnt delete master key file: %s", err.Error())
	}

	if err := device.Wallet.Delete(WALLET_INFO_FILE_NAME); err != nil {
		logger.Warnf("Couldnt delete wallet info file: %s", err.Error())
	}

	// keySize to be used to create master key
	keySize := sha256.Size

	// create a master lock to protect the master key
	// (salt is optional, if using one, ensure it is stored and passed in for future uses)
	masterLock, err := hkdf.NewMasterLock(pp, sha256.New, nil)
	if err != nil {
		logger.Errorf("Failed to create master lock: %s", err.Error())
		return err
	}

	// generate a random master key
	masterKeyContent := make([]byte, keySize)
	_, err = rand.Read(masterKeyContent)
	if err != nil {
		logger.Errorf("Failed to create master key: %s", err.Error())
		return err
	}

	// encrypt it
	masterKeyEnc, err := masterLock.Encrypt("", &secretlock.EncryptRequest{
		Plaintext: string(masterKeyContent)})
	if err != nil {
		logger.Errorf("Failed to encrypt master key: %s", err.Error())
		return err
	}

	err = device.Wallet.WriteFile(MASTER_KEY_FILE_NAME, []byte(masterKeyEnc.Ciphertext))
	if err != nil {
		logger.Errorf("Failed to write master key file: %s", err.Error())
		return err
	}
	return setupWallet(sname, pp, masterLock)
}

// Open opens the Sonr Wallet to interact with the blockchain
func Open(options ...Option) error {
	opts := defaultOptions()
	for _, opt := range options {
		opt(opts)
	}
	if device.Wallet.Exists(MASTER_KEY_FILE_NAME) {
		// create a master lock to protect the master key
		// (salt is optional, if using one, ensure it is stored and passed in for future uses)
		masterLock, err := hkdf.NewMasterLock(opts.passphrase, sha256.New, nil)
		if err != nil {
			logger.Errorf("Failed to create master lock: %s", err.Error())
			return err
		}
		return setupWallet(opts.sname, opts.passphrase, masterLock)
	} else {
		return errors.New("Wallet Master Key File does not exist")
	}
}

// setupWallet sets up the wallet
func setupWallet(sname, passphrase string, masterLock secretlock.Service) error {
	// create a master key reader from this file
	// Note: Now that the protected master key file is created, future calls to aries.New()
	//       start from the below lines + the masterLock creation above, ie the above code is for master key prep only.
	mkReader, err := local.MasterKeyFromPath(device.Wallet.JoinPath(MASTER_KEY_FILE_NAME))
	if err != nil {
		logger.Errorf("Failed to create master key reader: %s", err)
		return err
	}

	// finally create a new instance of local secret lock service using masterLock (key wrapper) and mkReader
	secLock, err := local.NewService(mkReader, masterLock)
	if err != nil {
		logger.Errorf("Failed to create local secret lock service: %s", err.Error())
		return err
	}

	// finally create the framework with custom secret lock service created above
	Framework, err = aries.New(aries.WithSecretLock(secLock))
	if err != nil {
		logger.Errorf("Failed to create framework: %s", err.Error())
		return err
	}

	if Instance == nil {
		provider, err := Framework.Context()
		if err != nil {
			logger.Errorf("Failed to create context: %s", err.Error())
			return err
		}

		if err := wallet.ProfileExists(sname, provider); err != nil {
			logger.Warnf("Profile does not exist, creating new...")
			// Check if wallet exists
			err = wallet.CreateProfile(sname, provider, wallet.WithPassphrase(passphrase))
			if err != nil {
				logger.Errorf("Failed to create profile: %s", err.Error())
				return err
			}
		} else {
			logger.Infof("Profile exists, updating...")
			err = wallet.UpdateProfile(sname, provider, wallet.WithPassphrase(passphrase))
			if err != nil {
				return err
			}
		}

		Instance, err = wallet.New(sname, provider)
		if err != nil {
			return err
		}
		token, err = Instance.Open(wallet.WithUnlockByPassphrase(passphrase))
		if err != nil {
			return err
		}
	}
	return createDefaultKeys(sname)
}

// SignedMetadata is a struct to be used for signing metadata.
type SignedMetadata struct {
	Timestamp int64
	PublicKey []byte
	NodeId    string
}

// SignedUUID is a struct to be converted into a UUID.
type SignedUUID struct {
	Timestamp int64
	Signature []byte
	Value     string
}

// CreateUUID makes a new UUID value signed by the local node's private key
func CreateUUID() (*SignedUUID, error) {
	// generate new UUID
	id := uuid.New().String()

	// sign UUID using local node's private key
	sig, err := Sign([]byte(id))
	if err != nil {
		logger.Errorf("%s - Failed to sign UUID", err)
		return nil, err
	}

	// Return UUID with signature
	return &SignedUUID{
		Value:     id,
		Signature: sig,
		Timestamp: time.Now().Unix(),
	}, nil
}

// CreateMetadata makes message data shared between all node's p2p protocols
func CreateMetadata(peerID peer.ID) (*SignedMetadata, error) {
	// Get local node's public key
	pubKey, err := DevicePubKey()
	if err != nil {
		logger.Errorf("%s - Failed to get local host's public key", err)
		return nil, err
	}

	// Marshal Public key into public key data
	nodePubKey, err := crypto.MarshalPublicKey(pubKey)
	if err != nil {
		logger.Errorf("%s - Failed to Extract Public Key", err)
		return nil, err
	}

	// Generate new Metadata
	return &SignedMetadata{
		Timestamp: time.Now().Unix(),
		PublicKey: nodePubKey,
		NodeId:    peer.Encode(peerID),
	}, nil
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
