package motor

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	mtr "github.com/sonr-io/sonr/pkg/motor"
	mt "github.com/sonr-io/sonr/pkg/motor/types"
	"github.com/sonr-io/sonr/pkg/motor/x/object"
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
	if err := json.Unmarshal(buf, &req); err != nil {
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
		return json.Marshal(resp)
	}
	return nil, errors.New("loading existing account not implemented")
}

func CreateAccount(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}
	// decode request
	var request mt.CreateAccountRequest
	if err := json.Unmarshal(buf, &request); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	if res, err := instance.CreateAccount(request); err == nil {
		return json.Marshal(res)
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
	if err := json.Unmarshal(buf, &request); err != nil {
		return nil, fmt.Errorf("error unmarshalling request: %s", err)
	}

	if res, err := instance.Login(request); err == nil {
		return json.Marshal(res)
	} else {
		return nil, err
	}
}

func CreateSchema(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}

	var request mt.CreateSchemaRequest
	if err := json.Unmarshal(buf, &request); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	if res, err := instance.CreateSchema(request); err == nil {
		return json.Marshal(res)
	} else {
		return nil, err
	}
}

func QueryWhatIs(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}

	var request mt.QueryWhatIsRequest
	if err := json.Unmarshal(buf, &request); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	if res, err := instance.QueryWhatIs(context.Background(), request); err == nil {
		return json.Marshal(res)
	} else {
		return nil, err
	}
}

// Address returns the address of the wallet.
func Address() string {
	if instance == nil {
		return ""
	}
	wallet := instance.GetWallet()
	if wallet == nil {
		return ""
	}
	addr, err := wallet.Address()
	if err != nil {
		return ""
	}
	return addr
}

// Balance returns the balance of the wallet.
func Balance() int {
	return int(instance.GetBalance())
}

// func Connect() error {
// 	if instance == nil {
// 		return errWalletNotExists
// 	}
// 	h, err := host.NewDefaultHost(context.Background(), config.DefaultConfig(config.Role_MOTOR))
// 	if err != nil {
// 		return err
// 	}
// 	instance.host = h
// 	return nil
// }

// DidDoc returns the DID document as JSON
func DidDoc() string {
	if instance == nil {
		return ""
	}
	doc := instance.GetDIDDocument()
	if doc == nil {
		return ""
	}
	buf, err := doc.MarshalJSON()
	if err != nil {
		return ""
	}
	return string(buf)
}
