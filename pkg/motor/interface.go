package motor

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	SendTx(routeUrl string, msg sdk.Msg) ([]byte, error)

	// Networking
	Connect() error
	GetDeviceID() string
	GetHost() host.SonrHost

	// Registry
	AddCredentialVerificationMethod(id string, cred *did.Credential) error
	CreateAccount(mt.CreateAccountRequest) (mt.CreateAccountResponse, error)
	CreateAccountWithKeys(mt.CreateAccountWithKeysRequest) (mt.CreateAccountWithKeysResponse, error)
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
