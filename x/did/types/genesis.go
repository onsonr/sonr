package types

import (
	"encoding/json"

	"cosmossdk.io/collections"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
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
		WhitelistedAssets:            DefaultAssets(),
		WhitelistedChains:            DefaultChains(),
		AllowedPublicKeys:            DefaultKeyInfos(),
		OpenidConfig:                 DefaultOpenIDConfig(),
		LocalhostRegistrationEnabled: true,
		ConveyancePreference:         "direct",
		AttestationFormats:           []string{"packed", "android-key", "fido-u2f", "apple"},
	}
}

// DefaultAssets returns the default asset infos: BTC, ETH, SNR, and USDC
func DefaultAssets() []*AssetInfo {
	return []*AssetInfo{
		{
			Name:      "Bitcoin",
			Symbol:    "BTC",
			Hrp:       "bc",
			Index:     0,
			AssetType: AssetType_ASSET_TYPE_NATIVE,
			IconUrl:   "https://cdn.sonr.land/BTC.svg",
		},
		{
			Name:      "Ethereum",
			Symbol:    "ETH",
			Hrp:       "eth",
			Index:     64,
			AssetType: AssetType_ASSET_TYPE_NATIVE,
			IconUrl:   "https://cdn.sonr.land/ETH.svg",
		},
		{
			Name:      "Sonr",
			Symbol:    "SNR",
			Hrp:       "idx",
			Index:     703,
			AssetType: AssetType_ASSET_TYPE_NATIVE,
			IconUrl:   "https://cdn.sonr.land/SNR.svg",
		},
	}
}

// DefaultChains returns the default chain infos: Bitcoin, Ethereum, and Sonr.
func DefaultChains() []*ChainInfo {
	return []*ChainInfo{}
}

// DefaultKeyInfos returns the default key infos: secp256k1, ed25519, keccak256, and bls12381.
func DefaultKeyInfos() []*KeyInfo {
	return []*KeyInfo{
		// Identity Key Info
		// Sonr Controller Key Info - From MPC
		{
			Role:      KeyRole_KEY_ROLE_INVOCATION,
			Curve:     KeyCurve_KEY_CURVE_P256,
			Algorithm: KeyAlgorithm_KEY_ALGORITHM_ECDSA,
			Encoding:  KeyEncoding_KEY_ENCODING_HEX,
			Type:      KeyType_KEY_TYPE_MPC,
		},

		// Sonr Vault Shared Key Info - From Registration
		{
			Role:      KeyRole_KEY_ROLE_ASSERTION,
			Curve:     KeyCurve_KEY_CURVE_BLS12381,
			Algorithm: KeyAlgorithm_KEY_ALGORITHM_UNSPECIFIED,
			Encoding:  KeyEncoding_KEY_ENCODING_MULTIBASE,
			Type:      KeyType_KEY_TYPE_ZK,
		},

		// Blockchain Key Info
		// Ethereum Key Info
		{
			Role:      KeyRole_KEY_ROLE_DELEGATION,
			Curve:     KeyCurve_KEY_CURVE_KECCAK256,
			Algorithm: KeyAlgorithm_KEY_ALGORITHM_ECDSA,
			Encoding:  KeyEncoding_KEY_ENCODING_HEX,
			Type:      KeyType_KEY_TYPE_BIP32,
		},
		// Bitcoin/IBC Key Info
		{
			Role:      KeyRole_KEY_ROLE_DELEGATION,
			Curve:     KeyCurve_KEY_CURVE_SECP256K1,
			Algorithm: KeyAlgorithm_KEY_ALGORITHM_ECDSA,
			Encoding:  KeyEncoding_KEY_ENCODING_HEX,
			Type:      KeyType_KEY_TYPE_BIP32,
		},

		// Authentication Key Info
		// Browser based WebAuthn
		{
			Role:      KeyRole_KEY_ROLE_AUTHENTICATION,
			Curve:     KeyCurve_KEY_CURVE_P256,
			Algorithm: KeyAlgorithm_KEY_ALGORITHM_ES256,
			Encoding:  KeyEncoding_KEY_ENCODING_RAW,
			Type:      KeyType_KEY_TYPE_WEBAUTHN,
		},
		// FIDO U2F
		{
			Role:      KeyRole_KEY_ROLE_AUTHENTICATION,
			Curve:     KeyCurve_KEY_CURVE_P256,
			Algorithm: KeyAlgorithm_KEY_ALGORITHM_ES256,
			Encoding:  KeyEncoding_KEY_ENCODING_RAW,
			Type:      KeyType_KEY_TYPE_WEBAUTHN,
		},
		// Cross-Platform Passkeys
		{
			Role:      KeyRole_KEY_ROLE_AUTHENTICATION,
			Curve:     KeyCurve_KEY_CURVE_ED25519,
			Algorithm: KeyAlgorithm_KEY_ALGORITHM_EDDSA,
			Encoding:  KeyEncoding_KEY_ENCODING_RAW,
			Type:      KeyType_KEY_TYPE_WEBAUTHN,
		},
	}
}

func DefaultOpenIDConfig() *OpenIDConfig {
	return &OpenIDConfig{
		Issuer:                 "https://sonr.id",
		AuthorizationEndpoint:  "https://api.sonr.id/auth",
		TokenEndpoint:          "https://api.sonr.id/token",
		UserinfoEndpoint:       "https://api.sonr.id/userinfo",
		ScopesSupported:        []string{"openid", "profile", "email", "web3", "sonr"},
		ResponseTypesSupported: []string{"code"},
		ResponseModesSupported: []string{"query", "form_post"},
		GrantTypesSupported:    []string{"authorization_code", "refresh_token"},
		AcrValuesSupported:     []string{"passkey"},
		SubjectTypesSupported:  []string{"public"},
	}
}

func (p Params) ActiveParams(ipfsActive bool) Params {
	p.IpfsActive = ipfsActive
	return p
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
