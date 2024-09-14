package builder

import (
	"fmt"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	didv1 "github.com/onsonr/sonr/api/did/v1"
	"github.com/onsonr/sonr/x/did/types"

	"github.com/go-webauthn/webauthn/protocol/webauthncose"
)

// PublicKey is an interface for a public key
type PublicKey interface {
	cryptotypes.PubKey
	Clone() cryptotypes.PubKey
	GetRaw() []byte
	GetRole() types.KeyRole
	GetAlgorithm() types.KeyAlgorithm
	GetEncoding() types.KeyEncoding
	GetCurve() types.KeyCurve
	GetKeyType() types.KeyType
}

// CreateAuthnVerification creates a new verification method for an authn method
func CreateAuthnVerification(namespace types.DIDNamespace, issuer string, controller string, pubkey *types.PubKey, identifier string) *types.VerificationMethod {
	return &types.VerificationMethod{
		Method:     namespace,
		Controller: controller,
		PublicKey:  pubkey,
		Id:         identifier,
		Issuer:     issuer,
	}
}

// CreateWalletVerification creates a new verification method for a wallet
func CreateWalletVerification(namespace types.DIDNamespace, controller string, pubkey *types.PubKey, identifier string) *didv1.VerificationMethod {
	return &didv1.VerificationMethod{
		Method:     APIFormatDIDNamespace(namespace),
		Controller: controller,
		PublicKey:  APIFormatPubKey(pubkey),
		Id:         identifier,
	}
}

// ExtractWebAuthnPublicKey parses the raw public key bytes and returns a JWK representation
func ExtractWebAuthnPublicKey(keyBytes []byte) (*types.PubKey_JWK, error) {
	key, err := webauthncose.ParsePublicKey(keyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	switch k := key.(type) {
	case *webauthncose.EC2PublicKeyData:
		return FormatEC2PublicKey(k)
	case *webauthncose.RSAPublicKeyData:
		return FormatRSAPublicKey(k)
	case *webauthncose.OKPPublicKeyData:
		return FormatOKPPublicKey(k)
	default:
		return nil, fmt.Errorf("unsupported key type")
	}
}

// NewInitialWalletAccounts creates a new set of verification methods for a wallet
func NewInitialWalletAccounts(controller string, pubkey *types.PubKey) ([]*didv1.VerificationMethod, error) {
	var verificationMethods []*didv1.VerificationMethod
	for method, chain := range types.InitialChainCodes {
		nk, err := computeBip32AccountPublicKey(pubkey, chain, 0)
		if err != nil {
			return nil, err
		}

		addr, err := chain.FormatAddress(nk)
		if err != nil {
			return nil, nil
		}
		verificationMethods = append(verificationMethods, CreateWalletVerification(method, controller, nk, method.FormatDID(addr)))
	}
	return verificationMethods, nil
}
