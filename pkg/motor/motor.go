package motor

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/mr-tron/base58"
	"github.com/sonr-io/multi-party-sig/pkg/party"
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/config"
	"github.com/sonr-io/sonr/pkg/crypto/mpc"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
	"github.com/sonr-io/sonr/pkg/host"
	dp "github.com/sonr-io/sonr/pkg/motor/x/discover"
	"github.com/sonr-io/sonr/pkg/tx"
	"github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
)

type MotorNodeImpl struct {
	DeviceID    string
	Cosmos      *client.Client
	Wallet      *mpc.Wallet
	Address     string
	PubKey      *secp256k1.PubKey
	DID         did.DID
	DIDDocument did.Document
	SonrHost    host.SonrHost

	// internal protocols
	isHostEnabled      bool
	isDiscoveryEnabled bool
	callback           common.MotorCallback
	discovery          *dp.DiscoverProtocol

	// configuration
	homeDir    string
	supportDir string
	tempDir    string
	clientMode mt.ClientMode

	// Sharding
	deviceShard   []byte
	sharedShard   []byte
	recoveryShard []byte
	unusedShards  [][]byte

	// resource management
	Resources *motorResources
	sh        *shell.Shell
}

func EmptyMotor(r *mt.InitializeRequest, cb common.MotorCallback) (*MotorNodeImpl, error) {
	if r.GetDeviceId() == "" {
		return nil, fmt.Errorf("DeviceID is required to initialize motor node")
	}
	
	return &MotorNodeImpl{
		isHostEnabled:      r.GetEnableHost(),
		isDiscoveryEnabled: r.GetEnableDiscovery(),
		DeviceID:           r.GetDeviceId(),
		homeDir:            r.GetHomeDir(),
		supportDir:         r.GetSupportDir(),
		tempDir:            r.GetTempDir(),
		clientMode:         r.GetClientMode(),
		callback:           cb,
	}, nil
}

func initMotor(mtr *MotorNodeImpl, options ...mpc.WalletOption) (err error) {
	// Create Client instance
	mtr.Cosmos = client.NewClient(mtr.clientMode)
	// Generate wallet
	log.Println("Generating wallet...")
	mtr.Wallet, err = mpc.GenerateWallet(mtr.callback, options...)
	if err != nil {
		return err
	}

	mtr.sh = shell.NewShell(mtr.Cosmos.GetIPFSApiAddress())
	mtr.Resources = newMotorResources(mtr.Cosmos, mtr.sh)

	// Get address
	if mtr.Address == "" {
		mtr.Address, err = mtr.Wallet.Address()
		if err != nil {
			return err
		}
	}

	shell := shell.NewShell(mtr.Cosmos.GetIPFSApiAddress())
	mtr.Resources = newMotorResources(mtr.Cosmos, shell)

	// Get public key
	mtr.PubKey, err = mtr.Wallet.PublicKeyProto()
	if err != nil {
		return err
	}

	// Set Base DID
	baseDid, err := did.ParseDID(fmt.Sprintf("did:snr:%s", strings.TrimPrefix(mtr.Address, "snr")))
	if err != nil {
		return err
	}
	mtr.DID = *baseDid
	log.Println("Wallet set to:", mtr.Address)
	mtr.GetClient().PrintConnectionEndpoints()
	log.Println("✅ Motor Wallet initialized")
	return nil
}

func (mtr *MotorNodeImpl) Connect() error {
	if mtr.Wallet == nil {
		return fmt.Errorf("wallet is not initialized")
	}

	if mtr.SonrHost != nil {
		log.Println("Host already connected")
		return nil
	}

	var err error
	// Create new host
	if mtr.isHostEnabled {
		log.Println("Creating host...")
		mtr.SonrHost, err = host.NewDefaultHost(context.Background(), config.DefaultConfig(config.Role_MOTOR, mtr.Address), mtr.callback)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("host is not enabled")
	}

	// Utilize discovery protocol
	if mtr.isDiscoveryEnabled {
		log.Println("Enabling Discovery...")
		mtr.discovery, err = dp.New(context.Background(), mtr.SonrHost, mtr.callback)
		if err != nil {
			return err
		}
	}
	log.Println("✅ Motor Host Connected")
	return nil
}

func (m *MotorNodeImpl) GetDeviceID() string {
	return m.DeviceID
}

func (m *MotorNodeImpl) GetAddress() string {
	return m.Address
}

func (m *MotorNodeImpl) GetWallet() *mpc.Wallet {
	return m.Wallet
}

func (m *MotorNodeImpl) GetPubKey() *secp256k1.PubKey {
	return m.PubKey
}

func (m *MotorNodeImpl) GetDID() did.DID {
	return m.DID
}

func (m *MotorNodeImpl) GetDIDDocument() did.Document {
	return m.DIDDocument
}

