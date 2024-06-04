package vault

import (
	"github.com/di-dao/sonr/crypto/kss"
	"github.com/di-dao/sonr/pkg/ipfs"
	"github.com/di-dao/sonr/pkg/vault/auth"
	"github.com/di-dao/sonr/pkg/vault/wallet"
	"github.com/ipfs/boxo/files"
)

type vaultFS struct {
	Wallet     *wallet.Wallet   `json:"wallet"`
	Creds      auth.Credentials `json:"credentials"`
	Properties auth.Properties  `json:"properties"`
}

func (v *vaultFS) GetInfoFile() *auth.InfoFile {
	return &auth.InfoFile{
		Creds:      v.Creds,
		Properties: v.Properties,
	}
}

func (v *vaultFS) ToFileMap() (map[string]files.File, error) {
	flMap := make(map[string]files.File)

	wallBz, err := v.Wallet.Marshal()
	if err != nil {
		return nil, err
	}
	walletFile := files.NewBytesFile(wallBz)
	flMap["wallet.json"] = walletFile

	info := v.GetInfoFile()
	infoBz, err := info.Marshal()
	if err != nil {
		return nil, err
	}
	infoFile := files.NewBytesFile(infoBz)
	flMap["info.json"] = infoFile

	return flMap, nil
}

func createVaultFS(set kss.Set) (*vaultFS, error) {
	wallet, err := wallet.New(set)
	if err != nil {
		return nil, err
	}

	return &vaultFS{
		Wallet:     wallet,
		Creds:      auth.NewCredentials(),
		Properties: auth.NewProperties(),
	}, nil
}

func loadVaultFS(vfs ipfs.VFS) (*vaultFS, error) {
	wallet := &wallet.Wallet{}
	walletBz, err := vfs.Get("wallet.json")
	if err != nil {
		return nil, err
	}

	err = wallet.Unmarshal(walletBz)
	if err != nil {
		return nil, err
	}

	info := &auth.InfoFile{}
	infoBz, err := vfs.Get("info.json")
	if err != nil {
		return nil, err
	}
	err = info.Unmarshal(infoBz)
	if err != nil {
		return nil, err
	}

	return &vaultFS{
		Wallet:     wallet,
		Creds:      info.Creds,
		Properties: info.Properties,
	}, nil
}
