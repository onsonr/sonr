package types

import (
	"encoding/json"
	fmt "fmt"

	"github.com/onsonr/sonr/x/did/types/orm/keyalgorithm"
	"github.com/onsonr/sonr/x/did/types/orm/keycurve"
	"github.com/onsonr/sonr/x/did/types/orm/keyencoding"
	"github.com/onsonr/sonr/x/did/types/orm/keyrole"
)

// DefaultParams returns default module parameters.
func DefaultParams() Params {
	return Params{
		AllowedPublicKeys:    DefaultKeyInfos(),
		ConveyancePreference: "direct",
		AttestationFormats:   []string{"packed", "android-key", "fido-u2f", "apple"},
	}
}

// DefaultSeedMessage returns the default seed message
func DefaultSeedMessage() string {
	l1 := "The Sonr Network shall make no protocol that respects the establishment of centralized authority,"
	l2 := "or prohibits the free exercise of decentralized identity; or abridges the freedom of data sovereignty,"
	l3 := "or of encrypted communication; or the right of the users to peaceally interact and transact,"
	l4 := "and to petition the Network for the redress of vulnerabilities."
	return fmt.Sprintf("%s %s %s %s", l1, l2, l3, l4)
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

// # Genesis Structures
//
// Equal returns true if two key infos are equal
func (k *KeyInfo) Equal(b *KeyInfo) bool {
	if k == nil && b == nil {
		return true
	}
	return false
}
