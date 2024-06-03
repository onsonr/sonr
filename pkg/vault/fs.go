package vault

import (
	"encoding/json"

	"github.com/di-dao/sonr/crypto/kss"
	"github.com/di-dao/sonr/pkg/vault/auth"
	"github.com/di-dao/sonr/pkg/vault/props"
	"github.com/di-dao/sonr/pkg/vault/wallet"
	"github.com/ipfs/boxo/files"
)

type InfoFile struct {
	Creds      auth.Credentials `json:"credentials"`
	Properties props.Properties `json:"properties"`
}

func (i *InfoFile) Marshal() ([]byte, error) {
	return json.Marshal(i)
}

func (i *InfoFile) Unmarshal(data []byte) error {
	return json.Unmarshal(data, i)
}

type vaultFS struct {
	Wallet     *wallet.Wallet
	Creds      auth.Credentials
	Properties props.Properties
}

func (v *vaultFS) GetInfoFile() *InfoFile {
	return &InfoFile{
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
		Properties: props.NewProperties(),
	}, nil
}
