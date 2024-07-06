package motor

import (
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/onsonr/hway/crypto/kss"
	"github.com/onsonr/hway/crypto/mpc"
	fs "github.com/onsonr/hway/internal/vfs"
)

const kSonrHRP = "idx"

// vfd is the struct implementation of an IPFS file system
type drive struct {
	kss    kss.Set
	folder fs.Folder
	addr   string
}

// NewVFS creates a new virtual file system.
func Spawn() (*drive, error) {
	kss, err := mpc.GenerateKss()
	if err != nil {
		return nil, err
	}

	addr, err := bech32.ConvertAndEncode(kSonrHRP, kss.PublicKey().Bytes())
	if err != nil {
		return nil, err
	}

	rootDir, err := fs.NewVaultFolder(addr)
	if err != nil {
		return nil, err
	}

	return &drive{
		folder: rootDir,
		addr:   addr,
		kss:    kss,
	}, nil
}

// // CreateFingerprint creates a fingerprint for the given database and public key
// func CreateFingerprint(dir fs.Folder, db dwn.Database, publicKey crypto.PublicKey) error {
// 	pk, err := secret.NewKey("credentials", publicKey)
// 	if err != nil {
// 		return err
// 	}
// 	creds, err := db.ListCredentials()
// 	if err != nil {
// 		return err
// 	}

// 	credIDStrs := make([]string, len(creds))
// 	for i, c := range creds {
// 		credIDStrs[i] = c.DID
// 	}

// 	acc, err := pk.CreateAccumulator(credIDStrs...)
// 	if err != nil {
// 		return err
// 	}

// 	bz, err := acc.MarshalBinary()
// 	if err != nil {
// 		return err
// 	}
// 	_, err = dir.WriteFile("fingerprint", bz, os.ModePerm)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // ValidateAndPurgeFingerprint validates the fingerprint and purges it if it's valid
// func ValidateAndPurgeFingerprint(dir fs.Folder, witness []byte, publicKey crypto.PublicKey) (bool, error) {
// 	pk, err := secret.NewKey("credentials", publicKey)
// 	if err != nil {
// 		return false, err
// 	}
// 	creds := new(accumulator.Accumulator)
// 	membership := new(accumulator.MembershipWitness)
// 	err = membership.UnmarshalBinary(witness)
// 	if err != nil {
// 		return false, err
// 	}

// 	bz, err := dir.ReadFile("fingerprint")
// 	if err != nil {
// 		return false, err
// 	}

// 	err = creds.UnmarshalBinary(bz)
// 	if err != nil {
// 		return false, err
// 	}

// 	if err := pk.VerifyWitness(creds, membership); err != nil {
// 		return false, err
// 	}

// 	err = dir.DeleteFile("fingerprint")
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }
