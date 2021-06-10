package auth

import (
	md "github.com/sonr-io/core/pkg/models"
)

type AuthService interface {
	BuildPrefix(val string) string
	CreateSName(req *md.UsernameRequest) *md.UsernameResponse
	GetUser(prefix string) *md.User
	PutUser(user *md.User)
	ValidateUser(sName string, mnemonic string) bool
}

type authService struct {
	AuthService
	keyPair  *md.KeyPair
	store    md.Store
	nbClient NamebaseClient
}

// Creates New Auth Service
func NewAuthService(cr *md.ConnectionRequest, keyPair *md.KeyPair, s md.Store) AuthService {
	// Create NBClient
	return &authService{
		store:    s,
		nbClient: newNambaseClient(cr.GetApiKeys()),
	}
}

func (as *authService) CreateSName(req *md.UsernameRequest) *md.UsernameResponse {
	
}
