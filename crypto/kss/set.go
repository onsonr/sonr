package kss

import (
	"io"

	"github.com/di-dao/sonr/crypto"
	"github.com/di-dao/sonr/crypto/daed"
	"github.com/ipfs/boxo/files"
)

// KssI is the interface for the keyshare set
type EncryptedSet interface {
	Decrypt(key []byte, kh *daed.AESSIV) (Set, error)
	PublicKey() crypto.PublicKey
	FileMap() map[string]files.File
}

type encryptedSet struct {
	publicKey crypto.PublicKey
	encValKey []byte
	encUsrKey []byte
}

func (es *encryptedSet) Decrypt(key []byte, kh *daed.AESSIV) (Set, error) {
	decValKs, err := kh.DecryptDeterministically(es.encValKey, key)
	if err != nil {
		return nil, err
	}

	decUsrKs, err := kh.DecryptDeterministically(es.encUsrKey, key)
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

func NewEncryptedSetFromFileMap(files map[string]files.File) (EncryptedSet, error) {
	usrKsFile := files["usr.ks"]
	valKsFile := files["val.ks"]
	encUsrKey, err := io.ReadAll(usrKsFile)
	if err != nil {
		return nil, err
	}

	encValKey, err := io.ReadAll(valKsFile)
	if err != nil {
		return nil, err
	}
	return &encryptedSet{
		encUsrKey: encUsrKey,
		encValKey: encValKey,
	}, nil
}
