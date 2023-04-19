package types

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/protocol/webauthncose"
	"github.com/shengdoushi/base58"
	"github.com/sonrhq/core/internal/crypto"
)

type Credential interface {
	// Controller returns the credential's controller
	Controller() string

	// Get the credential's DID
	Did() string

	// Descriptor returns the credential's descriptor
	Descriptor() protocol.CredentialDescriptor

	// GetWebauthnCredential returns the webauthn credential instance
	GetWebauthnCredential() *crypto.WebauthnCredential

	// Convert the credential to a DID VerificationMethod
	ToVerificationMethod() *VerificationMethod

	// Encrypt is used to encrypt a message for the credential
	Encrypt(msg []byte) ([]byte, error)

	// Decrypt is used to decrypt a message for the credential
	Decrypt(msg []byte) ([]byte, error)

	// Marshal is used to marshal the credential to JSON
	Marshal() ([]byte, error)
}

type didCredential struct {
	*crypto.WebauthnCredential `json:"credential,omitempty"`
	UserDid                    string `json:"controller,omitempty"`
}

func NewCredential(cred *crypto.WebauthnCredential, controller string) Credential {
	return &didCredential{
		WebauthnCredential: cred,
		UserDid:            controller,
	}
}

func LoadJSONCredential(bz []byte) (Credential, error) {
	vm := &VerificationMethod{}
	err := json.Unmarshal(bz, vm)
	if err != nil {
		return nil, err
	}
	return LoadCredential(vm)
}

func LoadCredential(vm *VerificationMethod) (Credential, error) {
	id := strings.Split(vm.Id, ":")
	// Decode the credential id
	credId, err := base64.RawURLEncoding.DecodeString(id[len(id)-1])
	if err != nil {
		return nil, fmt.Errorf("failed to decode credential id: %v", err)
	}
	// Extract the public key from PublicKeyMultibase
	pubKey, err := base58.Decode(vm.PublicKeyMultibase, base58.BitcoinAlphabet)
	if err != nil {
		return nil, fmt.Errorf("failed to decode public key: %v", err)
	}

	// Convert metadata to map and build the WebauthnAuthenticator
	authenticator := &crypto.WebauthnAuthenticator{}
	auth, _, err := webauthnFromMetadata(vm.GetMetadata())
	if err != nil {
		fmt.Println(err)
	} else {
		authenticator = auth
	}

	attestionType := ""
	for _, entry := range vm.GetMetadata() {
		if entry.Key == "attestion_type" {
			attestionType = entry.Value
		}
	}

	// Build the credential
	cred := &crypto.WebauthnCredential{
		Id:              credId,
		PublicKey:       pubKey,
		Authenticator:   authenticator,
		AttestationType: attestionType,
	}
	return NewCredential(cred, vm.Controller), nil
}

func (c *didCredential) Controller() string {
	return c.UserDid
}

// Descriptor returns the credential's descriptor
func (c *didCredential) Descriptor() protocol.CredentialDescriptor {
	return protocol.CredentialDescriptor{
		CredentialID:    c.WebauthnCredential.Id,
		Type:            protocol.PublicKeyCredentialType,
		AttestationType: c.WebauthnCredential.AttestationType,
	}
}

func (c *didCredential) GetWebauthnCredential() *crypto.WebauthnCredential {
	return c.WebauthnCredential
}

// MarshalJSON is used to marshal the credential to JSON
func (c *didCredential) Marshal() ([]byte, error) {
	vm := c.ToVerificationMethod()
	bz, err := json.Marshal(vm)
	if err != nil {
		return nil, err
	}
	return bz, nil
}

// ToVerificationMethod converts the credential to a DID VerificationMethod
func (c *didCredential) ToVerificationMethod() *VerificationMethod {
	did := fmt.Sprintf("did:key:%s", base64.RawURLEncoding.EncodeToString(c.WebauthnCredential.Id))
	pubMb := base58.Encode(c.WebauthnCredential.PublicKey, base58.BitcoinAlphabet)
	return &VerificationMethod{
		Id:                 did,
		Type:               "webauthn/alg-es256",
		PublicKeyMultibase: pubMb,
		Controller:         c.UserDid,
		Metadata:           webauthnToMetadata(c.WebauthnCredential.Authenticator, c.Descriptor()),
	}
}

// Did returns the credential's DID
func (c *didCredential) Did() string {
	return c.ToVerificationMethod().Id
}

