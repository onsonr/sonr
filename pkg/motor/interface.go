package motor

import (
	"context"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/sonr-io/sonr/internal/bucket"
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/crypto/mpc"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/host"
	"github.com/sonr-io/sonr/pkg/motor/x/object"
	mt "github.com/sonr-io/sonr/third_party/types/motor"
	bt "github.com/sonr-io/sonr/x/bucket/types"
)

type MotorNode interface {
	// Account
	GetAddress() string
	GetBalance() int64
	GetClient() *client.Client
	GetWallet() *mpc.Wallet
	GetPubKey() *secp256k1.PubKey
	SendTokens(req mt.PaymentRequest) (*mt.PaymentResponse, error)

	// Networking
	Connect() error
	GetDeviceID() string
	GetHost() host.SonrHost

	// Registry
	AddCredentialVerificationMethod(id string, cred *did.Credential) error
	CreateAccount(mt.CreateAccountRequest) (mt.CreateAccountResponse, error)
	GetDID() did.DID
	GetDIDDocument() did.Document
	Login(mt.LoginRequest) (mt.LoginResponse, error)

	// Schema
	CreateSchema(mt.CreateSchemaRequest) (mt.CreateSchemaResponse, error)
	NewObjectBuilder(schemaDid string) (*object.ObjectBuilder, error)

	// Buckets
	CreateBucket(context.Context, mt.CreateBucketRequest) (bucket.Bucket, error)
	GetBucket(did string) (bucket.Bucket, error)
	GetBuckets(ctx context.Context) ([]bucket.Bucket, error)
	UpdateBucketItems(ctx context.Context, did string, items []*bt.BucketItem) (bucket.Bucket, error)

	// Query
	QueryBucket(req mt.QueryWhereIsRequest) (*mt.QueryWhereIsResponse, error)
	QueryBucketGroup(req mt.QueryWhereIsByCreatorRequest) (*mt.QueryWhereIsByCreatorResponse, error)
	QueryWhoIs(req mt.QueryWhoIsRequest) (*mt.QueryWhoIsResponse, error)
	QuerySchema(req mt.QueryWhatIsRequest) (*mt.QueryWhatIsResponse, error)
	QuerySchemaByCreator(req mt.QueryWhatIsByCreatorRequest) (*mt.QueryWhatIsByCreatorResponse, error)
}
