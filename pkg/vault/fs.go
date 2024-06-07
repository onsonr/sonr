package vault

import (
	"time"

	"github.com/di-dao/sonr/crypto/kss"
	"github.com/di-dao/sonr/pkg/cache"
	"github.com/di-dao/sonr/pkg/ipfs"
	"github.com/di-dao/sonr/pkg/vault/types"
	"github.com/ipfs/boxo/files"
)

var vaultCache *cache.Cache[contextKey, vaultFS]

type contextKey string

func (c contextKey) String() string {
	return "vault/" + string(c)
}

func init() {
	// This is a placeholder
	vaultCache = cache.New[contextKey, vaultFS](time.Minute*5, time.Minute)
}

type vaultFS struct {
	Wallet     *types.Wallet     `json:"wallet"`
	Creds      types.Credentials `json:"credentials"`
	Properties types.Properties  `json:"properties"`
}

func (v *vaultFS) GetInfoFile() *types.InfoFile {
	return &types.InfoFile{
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
	wallet, err := types.NewWallet(set)
	if err != nil {
		return nil, err
	}

	return &vaultFS{
		Wallet:     wallet,
		Creds:      types.NewCredentials(),
		Properties: types.NewProperties(),
	}, nil
}

func loadVaultFS(vfs ipfs.VFS) (*vaultFS, error) {
	wallet := &types.Wallet{}
	walletBz, err := vfs.Get("wallet.json")
	if err != nil {
		return nil, err
	}

	err = wallet.Unmarshal(walletBz)
	if err != nil {
		return nil, err
	}

	info := &types.InfoFile{}
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
