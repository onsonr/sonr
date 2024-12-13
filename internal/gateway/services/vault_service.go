package services

import (
	"github.com/onsonr/sonr/pkg/ipfsapi"
	"gorm.io/gorm"
)

type VaultService struct {
	db         *gorm.DB
	tokenStore ipfsapi.IPFSTokenStore
}

func NewVaultService(db *gorm.DB, ipc ipfsapi.Client) *VaultService {
	return &VaultService{
		db:         db,
		tokenStore: ipfsapi.NewUCANStore(ipc),
	}
}
