package wallet

import (
	"crypto/rand"
	"crypto/sha256"
	"time"

	"github.com/google/uuid"
	"github.com/hyperledger/aries-framework-go/pkg/client/vcwallet"
	"github.com/hyperledger/aries-framework-go/pkg/framework/aries"
	"github.com/hyperledger/aries-framework-go/pkg/framework/context"
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
	Provider  *context.Provider
	Framework *aries.Aries
	Instance  *vcwallet.Client
	Info      *WalletInfo

	hasExistingWallet bool
)

func Open(options ...Option) error {
	opts := defaultOptions()
	for _, opt := range options {
		opt(opts)
	}
	if opts.reset {
		err := device.Wallet.Delete(MASTER_KEY_FILE_NAME)
		if err != nil {
			logger.Errorf("Failed to delete master key file: %s", err.Error())
			return err
		}

		err = device.Wallet.Delete(WALLET_INFO_FILE_NAME)
		if err != nil {
			logger.Errorf("Failed to delete wallet info file: %s", err.Error())
			return err
		}
	}

	if device.Wallet.Exists(MASTER_KEY_FILE_NAME) {
		hasExistingWallet = true
		err := loadWallet(opts.passphrase, opts.sname)
		if err != nil {
			logger.Errorf("Failed to load wallet: %s", err.Error())
			return err
		}
	} else {
		hasExistingWallet = false
		err := createWallet(opts.passphrase, opts.sname)
		if err != nil {
			logger.Errorf("Failed to create wallet: %s", err.Error())
			return err
		}
	}
	return nil
}

func loadWallet(passphrase string, sname string) error {
	// create a master lock to protect the master key
	// (salt is optional, if using one, ensure it is stored and passed in for future uses)
	masterLock, err := hkdf.NewMasterLock(passphrase, sha256.New, nil)
	if err != nil {
		logger.Errorf("Failed to create master lock: %s", err.Error())
		return err
	}
	return setupWallet(sname, passphrase, masterLock)
}

func createWallet(passphrase string, sname string) error {
	// keySize to be used to create master key
	keySize := sha256.Size

	// create a master lock to protect the master key
	// (salt is optional, if using one, ensure it is stored and passed in for future uses)
	masterLock, err := hkdf.NewMasterLock(passphrase, sha256.New, nil)
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
	return setupWallet(sname, passphrase, masterLock)
}

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

	Provider, err = Framework.Context()
	if err != nil {
		logger.Errorf("Failed to create context: %s", err.Error())
		return err
	}

	// Check if wallet exists
	err = vcwallet.CreateProfile(sname, Provider, wallet.WithPassphrase(passphrase))
	if err != nil {
		logger.Errorf("Failed to create profile: %s", err.Error())
		return err
	}

	Instance, err = vcwallet.New(sname, Provider, wallet.WithUnlockByPassphrase(passphrase))
	if err != nil {
		logger.Errorf("Failed to create vc wallet: %s", err.Error())
		return err
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
	privKey, err := DevicePrivKH()
	if err != nil {
		logger.Errorf("%s - Failed to get local host's private key", err)
		return nil, err
	}
	result, err := Provider.Crypto().Sign(msg, privKey)
	if err != nil {
		logger.Errorf("%s - Failed to Sign Data", err)
		return nil, err
	}
	return result, nil
}

// VerifyWith verifies a signature with specified pair
func Verify(msg []byte, sig []byte) error {
	privKey, err := DevicePrivKH()
	if err != nil {
		logger.Errorf("%s - Failed to get local host's private key", err)
		return err
	}
	return Provider.Crypto().Verify(msg, sig, privKey)
}
