package motor

import (
	"encoding/json"
	"fmt"

	"github.com/sonr-io/multi-party-sig/pkg/math/curve"
	"github.com/sonr-io/multi-party-sig/pkg/party"
	"github.com/sonr-io/multi-party-sig/protocols/cmp"
	"github.com/sonr-io/sonr/pkg/crypto"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/vault"
	rtmv1 "go.buf.build/grpc/go/sonr-io/motor/api/v1"
)

// Login creates a motor node from a LoginRequest
// TODO: calling balance does not seem to work after login
func Login(id string, requestBytes []byte) (*MotorNode, error) {
	// decode request
	var request rtmv1.LoginRequest
	if err := json.Unmarshal(requestBytes, &request); err != nil {
		return nil, fmt.Errorf("error unmarshalling request: %s", err)
	}

	// fetch vault shards
	fmt.Printf("fetching shards from vault... ")
	shards, err := vault.New().GetVaultShards(request.Did)
	if err != nil {
		return nil, fmt.Errorf("error getting vault shards: %s", err)
	}
	fmt.Println("done.")

	fmt.Printf("reconstructing wallet... ")
	cnfgs, err := createWalletConfigs(id, request, shards)
	if err != nil {
		return nil, fmt.Errorf("error creating preferred config: %s", err)
	}

	// generate wallet
	m, err := newMotor(id, crypto.WithConfigs(cnfgs))
	if err != nil {
		return nil, fmt.Errorf("error generating wallet: %s", err)
	}
	fmt.Println("done.")

	// TODO: fetch DID document from chain
	var didDoc did.Document
	m.DIDDocument = didDoc

	// assign shards
	m.deviceShard = shards.IssuedShards[m.DeviceID]
	m.sharedShard = shards.PskShard
	m.recoveryShard = shards.RecoveryShard
	m.unusedShards = destructureShards(shards.ShardBank)

	return m, nil
}

func createWalletConfigs(id string, req rtmv1.LoginRequest, shards vault.Vault) (map[party.ID]*cmp.Config, error) {
	configs := make(map[party.ID]*cmp.Config)

	// if a password is provided, prefer that over the DSC
	if req.Password != "" {
		// build recovery Config
		fmt.Println(req.Password)
		recShard, err := crypto.AesDecryptWithPassword(req.Password, shards.RecoveryShard)
		if err != nil {
			return nil, fmt.Errorf("error decrypting recovery shard (%s): %s", recShard, err)
		}

		configs["recovery"], err = hydrateConfig(recShard)
		if err != nil {
			return nil, fmt.Errorf("recovery shard: %s", err)
		}
	} else {
		// build DSC Config if password is not provided
		deviceShard, ok := shards.IssuedShards[id]
		if !ok {
			return nil, fmt.Errorf("could not find device shard with key '%s'", id)
		}
		dscShard, err := crypto.AesDecryptWithKey(req.AesDscKey, deviceShard)
		if err != nil {
			return nil, fmt.Errorf("error decrypting DSC shard: %s", err)
		}

		configs["dsc"], err = hydrateConfig(dscShard)
		if err != nil {
			return nil, fmt.Errorf("dsc shard: %s", err)
		}
	}

	// in all cases, use the PSK
	pskShard, err := crypto.AesDecryptWithKey(req.AesPskKey, shards.PskShard)
	if err != nil {
		return nil, fmt.Errorf("error decrypting PSK shard: %s", err)
	}

	configs["psk"], err = hydrateConfig(pskShard)
	if err != nil {
		return nil, fmt.Errorf("psk shard: %s", err)
	}

	return configs, nil
}

func hydrateConfig(c []byte) (*cmp.Config, error) {
	cnf := cmp.EmptyConfig(curve.Secp256k1{})
	if err := cnf.UnmarshalBinary(c); err != nil {
		return nil, fmt.Errorf("error unmarshalling shard: %s", err)
	}
	return cnf, nil
}

func destructureShards(s []vault.Shard) [][]byte {
	result := make([][]byte, len(s))
	for i, v := range s {
		result[i] = v
	}
	return result
}
