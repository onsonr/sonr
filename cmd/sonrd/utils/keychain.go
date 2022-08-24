package utils

import (
	"bytes"
	"encoding/gob"

	"github.com/pkg/errors"
	"github.com/sonr-io/keyring"
	"github.com/sonr-io/sonr/pkg/crypto/mpc"
	"github.com/sonr-io/sonr/pkg/fs"
	mt "github.com/sonr-io/sonr/third_party/types/motor"
)

const K_SERVICE_NAME = "sonr-dev"

type UserAuth struct {
	Password  string
	AesDSCKey []byte
	AesPSKKey []byte
}

func (i UserAuth) Validate() bool {
	if len(i.Password) < 12 {
		return false
	}
	if i.AesDSCKey == nil {
		return false
	}
	if len(i.AesDSCKey) < 32 {
		return false
	}
	return true
}

func (i UserAuth) GenAccountCreateRequest() (*mt.CreateAccountRequest, error) {
	if i.Validate() {
		return &mt.CreateAccountRequest{
			Password:  i.Password,
			AesDscKey: i.AesDSCKey,
		}, nil
	}
	return nil, errors.New("Invalid User Auth Object")
}

func NewUserAuth(pwd string) (UserAuth, error) {
	if len(pwd) < 12 {
		return UserAuth{}, errors.New("Password must be atleast 12 characters")
	}

	aesKey, err := mpc.NewAesKey()
	if err != nil {
		return UserAuth{}, err
	}
	return UserAuth{
		Password:  pwd,
		AesDSCKey: aesKey,
	}, nil
}

func (i UserAuth) StoreAuth(addr string, psk []byte) error {
	kc, err := fetchKCService()
	if err != nil {
		return errors.Wrap(err, "Failed to initialize keychain service")
	}
	if !i.Validate() {
		return errors.New("Invalid UserAuth Object")
	}
	i.AesPSKKey = psk

	var b bytes.Buffer
	e := gob.NewEncoder(&b)
	if err := e.Encode(i); err != nil {
		return err
	}

	err = kc.Set(keyring.Item{
		Key:  addr,
		Data: b.Bytes(),
	})
	if err != nil {
		return err
	}
	return nil
}

func GetUserAuth(addr string) (UserAuth, error) {
	kc, err := fetchKCService()
	if err != nil {
		return UserAuth{}, errors.Wrap(err, "Failed to initialize keychain service")
	}
	i, err := kc.Get(addr)
	if err != nil {
		return UserAuth{}, err
	}
	if i.Data == nil || len(i.Data) == 0 {
		return UserAuth{}, errors.New("Keychain Item data is invalid (empty or nil)")
	}
	var ua UserAuth
	b := bytes.NewBuffer(i.Data)
	d := gob.NewDecoder(b)
	if err := d.Decode(&ua); err != nil {
		return UserAuth{}, err
	}
	return ua, nil
}

func fetchKCService() (keyring.Keyring, error) {
	return keyring.Open(keyring.Config{
		ServiceName:              K_SERVICE_NAME,
		KeychainTrustApplication: true,
		AllowedBackends: []keyring.BackendType{
			keyring.KeychainBackend,
		},
	})
}

func getSecureFolderPath() string {
	folder, err := fs.Support.CreateFolder("io.sonr.blockchain")
	if err != nil {
		return ""
	}
	return folder.Path()
}
