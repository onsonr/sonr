package service

import (
	"encoding/base64"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/patrickmn/go-cache"
	"github.com/sonrhq/core/pkg/wallet"
	"lukechampine.com/blake3"
)

// VerifyChallenge verifies the challenge signature
func (s *serviceHandlerImpl) IssueChallenge(uuid string) (protocol.URLEncodedBase64, error) {
	// Marshal the service into JSON.
	sbz, err := s.service.Marshal()
	if err != nil {
		return nil, err
	}

	// Append the UUID to the service.
	sbz = append(sbz, uuid...)

	// Blake 3 hash the service.
	bz := blake3.Sum256(sbz)
	chal := base64.RawURLEncoding.EncodeToString(bz[:])
	return protocol.URLEncodedBase64(chal), nil
}

// VerifyChallenge verifies the challenge signature
func (s *serviceHandlerImpl) VerifyChallenge(pcc *protocol.ParsedCredentialCreationData, uuid string) error {
	// Marshal the service into JSON.
	sbz, err := s.service.Marshal()
	if err != nil {
		return err
	}

	// Append the UUID to the service.
	sbz = append(sbz, uuid...)

	// Blake 3 hash the service.
	bz := blake3.Sum256(sbz)
	raw := base64.RawURLEncoding.EncodeToString(bz[:])
	chal := protocol.URLEncodedBase64(raw)

	// Verify the challenge.
	err = pcc.Verify(chal.String(), false, s.service.Id, []string{s.service.Origin})
	if err != nil {
		return err
	}
	return nil
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
		s.cache.Set(currId, bz, cache.DefaultExpiration)
	case err := <-errChan:
		panic(err)
	}
}
