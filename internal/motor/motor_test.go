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
	aesKey := loadKey()
	if aesKey == nil || len(aesKey) != 32 {
		key, err := crypto.NewAesKey()
		assert.NoError(t, err, "generates aes key")
		aesKey = key

		// store the key
		fmt.Printf("stored key? %v\n", storeKey(key))
	} else {
		fmt.Println("loaded key")
	}

	req, err := json.Marshal(prt.CreateAccountRequest{
		Password:  "password123",
		AesDscKey: aesKey,
	})
	assert.NoError(t, err, "create account request marshals")
	m, _, err := New("test_device")
	assert.NoError(t, err, "creates motor node")

	_, err = m.CreateAccount(req)
	assert.NoError(t, err, "wallet generation succeeds")

	b := m.Balance()
	log.Println("balance:", b)

	// Print the address of the wallet
	log.Println("address:", m.Address)
}

func Test_Login(t *testing.T) {
  aesKey := loadKey()
  if aesKey == nil || len(aesKey) != 32 {
    t.Errorf("could not load key.")
    return
  }

  req, err := json.Marshal(prt.LoginRequest{
    Did: "snr1ary8qqfwrjw8d696qge0chc9azpzyhf2g26j4f",
    Password: "password123",
    AesDscKey: aesKey,
  })
  assert.NoError(t, err, "request marshals")

  m, err := Login("test_device", req)
  assert.NoError(t, err, "login succeeds")

  fmt.Println("logged in")
  fmt.Println("balance: ", m.Balance())
  fmt.Println("address: ", m.Address)
}

func storeKey(aesKey []byte) bool {
	file, err := os.Create("aes.key")
	if err != nil {
		return false
	}
	defer file.Close()

	_, err = file.Write(aesKey)
	return err == nil
}

func loadKey() []byte {
	var file *os.File
	if _, err := os.Stat("aes.key"); os.IsNotExist(err) {
		file, err = os.Create("aes.key")
	} else {
		file, err = os.Open("aes.key")
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
