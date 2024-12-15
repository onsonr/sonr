package middleware

import (
	"context"
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/models"
	"github.com/onsonr/sonr/internal/gateway/models/repository"
	"github.com/onsonr/sonr/pkg/common"
)

type SessionsContext struct {
	echo.Context
	dbq *repository.Queries
	id  string
}

func UseSessions(conn *sql.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &SessionsContext{dbq: repository.New(conn), Context: c}
			baseSessionCreateParams := models.BaseSessionCreateParams(cc)
			cc.id = baseSessionCreateParams.ID
			if _, err := cc.dbq.CreateSession(bgCtx(), baseSessionCreateParams); err != nil {
				return err
			}
			// Set Cookie
			if err := common.WriteCookie(c, common.SessionID, cc.id); err != nil {
				return err
			}
			return next(cc)
		}
	}
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
func bgCtx() context.Context {
	ctx := context.Background()
	return ctx
}
