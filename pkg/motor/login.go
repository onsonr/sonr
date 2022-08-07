package motor

import (
	"fmt"

	"github.com/sonr-io/multi-party-sig/pkg/math/curve"
	"github.com/sonr-io/multi-party-sig/pkg/party"
	"github.com/sonr-io/multi-party-sig/protocols/cmp"
	"github.com/sonr-io/sonr/pkg/crypto/mpc"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/vault"
	rtmv1 "go.buf.build/grpc/go/sonr-io/motor/api/v1"
)

// Login creates a motor node from a LoginRequest
func (mtr *motorNodeImpl) Login(request rtmv1.LoginRequest) (rtmv1.LoginResponse, error) {
	if request.Did == "" {
		return rtmv1.LoginResponse{}, fmt.Errorf("did must be provided")
	}

	mtr.Address = request.Did

	// fetch vault shards
	fmt.Printf("fetching shards from vault... ")
	shards, err := vault.New().GetVaultShards(request.Did)
	if err != nil {
		return rtmv1.LoginResponse{}, fmt.Errorf("error getting vault shards: %s", err)
	}
	fmt.Println("done.")

	fmt.Printf("reconstructing wallet... ")
	cnfgs, err := createWalletConfigs(mtr.DeviceID, request, shards)
	if err != nil {
		return rtmv1.LoginResponse{}, fmt.Errorf("error creating preferred config: %s", err)
	}

	// generate wallet
	if err = initMotor(mtr, mpc.WithConfigs(cnfgs)); err != nil {
		return rtmv1.LoginResponse{}, fmt.Errorf("error generating wallet: %s", err)
	}
	fmt.Println("done.")

	// fetch DID document from chain
	whoIs, err := mtr.Cosmos.QueryWhoIs(request.Did)
	if err != nil {
		return rtmv1.LoginResponse{}, fmt.Errorf("error fetching whois: %s", err)
	}

	// TODO: this is a hacky workaround for the Id not being populated in the DID document
	whoIs.DidDocument.Id = did.CreateDIDFromAccount(whoIs.Owner)
	mtr.DIDDocument, err = whoIs.DidDocument.ToPkgDoc()
	if err != nil {
		return rtmv1.LoginResponse{}, fmt.Errorf("error getting DID Document: %s", err)
	}

	// assign shards
	mtr.deviceShard = shards.IssuedShards[mtr.DeviceID]
	mtr.sharedShard = shards.PskShard
	mtr.recoveryShard = shards.RecoveryShard
	mtr.unusedShards = destructureShards(shards.ShardBank)

	return rtmv1.LoginResponse{
		Success: true,
	}, nil
}

func createWalletConfigs(id string, req rtmv1.LoginRequest, shards vault.Vault) (map[party.ID]*cmp.Config, error) {
	configs := make(map[party.ID]*cmp.Config)

	// if a password is provided, prefer that over the DSC
	if req.Password != "" {
		// build recovery Config
		recShard, err := mpc.AesDecryptWithPassword(req.Password, shards.RecoveryShard)
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
		dscShard, err := mpc.AesDecryptWithKey(req.AesDscKey, deviceShard)
		if err != nil {
			return nil, fmt.Errorf("error decrypting DSC shard: %s", err)
		}

		configs["dsc"], err = hydrateConfig(dscShard)
		if err != nil {
			return nil, fmt.Errorf("dsc shard: %s", err)
		}
	}

	// in all cases, use the PSK
	pskShard, err := mpc.AesDecryptWithKey(req.AesPskKey, shards.PskShard)
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
