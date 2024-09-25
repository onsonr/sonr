package types

import (
	"encoding/json"
)

// DefaultParams returns default module parameters.
func DefaultParams() Params {
	// TODO:
	return Params{
		IpfsActive: true,
		Schema:     DefaultSchema(),
	}
}

// Stringer method for Params.
func (p Params) String() string {
	bz, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	return string(bz)
}

// Validate does the sanity check on the params.
func (p Params) Validate() error {
	// TODO:
	return nil
}

// DefaultSchema returns the default schema
func DefaultSchema() *Schema {
	return &Schema{
		Version:    1,
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
