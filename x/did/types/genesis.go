package types

import (
	"encoding/json"
	fmt "fmt"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/collections"
	"cosmossdk.io/x/nft"
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
		GlobalIntegrity: DefaultGlobalIntegrity(),
		Params:          DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate
	if gs.GlobalIntegrity == nil {
		return fmt.Errorf("global integrity proof is nil")
	}
	return gs.Params.Validate()
}

// DefaultNFTClasses configures the Initial DIDNamespace NFT classes
func DefaultNFTClasses(nftGenesis *nft.GenesisState) error {
	for _, n := range DIDNamespace_value {
		nftGenesis.Classes = append(nftGenesis.Classes, DIDNamespace(n).GetNFTClass())
	}
	return nil
}

// DefaultParams returns default module parameters.
func DefaultParams() Params {
	return Params{
		WhitelistedAssets:            DefaultAssets(),
		AllowedPublicKeys:            DefaultKeyInfos(),
		LocalhostRegistrationEnabled: true,
		ConveyancePreference:         "direct",
		AttestationFormats:           []string{"packed", "android-key", "fido-u2f", "apple"},
	}
}

// DefaultGlobalIntegrity returns the default global integrity proof
func DefaultGlobalIntegrity() *GlobalIntegrity {
	return &GlobalIntegrity{
		Controller:  "did:sonr:0x0",
		Seed:        DefaultSeedMessage(),
		Accumulator: []byte{},
		Count:       0,
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

// DefaultKeyInfos returns the default key infos: secp256k1, ed25519, keccak256, and bls12381.
func DefaultKeyInfos() map[string]*KeyInfo {
	return map[string]*KeyInfo{
		// Identity Key Info
		// Sonr Controller Key Info - From MPC
		"auth.dwn": {
			Role:      KeyRole_KEY_ROLE_INVOCATION,
			Curve:     KeyCurve_KEY_CURVE_P256,
			Algorithm: KeyAlgorithm_KEY_ALGORITHM_ECDSA,
			Encoding:  KeyEncoding_KEY_ENCODING_HEX,
			Type:      KeyType_KEY_TYPE_MPC,
		},

		// Sonr Vault Shared Key Info - From Registration
		"auth.zk": {
			Role:      KeyRole_KEY_ROLE_ASSERTION,
			Curve:     KeyCurve_KEY_CURVE_BLS12381,
			Algorithm: KeyAlgorithm_KEY_ALGORITHM_UNSPECIFIED,
			Encoding:  KeyEncoding_KEY_ENCODING_MULTIBASE,
			Type:      KeyType_KEY_TYPE_ZK,
		},

		// Blockchain Key Info
		// Ethereum Key Info
		"auth.ethereum": {
			Role:      KeyRole_KEY_ROLE_DELEGATION,
			Curve:     KeyCurve_KEY_CURVE_KECCAK256,
			Algorithm: KeyAlgorithm_KEY_ALGORITHM_ECDSA,
			Encoding:  KeyEncoding_KEY_ENCODING_HEX,
			Type:      KeyType_KEY_TYPE_BIP32,
		},
		// Bitcoin/IBC Key Info
		"auth.bitcoin": {
			Role:      KeyRole_KEY_ROLE_DELEGATION,
			Curve:     KeyCurve_KEY_CURVE_SECP256K1,
			Algorithm: KeyAlgorithm_KEY_ALGORITHM_ECDSA,
			Encoding:  KeyEncoding_KEY_ENCODING_HEX,
			Type:      KeyType_KEY_TYPE_BIP32,
		},

		// Authentication Key Info
		// Browser based WebAuthn
		"webauthn.browser": {
			Role:      KeyRole_KEY_ROLE_AUTHENTICATION,
			Curve:     KeyCurve_KEY_CURVE_P256,
			Algorithm: KeyAlgorithm_KEY_ALGORITHM_ES256,
			Encoding:  KeyEncoding_KEY_ENCODING_RAW,
			Type:      KeyType_KEY_TYPE_WEBAUTHN,
		},
		// FIDO U2F
		"webauthn.fido": {
			Role:      KeyRole_KEY_ROLE_AUTHENTICATION,
			Curve:     KeyCurve_KEY_CURVE_P256,
			Algorithm: KeyAlgorithm_KEY_ALGORITHM_ES256,
			Encoding:  KeyEncoding_KEY_ENCODING_RAW,
			Type:      KeyType_KEY_TYPE_WEBAUTHN,
		},
		// Cross-Platform Passkeys
		"webauthn.passkey": {
			Role:      KeyRole_KEY_ROLE_AUTHENTICATION,
			Curve:     KeyCurve_KEY_CURVE_ED25519,
			Algorithm: KeyAlgorithm_KEY_ALGORITHM_EDDSA,
			Encoding:  KeyEncoding_KEY_ENCODING_RAW,
			Type:      KeyType_KEY_TYPE_WEBAUTHN,
		},
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

// Equal returns true if two validator infos are equal
func (v *ValidatorInfo) Equal(b *ValidatorInfo) bool {
	if v == nil && b == nil {
		return true
	}
	return false
}
