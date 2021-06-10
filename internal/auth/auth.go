package auth

import (
	md "github.com/sonr-io/core/pkg/models"
)

type AuthService interface {
	BuildPrefix(val string) string
	CreateSName(name string)
	GetUser(prefix string) *md.User
	PutUser(user *md.User)
	ValidateUser(val string) bool
}

type authService struct {
	AuthService
	store    md.Store
	nbClient NamebaseClient
}

// Creates New Auth Service
func NewAuthService(cr *md.ConnectionRequest, s md.Store) AuthService {
	// Create NBClient
	return &authService{
		store:    s,
		nbClient: newNambaseClient(cr.GetApiKeys()),
	}
}
