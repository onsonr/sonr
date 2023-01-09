package types

import (
	"crypto"
	"crypto/ed25519"
	"encoding/json"
	"errors"
	fmt "fmt"

	"github.com/shengdoushi/base58"
	common "github.com/sonr-hq/sonr/pkg/common"
	"github.com/sonr-hq/sonr/pkg/common/jwx"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

// VerificationMethodOption is used to define options that modify the creation of the verification method
type VerificationMethodOption func(vm *VerificationMethod) error

// WithController sets the controller of a verificationMethod
func WithController(v string) VerificationMethodOption {
	return func(vm *VerificationMethod) error {
		_, err := ParseDID(v)
		if err != nil {
			return err
		}
		vm.Controller = v
		return nil
	}
}

//
// VerificationMethod Creation Functions
//

// NewWebAuthnVM creates a new WebAuthn VerificationMethod
func NewWebAuthnVM(webauthnCredential *common.WebauthnCredential, options ...VerificationMethodOption) (*VerificationMethod, error) {
	// Add Base58 ID to Credential
	id := fmt.Sprintf("did:webauth:%s", base58.Encode(webauthnCredential.Id, base58.BitcoinAlphabet))

	// Configure base Verification MEthod
	vm := &VerificationMethod{
		ID:                 id,
		Type:               KeyType_KeyType_WEB_AUTHN_AUTHENTICATION_2018,
		WebauthnCredential: webauthnCredential,
	}

	// Apply VerificationMethod Options
	for _, opt := range options {
		err := opt(vm)
		if err != nil {
			return nil, err
		}
	}
	return vm, nil
}

// NewVerificationMethod is a convenience method to easily create verificationMethods based on a set of given params.
// It automatically encodes the provided public key based on the keyType.
func NewVerificationMethod(id string, keyType KeyType, controller string, key interface{}) (*VerificationMethod, error) {
	vm := &VerificationMethod{
		ID:         id,
		Type:       keyType,
		Controller: controller,
	}

	// Check for JWK
	if keyType == KeyType_KeyType_JSON_WEB_KEY_2020 {
		// Check if Key is crypto.PublicKey
		if _, ok := key.(crypto.PublicKey); !ok {
			return nil, fmt.Errorf("key is not a crypto.PublicKey")
		}

		// Convert to jwk.Key
		keyAsJWK, err := jwx.New(key).CreateEncJWK()
		if err != nil {
			return nil, err
		}
		// Convert to JSON and back to fix encoding of key material to make sure
		// an unmarshalled and newly created VerificationMethod are equal on object level.
		// The format of PublicKeyJwk in verificationMethod is a map[string]interface{}.
		// We can't use the Key.AsMap since the values of the map will all be internal jwk lib structs.
		// After unmarshalling all the fields will be map[string]string.
		keyAsJSON, err := json.Marshal(keyAsJWK)
		if err != nil {
			return nil, err
		}
		keyAsMap := map[string]string{}
		json.Unmarshal(keyAsJSON, &keyAsMap)

		vm.PublicKeyJwk = keyAsMap
	}

	// Check for Validator/P2P key
	if keyType == KeyType_KeyType_ED25519_VERIFICATION_KEY_2018 {
		// Check if Key is crypto.PublicKey
		if _, ok := key.(crypto.PublicKey); !ok {
			return nil, fmt.Errorf("key is not a crypto.PublicKey")
		}

		ed25519Key, ok := key.(ed25519.PublicKey)
		if !ok {
			return nil, errors.New("wrong key type")
		}
		encodedKey := base58.Encode(ed25519Key, base58.BitcoinAlphabet)
		vm.PublicKeyMultibase = encodedKey
	}

	// Check for Webauthn key
	if keyType == KeyType_KeyType_WEB_AUTHN_AUTHENTICATION_2018 {
		// Check if Key is WebauthnCredential
		if _, ok := key.(*common.WebauthnCredential); !ok {
			return nil, fmt.Errorf("key is not a WebauthnCredential")
		}
		vm.WebauthnCredential = key.(*common.WebauthnCredential)
	}

	// Check for Secp256k1 key
	if keyType == KeyType_KeyType_ECDSA_SECP256K1_VERIFICATION_KEY_2019 {
		// Switch Interface to *secp256k1.PublicKey or string
		switch key.(type) {
		case *secp256k1.PubKey:
			vm.BlockchainAccountId = key.(*secp256k1.PubKey).Address().String()
		case string:
			vm.BlockchainAccountId = key.(string)
		default:
			return nil, fmt.Errorf("key is not a secp256k1.PublicKey or string")
		}
	}
	return vm, nil
}

func (v VerificationRelationship) MarshalJSON() ([]byte, error) {
	if v.Reference != "" {
		return json.Marshal(*v.VerificationMethod)
	} else {
		return json.Marshal(v.Reference)
	}
}

func (v *VerificationRelationship) UnmarshalJSON(b []byte) error {
	// try to figure out if the item is an object of a string
	type Alias VerificationRelationship
	switch b[0] {
	case '{':
		tmp := Alias{VerificationMethod: &VerificationMethod{}}
		err := json.Unmarshal(b, &tmp)
		if err != nil {
			return fmt.Errorf("could not parse verificationRelation method: %w", err)
		}
		*v = (VerificationRelationship)(tmp)
	case '"':
		err := json.Unmarshal(b, &v.Reference)
		if err != nil {
			return fmt.Errorf("could not parse verificationRelation key relation DID: %w", err)
		}
	default:
		return errors.New("verificationRelation is invalid")
	}
	return nil
}

func (v *VerificationMethod) UnmarshalJSON(bytes []byte) error {
	type Alias VerificationMethod
	tmp := Alias{}
	err := json.Unmarshal(bytes, &tmp)
	if err != nil {
		return err
	}
	*v = (VerificationMethod)(tmp)
	return nil
}

func resolveVerificationRelationships(relationships []*VerificationRelationship, methods []*VerificationMethod) error {
	for i, relationship := range relationships {
		if relationship.Reference != "" {
			continue
		}
		if resolved := resolveVerificationRelationship(relationship.Reference, methods); resolved == nil {
			return fmt.Errorf("unable to resolve %s: %s", verificationMethodKey, relationship.Reference)
		} else {
			relationships[i] = resolved
			relationships[i].Reference = relationship.Reference
		}
	}
	return nil
}

func resolveVerificationRelationship(reference string, methods []*VerificationMethod) *VerificationRelationship {
	for _, method := range methods {
		if method.ID == reference {
			return &VerificationRelationship{VerificationMethod: method}
		}
	}
	return nil
}