func (m *MotorNodeImpl) GetHost() host.SonrHost {
	return m.SonrHost
}

// Checking the balance of the wallet.
func (m *MotorNodeImpl) GetBalance() int64 {
	cs, err := m.Cosmos.CheckBalance(m.Address)
	if err != nil {
		return 0
	}
	if len(cs) <= 0 {
		return 0
	}
	return cs[0].Amount.Int64()
}

func (m *MotorNodeImpl) GetClient() *client.Client {
	return m.Cosmos
}

// GetVerificationMethod returns the VerificationMethod for the given party.
func (w *MotorNodeImpl) GetVerificationMethod(id party.ID) (*did.VerificationMethod, error) {
	vmdid, err := did.ParseDID(fmt.Sprintf("did:snr:%s#%s", strings.TrimPrefix(w.Address, "snr"), id))
	if err != nil {
		return nil, err
	}

	// Get base58 encoded public key.
	pub, err := w.Wallet.PublicKeyBase58()
	if err != nil {
		return nil, err
	}

	// Return the shares VerificationMethod
	return &did.VerificationMethod{
		ID:              *vmdid,
		Type:            ssi.ECDSASECP256K1VerificationKey2019,
		Controller:      w.DID,
		PublicKeyBase58: pub,
	}, nil
}

/*
Adds a Credential to the DidDocument of the account
*/
func (w *MotorNodeImpl) AddCredentialVerificationMethod(id string, cred *did.Credential) error {
	if w.DIDDocument == nil {
		return fmt.Errorf("cannot create verification method did document not found")
	}

	vmdid, err := did.ParseDID(fmt.Sprintf("did:snr:%s#%s", strings.TrimPrefix(w.Address, "snr"), id))
	if err != nil {
		return err
	}

	enc := base58.Encode(cred.PublicKey)

	// Return the shares VerificationMethod
	vm := &did.VerificationMethod{
		ID:              *vmdid,
		Type:            ssi.ECDSASECP256K1VerificationKey2019,
		Controller:      w.DID,
		PublicKeyBase58: enc,
		Credential:      cred,
	}
	w.DIDDocument.AddAssertionMethod(vm)

	// does not seem to be needed to check on the response if there is no err present.
	_, err = updateWhoIs(w)

	if err != nil {
		return err
	}

	return nil
}

func (w *MotorNodeImpl) SendTx(routeUrl string, msg sdk.Msg) ([]byte, error) {
	cleanMsgRoute := strings.TrimLeft(routeUrl, "/")
	typeUrl := fmt.Sprintf("/sonrio.sonr.%s", cleanMsgRoute)
	txRaw, err := tx.SignTxWithWallet(w.Wallet, typeUrl, msg)
	if err != nil {
		return nil, fmt.Errorf("failed to sign tx (%s) with wallet: %s", typeUrl, err)
	}

	resp, err := w.Cosmos.BroadcastTx(txRaw)
	if err != nil {
		return nil, fmt.Errorf("failed to broadcast tx (%s): %s", typeUrl, err)
	}
	return resp.GetTxResponse().Marshal()
}

func SetupTestAddressWithKeys(motor *MotorNodeImpl) (error) {
	aesKey := loadKey("aes.key")
	if aesKey == nil || len(aesKey) != 32 {
		key, err := mpc.NewAesKey()
		if err != nil {
			return err
		}
		aesKey = key
	}

	psk, err := mpc.NewAesKey()
	if err != nil {
		return err
	}

	req := mt.CreateAccountWithKeysRequest{
		Password:  "password123",
		AesDscKey: aesKey,
		AesPskKey: psk,
	}

	_, err = motor.CreateAccountWithKeys(req)
	if err != nil {
		return err
	}

	storeKey(fmt.Sprintf("psk%s", motor.Address), psk)

	return nil
}

func SetupTestAddress(motor *MotorNodeImpl) (error) {
	req := mt.CreateAccountRequest{
		Password: "password123",
	}
	_, err := motor.CreateAccount(req)
	if err != nil {
		return err
	}

	return nil
}

func loadKey(n string) []byte {
	name := fmt.Sprintf("./test_keys/%s", n)
	var file *os.File
	if _, err := os.Stat(name); os.IsNotExist(err) {
		file, err = os.Create(name)
		if err != nil {
			return nil
		}
	} else if err != nil {
		fmt.Printf("load err: %s\n", err)
	} else {
		file, err = os.Open(name)
		if err != nil {
			return nil
		}
	}
	defer file.Close()

	key, err := ioutil.ReadAll(file)
	if err != nil {
		return nil
	}
	return key
}

func storeKey(n string, aesKey []byte) bool {
	name := fmt.Sprintf("./test_keys/%s", n)
	file, err := os.Create(name)
	if err != nil {
		return false
	}
	defer file.Close()

	_, err = file.Write(aesKey)
	return err == nil
}