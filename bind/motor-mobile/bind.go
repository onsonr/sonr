package motor

import (
	"context"
	"errors"
	"fmt"

	mtr "github.com/sonr-io/sonr/pkg/motor"
	"github.com/sonr-io/sonr/pkg/motor/x/object"
	mt "github.com/sonr-io/sonr/thirdparty/types/motor"
	"github.com/sonr-io/sonr/x/registry/types"
	_ "golang.org/x/mobile/bind"
)

var (
	errWalletExists    = errors.New("mpc wallet already exists")
	errWalletNotExists = errors.New("mpc wallet does not exist")
)

var (
	instance       mtr.MotorNode
	objectBuilders map[string]*object.ObjectBuilder
)

func Init(buf []byte) ([]byte, error) {
	// Unmarshal the request
	var req mt.InitializeRequest
	if err := req.Unmarshal(buf); err != nil {
		return nil, err
	}

	// Check if public key provided
	if req.DeviceKeyprintPub == nil {
		// Create Motor instance
		instance = mtr.EmptyMotor(req.DeviceId)

		// init objectBuilders
		objectBuilders = make(map[string]*object.ObjectBuilder)

		// Return Initialization Response
		resp := mt.InitializeResponse{
			Success: true,
		}
		return resp.Marshal()
	}
	return nil, errors.New("loading existing account not implemented")
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

// QueryWhatIs returns the Document of the specified Schema.
func QueryWhatIs(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}

	var request mt.QueryWhatIsRequest
	if err := request.Unmarshal(buf); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	if res, err := instance.QueryWhatIs(context.Background(), request); err == nil {
		return res.Marshal()
	} else {
		return nil, err
	}
}

// SendTokens sends tokens to the specified address.
func SendTokens(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}

	var request mt.SendTokenRequest
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
	// Check if instance is initialized
	if instance == nil {
		return nil, errWalletNotExists
	}

	// Get Wallet Address
	bal := int32(instance.GetBalance())
	wallet := instance.GetWallet()
	if wallet == nil {
		return nil, errWalletNotExists
	}
	addr, err := wallet.Address()
	if err != nil {
		return nil, err
	}

	// Get Account DID Document
	doc := instance.GetDIDDocument()
	if doc == nil {
		return nil, errWalletNotExists
	}
	diddoc, err := types.NewDIDDocumentFromPkg(doc)
	if err != nil {
		return nil, err
	}

	// Return response
	resp := mt.StatResponse{
		Address:     addr,
		Balance:     bal,
		DidDocument: diddoc,
	}
	return resp.Marshal()
}
