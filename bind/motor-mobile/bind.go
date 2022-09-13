package motor

import (
	"fmt"

	mtr "github.com/sonr-io/sonr/pkg/motor"
	"github.com/sonr-io/sonr/pkg/motor/x/object"
	ct "github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	rt "github.com/sonr-io/sonr/x/registry/types"
	_ "golang.org/x/mobile/bind"
)

type MotorCallback interface {
	OnDiscover(data []byte)
	OnMotorEvent(msg string, isDone bool)
}

var (
	objectBuilders map[string]*object.ObjectBuilder
	instance       mtr.MotorNode
	callback       ct.MotorCallback
)

func Init(buf []byte, cb MotorCallback) ([]byte, error) {
	// Unmarshal the request
	var req mt.InitializeRequest
	if err := req.Unmarshal(buf); err != nil {
		return nil, err
	}

	// Create Motor instance
	mtr, err := mtr.EmptyMotor(&req, cb)
	if err != nil {
		return nil, err
	}
	instance = mtr
	callback = cb

	// init objectBuilders
	objectBuilders = make(map[string]*object.ObjectBuilder)

	// Return Initialization Response
	resp := mt.InitializeResponse{
		Success: true,
	}

	if req.AuthInfo != nil {
		if res, err := instance.Login(mt.LoginRequest{
			Did:      req.AuthInfo.Did,
			Password: req.AuthInfo.Password,
		}); err == nil {
			return res.Marshal()
		}
	}
	return resp.Marshal()
}

func CreateAccount(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}
	// decode request
	request := mt.CreateAccountRequest{}
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	if res, err := instance.CreateAccount(request); err == nil {
		return res.Marshal()
	} else {
		return nil, err
	}
}

func Login(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	// decode request
	var request mt.LoginRequest
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("error unmarshalling request: %s", err)
	}

	if res, err := instance.Login(request); err == nil {
		return res.Marshal()
	} else {
		return nil, err
	}
}

func Connect() error {
	if instance == nil {
		return ct.ErrMotorWalletNotInitialized
	}
	return instance.Connect()
}

func CreateSchema(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	var request mt.CreateSchemaRequest
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	if res, err := instance.CreateSchema(request); err == nil {
		return res.Marshal()
	} else {
		return nil, err
	}
}

func QuerySchema(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	var request mt.QueryWhatIsRequest
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	res, err := instance.QueryWhatIs(request)
	if err != nil {
		return nil, err
	}
	return res.Marshal()
}

func QuerySchemaByCreator(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	var request mt.QueryWhatIsByCreatorRequest
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	res, err := instance.QueryWhatIsByCreator(request)
	if err != nil {
		return nil, err
	}
	return res.Marshal()
}

func QuerySchemaByDid(did string) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	res, err := instance.QueryWhatIsByDid(did)
	if err != nil {
		return nil, err
	}
	return res.Marshal()
}

func QueryBucket(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	var request mt.QueryWhereIsRequest
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	res, err := instance.QueryWhereIs(request)
	if err != nil {
		return nil, err
	}
	return res.Marshal()
}

func QueryBucketByCreator(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	var request mt.QueryWhereIsByCreatorRequest
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	res, err := instance.QueryWhereIsByCreator(request)
	if err != nil {
		return nil, err
	}
	return res.Marshal()
}

// IssuePayment creates a send/receive token request to the specified address.
func IssuePayment(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	var request mt.PaymentRequest
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	if res, err := instance.SendTokens(request); err == nil {
		return res.Marshal()
	} else {
		return nil, err
	}
}

// Stat returns general information about the Motor node its wallet and accompanying Account.
func Stat() ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	doc := instance.GetDIDDocument()
	if doc == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}
	didDoc, err := rt.NewDIDDocumentFromPkg(doc)
	if err != nil {
		return nil, err
	}

	resp := mt.StatResponse{
		Address:     instance.GetAddress(),
		Balance:     int32(instance.GetBalance()),
		DidDocument: didDoc,
	}
	return resp.Marshal()
}

func BuyAlias(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	var request rt.MsgBuyAlias
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	resp, err := instance.BuyAlias(request)
	if err != nil {
		return nil, err
	}

	return resp.Marshal()
}

func SellAlias(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	var request rt.MsgSellAlias
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	resp, err := instance.SellAlias(request)
	if err != nil {
		return nil, err
	}

	return resp.Marshal()
}

func TransferAlias(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, ct.ErrMotorWalletNotInitialized
	}

	var request rt.MsgTransferAlias
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	resp, err := instance.TransferAlias(request)
	if err != nil {
		return nil, err
	}

	return resp.Marshal()
}
