package types

import (
	"encoding/json"
	fmt "fmt"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/collections"
	"github.com/onsonr/sonr/x/did/types/orm/assettype"
	"github.com/onsonr/sonr/x/did/types/orm/keyalgorithm"
	"github.com/onsonr/sonr/x/did/types/orm/keycurve"
	"github.com/onsonr/sonr/x/did/types/orm/keyencoding"
	"github.com/onsonr/sonr/x/did/types/orm/keyrole"
	"github.com/onsonr/sonr/x/did/types/orm/keytype"
)

// ParamsKey saves the current module params.
var ParamsKey = collections.NewPrefix(0)

const (
	ModuleName = "did"

	StoreKey = ModuleName

	QuerierRoute = ModuleName
)

var ORMModuleSchema = ormv1alpha1.ModuleSchemaDescriptor{
	SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: "did/v1/state.proto"},
	},
	Prefix: []byte{0},
}

// this line is used by starport scaffolding # genesis/types/import

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

// DefaultParams returns default module parameters.
func DefaultParams() Params {
	return Params{
		WhitelistedAssets:    DefaultAssets(),
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

// DefaultAssets returns the default asset infos: BTC, ETH, SNR, and USDC
func DefaultAssets() []*AssetInfo {
	return []*AssetInfo{
		{
			Name:      "Bitcoin",
			Symbol:    "BTC",
			Hrp:       "bc",
			Index:     0,
			AssetType: assettype.Native.String(),
			IconUrl:   "https://cdn.sonr.land/BTC.svg",
		},
		{
			Name:      "Ethereum",
			Symbol:    "ETH",
			Hrp:       "eth",
			Index:     64,
			AssetType: assettype.Native.String(),
			IconUrl:   "https://cdn.sonr.land/ETH.svg",
		},
		{
			Name:      "Sonr",
			Symbol:    "SNR",
			Hrp:       "idx",
			Index:     703,
			AssetType: assettype.Native.String(),
			IconUrl:   "https://cdn.sonr.land/SNR.svg",
		},
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
			Type:      keytype.Mpc.String(),
		},

		// Sonr Vault Shared Key Info - From Registration
		"auth.zk": {
			Role:      keyrole.Assertion.String(),
			Curve:     keycurve.Bls12381.String(),
			Algorithm: keyalgorithm.Es256k.String(),
			Encoding:  keyencoding.Multibase.String(),
			Type:      keytype.Zk.String(),
		},

		// Blockchain Key Info
		// Ethereum Key Info
		"auth.ethereum": {
			Role:      keyrole.Delegation.String(),
			Curve:     keycurve.Keccak256.String(),
			Algorithm: keyalgorithm.Ecdsa.String(),
			Encoding:  keyencoding.Hex.String(),
			Type:      keytype.Bip32.String(),
		},
		// Bitcoin/IBC Key Info
		"auth.bitcoin": {
			Role:      keyrole.Delegation.String(),
			Curve:     keycurve.Secp256k1.String(),
			Algorithm: keyalgorithm.Ecdsa.String(),
			Encoding:  keyencoding.Hex.String(),
			Type:      keytype.Bip32.String(),
		},

		// Authentication Key Info
		// Browser based WebAuthn
		"webauthn.browser": {
			Role:      keyrole.Authentication.String(),
			Curve:     keycurve.P256.String(),
			Algorithm: keyalgorithm.Es256.String(),
			Encoding:  keyencoding.Raw.String(),
			Type:      keytype.Webauthn.String(),
		},
		// FIDO U2F
		"webauthn.fido": {
			Role:      keyrole.Authentication.String(),
			Curve:     keycurve.P256.String(),
			Algorithm: keyalgorithm.Es256.String(),
			Encoding:  keyencoding.Raw.String(),
			Type:      keytype.Webauthn.String(),
		},
		// Cross-Platform Passkeys
		"webauthn.passkey": {
			Role:      keyrole.Authentication.String(),
			Curve:     keycurve.Ed25519.String(),
			Algorithm: keyalgorithm.Eddsa.String(),
			Encoding:  keyencoding.Raw.String(),
			Type:      keytype.Webauthn.String(),
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

//
// # Genesis Structures
//

// Equal returns true if two asset infos are equal
func (a *AssetInfo) Equal(b *AssetInfo) bool {
	if a == nil && b == nil {
		return true
	}
	return false
}

// Equal returns true if two key infos are equal
func (k *KeyInfo) Equal(b *KeyInfo) bool {
	if k == nil && b == nil {
		return true
	}
	return false
}
