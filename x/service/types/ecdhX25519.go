package types

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	fmt "fmt"
	io "io"
	"math/big"

	"github.com/go-webauthn/webauthn/protocol/webauthncose"
	"golang.org/x/crypto/hkdf"
)

func eciesEncrypt(publicKeyData webauthncose.EC2PublicKeyData, plaintext []byte) ([]byte, error) {
	// Convert the X and Y coordinates of the public key to big.Int values
	x := new(big.Int).SetBytes(publicKeyData.XCoord)
	y := new(big.Int).SetBytes(publicKeyData.YCoord)

	// Create an ECDSA public key from the X and Y coordinates and the curve identifier
	curve := getCurve(publicKeyData.Algorithm)
	if curve == nil {
		return nil, fmt.Errorf("unsupported curve identifier: %d", publicKeyData.Algorithm)
	}
	publicKey := ecdsa.PublicKey{Curve: curve, X: x, Y: y}
	// Generate a random key pair for the ephemeral keys
	ephemeralPrivateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate ephemeral private key: %w", err)
	}

	// Derive the shared secret
	sharedSecretX, _ := publicKey.Curve.ScalarMult(publicKey.X, publicKey.Y, ephemeralPrivateKey.D.Bytes())
	sharedSecret := sharedSecretX.Bytes()

	// Derive the encryption key and MAC key from the shared secret
	encryptionKey, macKey := deriveKeys(sharedSecret)

	// Generate a random IV
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, fmt.Errorf("failed to generate IV: %w", err)
	}

	// Encrypt the plaintext using AES-256-GCM with the encryption key and IV
	ciphertext := make([]byte, len(plaintext))
	c, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}
	gcm, err := cipher.NewGCMWithNonceSize(c, aes.BlockSize)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM cipher: %w", err)
	}
	gcm.Seal(ciphertext[:0], iv, plaintext, nil)

	// Compute the MAC tag
	mac := hmac.New(sha256.New, macKey)
	mac.Write(iv)
	mac.Write(ciphertext)
	tag := mac.Sum(nil)

	// Encode the ephemeral public key
	ephemeralPublicKeyBytes := elliptic.Marshal(elliptic.P256(), ephemeralPrivateKey.PublicKey.X, ephemeralPrivateKey.PublicKey.Y)
	ephemeralPublicKey, err := x509.ParsePKIXPublicKey(ephemeralPublicKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to create ephemeral public key: %w", err)
	}
	encodedPublicKey := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(ephemeralPublicKey.(*rsa.PublicKey)),
	})

	// Concatenate the public key, IV, ciphertext, and MAC tag
	result := make([]byte, len(encodedPublicKey)+len(iv)+len(ciphertext)+len(tag))
	copy(result[:len(encodedPublicKey)], encodedPublicKey)
	copy(result[len(encodedPublicKey):len(encodedPublicKey)+len(iv)], iv)
	copy(result[len(encodedPublicKey)+len(iv):len(encodedPublicKey)+len(iv)+len(ciphertext)], ciphertext)
	copy(result[len(encodedPublicKey)+len(iv)+len(ciphertext):], tag)
	return result, nil
}

func deriveKeys(sharedSecret []byte) ([]byte, []byte) {
	// Use HKDF to derive the encryption key and MAC key from the shared secret
	info := []byte("encryption key")
	encryptionKey := hkdf.New(sha256.New, sharedSecret, nil, info)
	info = []byte("MAC key")
	macKey := hkdf.New(sha256.New, sharedSecret, nil, info)
	encKeyBytes := make([]byte, 32)
	if _, err := io.ReadFull(encryptionKey, encKeyBytes); err != nil {
		return nil, nil
	}
	macKeyBytes := make([]byte, 32)
	if _, err := io.ReadFull(macKey, macKeyBytes); err != nil {
		return nil, nil
	}
	return encKeyBytes, macKeyBytes
}

func derivePrivateKey(credential *WebauthnCredential) (*ecdsa.PrivateKey, error) {
	// Parse the public key from the credential
	pubKeyFace, err := webauthncose.ParsePublicKey(credential.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}
	pubKey, ok := pubKeyFace.(webauthncose.EC2PublicKeyData)
	if !ok {
		return nil, fmt.Errorf("public key is not an EC2 key")
	}

	// Convert the x and y coordinates to *big.Int values
	xCoord := new(big.Int).SetBytes(pubKey.XCoord)
	yCoord := new(big.Int).SetBytes(pubKey.YCoord)

	// Generate a new ephemeral private key
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate ephemeral private key: %w", err)
	}

	// Derive the shared secret using ECDH
	sharedX, _ := privKey.Curve.ScalarMult(xCoord, yCoord, privKey.D.Bytes())
	sharedSecret := sharedX.Bytes()

	// Derive a 256-bit key using HKDF
	keyBytes := make([]byte, 32)
	info := []byte("webauthn-secret")
	hkdf := hkdf.New(sha256.New, sharedSecret, nil, info)
	if _, err := hkdf.Read(keyBytes); err != nil {
		return nil, fmt.Errorf("failed to derive key: %w", err)
	}

	// Create a new private key from the derived key
	return &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: elliptic.P256(),
			X:     xCoord,
			Y:     yCoord,
		},
		D: new(big.Int).SetBytes(keyBytes),
	}, nil
}

func sharedSecret(privateKey *ecdsa.PrivateKey, publicKey webauthncose.EC2PublicKeyData) ([]byte, error) {
	// Convert the X and Y coordinates of the public key to big.Int values
	x := new(big.Int).SetBytes(publicKey.XCoord)
	y := new(big.Int).SetBytes(publicKey.YCoord)

	// Create an ECDSA public key from the X and Y coordinates and the curve identifier
	curve := getCurve(publicKey.Algorithm)
	if curve == nil {
		return nil, fmt.Errorf("unsupported curve identifier: %d", publicKey.Algorithm)
	}
	publicKeyEcdsa := ecdsa.PublicKey{Curve: curve, X: x, Y: y}

	// Calculate the shared secret using ECDH
	x, _ = curve.ScalarMult(publicKeyEcdsa.X, publicKeyEcdsa.Y, privateKey.D.Bytes())
	return x.Bytes(), nil
}

func getCurve(curveID int64) elliptic.Curve {
	var curve elliptic.Curve
	switch webauthncose.COSEAlgorithmIdentifier(curveID) {
	case webauthncose.AlgES512: // IANA COSE code for ECDSA w/ SHA-512
		curve = elliptic.P521()
	case webauthncose.AlgES384: // IANA COSE code for ECDSA w/ SHA-384
		curve = elliptic.P384()
	case webauthncose.AlgES256: // IANA COSE code for ECDSA w/ SHA-256
		curve = elliptic.P256()
	default:
		return nil
	}
	return curve
}
