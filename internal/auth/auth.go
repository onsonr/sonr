package auth

import (
	"log"

	md "github.com/sonr-io/core/pkg/models"
)

type AuthService interface {
	BuildPrefix(val string) string
	CheckSName(req *md.AuthenticationRequest) *md.AuthenticationResponse
	GetUser(prefix string) *md.User
	PutUser(user *md.User)
	SaveSName(req *md.AuthenticationRequest) *md.AuthenticationResponse
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
func NewAuthService(req *md.ConnectionRequest, s md.Store, cb md.NodeCallback) AuthService {
	// Create NBClient
	return &authService{
		store:    s,
		nbClient: newNambaseClient(req.GetApiKeys()),
		device:   req.GetDevice(),
		callback: cb,
	}
}

// Checks if User can use SName
func (as *authService) CheckSName(req *md.AuthenticationRequest) *md.AuthenticationResponse {
	// Initialze Response
	resp := &md.AuthenticationResponse{
		SName:    req.GetSName(),
		Mnemonic: req.GetMnemonic(),
		IsValid:  true,
		IsSaved:  false,
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

// Saves SName after Validation
func (as *authService) SaveSName(req *md.AuthenticationRequest) *md.AuthenticationResponse {
	// Initialze Response
	resp := as.CheckSName(req)

	// Get Prefix
	prefix, err := req.Device.Prefix(req.SName)
	if err != nil {
		return nil
	}

	// Get FingerPrint
	fingerprint, err := req.Device.Fingerprint(req.GetMnemonic())
	if err != nil {
		return nil
	}

	// Validate
	if resp.GetIsValid() {
		// Save Record, Check Success
		if ok := as.nbClient.AddRecord(req.ToHSRecord(prefix, fingerprint)); ok {
			// Update Response
			resp.IsSaved = true
			resp.Prefix = prefix

			// Save to Store
			err := as.store.PutCrypto(resp.ToUserCrypto())
			if err != nil {
				log.Println(err.String())
			}
		} else {
			resp.IsSaved = false
		}
	}
	return resp
}
