package motor

import (
	"encoding/json"
	"errors"
	"fmt"

	mtr "github.com/sonr-io/sonr/internal/motor"
	"github.com/sonr-io/sonr/pkg/crypto"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
	_ "golang.org/x/mobile/bind"
)

var (
	errWalletExists    = errors.New("mpc wallet already exists")
	errWalletNotExists = errors.New("mpc wallet does not exist")
)

type motor struct {
	//	node *host.SonrHost
	wallet *crypto.MPCWallet
	// host   host.SonrHost
}

var instance *motor

func CreateAccount(buf []byte) ([]byte, error) {
	if res, err := mtr.CreateAccount(buf); err == nil {
		return json.Marshal(res)
	} else {
		return nil, err
	}
}

// NewWallet creates a new mpc based wallet.
func NewWallet() error {
	if instance != nil {
		return errWalletExists
	}
	w, err := crypto.GenerateWallet()
	if err != nil {
		return err
	}
	instance = &motor{
		wallet: w,
	}
	return nil
}

// LoadWallet unmarshals the given JSON into the wallet.
func LoadWallet(buf []byte) error {
	if instance != nil {
		return errWalletExists
	}
	w := &crypto.MPCWallet{}
	err := w.Unmarshal(buf)
	if err != nil {
		return err
	}
	instance = &motor{
		wallet: w,
	}
	return nil
}

// Address returns the address of the wallet.
func Address() string {
	if instance == nil {
		return ""
	}
	addr, err := instance.wallet.Address()
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
	buf, err := instance.wallet.DIDDocument.MarshalJSON()
	if err != nil {
		return ""
	}
	return string(buf)
}

// ImportCredentials imports the given credentials into the wallet.
func ImportCredential(buf []byte) error {
	if instance == nil {
		return errWalletNotExists
	}
	var cred did.Credential
	err := json.Unmarshal(buf, &cred)
	if err != nil {
		return err
	}
	vmdid, err := did.ParseDID(fmt.Sprintf("%s#%s", instance.wallet.DID, cred.ID))
	if err != nil {
		return err
	}
	vm := &did.VerificationMethod{
		ID:         *vmdid,
		Type:       ssi.ECDSASECP256K1VerificationKey2019,
		Controller: instance.wallet.DID,
		Credential: &cred,
	}
	instance.wallet.DIDDocument.AddAssertionMethod(vm)
	return nil
}

// MarshalWallet returns the JSON representation of the wallet.
func MarshalWallet() []byte {
	if instance == nil {
		return nil
	}
	buf, err := instance.wallet.Marshal()
	if err != nil {
		return nil
	}
	return buf
}
