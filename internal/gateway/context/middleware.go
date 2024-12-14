package context

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/medama-io/go-useragent"
	"github.com/onsonr/sonr/internal/gateway/models"
	"github.com/onsonr/sonr/internal/gateway/providers"
	config "github.com/onsonr/sonr/pkg/config/hway"
	"github.com/onsonr/sonr/pkg/ipfsapi"
	"gorm.io/gorm"
)

// Middleware creates a new session middleware
func Middleware(db *gorm.DB, env config.Hway, ipc ipfsapi.Client) echo.MiddlewareFunc {
	ua := useragent.NewParser()
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			agent := ua.Parse(c.Request().UserAgent())
			cc := NewHTTPContext(c, db, agent, env.GetSonrGrpcUrl(), ipc)
			if err := cc.initSession(); err != nil {
				return err
			}
			return next(cc)
		}
	}
}

// HTTPContext is the context for HTTP endpoints.
type HTTPContext struct {
	echo.Context
	providers.VaultProvider
	providers.Resolver
	db   *gorm.DB
	sess *models.Session
	user *models.User
	env  config.Hway
	useragent.UserAgent
}

// Get returns the HTTPContext from the echo context
func Get(c echo.Context) (*HTTPContext, error) {
	ctx, ok := c.(*HTTPContext)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Session Context not found")
	}
	return ctx, nil
}

// NewHTTPContext creates a new session context
func NewHTTPContext(c echo.Context, db *gorm.DB, a useragent.UserAgent, grpcAddr string, ipc ipfsapi.Client) *HTTPContext {
	return &HTTPContext{
		Context:       c,
		db:            db,
		Resolver:      providers.NewResolverService(grpcAddr),
		UserAgent:     a,
		VaultProvider: providers.NewVaultService(ipc),
	}
}

// Session returns the current session
func (s *HTTPContext) Session() *models.Session {
	return s.sess
}
