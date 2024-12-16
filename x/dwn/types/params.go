package types

import (
	"encoding/json"

	"github.com/onsonr/sonr/internal/models"
	"github.com/onsonr/sonr/internal/models/keyalgorithm"
	"github.com/onsonr/sonr/internal/models/keycurve"
	"github.com/onsonr/sonr/internal/models/keyencoding"
	"github.com/onsonr/sonr/internal/models/keyrole"
)

// DefaultParams returns default module parameters.
func DefaultParams() Params {
	return Params{
		ConveyancePreference: "direct",
		AttestationFormats:   []string{"packed", "android-key", "fido-u2f", "apple"},
		Schema:               DefaultSchema(),
		AllowedOperators: []string{ // TODO:
			"localhost",
			"didao.xyz",
			"sonr.id",
		},
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
		Version:    SchemaVersion,
		Account:    GetSchema(&models.Account{}),
		Asset:      GetSchema(&models.Asset{}),
		Chain:      GetSchema(&models.Chain{}),
		Credential: GetSchema(&models.Credential{}),
		Grant:      GetSchema(&models.Grant{}),
		Keyshare:   GetSchema(&models.Keyshare{}),
		Profile:    GetSchema(&models.Profile{}),
	}
}

func DefaultKeyInfos() map[string]*KeyInfo {
	return map[string]*KeyInfo{
		// Identity Key Info
		// Sonr Controller Key Info - From MPC
		"auth.dwn": {
			Role:      keyrole.Invocation.String(),
			Curve:     keycurve.P256.String(),
			Algorithm: keyalgorithm.Ecdsa.String(),
			Encoding:  keyencoding.Hex.String(),
		},

		// Sonr Vault Shared Key Info - From Registration
		"auth.zk": {
			Role:      keyrole.Assertion.String(),
			Curve:     keycurve.Bls12381.String(),
			Algorithm: keyalgorithm.Es256k.String(),
			Encoding:  keyencoding.Multibase.String(),
		},

		// Blockchain Key Info
		// Ethereum Key Info
		"auth.ethereum": {
			Role:      keyrole.Delegation.String(),
			Curve:     keycurve.Keccak256.String(),
			Algorithm: keyalgorithm.Ecdsa.String(),
			Encoding:  keyencoding.Hex.String(),
		},
		// Bitcoin/IBC Key Info
		"auth.bitcoin": {
			Role:      keyrole.Delegation.String(),
			Curve:     keycurve.Secp256k1.String(),
			Algorithm: keyalgorithm.Ecdsa.String(),
			Encoding:  keyencoding.Hex.String(),
		},

		// Authentication Key Info
		// Browser based WebAuthn
		"webauthn.browser": {
			Role:      keyrole.Authentication.String(),
			Curve:     keycurve.P256.String(),
			Algorithm: keyalgorithm.Es256.String(),
			Encoding:  keyencoding.Raw.String(),
		},
		// FIDO U2F
		"webauthn.fido": {
			Role:      keyrole.Authentication.String(),
			Curve:     keycurve.P256.String(),
			Algorithm: keyalgorithm.Es256.String(),
			Encoding:  keyencoding.Raw.String(),
		},
		// Cross-Platform Passkeys
		"webauthn.passkey": {
			Role:      keyrole.Authentication.String(),
			Curve:     keycurve.Ed25519.String(),
			Algorithm: keyalgorithm.Eddsa.String(),
			Encoding:  keyencoding.Raw.String(),
		},
	}
}

// # Genesis Structures
//
// Equal returns true if two key infos are equal
func (k *KeyInfo) Equal(b *KeyInfo) bool {
	if k == nil && b == nil {
		return true
	}
	return false
}
