package motor

import (
	"encoding/json"
	"errors"
	"log"

	mtr "github.com/sonr-io/sonr/internal/motor"
	_ "golang.org/x/mobile/bind"
)

var (
	errWalletExists    = errors.New("mpc wallet already exists")
	errWalletNotExists = errors.New("mpc wallet does not exist")
)

var instance *mtr.MotorNode

func Init(buf []byte) ([]byte, error) {
	if buf == nil {
		log.Println("no dsc shard provided")
	}
	n, err := mtr.New()
	if err != nil {
		log.Println("[FATAL] motor:", err)
		return err
	}
	instance = n
	return nil
}

func CreateAccount(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}
	if res, err := instance.CreateAccount(buf); err == nil {
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
	addr, err := instance.Wallet.Address()
	if err != nil {
		return ""
	}
	return addr
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
	buf, err := instance.DIDDoc.MarshalJSON()
	if err != nil {
		return ""
	}
	return string(buf)
}
