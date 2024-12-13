package services

import (
	"github.com/onsonr/sonr/pkg/ipfsapi"
)

type VaultService struct {
	tokenStore ipfsapi.IPFSTokenStore
}

func NewVaultService(ipc ipfsapi.Client) *VaultService {
	return &VaultService{
		tokenStore: ipfsapi.NewUCANStore(ipc),
	}
}
