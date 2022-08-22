package motor

import (
	"context"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/mr-tron/base58"
	"github.com/sonr-io/multi-party-sig/pkg/party"
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/config"
	"github.com/sonr-io/sonr/pkg/crypto/mpc"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
	"github.com/sonr-io/sonr/pkg/host"
	bt "github.com/sonr-io/sonr/x/bucket/types"
	st "github.com/sonr-io/sonr/x/schema/types"
	"google.golang.org/grpc"
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

	// Sharding
	deviceShard   []byte
	sharedShard   []byte
	recoveryShard []byte
	unusedShards  [][]byte

	// query clients
	schemaQueryClient st.QueryClient
	bucketQueryClient bt.QueryClient

	// resource management
	Resources *motorResources
}

func EmptyMotor(id string) *motorNodeImpl {
	return &motorNodeImpl{
		DeviceID: id,
	}
}

func initMotor(mtr *motorNodeImpl, options ...mpc.WalletOption) (err error) {
	// Create Client instance
	mtr.Cosmos = client.NewClient(client.ConnEndpointType_DEV)

	grpcConn, err := grpc.Dial(
		mtr.Cosmos.GetRPCAddress(),
		grpc.WithInsecure(),
	)
	if err != nil {
		return err
	}

	shell := shell.NewShell(mtr.Cosmos.GetIPFSApiAddress())
	mtr.schemaQueryClient = st.NewQueryClient(grpcConn)
	mtr.bucketQueryClient = bt.NewQueryClient(grpcConn)

	mtr.Resources = newMotorResources(mtr.Cosmos, mtr.bucketQueryClient, mtr.schemaQueryClient, shell)

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

	// Set Base DID
	baseDid, err := did.ParseDID(fmt.Sprintf("did:snr:%s", strings.TrimPrefix(mtr.Address, "snr")))
	if err != nil {
		return err
	}
	mtr.DID = *baseDid

	// It creates a new host.
	mtr.SonrHost, err = host.NewDefaultHost(context.Background(), config.DefaultConfig(config.Role_MOTOR))
	if err != nil {
		return err
	}

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
