package motor

import (
	"fmt"

	kr "github.com/sonr-io/sonr/internal/keyring"
	"github.com/sonr-io/sonr/pkg/crypto/mpc"
	"github.com/sonr-io/sonr/pkg/vault"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
)

func (mtr *motorNodeImpl) OnboardDevice(req mt.OnboardDeviceRequest) (mt.OnboardDeviceResponse, error) {
	// login using creds given during device linking
	_, err := mtr.LoginWithKeys(mt.LoginWithKeysRequest{
		AccountId: req.AccountId,
		Password:  req.Password,
		AesPskKey: req.AesPskKey,
	})
	if err != nil {
		return mt.OnboardDeviceResponse{}, fmt.Errorf("login: %s", err)
	}

	// get a new shard from the vault
	vc := vault.New()
	newShard, err := vc.PopShard(mtr.Address)
	if err != nil {
		return mt.OnboardDeviceResponse{}, fmt.Errorf("pop shard: %s", err)
	}

	// create a DSC to encrypt new shard
	dsc, err := kr.CreateDSC()
	if err != nil {
		return mt.OnboardDeviceResponse{}, fmt.Errorf("create DSC: %s", err)
	}

	cipherShard, err := mpc.AesEncryptWithKey(dsc, newShard)
	if err != nil {
		return mt.OnboardDeviceResponse{}, fmt.Errorf("encrypt shard: %s", err)
	}

	// update the vault
	service, err := vc.IssueShard(
		mtr.Address,
		string(newShard[len(newShard)-3:]),
		mtr.DeviceID,
		string(cipherShard),
	)
	if err != nil {
		return mt.OnboardDeviceResponse{}, fmt.Errorf("issue shard: %s", err)
	}

	// update the WhoIs with new vault
	mtr.DIDDocument.AddService(service)
	r, err := updateWhoIs(mtr)
	if err != nil {
		return mt.OnboardDeviceResponse{}, fmt.Errorf("update WhoIs: %s", err)
	}

	return mt.OnboardDeviceResponse{
		Success: true,
		WhoIs:   r.WhoIs,
	}, nil
}
