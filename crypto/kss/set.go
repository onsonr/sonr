package kss

import (
	"github.com/di-dao/sonr/crypto"
	"github.com/ipfs/boxo/files"

	"github.com/tink-crypto/tink-go/v2/daead"
	"github.com/tink-crypto/tink-go/v2/keyset"
)

// KssI is the interface for the keyshare set
type EncryptedSet interface {
	Decrypt(key []byte, kh *keyset.Handle) (Set, error)
	PublicKey() crypto.PublicKey
	FileMap() map[string]files.File
}

type encryptedSet struct {
	publicKey crypto.PublicKey
	encValKey []byte
	encUsrKey []byte
}

func (es *encryptedSet) Decrypt(key []byte, kh *keyset.Handle) (Set, error) {
	d, err := daead.New(kh)
	if err != nil {
		return nil, err
	}

	decValKs, err := d.DecryptDeterministically(es.encValKey, key)
	if err != nil {
		return nil, err
	}

	decUsrKs, err := d.DecryptDeterministically(es.encUsrKey, key)
	if err != nil {
		return nil, err
	}
	return LoadKeyshareSet(decValKs, decUsrKs)
}

func (es *encryptedSet) PublicKey() crypto.PublicKey {
	return es.publicKey
}

func (es *encryptedSet) FileMap() map[string]files.File {
	fileMap := make(map[string]files.File)
	usrKsFile := files.NewBytesFile(es.encUsrKey)
	valKsFile := files.NewBytesFile(es.encValKey)
	fileMap["usr.ks"] = usrKsFile
	fileMap["val.ks"] = valKsFile
	return nil
}
