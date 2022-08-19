package motor

import (
	"context"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/sonr-io/sonr/internal/bucket"
	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/crypto/mpc"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/host"
	mt "github.com/sonr-io/sonr/pkg/motor/types"
	"github.com/sonr-io/sonr/pkg/motor/x/object"
	bt "github.com/sonr-io/sonr/x/bucket/types"
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

	CreateBucket(context.Context, mt.CreateBucketRequest) (bucket.Bucket, error)
	QueryWhereIs(ctx context.Context, did string) error
	QueryWhereIsByCreator(ctx context.Context) error

	UpdateBucketItems(ctx context.Context, did string, items []*bt.BucketItem) (bucket.Bucket, error)
	GetBucket(ctx context.Context, did string) (bucket.Bucket, error)
	GetBuckets(ctx context.Context) ([]bucket.Bucket, error)
}
