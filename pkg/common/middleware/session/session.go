package session

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/common"
	"github.com/onsonr/sonr/pkg/common/middleware/cookie"
	"github.com/onsonr/sonr/pkg/common/middleware/header"
	"github.com/onsonr/sonr/pkg/common/types"
)

// HTTPContext is the context for DWN endpoints.
type HTTPContext struct {
	echo.Context
	ctx    context.Context
	role   common.PeerRole
	client *types.ClientConfig
	peer   *types.PeerInfo
	user   *types.UserAgent
	vault  *types.VaultDetails
}

// Ensure HTTPContext implements context.Context
func (s *HTTPContext) Deadline() (deadline time.Time, ok bool) {
	return s.ctx.Deadline()
}

func (s *HTTPContext) Done() <-chan struct{} {
	return s.ctx.Done()
}

func (s *HTTPContext) Err() error {
	return s.ctx.Err()
}

func (s *HTTPContext) Value(key interface{}) interface{} {
	return s.ctx.Value(key)
}

// initHTTPContext loads the headers from the request.
func initHTTPContext(c echo.Context) *HTTPContext {
	var err error
	pi := extractPeerInfo(c)
	clc := extractConfigClient(c)
	ua := extractUserAgent(c)

	// Create a new context with the base context and the session ID
	baseCtx := c.Request().Context()
	baseCtx = WithSessionID(baseCtx, pi.Id)
	baseCtx = WithHasAuthorization(baseCtx, header.Exists(c, header.Authorization))
	baseCtx = WithIsMobile(baseCtx, ua.IsMobile)
	baseCtx = WithChainID(baseCtx, clc.ChainID)
	baseCtx = WithIPFSHost(baseCtx, clc.IpfsHost)
	baseCtx = WithSonrAPI(baseCtx, clc.SonrAPIURL)
	baseCtx = WithSonrRPC(baseCtx, clc.SonrRPCURL)
	baseCtx = WithSonrWS(baseCtx, clc.SonrWSURL)

	// Add the user handle to the context if it exists
	if ok := cookie.Exists(c, cookie.UserHandle); ok {
		uh, err := cookie.Read(c, cookie.UserHandle)
		if err != nil {
			c.Logger().Error(err)
		}
		baseCtx = WithUserHandle(baseCtx, uh)
	}

	// Add the user to the context if it exists
	cc := &HTTPContext{
		Context: c,
		ctx:     baseCtx,
		role:    extractPeerRole(c),
		client:  clc,
		peer:    pi,
		user:    ua,
	}

	// Add the vault to the context if it exists
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

// HasHandle returns true if the user has a handle.
func (s *HTTPContext) HasHandle() bool {
	return cookie.Exists(s, cookie.UserHandle)
}

// ID returns the ksuid http cookie.
func (s *HTTPContext) ID() string {
	return s.peer.Id
}

// IsAuthenticated returns true if the user is authenticated.
func (s *HTTPContext) IsAuthenticated() bool {
	return header.Exists(s, header.Authorization)
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
