package context

import (
	"context"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/gateway/context/repository"
	"github.com/onsonr/sonr/pkg/common"
	"github.com/segmentio/ksuid"
	"golang.org/x/exp/rand"
)

// isUnavailableDevice returns true if the device is unavailable
func IsUnavailableDevice(c echo.Context) bool {
	s, err := Get(c)
	if err != nil {
		return true
	}
	return s.IsBot() || s.IsTV()
}

// initSession initializes or loads an existing session
func (s *HTTPContext) initSession() error {
	sessionID := s.getOrCreateSessionID()
	f := rand.Intn(5) + 1
	l := rand.Intn(4) + 1
	challenge, err := protocol.CreateChallenge()
	if err != nil {
		return err
	}
	s.id = sessionID

	// Try to load existing session
	dbSession, err := s.GetSessionByID(context.Background(), sessionID)
	if err != nil {
		// Create new session if not found
		params := repository.CreateSessionParams{
			ID:             sessionID,
			BrowserName:    s.GetBrowser(),
			BrowserVersion: s.GetMajorVersion(),
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
		dbSession, err = s.CreateSession(context.Background(), params)
		if err != nil {
			return err
		}
	}

	s.sess = &dbSession
	return nil
}

func (s *HTTPContext) getOrCreateSessionID() string {
	if ok := common.CookieExists(s.Context, common.SessionID); !ok {
		sessionID := ksuid.New().String()
		common.WriteCookie(s.Context, common.SessionID, sessionID)
		return sessionID
	}

	sessionID, err := common.ReadCookie(s.Context, common.SessionID)
	if err != nil {
		sessionID = ksuid.New().String()
		common.WriteCookie(s.Context, common.SessionID, sessionID)
	}
	return sessionID
}

func boolToInt64(b bool) int64 {
	if b {
		return 1
	}
	return 0
}
