package motor

import (
	"errors"
	"fmt"

	mtr "github.com/sonr-io/sonr/pkg/motor"
	"github.com/sonr-io/sonr/pkg/motor/x/object"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	rt "github.com/sonr-io/sonr/x/registry/types"
	_ "golang.org/x/mobile/bind"
)

var (
	errWalletNotExists = errors.New("mpc wallet does not exist")
)

var (
	instance mtr.MotorNode
	callback MotorCallback
)

type MotorCallback interface {
	OnDiscover(data []byte)
	OnWalletEvent(msg string, isDone bool)
}

func Init(buf []byte, cb MotorCallback) ([]byte, error) {
	// Unmarshal the request
	var req mt.InitializeRequest
	if err := req.Unmarshal(buf); err != nil {
		return nil, err
	}

	// Check if public key provided

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
			Did:       req.AuthInfo.Did,
			Password:  req.AuthInfo.Password,
			AesDscKey: req.AuthInfo.AesDscKey,
			AesPskKey: req.AuthInfo.AesPskKey,
		}); err == nil {
			return res.Marshal()
		}
	}
	return resp.Marshal()
}

func CreateAccount(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, errWalletNotExists
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
		return nil, errWalletNotExists
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
		return errWalletNotExists
	}
	return instance.Connect()
}

func CreateSchema(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, errWalletNotExists
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
		return nil, errWalletNotExists
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
		return nil, errWalletNotExists
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
		return nil, errWalletNotExists
	}

	res, err := instance.QueryWhatIsByDid(did)
	if err != nil {
		return nil, err
	}
	return res.Marshal()
}

func QueryBucket(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, errWalletNotExists
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
		return nil, errWalletNotExists
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
		return nil, errWalletNotExists
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
		return nil, errWalletNotExists
	}

	doc := instance.GetDIDDocument()
	if doc == nil {
		return nil, errWalletNotExists
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
		return nil, errWalletNotExists
	}
	var request rt.MsgBuyAlias
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}
	return instance.SendTx(request.Route(), &request)
}

func SellAlias(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}
	var request rt.MsgSellAlias
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}
	return instance.SendTx(request.Route(), &request)
}

func TransferAlias(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}
	var request rt.MsgTransferAlias
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}
	return instance.SendTx(request.Route(), &request)
}
