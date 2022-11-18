package motor

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sonr-io/sonr/pkg/crypto/mpc"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
)

type MotorNode interface {
	CreateAccount(request mt.CreateAccountRequest) (mt.CreateAccountResponse, error)
	CreateAccountWithKeys(request mt.CreateAccountWithKeysRequest) (mt.CreateAccountWithKeysResponse, error)
	GetAddress() string
}

func SetupTestAddressWithKeys(m MotorNode) error {
	aesKey := LoadKey("aes.key")
	if aesKey == nil || len(aesKey) != 32 {
		key, err := mpc.NewAesKey()
		if err != nil {
			return err
		}
		aesKey = key
	}

	psk, err := mpc.NewAesKey()
	if err != nil {
		return err
	}

	req := mt.CreateAccountWithKeysRequest{
		Password:  "password123",
		AesDscKey: aesKey,
		AesPskKey: psk,
	}

	_, err = m.CreateAccountWithKeys(req)
	if err != nil {
		return err
	}

	StoreKey(fmt.Sprintf("psk%s", m.GetAddress()), psk)

	return nil
}

func SetupTestAddress(m MotorNode) error {
	req := mt.CreateAccountRequest{
		Password: "password123",
	}
	_, err := m.CreateAccount(req)
	if err != nil {
		return err
	}

	return nil
}

func LoadKey(n string) []byte {
	name := fmt.Sprintf("./test_keys/%s", n)
	var file *os.File
	if _, err := os.Stat(name); os.IsNotExist(err) {
		file, err = os.Create(name)
		if err != nil {
			return nil
		}
	} else if err != nil {
		fmt.Printf("load err: %s\n", err)
	} else {
		file, err = os.Open(name)
		if err != nil {
			return nil
		}
	}
	defer file.Close()

	key, err := ioutil.ReadAll(file)
	if err != nil {
		return nil
	}
	return key
}

func StoreKey(n string, aesKey []byte) bool {
	name := fmt.Sprintf("./test_keys/%s", n)
	file, err := os.Create(name)
	if err != nil {
		return false
	}
	defer file.Close()

	_, err = file.Write(aesKey)
	return err == nil
}
