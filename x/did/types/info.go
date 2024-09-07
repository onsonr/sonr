package types

import (
	"encoding/hex"

	"github.com/mr-tron/base58/base58"
)

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

// Equal returns true if two chain infos are equal
func (c *ChainInfo) Equal(b *ChainInfo) bool {
	if c == nil && b == nil {
		return true
	}
	return false
}

// Equal returns true if two OpenID config infos are equal
func (o *OpenIDConfig) Equal(b *OpenIDConfig) bool {
	if o == nil && b == nil {
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

// DecodePublicKey extracts the public key from the given data
func (k *KeyInfo) DecodePublicKey(data interface{}) ([]byte, error) {
	var bz []byte
	switch v := data.(type) {
	case string:
		bz = []byte(v)
	case []byte:
		bz = v
	default:
		return nil, ErrUnsupportedKeyEncoding
	}

	if k.Encoding == KeyEncoding_KEY_ENCODING_RAW {
		return bz, nil
	}
	if k.Encoding == KeyEncoding_KEY_ENCODING_HEX {
		return hex.DecodeString(string(bz))
	}
	if k.Encoding == KeyEncoding_KEY_ENCODING_MULTIBASE {
		return base58.Decode(string(bz))
	}
	return nil, ErrUnsupportedKeyEncoding
}

// EncodePublicKey encodes the public key according to the KeyInfo's encoding
func (k *KeyInfo) EncodePublicKey(data []byte) (string, error) {
	if k.Encoding == KeyEncoding_KEY_ENCODING_RAW {
		return string(data), nil
	}
	if k.Encoding == KeyEncoding_KEY_ENCODING_HEX {
		return hex.EncodeToString(data), nil
	}
	if k.Encoding == KeyEncoding_KEY_ENCODING_MULTIBASE {
		return base58.Encode(data), nil
	}
	return "", ErrUnsupportedKeyEncoding
}

// DiscoveryDocument represents the OIDC discovery document.
type DiscoveryDocument struct {
	Issuer                            string   `json:"issuer"`
	AuthorizationEndpoint             string   `json:"authorization_endpoint"`
	TokenEndpoint                     string   `json:"token_endpoint"`
	UserinfoEndpoint                  string   `json:"userinfo_endpoint"`
	JwksURI                           string   `json:"jwks_uri"`
	RegistrationEndpoint              string   `json:"registration_endpoint"`
	ScopesSupported                   []string `json:"scopes_supported"`
	ResponseTypesSupported            []string `json:"response_types_supported"`
	SubjectTypesSupported             []string `json:"subject_types_supported"`
	IDTokenSigningAlgValuesSupported  []string `json:"id_token_signing_alg_values_supported"`
	ClaimsSupported                   []string `json:"claims_supported"`
	GrantTypesSupported               []string `json:"grant_types_supported"`
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported"`
}

var WalletKeyInfo = &KeyInfo{
	Role:      KeyRole_KEY_ROLE_DELEGATION,
	Curve:     KeyCurve_KEY_CURVE_SECP256K1,
	Algorithm: KeyAlgorithm_KEY_ALGORITHM_ECDSA,
	Encoding:  KeyEncoding_KEY_ENCODING_HEX,
	Type:      KeyType_KEY_TYPE_BIP32,
}

var EthKeyInfo = &KeyInfo{
	Role:      KeyRole_KEY_ROLE_DELEGATION,
	Curve:     KeyCurve_KEY_CURVE_KECCAK256,
	Algorithm: KeyAlgorithm_KEY_ALGORITHM_ECDSA,
	Encoding:  KeyEncoding_KEY_ENCODING_HEX,
	Type:      KeyType_KEY_TYPE_BIP32,
}

var SonrKeyInfo = &KeyInfo{
	Role:      KeyRole_KEY_ROLE_INVOCATION,
	Curve:     KeyCurve_KEY_CURVE_P256,
	Algorithm: KeyAlgorithm_KEY_ALGORITHM_ECDSA,
	Encoding:  KeyEncoding_KEY_ENCODING_HEX,
	Type:      KeyType_KEY_TYPE_MPC,
}

var ChainCodeKeyInfos = map[ChainCode]*KeyInfo{
	ChainCodeBTC: WalletKeyInfo,
	ChainCodeETH: EthKeyInfo,
	ChainCodeSNR: SonrKeyInfo,
	ChainCodeIBC: WalletKeyInfo,
}
