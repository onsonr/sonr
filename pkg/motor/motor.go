package motor

import (
	"context"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/kataras/golog"
	"github.com/mr-tron/base58"
	"github.com/sonr-io/multi-party-sig/pkg/party"
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/config"
	"github.com/sonr-io/sonr/pkg/crypto/mpc"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
	"github.com/sonr-io/sonr/pkg/host"
	mt "github.com/sonr-io/sonr/pkg/motor/types"
	"github.com/sonr-io/sonr/pkg/motor/x/object"
	st "github.com/sonr-io/sonr/x/schema/types"
	"google.golang.org/grpc"
)

type MotorNode interface {
	GetDeviceID() string

	GetAddress() string
	GetBalance() int64

	GetClient() *client.Client
	GetWallet() *mpc.Wallet
	GetPubKey() *secp256k1.PubKey
	GetDID() did.DID
	GetDIDDocument() did.Document
	GetHost() host.SonrHost
	AddCredentialVerificationMethod(id string, cred *did.Credential) error
	CreateAccount(mt.CreateAccountRequest) (mt.CreateAccountResponse, error)
	Login(mt.LoginRequest) (mt.LoginResponse, error)

	CreateSchema(mt.CreateSchemaRequest) (mt.CreateSchemaResponse, error)
	QueryWhatIs(context.Context, mt.QueryWhatIsRequest) (mt.QueryWhatIsResponse, error)

	NewObjectBuilder(schemaDid string) (*object.ObjectBuilder, error)
}

type motorNodeImpl struct {
	DeviceID    string
	Cosmos      *client.Client
	Wallet      *mpc.Wallet
	Address     string
	PubKey      *secp256k1.PubKey
	DID         did.DID
	DIDDocument did.Document
	SonrHost    host.SonrHost
	Logger      *golog.Logger
	// Sharding
	deviceShard   []byte
	sharedShard   []byte
	recoveryShard []byte
	unusedShards  [][]byte

	// query clients
	schemaQueryClient st.QueryClient

	// resource management
	Resources *motorResources
}

func EmptyMotor(id string, logLevel string) *motorNodeImpl {
	logger := golog.New()
	logger.SetLevel(logLevel)
	return &motorNodeImpl{
		DeviceID: id,
		Logger:   golog.New(),
	}
}

func initMotor(mtr *motorNodeImpl, options ...mpc.WalletOption) (err error) {
	// Create Client instance
	mtr.Logger.Debug("Intializing motor with account options")
	mtr.Cosmos = client.NewClient(client.ConnEndpointType_BETA)
	mtr.Logger.Debug("Creating grpc connection type %d", client.ConnEndpointType_BETA)
	mtr.Logger.Info("starting creating motor resources")

	grpcConn, err := grpc.Dial(
		mtr.Cosmos.GetRPCAddress(),
		grpc.WithInsecure(),
	)
	if err != nil {
		return err
	}
	mtr.Logger.Info("Done creating grpc client")
	mtr.schemaQueryClient = st.NewQueryClient(grpcConn)
	mtr.Resources = newMotorResources(mtr.Cosmos, mtr.Logger, mtr.schemaQueryClient)
	mtr.Logger.Info("Done creating query cleints")
	mtr.Logger.Info("Start generating user wallet")

	// Generate wallet
	mtr.Wallet, err = mpc.GenerateWallet(options...)
	if err != nil {
		return err
	}

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
	mtr.Logger.Info("Done Configuring user wallet")

	// Set Base DID
	baseDid, err := did.ParseDID(fmt.Sprintf("did:snr:%s", strings.TrimPrefix(mtr.Address, "snr")))
	if err != nil {
		return err
	}
	mtr.DID = *baseDid

	mtr.Logger.Info("Start creating host with motor context")
	// It creates a new host.
	mtr.SonrHost, err = host.NewDefaultHost(context.Background(), config.DefaultConfig(config.Role_MOTOR))
	if err != nil {
		return err
	}

	mtr.Logger.Info("Done creating motor resources")

	// Create motorNodeImpl
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
