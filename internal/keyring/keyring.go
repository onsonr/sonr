package keyring

import (
	"errors"
	"fmt"

	kr "github.com/99designs/keyring"
	"github.com/sonr-io/sonr/pkg/crypto/mpc"
)

const (
	SERVICE_NAME_KEY = "SONR_MOTOR_KEYRING"

	DSC_KEY = "DEVICE_SPECIFIC_CREDENTIAL"
	PSK_KEY = "PRESHARED_KEY"
)

func CreateDSC() ([]byte, error) {
	newDsc, err := mpc.NewAesKey()
	if err != nil {
		return nil, err
	}

	ring, err := openKeyring()
	if err != nil {
		return nil, fmt.Errorf("open keyring: %s", err)
	}

	if err := ring.Set(kr.Item{
		Key:  DSC_KEY,
		Data: newDsc,
	}); err != nil {
		return nil, fmt.Errorf("store DSC: %s", err)
	}

	return newDsc, nil
}

func GetDSC() ([]byte, error) {
	ring, err := openKeyring()
	if err != nil {
		return nil, fmt.Errorf("open keyring: %s", err)
	}

	item, err := ring.Get(DSC_KEY)
	if err != nil {
		return nil, fmt.Errorf("get DSC: %s", err)
	}

	if len(item.Data) != 32 {
		return nil, errors.New("DSC not 32 bytes")
	}

	return item.Data, nil
}

func CreatePSK() ([]byte, error) {
	newPsk, err := mpc.NewAesKey()
	if err != nil {
		return nil, err
	}

	ring, err := openKeyring()
	if err != nil {
		return nil, fmt.Errorf("open keyring: %s", err)
	}

	if err := ring.Set(kr.Item{
		Key:  PSK_KEY,
		Data: newPsk,
	}); err != nil {
		return nil, fmt.Errorf("store PSK: %s", err)
	}

	return newPsk, nil
}

func GetPSK() ([]byte, error) {
	ring, err := openKeyring()
	if err != nil {
		return nil, fmt.Errorf("open keyring: %s", err)
	}

	item, err := ring.Get(PSK_KEY)
	if err != nil {
		return nil, fmt.Errorf("get PSK: %s", err)
	}

	if len(item.Data) != 32 {
		return nil, errors.New("DSC not 32 bytes")
	}

	return item.Data, nil
}

func openKeyring() (kr.Keyring, error) {
	return kr.Open(kr.Config{
		ServiceName: SERVICE_NAME_KEY,
	})
}
