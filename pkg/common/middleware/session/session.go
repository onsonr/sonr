package session

import (
	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/common"
	"github.com/onsonr/sonr/pkg/common/types"
)

// HTTPContext is the context for DWN endpoints.
type HTTPContext struct {
	echo.Context

	role   common.PeerRole
	client *types.ClientConfig
	peer   *types.PeerInfo
	user   *types.UserAgent
	vault  *types.VaultDetails
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

	if ok := cc.role.Is(common.RoleMotr); ok {
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

	if ok := cc.role.Is(common.RoleMotr); ok {
		cc.vault, err = extractConfigVault(cc.Context)
		if err != nil {
			cc.Logger().Error(err)
		}
	}
	return cc
}

// ID returns the ksuid http cookie.
func (s *HTTPContext) ID() string {
	return s.peer.Id
}

func (s *HTTPContext) LoginOptions(credentials []common.CredDescriptor) *common.LoginOptions {
	ch, _ := common.Base64Decode(s.peer.Challenge)
	return &common.LoginOptions{
		Challenge:          ch,
		Timeout:            10000,
		AllowedCredentials: credentials,
	}
}

func (s *HTTPContext) RegisterOptions(subject string) *common.RegisterOptions {
	ch, _ := common.Base64Decode(s.peer.Challenge)
	opts := baseRegisterOptions()
	opts.Challenge = ch
	opts.User = buildUserEntity(subject)
	return opts
}

// Address returns the sonr address from the cookies.
func (s *HTTPContext) ClientConfig() *types.ClientConfig {
	return s.client
}

// IPFSGateway returns the IPFS gateway URL from the headers.
func (s *HTTPContext) UserAgent() *types.UserAgent {
	return s.user
}

// ChainID returns the chain ID from the headers.
func (s *HTTPContext) VaultDetails() *types.VaultDetails {
	return s.vault
}
