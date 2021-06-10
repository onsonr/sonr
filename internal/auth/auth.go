package auth

import (
	md "github.com/sonr-io/core/pkg/models"
)

type AuthService interface {
	BuildPrefix(val string) string
	CreateSName(req *md.AuthenticationRequest) *md.AuthenticationResponse
	GetUser(prefix string) *md.User
	PutUser(user *md.User)
	ValidateUser(sName string, mnemonic string) bool
}

type authService struct {
	AuthService
	device   *md.Device
	keyPair  *md.KeyPair
	store    md.Store
	nbClient NamebaseClient
	callback md.NodeCallback
}

// Creates New Auth Service
func NewAuthService(req *md.AuthenticationRequest, s md.Store, cb md.NodeCallback) AuthService {
	// Create NBClient
	return &authService{
		store:    s,
		nbClient: newNambaseClient(req.GetApiKeys()),
		device:   req.GetDevice(),
		callback: cb,
	}
}

func (as *authService) CreateSName(req *md.AuthenticationRequest) *md.AuthenticationResponse {
	// Initialze Response
	resp := &md.AuthenticationResponse{
		SName:    req.GetSName(),
		Mnemonic: req.GetMnemonic(),
		IsValid:  true,
	}

	// Fetch Records
	records, err := as.nbClient.Refresh()
	if err != nil {
		return nil
	}

	// Validate Name
	for _, r := range records {
		if r.IsName(req.GetSName()) {
			resp.IsValid = false
			break
		}
	}
	return resp
}
