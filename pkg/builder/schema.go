package builder

import "github.com/onsonr/sonr/config/dwn"

type SchemaVersion = int

var CurrentSchemaVersion SchemaVersion = 1

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
