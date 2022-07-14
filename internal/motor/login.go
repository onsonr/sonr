package motor

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sonr-io/multi-party-sig/pkg/math/curve"
	"github.com/sonr-io/multi-party-sig/pkg/party"
	"github.com/sonr-io/multi-party-sig/protocols/cmp"
	"github.com/sonr-io/sonr/pkg/client"
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
	shards, err := vault.New().GetVaultShards(request.Did)
	if err != nil {
		return nil, fmt.Errorf("error getting vault shards: %s", err)
	}

	// build recovery Config
	cnfPw := cmp.EmptyConfig(curve.Secp256k1{})
	recShard, err := crypto.AesDecryptWithPassword(request.Password, []byte(shards.RecoveryShard.Value))
	if err != nil {
		return nil, fmt.Errorf("error decrypting recovery shard (%s): %s", recShard, err)
	}
	if err := cnfPw.UnmarshalBinary([]byte(recShard)); err != nil {
		return nil, fmt.Errorf("error unmarshalling recovery shard: %s", err)
	}

	// build DSC Config
	// TODO: get the actual dsc key from keychain
	cnfDsc := cmp.EmptyConfig(curve.Secp256k1{})
	deviceShard, ok := shards.IssuedShards[id]
	if !ok {
		return nil, fmt.Errorf("could not find device shard with key '%s'", id)
	}
	dscShard, err := dscDecrypt([]byte(deviceShard.Value), request.AesDscKey)
	if err != nil {
		return nil, fmt.Errorf("error decrypting DSC shard: %s", err)
	}
	if err := cnfDsc.UnmarshalBinary([]byte(dscShard)); err != nil {
		return nil, fmt.Errorf("error unmarshalling DSC shard: %s", err)
	}

	// generate wallet
	w, err := crypto.GenerateWallet(crypto.WithConfigs(map[party.ID]*cmp.Config{
		"dsc":      cnfDsc,
		"recovery": cnfPw,
	}))
	if err != nil {
		return nil, fmt.Errorf("error generating wallet: %s", err)
	}

	// get address
	bechAddr, err := w.Address()
	if err != nil {
		return nil, fmt.Errorf("error getting bech address: %s", err)
	}

	// get public key
	pk, err := w.PublicKeyProto()
	if err != nil {
		return nil, fmt.Errorf("error getting public key: %s", err)
	}

	// TODO: fetch DID document from chain
	var didDoc did.Document

  
  d, err := did.ParseDID(did.CreateDIDFromAccount(request.Did))
  if err != nil {
    return nil, fmt.Errorf("error parsing did: %s", err)
  }

	return &MotorNode{
		DeviceID:    id,
		Cosmos:      client.NewClient(client.ConnEndpointType_BETA),
		Wallet:      w,
		Address:     bechAddr,
		PubKey:      pk,
		DID:         *d,
		DIDDocument: didDoc,

		// TODO: some state should exist to say if these are encrypted or not
		// (in this case they are, but during create they aren't)
		deviceShard:   []byte(deviceShard.Value),
		sharedShard:   []byte(shards.PskShard.Value),
		recoveryShard: []byte(shards.RecoveryShard.Value),
		unusedShards:  destructureShards(shards.ShardBank),
	}, nil
}

func dscDecrypt(ciphershard, dsc []byte) ([]byte, error) {
	if len(dsc) != 32 {
		return nil, errors.New("dsc must be 32 bytes")
	}
	return crypto.AesDecryptWithKey(dsc, ciphershard)
}

func destructureShards(s []vault.Shard) [][]byte {
	result := make([][]byte, len(s))
	for i, v := range s {
		result[i] = v.Value
	}
	return result
}
