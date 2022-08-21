package motor

import (
	"fmt"
	"log"
	"testing"

	"github.com/sonr-io/sonr/pkg/crypto/mpc"
	"github.com/sonr-io/sonr/thirdparty/types/common"
	mt "github.com/sonr-io/sonr/thirdparty/types/motor"
	"github.com/stretchr/testify/assert"
)

func Test_CreateAccount(t *testing.T) {
	aesKey := loadKey("aes.key")
	if aesKey == nil || len(aesKey) != 32 {
		key, err := mpc.NewAesKey()
		assert.NoError(t, err, "generates aes key")
		aesKey = key

		// store the key
		fmt.Printf("stored key? %v\n", storeKey("aes.key", key))
	} else {
		fmt.Println("loaded key")
	}

	req := mt.CreateAccountRequest{
		Password:  "password123",
		AesDscKey: aesKey,
	}

	m := EmptyMotor(&mt.InitializeRequest{
		DeviceId: "test_device",
	}, common.DefaultCallback())
	res, err := m.CreateAccount(req)
	assert.NoError(t, err, "wallet generation succeeds")

	// write the PSK to local file system for later use
	if err == nil {
		fmt.Printf("stored psk? %v\n", storeKey(fmt.Sprintf("psk%s", m.Address), res.AesPsk))
	}

	b := m.GetBalance()
	log.Println("balance:", b)

	// Print the address of the wallet
	log.Println("address:", m.Address)
}

func Test_Login(t *testing.T) {
	t.Run("with password", func(t *testing.T) {
		pskKey := loadKey(fmt.Sprintf("psk%s", ADDR))
		if pskKey == nil || len(pskKey) != 32 {
			t.Errorf("could not load psk key")
			return
		}

		req := mt.LoginRequest{
			Did:       ADDR,
			Password:  "password123",
			AesPskKey: pskKey,
		}

		m := EmptyMotor(&mt.InitializeRequest{
			DeviceId: "test_device",
		}, common.DefaultCallback())
		_, err := m.Login(req)
		assert.NoError(t, err, "login succeeds")

		if err == nil {
			fmt.Println("balance: ", m.GetBalance())
			fmt.Println("address: ", m.Address)
		}
	})

	t.Run("with DSC", func(t *testing.T) {
		aesKey := loadKey("aes.key")
		fmt.Printf("aes: %x\n", aesKey)
		if aesKey == nil || len(aesKey) != 32 {
			t.Errorf("could not load key.")
			return
		}

		pskKey := loadKey(fmt.Sprintf("psk%s", ADDR))
		if pskKey == nil || len(pskKey) != 32 {
			t.Errorf("could not load psk key")
			return
		}

		req := mt.LoginRequest{
			Did:       ADDR,
			AesDscKey: aesKey,
			AesPskKey: pskKey,
		}

		m := EmptyMotor(&mt.InitializeRequest{
			DeviceId: "test_device",
		}, common.DefaultCallback())
		_, err := m.Login(req)
		assert.NoError(t, err, "login succeeds")

		if err == nil {
			fmt.Println("balance: ", m.GetBalance())
			fmt.Println("address: ", m.Address)
		}
	})
}

func Test_LoginAndMakeRequest(t *testing.T) {
	pskKey := loadKey(fmt.Sprintf("psk%s", ADDR))
	if pskKey == nil || len(pskKey) != 32 {
		t.Errorf("could not load psk key")
		return
	}

	req := mt.LoginRequest{
		Did:       ADDR,
		Password:  "password123",
		AesPskKey: pskKey,
	}

	m := EmptyMotor(&mt.InitializeRequest{
		DeviceId: "test_device",
	}, common.DefaultCallback())
	_, err := m.Login(req)
	assert.NoError(t, err, "login succeeds")

	// do something with the logged in account
	m.DIDDocument.AddAlias("gotest.snr")
	_, err = updateWhoIs(m)
	assert.NoError(t, err, "updates successfully")
}
