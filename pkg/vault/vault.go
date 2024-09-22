package vault

import (
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/ipfs/boxo/files"
	"github.com/onsonr/crypto/mpc"
	"github.com/onsonr/sonr/config/dwn"
)

type SchemaVersion = int

var CurrentSchemaVersion SchemaVersion = 1

type Vault struct {
	FS    files.Node
	ValKs mpc.Share
}

func New(subject string, origin string, chainID string) (*Vault, error) {
	shares, err := mpc.GenerateKeyshares()
	var (
		valKs = shares[0]
		usrKs = shares[1]
	)
	usrKsJSON, err := usrKs.Marshal()
	if err != nil {
		return nil, err
	}

	sonrAddr, err := bech32.ConvertAndEncode("idx", valKs.GetPublicKey())
	if err != nil {
		return nil, err
	}

	cnfg := NewConfig(usrKsJSON, sonrAddr, chainID, DefaultSchema())

	fileMap, err := NewVaultDirectory(cnfg)
	if err != nil {
		return nil, err
	}

	return &Vault{
		FS:    fileMap,
		ValKs: valKs,
	}, nil
}

func DefaultSchema() *dwn.Schema {
	return &dwn.Schema{
		Version:    CurrentSchemaVersion,
		Account:    "++, id, name, address, publicKey, chainCode, index, controller, createdAt",
		Asset:      "++, id, name, symbol, decimals, chainCode, createdAt",
		Chain:      "++, id, name, networkId, chainCode, createdAt",
		Credential: "++, id, subject, controller, attestationType, origin, label, deviceId, credentialId, publicKey, transport, signCount, userPresent, userVerified, backupEligible, backupState, cloneWarning, createdAt, updatedAt",
		Jwk:        "++, kty, crv, x, y, n, e",
		Grant:      "++, subject, controller, origin, token, scopes, createdAt, updatedAt",
		Keyshare:   "++, id, data, role, createdAt, lastRefreshed",
		PublicKey:  "++, role, algorithm, encoding, curve, key_type, raw, jwk",
		Profile:    "++, id, subject, controller, originUri, publicMetadata, privateMetadata, createdAt, updatedAt",
	}
}
