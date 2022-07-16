package motor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/sonr-io/sonr/pkg/crypto"
	"github.com/stretchr/testify/assert"
	prt "go.buf.build/grpc/go/sonr-io/motor/api/v1"
)

func Test_CreateAccount(t *testing.T) {
	aesKey := loadKey("aes.key")
	if aesKey == nil || len(aesKey) != 32 {
		key, err := crypto.NewAesKey()
		assert.NoError(t, err, "generates aes key")
		aesKey = key

		// store the key
		fmt.Printf("stored key? %v\n", storeKey("aes.key", key))
	} else {
		fmt.Println("loaded key")
	}

	req, err := json.Marshal(prt.CreateAccountRequest{
		Password:  "password123",
		AesDscKey: aesKey,
	})
	assert.NoError(t, err, "create account request marshals")

	m, psk, err := CreateAccount("test_device", req)
	assert.NoError(t, err, "wallet generation succeeds")

	// write the PSK to local file system for later use
	if err == nil {
		fmt.Printf("stored psk? %v\n", storeKey(fmt.Sprintf("psk%s", m.Address), psk))
	}

	b := m.Balance()
	log.Println("balance:", b)

	// Print the address of the wallet
	log.Println("address:", m.Address)
}

func Test_Login(t *testing.T) {
	did := "snr1k97t92pntjnzyxlplyaxs432tgnf32e7mm5ga8"
	t.Run("with password", func(t *testing.T) {
		pskKey := loadKey(fmt.Sprintf("psk%s", did))
		if pskKey == nil || len(pskKey) != 32 {
			t.Errorf("could not load psk key")
			return
		}

		req, err := json.Marshal(prt.LoginRequest{
			Did:       did,
			Password:  "password123",
			AesPskKey: pskKey,
		})
		assert.NoError(t, err, "request marshals")

		m, err := Login("test_device", req)
		assert.NoError(t, err, "login succeeds")

		if err == nil {
			fmt.Println("balance: ", m.Balance())
			fmt.Println("address: ", m.Address)
		}
	})

	t.Run("with DSC", func(t *testing.T) {
		aesKey := loadKey("aes.key")
		if aesKey == nil || len(aesKey) != 32 {
			t.Errorf("could not load key.")
			return
		}

		pskKey := loadKey(fmt.Sprintf("psk%s", did))
		if pskKey == nil || len(pskKey) != 32 {
			t.Errorf("could not load psk key")
			return
		}

		req, err := json.Marshal(prt.LoginRequest{
			Did:       did,
			AesDscKey: aesKey,
			AesPskKey: pskKey,
		})
		assert.NoError(t, err, "request marshals")

		m, err := Login("test_device", req)
		assert.NoError(t, err, "login succeeds")

		if err == nil {
			fmt.Println("balance: ", m.Balance())
			fmt.Println("address: ", m.Address)
		}
	})
}

func storeKey(name string, aesKey []byte) bool {
	file, err := os.Create(name)
	if err != nil {
		return false
	}
	defer file.Close()

	_, err = file.Write(aesKey)
	return err == nil
}

func loadKey(name string) []byte {
	var file *os.File
	if _, err := os.Stat(name); os.IsNotExist(err) {
		file, err = os.Create(name)
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
