package middleware

import (
	gocontext "context"
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/medama-io/go-useragent"
	"github.com/onsonr/sonr/internal/gateway/context"
	"github.com/onsonr/sonr/internal/gateway/models"
	"github.com/onsonr/sonr/internal/gateway/models/repository"
	"github.com/onsonr/sonr/pkg/common"
)

type SessionsContext struct {
	echo.Context
	dbq   *repository.Queries
	id    string
	agent useragent.UserAgent
}

func UseSessions(conn *sql.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ua := useragent.NewParser()
			agent := ua.Parse(c.Request().UserAgent())
			cc := &SessionsContext{dbq: repository.New(conn), Context: c, agent: agent}
			return next(cc)
		}
	}
}

func NewSession(c echo.Context) error {
	cc, ok := c.(*SessionsContext)
	if !ok {
		return nil
	}
	baseSessionCreateParams := models.BaseSessionCreateParams(cc)
	cc.id = baseSessionCreateParams.ID
	if _, err := cc.dbq.CreateSession(bgCtx(), baseSessionCreateParams); err != nil {
		return err
	}
	// Set Cookie
	if err := context.WriteCookie(c, context.SessionID, cc.id); err != nil {
		return err
	}
	return nil
}

// ForbiddenDevice returns true if the device is unavailable
func ForbiddenDevice(c echo.Context) bool {
	cc, ok := c.(*SessionsContext)
	if !ok {
		return true
	}
	return cc.agent.IsBot() || cc.agent.IsTV()
}

func GetOrigin(c echo.Context) string {
	return c.Request().Host
}

func GetSessionID(c echo.Context) string {
	// Check from context
	cc, ok := c.(*SessionsContext)
	if !ok {
		return ""
	}
	// check from cookie
	if cc.id == "" {
		if ok := common.CookieExists(c, common.SessionID); !ok {
			return ""
		}
		cc.id = common.ReadCookieUnsafe(c, common.SessionID)
	}
	return cc.id
}

func GetSessionChallenge(c echo.Context) string {
	cc, ok := c.(*SessionsContext)
	if !ok {
		return ""
	}
	s, err := cc.dbq.GetChallengeBySessionID(bgCtx(), cc.id)
	if err != nil {
		return ""
	}
	return s
}

func GetHandle(c echo.Context) string {
	// First check for the cookie
	handle := common.ReadCookieUnsafe(c, common.UserHandle)
	if handle != "" {
		return handle
	}

	// Then check the session
	cc, ok := c.(*SessionsContext)
	if !ok {
		return ""
	}
	s, err := cc.dbq.GetSessionByID(bgCtx(), cc.id)
	if err != nil {
		return ""
	}
	profile, err := cc.dbq.GetProfileByID(bgCtx(), s.ProfileID)
	if err != nil {
		return ""
	}
	return profile.Handle
}

func GetHumanVerificationNumbers(c echo.Context) (int64, int64) {
	cc, ok := c.(*SessionsContext)
	if !ok {
		return 0, 0
	}
	s, err := cc.dbq.GetHumanVerificationNumbers(bgCtx(), cc.id)
	if err != nil {
		return 0, 0
	}
	return s.IsHumanFirst, s.IsHumanLast
}

// utility function to get a context
func bgCtx() gocontext.Context {
	ctx := gocontext.Background()
	return ctx
}
