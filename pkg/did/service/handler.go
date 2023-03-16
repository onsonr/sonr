package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/sonrhq/core/pkg/client/chain"
	"github.com/sonrhq/core/pkg/wallet"
	snrcrypto "github.com/sonrhq/core/types/crypto"
	v1 "github.com/sonrhq/core/types/vault/v1"
	"github.com/sonrhq/core/x/identity/types"
)

type ServiceHandler interface {
	// GetRPID returns the RPID for the service.
	GetRPID() string

	// GetRPName returns the RPName for the service.
	GetRPName() string

	// This method is used to get the challenge response from the DID controller.
	BeginRegistration(req *v1.RegisterStartRequest) ([]byte, error)

	// This is the method that will be called when the user clicks on the "Register" button.
	FinishRegistration(req *v1.RegisterFinishRequest) (*snrcrypto.PubKey, error)

	// This method is used to get the options for the assertion.
	BeginLogin(req *v1.LoginStartRequest) ([]byte, error)

	// This is the method that will be called when the user clicks the "Login" button on the login page.
	FinishLogin(req *v1.LoginFinishRequest) (*snrcrypto.PubKey, error)
}

type serviceHandlerImpl struct {
	// Service is the service that will be used to handle the requests.
	service *types.Service

	// Client is the client that will be used to interact with the chain.
	sonrQueryClient *chain.SonrQueryClient

	cache      *cache.Cache
	newWallets chan wallet.Wallet
}

func LoadHandler(origin string, apiEndpoint chain.APIEndpoint) (ServiceHandler, error) {
	// Get the service from the chain.
	sonrQueryClient := chain.NewClient(apiEndpoint)
	service, err := sonrQueryClient.GetService(context.Background(), origin)
	if err != nil {
		return nil, fmt.Errorf("failed to get service from chain: %w", err)
	}
	c := cache.New(30*time.Second, 1*time.Minute)
	return &serviceHandlerImpl{
		service:         service,
		sonrQueryClient: sonrQueryClient,
		cache:           c,
		newWallets:      make(chan wallet.Wallet),
	}, nil
}

// GetRPID returns the RPID for the service.
func (s *serviceHandlerImpl) GetRPID() string {
	return s.service.GetOrigin()
}

// GetRPName returns the RPName for the service.
func (s *serviceHandlerImpl) GetRPName() string {
	return s.service.GetName()
}

// BeginRegistration is the method that will be called when the user clicks on the "Register" button.
func (s *serviceHandlerImpl) BeginRegistration(req *v1.RegisterStartRequest) ([]byte, error) {
	// Get the parameters from the chain.
	params := types.NewParams()

	// Issue the challenge.
	resp, err := params.NewWebauthnCreationOptions(s.service, req.Uuid, req.DeviceLabel)
	if err != nil {
		return nil, fmt.Errorf("failed to issue challenge: %w", err)
	}
	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}
	go s.GenerateWallet(req.Uuid, 1)
	return jsonResponse, nil
}

// FinishRegistration is the method that will be called when the user clicks on the "Register" button.
func (s *serviceHandlerImpl) FinishRegistration(req *v1.RegisterFinishRequest) (*snrcrypto.PubKey, error) {
	// Verify the challenge.
	cred, err := s.service.VerifyCreationChallenge(req.CredentialResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to verify challenge: %w", err)
	}
	exportedWall, ok := s.cache.Get(req.Uuid)
	if !ok {
		return nil, fmt.Errorf("failed to get wallet from cache")
	}

	wall, err := wallet.Import(exportedWall.([]byte))
	if err != nil {
		return nil, fmt.Errorf("failed to import wallet: %w", err)
	}
	err = wall.SetAuthentication(cred)
	if err != nil {
		return nil, fmt.Errorf("failed to set authentication: %w", err)
	}
	return cred.PubKey(), nil
}

// BeginLogin is the method that will be called when the user clicks on the "Login" button.
func (s *serviceHandlerImpl) BeginLogin(req *v1.LoginStartRequest) ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

// FinishLogin is the method that will be called when the user clicks on the "Login" button.
func (s *serviceHandlerImpl) FinishLogin(req *v1.LoginFinishRequest) (*snrcrypto.PubKey, error) {
	return nil, fmt.Errorf("not implemented")
}

// GenerateWallet generates a new wallet
func (s *serviceHandlerImpl) GenerateWallet(currId string, threshold int) {
	wallChan := make(chan wallet.Wallet)
	errChan := make(chan error)
	go func() {
		wall, err := wallet.NewWallet(currId, threshold)
		if err != nil {
			errChan <- err
			return
		}
		wallChan <- wall
	}()

	select {
	case wall := <-wallChan:
		bz, err := wall.Export()
		if err != nil {
			errChan <- err
			return
		}
		s.cache.Set(currId, bz, 15*time.Second)
	case err := <-errChan:
		panic(err)
	}
}
