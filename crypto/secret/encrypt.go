package secret

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"

	"github.com/cloudflare/circl/kem/kyber/kyber768"
	"github.com/onsonr/hway/crypto/accumulator"
	"github.com/ipfs/go-cid"
)

const AccumulatorMarshalledSize = 60

func (s *PrimaryKey) Encrypt(acc *accumulator.Accumulator, vaultCID string, message []byte) ([]byte, error) {
	pub, _, err := deriveKyberKeypair(acc, vaultCID)
	if err != nil {
		return nil, err
	}

	ct := make([]byte, kyber768.CiphertextSize)
	ss := make([]byte, kyber768.SharedKeySize)
	pub.EncapsulateTo(ct, ss, nil)

	block, err := aes.NewCipher(ss)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("failed to generate nonce: %w", err)
	}

	accBytes, err := acc.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal accumulator: %w", err)
	}

	if len(accBytes) != AccumulatorMarshalledSize {
		return nil, fmt.Errorf("unexpected accumulator marshalled size: got %d, want %d", len(accBytes), AccumulatorMarshalledSize)
	}

	paddedMessage := append(accBytes, message...)
	encryptedMessage := gcm.Seal(nil, nonce, paddedMessage, nil)

	result := append(ct, nonce...)
	result = append(result, encryptedMessage...)

	return result, nil
}

// Decrypt is
func (s *PrimaryKey) Decrypt(vaultCID string, encryptedData []byte, witness *accumulator.MembershipWitness, pubKey *accumulator.PublicKey) ([]byte, error) {
	if len(encryptedData) < kyber768.CiphertextSize+AccumulatorMarshalledSize {
		return nil, fmt.Errorf("invalid encrypted data: too short")
	}

	// Extract and unmarshal the accumulator from the first 60 bytes
	var decryptedAcc accumulator.Accumulator
	err := decryptedAcc.UnmarshalBinary(encryptedData[:AccumulatorMarshalledSize])
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal accumulator: %w", err)
	}

	// Derive Kyber keypair using the unmarshalled accumulator
	_, priv, err := deriveKyberKeypair(&decryptedAcc, vaultCID)
	if err != nil {
		return nil, err
	}

	// Decapsulate the shared secret
	ct := encryptedData[AccumulatorMarshalledSize : AccumulatorMarshalledSize+kyber768.CiphertextSize]
	ss := make([]byte, kyber768.SharedKeySize)
	priv.DecapsulateTo(ss, ct)

	// Set up AES-GCM decryption
	block, err := aes.NewCipher(ss)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}

	nonceSize := gcm.NonceSize()
	if len(encryptedData) < AccumulatorMarshalledSize+kyber768.CiphertextSize+nonceSize {
		return nil, fmt.Errorf("invalid encrypted data: too short for nonce")
	}

	nonce := encryptedData[AccumulatorMarshalledSize+kyber768.CiphertextSize : AccumulatorMarshalledSize+kyber768.CiphertextSize+nonceSize]
	ciphertext := encryptedData[AccumulatorMarshalledSize+kyber768.CiphertextSize+nonceSize:]

	// Decrypt the message
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt: %w", err)
	}

	// Verify the witness using the decrypted accumulator and provided secret key
	if err := witness.Verify(pubKey, &decryptedAcc); err != nil {
		return nil, fmt.Errorf("unauthorized witness: %w", err)
	}

	return plaintext, nil
}

func deriveKyberKeypair(acc *accumulator.Accumulator, vaultCID string) (*kyber768.PublicKey, *kyber768.PrivateKey, error) {
	seed, err := generateDeterministicSeed(acc, vaultCID)
	if err != nil {
		return nil, nil, err
	}

	// Ensure the seed is the correct size for Kyber768
	if len(seed) < kyber768.KeySeedSize {
		expandedSeed := make([]byte, kyber768.KeySeedSize)
		copy(expandedSeed, seed)
		seed = expandedSeed
	}

	pub, priv := kyber768.NewKeyFromSeed(seed[:kyber768.KeySeedSize])
	return pub, priv, nil
}

func generateDeterministicSeed(acc *accumulator.Accumulator, vaultCID string) ([]byte, error) {
	_, err := cid.Decode(vaultCID)
	if err != nil {
		return nil, fmt.Errorf("invalid IPFS CID: %w", err)
	}

	accBytes, err := acc.MarshalBinary()
	if err != nil {
		return nil, err
	}
	data := append(accBytes, []byte(vaultCID)...)

	hash := sha256.Sum256(data)
	return hash[:], nil
}
