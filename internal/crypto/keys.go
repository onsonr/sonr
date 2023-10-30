package crypto

// EncryptionKey is an interface for encryption keys
type EncryptionKey interface {
	// Bytes returns the bytes of the encryption key
	Bytes() []byte

	// Encrypt encrypts the message with the encryption key
	Encrypt(msg []byte) ([]byte, error)

	// Decrypt decrypts the message with the encryption key
	Decrypt(msg []byte) ([]byte, error)
}
