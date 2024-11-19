package session

import (
	"github.com/labstack/echo/v4"

	commonv1 "github.com/onsonr/sonr/pkg/common/types"
)

// HTTPContext is the context for DWN endpoints.
type HTTPContext struct {
	echo.Context

	role   commonv1.PeerRole
	client *commonv1.ClientConfig
	peer   *commonv1.PeerInfo
	user   *commonv1.UserAgent
	vault  *commonv1.VaultDetails
}

// initHTTPContext loads the headers from the request.
func initHTTPContext(c echo.Context) *HTTPContext {
	var err error
	cc := &HTTPContext{
		Context: c,
		role:    extractPeerRole(c),
		client:  extractConfigClient(c),
		peer:    extractPeerInfo(c),
		user:    extractUserAgent(c),
	}

	if ok := cc.role.Is(commonv1.RoleMotr); ok {
		cc.vault, err = extractConfigVault(c)
		if err != nil {
			c.Logger().Error(err)
		}
	}
	return cc
}

// loadHTTPContext loads the headers into an existing context
func loadHTTPContext(cc *HTTPContext) *HTTPContext {
	var err error

	cc.role = extractPeerRole(cc.Context)
	cc.client = extractConfigClient(cc.Context)
	cc.peer = extractPeerInfo(cc.Context)
	cc.user = extractUserAgent(cc.Context)

	if ok := cc.role.Is(commonv1.RoleMotr); ok {
		cc.vault, err = extractConfigVault(cc.Context)
		if err != nil {
			cc.Logger().Error(err)
		}
	}
	return cc
}

// ID returns the ksuid http cookie.
func (s *HTTPContext) ID() string {
	return s.peer.ID
}

func (s *HTTPContext) LoginOptions(credentials []commonv1.CredDescriptor) *commonv1.LoginOptions {
	return &commonv1.LoginOptions{
		Challenge:          s.peer.Challenge,
		Timeout:            10000,
		AllowedCredentials: credentials,
	}
}

func (s *HTTPContext) RegisterOptions(subject string) *commonv1.RegisterOptions {
	opts := baseRegisterOptions()
	opts.Challenge = s.peer.Challenge
	opts.User = buildUserEntity(subject)
	return opts
}

// Address returns the sonr address from the cookies.
func (s *HTTPContext) ClientConfig() *commonv1.ClientConfig {
	return s.client
}

// IPFSGateway returns the IPFS gateway URL from the headers.
func (s *HTTPContext) UserAgent() *commonv1.UserAgent {
	return s.user
}

// ChainID returns the chain ID from the headers.
func (s *HTTPContext) VaultDetails() *commonv1.VaultDetails {
	return s.vault
}
