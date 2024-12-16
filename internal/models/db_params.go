package models

import (
	ctx "github.com/onsonr/sonr/internal/context"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"
	"github.com/medama-io/go-useragent"
	"github.com/onsonr/sonr/internal/models/repository"
	"github.com/segmentio/ksuid"
	"golang.org/x/exp/rand"
)

func BaseSessionCreateParams(e echo.Context) repository.CreateSessionParams {
	f := rand.Intn(5) + 1
	l := rand.Intn(4) + 1
	challenge, _ := protocol.CreateChallenge()
	id := getOrCreateSessionID(e)
	ua := useragent.NewParser()
	s := ua.Parse(e.Request().UserAgent())

	return repository.CreateSessionParams{
		ID:             id,
		BrowserName:    s.GetBrowser(),
		BrowserVersion: s.GetMajorVersion(),
		ClientIpaddr:   e.RealIP(),
		Platform:       s.GetOS(),
		IsMobile:       boolToInt64(s.IsMobile()),
		IsTablet:       boolToInt64(s.IsTablet()),
		IsDesktop:      boolToInt64(s.IsDesktop()),
		IsBot:          boolToInt64(s.IsBot()),
		IsTv:           boolToInt64(s.IsTV()),
		IsHumanFirst:   int64(f),
		IsHumanLast:    int64(l),
		Challenge:      challenge.String(),
	}
}

func getOrCreateSessionID(c echo.Context) string {
	if ok := ctx.CookieExists(c, ctx.SessionID); !ok {
		sessionID := ksuid.New().String()
		ctx.WriteCookie(c, ctx.SessionID, sessionID)
		return sessionID
	}

	sessionID, err := ctx.ReadCookie(c, ctx.SessionID)
	if err != nil {
		sessionID = ksuid.New().String()
		ctx.WriteCookie(c, ctx.SessionID, sessionID)
	}
	return sessionID
}

func boolToInt64(b bool) int64 {
	if b {
		return 1
	}
	return 0
}
