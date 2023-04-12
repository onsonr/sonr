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
	"github.com/sonrhq/core/internal/crypto"
)

type Credential interface {
	// Controller returns the credential's controller
	Controller() string

	// Get the credential's DID
	Did() string

	// Descriptor returns the credential's descriptor
	Descriptor() protocol.CredentialDescriptor

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
	Credential *crypto.WebauthnCredential `json:"credential,omitempty"`
	UserDid    string                     `json:"controller,omitempty"`
}

func NewCredential(cred *crypto.WebauthnCredential, controller string) Credential {
	return &didCredential{
		Credential: cred,
		UserDid:    controller,
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
	// Extract the public key from PublicKeyMultibase
	pubKey, err := base64.RawURLEncoding.DecodeString(vm.PublicKeyMultibase)
	if err != nil {
		return nil, fmt.Errorf("failed to decode public key: %v", err)
	}

	// Extract the credential ID from the verification method ID
	id := strings.Split(vm.Id, "#")[1]
	credID, err := base64.RawURLEncoding.DecodeString(id)
	if err != nil {
		return nil, fmt.Errorf("failed to decode credential ID: %v", err)
	}

	// Convert metadata to map and build the WebauthnAuthenticator
	authenticator := &crypto.WebauthnAuthenticator{}
	auth, err := authenticatorFromMetadata(vm.GetMetadata())
	if err != nil {
		fmt.Println(err)
	} else {
		authenticator = auth
	}

	// Build the credential
	cred := &crypto.WebauthnCredential{
		Id:            credID,
		PublicKey:     pubKey,
		Authenticator: authenticator,
	}

	return NewCredential(cred, vm.Controller), nil
}

func (c *didCredential) Controller() string {
	return c.UserDid
}

// Descriptor returns the credential's descriptor
func (c *didCredential) Descriptor() protocol.CredentialDescriptor {
	return protocol.CredentialDescriptor{
		CredentialID:    c.Credential.Id,
		Type:            protocol.PublicKeyCredentialType,
		Transport:       []protocol.AuthenticatorTransport{protocol.Internal},
		AttestationType: "direct",
	}
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
	did := fmt.Sprintf("did:key:%s#%s", base64.RawURLEncoding.EncodeToString(c.Credential.PublicKey), base64.RawURLEncoding.EncodeToString(c.Credential.Id))
	pubMb := base64.RawURLEncoding.EncodeToString(c.Credential.PublicKey)
	vmType := crypto.Ed25519KeyType.FormatString()
	return &VerificationMethod{
		Id:                 did,
		Type:               vmType,
		PublicKeyMultibase: pubMb,
		Controller:         c.UserDid,
		Metadata:           authenticatorToMetadata(c.Credential.Authenticator),
	}
}

// Did returns the credential's DID
func (c *didCredential) Did() string {
	return c.ToVerificationMethod().Id
}

// Encrypt is used to encrypt a message for the credential
func (c *didCredential) Encrypt(data []byte) ([]byte, error) {
	// Get the public key from the credential
	keyFace, err := webauthncose.ParsePublicKey(c.Credential.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}
	publicKey, ok := keyFace.(webauthncose.EC2PublicKeyData)
	if !ok {
		return nil, errors.New("public key is not an EC2 key")
	}
	// Derive a shared secret using ECDH
	privateKey, err := derivePrivateKey(c.Credential)
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
	keyFace, err := webauthncose.ParsePublicKey(c.Credential.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}
	publicKey, ok := keyFace.(webauthncose.EC2PublicKeyData)
	if !ok {
		return nil, errors.New("public key is not an EC2 key")
	}
	// Derive a shared secret using ECDH
	privateKey, err := derivePrivateKey(c.Credential)
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

func authenticatorToMetadata(authenticator *crypto.WebauthnAuthenticator) []*KeyValuePair {
	authenticatorMap := make(map[string]string)
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

func authenticatorFromMetadata(metadata []*KeyValuePair) (*crypto.WebauthnAuthenticator, error) {
	authenticatorMap := KeyValueListToMap(metadata)
	aaguid, err := base64.StdEncoding.DecodeString(authenticatorMap["aaguid"])
	if err != nil {
		return nil, err
	}
	signCount, err := strconv.ParseUint(authenticatorMap["sign_count"], 10, 32)
	if err != nil {
		return nil, err
	}
	cloneWarning, err := strconv.ParseBool(authenticatorMap["clone_warning"])
	if err != nil {
		return nil, err
	}
	authenticator := &crypto.WebauthnAuthenticator{
		Aaguid:       aaguid,
		SignCount:    uint32(signCount),
		CloneWarning: cloneWarning,
	}
	return authenticator, nil
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
