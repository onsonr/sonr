package local

import (
	"os"
	"path"

	"github.com/tink-crypto/tink-go/v2/daead"
	"github.com/tink-crypto/tink-go/v2/keyset"
)

// These constants are derived from the above variables.
// These are the ones we will want to use in the code, based on
// any overrides above
var (
	nodeDir = ".sonr"

	defaultNodeHome = os.ExpandEnv("$HOME/") + nodeDir

	kh *keyset.Handle
)

func keysetFile() string {
	return path.Join(defaultNodeHome, "daead_keyset.json")
}

func setupKeyHandle() {
	if _, err := os.Stat(keysetFile()); os.IsNotExist(err) {
		// If the keyset file doesn't exist, generate a new key handle
		kh, _ = NewKeyHandle()
	} else {
		// If the keyset file exists, load the key handle from the file
		kh, _ = ReadKeyHandle()
	}
}

// NewKeyHandle creates a new keyset, writes it to a file, and returns the keyset handle
func NewKeyHandle() (*keyset.Handle, error) {
	kh, err := keyset.NewHandle(daead.AESSIVKeyTemplate())
	if err != nil {
		return nil, err
	}

	// Write the keyset handle to a file
	f, err := os.Create(keysetFile())
	if err != nil {
		return nil, err
	}
	defer f.Close()

	writer := keyset.NewJSONWriter(f)
	if err := kh.WriteWithNoSecrets(writer); err != nil {
		return nil, err
	}
	return kh, nil
}

// ReadKeyHandle reads the keyset handle from a file and returns it
func ReadKeyHandle() (*keyset.Handle, error) {
	// Read the keyset handle from the file
	f, err := os.Open(keysetFile())
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := keyset.NewJSONReader(f)
	kh, err := keyset.ReadWithNoSecrets(reader)
	if err != nil {
		return nil, err
	}

	return kh, nil
}

// Encrypt takes plaintext and associated data and returns ciphertext
func Encrypt(plaintext []byte, associatedData []byte) ([]byte, error) {
	d, err := daead.New(kh)
	if err != nil {
		return nil, err
	}

	ciphertext, err := d.EncryptDeterministically(plaintext, associatedData)
	if err != nil {
		return nil, err
	}

	return ciphertext, nil
}

// Decrypt takes ciphertext and associated data and returns plaintext
func Decrypt(ciphertext []byte, associatedData []byte) ([]byte, error) {
	d, err := daead.New(kh)
	if err != nil {
		return nil, err
	}

	decrypted, err := d.DecryptDeterministically(ciphertext, associatedData)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}
