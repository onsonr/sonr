package motor

import (
	"context"
	"fmt"
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
	"github.com/sonr-io/sonr/pkg/logger"
	dp "github.com/sonr-io/sonr/pkg/motor/x/discover"
	"github.com/sonr-io/sonr/pkg/tx"
	"github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
)

type motorNodeImpl struct {
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

	//Logging
	log      *logger.Logger
	logLevel string
}

func EmptyMotor(r *mt.InitializeRequest, cb common.MotorCallback) (*motorNodeImpl, error) {
	if r.GetDeviceId() == "" {
		return nil, fmt.Errorf("DeviceID is required to initialize motor node")
	}

	if r.GetLogLevel() == "" {
		r.LogLevel = "warn"
	}

	return &motorNodeImpl{
		isHostEnabled:      r.GetEnableHost(),
		isDiscoveryEnabled: r.GetEnableDiscovery(),
		DeviceID:           r.GetDeviceId(),
		homeDir:            r.GetHomeDir(),
		supportDir:         r.GetSupportDir(),
		tempDir:            r.GetTempDir(),
		clientMode:         r.GetClientMode(),
		callback:           cb,
		logLevel:           r.LogLevel,
	}, nil
}

func initMotor(mtr *motorNodeImpl, options ...mpc.WalletOption) (err error) {
	mtr.log = logger.New(mtr.logLevel, "motor")
	// Generate wallet
	mtr.log.Info("Generating wallet...")
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

	mtr.log.Info("Connection Endpoints")
	mtr.log.Info("REST: %s\n", mtr.GetClient().GetAPIAddress())
	mtr.log.Info("RPC: %s\n", mtr.GetClient().GetRPCAddress())
	mtr.log.Info("Faucet: %s\n", mtr.GetClient().GetFaucetAddress())
	mtr.log.Info("IPFS: %s\n", mtr.GetClient().GetIPFSAddress())
	mtr.log.Info("✅ Motor Wallet initialized")
	return nil
}

func (mtr *motorNodeImpl) Connect() error {
	if mtr.Wallet == nil {
		return fmt.Errorf("wallet is not initialized")
	}

	if mtr.SonrHost != nil {
		mtr.log.Warn("Host already connected")
		return nil
	}

	var err error
	// Create new host
	if mtr.isHostEnabled {
		mtr.log.Info("Creating host...")
		mtr.SonrHost, err = host.NewDefaultHost(context.Background(), config.DefaultConfig(config.Role_MOTOR, mtr.Address), mtr.callback)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("host is not enabled")
	}

	// Utilize discovery protocol
	if mtr.isDiscoveryEnabled {
		mtr.log.Info("Enabling Discovery...")
		mtr.discovery, err = dp.New(context.Background(), mtr.SonrHost, mtr.callback)
		if err != nil {
			return err
		}
	}
	mtr.log.Info("✅ Motor Host Connected")
	return nil
}

func (m *motorNodeImpl) GetDeviceID() string {
	return m.DeviceID
}

func (m *motorNodeImpl) GetAddress() string {
	return m.Address
}

func (m *motorNodeImpl) GetWallet() *mpc.Wallet {
	return m.Wallet
}

func (m *motorNodeImpl) GetPubKey() *secp256k1.PubKey {
	return m.PubKey
}

func (m *motorNodeImpl) GetDID() did.DID {
	return m.DID
}

func (m *motorNodeImpl) GetDIDDocument() did.Document {
	return m.DIDDocument
}

func (m *motorNodeImpl) GetHost() host.SonrHost {
	return m.SonrHost
}

// Checking the balance of the wallet.
func (m *motorNodeImpl) GetBalance() int64 {
	cs, err := m.Cosmos.CheckBalance(m.Address)
	if err != nil {
		return 0
	}
	if len(cs) <= 0 {
		return 0
	}
	return cs[0].Amount.Int64()
}

func (m *motorNodeImpl) GetClient() *client.Client {
	return m.Cosmos
}

// GetVerificationMethod returns the VerificationMethod for the given party.
func (w *motorNodeImpl) GetVerificationMethod(id party.ID) (*did.VerificationMethod, error) {
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
func (w *motorNodeImpl) AddCredentialVerificationMethod(id string, cred *did.Credential) error {
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

func (w *motorNodeImpl) SendTx(routeUrl string, msg sdk.Msg) ([]byte, error) {
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