// Encrypt is used to encrypt a message for the credential
func (c *didCredential) Encrypt(data []byte) ([]byte, error) {
	// Get the public key from the credential
	keyFace, err := webauthncose.ParsePublicKey(c.WebauthnCredential.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}
	publicKey, ok := keyFace.(webauthncose.EC2PublicKeyData)
	if !ok {
		return nil, errors.New("public key is not an EC2 key")
	}
	// Derive a shared secret using ECDH
	privateKey, err := derivePrivateKey(c.WebauthnCredential)
	if err != nil {
		return nil, fmt.Errorf("failed to derive private key: %w", err)
	}
	sharedSecret, err := sharedSecret(privateKey, publicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to derive shared secret: %w", err)
	}
	// Use the shared secret as the encryption key
	block, err := aes.NewCipher(sharedSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}

	// Generate a random IV
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, fmt.Errorf("failed to generate IV: %w", err)
	}

	// Encrypt the data using AES-GCM
	ciphertext := make([]byte, len(data))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM cipher: %w", err)
	}
	gcm.Seal(ciphertext[:0], iv, data, nil)

	// Encrypt the AES-GCM key using ECIES
	encryptedKey, err := eciesEncrypt(publicKey, sharedSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt key: %w", err)
	}

	// Concatenate the IV and ciphertext into a single byte slice
	result := make([]byte, len(iv)+len(ciphertext)+len(encryptedKey))
	copy(result[:len(iv)], iv)
	copy(result[len(iv):len(iv)+len(ciphertext)], ciphertext)
	copy(result[len(iv)+len(ciphertext):], encryptedKey)

	return result, nil
}

// Decrypt is used to decrypt a message for the credential
func (c *didCredential) Decrypt(data []byte) ([]byte, error) {
	// Get the public key from the credential
	keyFace, err := webauthncose.ParsePublicKey(c.WebauthnCredential.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}
	publicKey, ok := keyFace.(webauthncose.EC2PublicKeyData)
	if !ok {
		return nil, errors.New("public key is not an EC2 key")
	}
	// Derive a shared secret using ECDH
	privateKey, err := derivePrivateKey(c.WebauthnCredential)
	if err != nil {
		return nil, fmt.Errorf("failed to derive private key: %w", err)
	}

	// Derive the shared secret using ECDH and the WebAuthn credential
	sharedSecret, err := sharedSecret(privateKey, publicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to derive shared secret: %w", err)
	}

	// Use the shared secret as the decryption key
	block, err := aes.NewCipher(sharedSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}

	// Split the IV and ciphertext from the encrypted data
	iv := data[:aes.BlockSize]
	ciphertext := data[aes.BlockSize:]

	// Decrypt the ciphertext using AES-GCM
	plaintext := make([]byte, len(ciphertext))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM cipher: %w", err)
	}
	if _, err := gcm.Open(plaintext[:0], iv, ciphertext, nil); err != nil {
		return nil, fmt.Errorf("failed to decrypt data: %w", err)
	}
	return plaintext, nil
}

func webauthnToMetadata(authenticator *crypto.WebauthnAuthenticator, desc protocol.CredentialDescriptor) []*KeyValuePair {
	authenticatorMap := make(map[string]string)
	authenticatorMap["attestion_type"] = desc.AttestationType

	if authenticator == nil {
		return MapToKeyValueList(authenticatorMap)
	}
	aaguid := base64.StdEncoding.EncodeToString(authenticator.Aaguid)
	authenticatorMap["aaguid"] = aaguid
	signCount := strconv.FormatUint(uint64(authenticator.SignCount), 10)
	authenticatorMap["sign_count"] = signCount
	cloneWarning := strconv.FormatBool(authenticator.CloneWarning)
	authenticatorMap["clone_warning"] = cloneWarning
	return MapToKeyValueList(authenticatorMap)
}

func webauthnFromMetadata(metadata []*KeyValuePair) (*crypto.WebauthnAuthenticator, protocol.CredentialDescriptor, error) {
	authenticatorMap := KeyValueListToMap(metadata)
	desc := protocol.CredentialDescriptor{
		AttestationType: authenticatorMap["attestion_type"],
	}

	aaguid, err := base64.StdEncoding.DecodeString(authenticatorMap["aaguid"])
	if err != nil {
		return nil, desc, err
	}
	signCount, err := strconv.ParseUint(authenticatorMap["sign_count"], 10, 32)
	if err != nil {
		return nil, desc, err
	}
	cloneWarning, err := strconv.ParseBool(authenticatorMap["clone_warning"])
	if err != nil {
		return nil, desc, err
	}
	authenticator := &crypto.WebauthnAuthenticator{
		Aaguid:       aaguid,
		SignCount:    uint32(signCount),
		CloneWarning: cloneWarning,
	}

	return authenticator, desc, nil
}

func ValidateWebauthnCredential(credential *crypto.WebauthnCredential, controller string) (Credential, error) {
	// Check for nil credential
	if credential == nil {
		return nil, errors.New("credential is nil")
	}

	// Check for nil credential id
	if credential.Id == nil {
		return nil, errors.New("credential id is nil")
	}

	// Check for nil credential public key
	if credential.PublicKey == nil {
		return nil, errors.New("credential public key is nil")
	}
	return NewCredential(credential, controller), nil
}
