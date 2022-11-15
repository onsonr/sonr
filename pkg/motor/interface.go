package motor

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	cosmostx "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/crypto/mpc"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/host"
	"github.com/sonr-io/sonr/pkg/motor/x/document"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	rt "github.com/sonr-io/sonr/x/registry/types"
)

type MotorNode interface {
	// Account
	GetAddress() string
	GetBalance() int64
	GetClient() *client.Client
	GetWallet() *mpc.Wallet
	GetPubKey() *secp256k1.PubKey
	SendTokens(req mt.PaymentRequest) (*mt.PaymentResponse, error)
	SendTx(routeURL string, msg sdk.Msg) (*cosmostx.BroadcastTxResponse, error)

	// Networking
	Connect(request mt.ConnectRequest) (*mt.ConnectResponse, error)
	GetDeviceID() string
	GetHost() host.SonrHost
	IsHostActive() bool
	OpenLinking(request mt.LinkingRequest) (*mt.LinkingResponse, error)
	PairDevice(request mt.PairingRequest) (*mt.PairingResponse, error)

	// Registry
	AddCredentialVerificationMethod(id string, cred *did.Credential) error
	CreateAccount(request mt.CreateAccountRequest, waitForVault bool) (mt.CreateAccountResponse, error)
	CreateAccountWithKeys(request mt.CreateAccountWithKeysRequest, waitForVault bool) (mt.CreateAccountWithKeysResponse, error)
	OnboardDevice(req mt.OnboardDeviceRequest) (mt.OnboardDeviceResponse, error)
	GetDID() did.DID
	GetDIDDocument() did.Document
	Login(mt.LoginRequest) (mt.LoginResponse, error)
	LoginWithKeys(mt.LoginWithKeysRequest) (mt.LoginResponse, error)
	BuyAlias(rt.MsgBuyAlias) (rt.MsgBuyAliasResponse, error)
	SellAlias(rt.MsgSellAlias) (rt.MsgSellAliasResponse, error)
	TransferAlias(rt.MsgTransferAlias) (rt.MsgTransferAliasResponse, error)

	// Schema
	CreateSchema(mt.CreateSchemaRequest) (mt.CreateSchemaResponse, error)
	NewDocumentBuilder(schemaDid string) (*document.DocumentBuilder, error)
	UploadDocument(req mt.UploadDocumentRequest) (*mt.UploadDocumentResponse, error)

	// Query
	QueryBuckets(req mt.FindBucketConfigRequest) (*mt.FindBucketConfigResponse, error)
	QueryWhoIs(req mt.QueryWhoIsRequest) (*mt.QueryWhoIsResponse, error)
	QueryWhatIs(req mt.QueryWhatIsRequest) (*mt.QueryWhatIsResponse, error)
	QueryWhatIsByCreator(req mt.QueryWhatIsByCreatorRequest) (*mt.QueryWhatIsByCreatorResponse, error)
	QueryWhatIsByDid(did string) (*mt.QueryWhatIsResponse, error)

	// Bucket
	GenerateBucket(req mt.GenerateBucketRequest) (*mt.GenerateBucketResponse, error)
	AddBucketItems(req mt.AddBucketItemsRequest) (*mt.AddBucketItemsResponse, error)
	GetBucketItems(req mt.GetBucketItemsRequest) (*mt.GetBucketItemsResponse, error)
	BurnBucket(req mt.BurnBucketRequest) (*mt.BurnBucketResponse, error)
}
