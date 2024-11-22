package session

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/onsonr/sonr/pkg/common"
)

type contextKey string

// Context keys
const (
	SessionIDKey        contextKey = "session_id"
	HasAuthorizationKey contextKey = "has_authorization"
	UserHandleKey       contextKey = "user_handle"
	IsMobileKey         contextKey = "is_mobile"
	ChainIDKey          contextKey = "chain_id"
	IPFSHostKey         contextKey = "ipfs_host"
	SonrAPIKey          contextKey = "sonr_api"
	SonrRPCKey          contextKey = "sonr_rpc"
	SonrWSKey           contextKey = "sonr_ws"
)

type Context = common.SessionCtx

// Get returns the session.Context from the echo context.
func Get(c echo.Context) (Context, error) {
	ctx, ok := c.(*HTTPContext)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "DWN Context not found")
	}
	return loadHTTPContext(ctx), nil
}

// WithSessionID sets the session ID in the context
func WithSessionID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, SessionIDKey, id)
}

// WithHasAuthorization sets the has authorization in the context
func WithHasAuthorization(ctx context.Context, has bool) context.Context {
	return context.WithValue(ctx, HasAuthorizationKey, has)
}

// WithUserHandle sets the user handle in the context
func WithUserHandle(ctx context.Context, handle string) context.Context {
	return context.WithValue(ctx, UserHandleKey, handle)
}

// WithIsMobile sets the is mobile in the context
func WithIsMobile(ctx context.Context, is bool) context.Context {
	return context.WithValue(ctx, IsMobileKey, is)
}

func WithChainID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, ChainIDKey, id)
}

func WithIPFSHost(ctx context.Context, host string) context.Context {
	return context.WithValue(ctx, IPFSHostKey, host)
}

func WithSonrAPI(ctx context.Context, api string) context.Context {
	return context.WithValue(ctx, SonrAPIKey, api)
}

func WithSonrRPC(ctx context.Context, rpc string) context.Context {
	return context.WithValue(ctx, SonrRPCKey, rpc)
}

func WithSonrWS(ctx context.Context, ws string) context.Context {
	return context.WithValue(ctx, SonrWSKey, ws)
}

// GetSessionID gets the session ID from the context
func GetSessionID(ctx context.Context) string {
	if id, ok := ctx.Value(SessionIDKey).(string); ok {
		return id
	}
	return ""
}

// GetHasAuthorization gets the has authorization from the context
func GetHasAuthorization(ctx context.Context) bool {
	if has, ok := ctx.Value(HasAuthorizationKey).(bool); ok {
		return has
	}
	return false
}

// GetUserHandle gets the user handle from the context
func GetUserHandle(ctx context.Context) string {
	if handle, ok := ctx.Value(UserHandleKey).(string); ok {
		return handle
	}
	return ""
}

// GetIsMobile gets the is mobile from the context
func GetIsMobile(ctx context.Context) bool {
	if is, ok := ctx.Value(IsMobileKey).(bool); ok {
		return is
	}
	return false
}

// GetChainID gets the chain ID from the context
func GetChainID(ctx context.Context) string {
	if id, ok := ctx.Value(ChainIDKey).(string); ok {
		return id
	}
	return ""
}

// GetIPFSHost gets the IPFS host from the context
func GetIPFSHost(ctx context.Context) string {
	if host, ok := ctx.Value(IPFSHostKey).(string); ok {
		return host
	}
	return ""
}

// GetSonrAPI gets the Sonr API from the context
func GetSonrAPI(ctx context.Context) string {
	if api, ok := ctx.Value(SonrAPIKey).(string); ok {
		return api
	}
	return ""
}

// GetSonrRPC gets the Sonr RPC from the context
func GetSonrRPC(ctx context.Context) string {
	if rpc, ok := ctx.Value(SonrRPCKey).(string); ok {
		return rpc
	}
	return ""
}

// GetSonrWS gets the Sonr WS from the context
func GetSonrWS(ctx context.Context) string {
	if ws, ok := ctx.Value(SonrWSKey).(string); ok {
		return ws
	}
	return ""
}
