package motor

import (
	"context"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/internal/bucket"
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/crypto/mpc"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/host"
	"github.com/sonr-io/sonr/pkg/motor/x/object"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	bt "github.com/sonr-io/sonr/x/bucket/types"
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
	NewObjectBuilder(schemaDid string) (*object.ObjectBuilder, error)

	// Buckets

	// Creates a new bucket with the defined properties in the request. Returns and instance of `bucket`. before returning both content and buckets are resolved.
	CreateBucket(context.Context, mt.CreateBucketRequest) (bucket.Bucket, error)

	GetBucket(did string) (bucket.Bucket, error)

	GetBuckets(ctx context.Context) ([]bucket.Bucket, error)

	GetDocument(req mt.GetDocumentRequest) (*mt.GetDocumentResponse, error)
	/*
		Updates a pre existing Bucket's label. before calling update the bucket must already be resolved using `GetBucket`
	*/
	UpdateBucketLabel(context context.Context, did string, label string) (bucket.Bucket, error)

	/*
		Updates a pre existing Bucket's visibility. Before calling update the bucket must already be resolved using `GetBucket`.
		Note that changing a buckets visibility can cause issues for other applications using a previously public bucket.
	*/
	UpdateBucketVisibility(context context.Context, did string, visibility bt.BucketVisibility) (bucket.Bucket, error)

	/*
		Updates a pre existing Bucket's BucketItems. Before calling update the bucket must already be resolved using `GetBucket`.
		Should be used for both adding and removing content. Once a buckets content is updated, content is updated to reflect the updated items.
	*/
	UpdateBucketItems(context context.Context, did string, items []*bt.BucketItem) (bucket.Bucket, error)
	SeachBucketBySchema(req mt.SeachBucketContentBySchemaRequest) (mt.SearchBucketContentBySchemaResponse, error)

	// Query
	QueryWhoIs(req mt.QueryWhoIsRequest) (*mt.QueryWhoIsResponse, error)
	QueryWhatIs(req mt.QueryWhatIsRequest) (*mt.QueryWhatIsResponse, error)
	QueryWhatIsByCreator(req mt.QueryWhatIsByCreatorRequest) (*mt.QueryWhatIsByCreatorResponse, error)
	QueryWhatIsByDid(did string) (*mt.QueryWhatIsResponse, error)
	QueryWhereIs(req mt.QueryWhereIsRequest) (*mt.QueryWhereIsResponse, error)
	QueryWhereIsByCreator(req mt.QueryWhereIsByCreatorRequest) (*mt.QueryWhereIsByCreatorResponse, error)
	QueryObject(cid string) (map[string]interface{}, error)

	UploadDocument(req mt.UploadDocumentRequest) (*mt.UploadDocumentResponse, error)
}
