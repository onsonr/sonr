package mpc

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/golang-jwt/jwt"
)

// MPCSigningMethod implements the SigningMethod interface for MPC-based signing
type MPCSigningMethod struct {
	Name string
	ks   ucanKeyshare
}

// NewJWTSigningMethod creates a new MPC signing method with the given keyshare source
func NewJWTSigningMethod(name string, ks ucanKeyshare) *MPCSigningMethod {
	return &MPCSigningMethod{
		Name: name,
		ks:   ks,
	}
}

// Alg returns the signing method's name
func (m *MPCSigningMethod) Alg() string {
	return m.Name
}

// Verify verifies the signature using the MPC public key
func (m *MPCSigningMethod) Verify(signingString, signature string, key interface{}) error {
	// Decode the signature
	sig, err := base64.RawURLEncoding.DecodeString(signature)
	if err != nil {
		return err
	}

	// Hash the signing string
	hasher := sha256.New()
	hasher.Write([]byte(signingString))
	digest := hasher.Sum(nil)

	// Verify using the keyshare's public key
	valid, err := VerifySignature(m.ks.valShare.PublicKey, digest, sig)
	if err != nil {
		return fmt.Errorf("failed to verify signature: %w", err)
	}
	if !valid {
		return fmt.Errorf("invalid signature")
	}

	return nil
}

// Sign signs the data using MPC
func (m *MPCSigningMethod) Sign(signingString string, key interface{}) (string, error) {
	// Hash the signing string
	hasher := sha256.New()
	hasher.Write([]byte(signingString))
	digest := hasher.Sum(nil)

	// Create signing functions
	signFunc, err := m.ks.userShare.SignFunc(digest)
	if err != nil {
		return "", fmt.Errorf("failed to create sign function: %w", err)
	}

	valSignFunc, err := m.ks.valShare.SignFunc(digest)
	if err != nil {
		return "", fmt.Errorf("failed to create validator sign function: %w", err)
	}

	// Run the signing protocol
	sig, err := RunSignProtocol(valSignFunc, signFunc)
	if err != nil {
		return "", fmt.Errorf("failed to run sign protocol: %w", err)
	}

	// Serialize the signature
	sigBytes, err := SerializeSignature(sig)
	if err != nil {
		return "", fmt.Errorf("failed to serialize signature: %w", err)
	}

	// Encode the signature
	encoded := base64.RawURLEncoding.EncodeToString(sigBytes)
	return encoded, nil
}

func init() {
	// Register the MPC signing method
	jwt.RegisterSigningMethod("MPC256", func() jwt.SigningMethod {
		return &MPCSigningMethod{
			Name: "MPC256",
		}
	})
}
