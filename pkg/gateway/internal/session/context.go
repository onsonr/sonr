package session

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common"
	"github.com/onsonr/sonr/pkg/gateway/internal/database"
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

// Get returns the HTTPContext from the echo context
func Get(c echo.Context) (*HTTPContext, error) {
	ctx, ok := c.(*HTTPContext)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Session Context not found")
	}
	return ctx, nil
}

// HTTPContext is the context for HTTP endpoints.
type HTTPContext struct {
	echo.Context
	db   *gorm.DB
	sess *database.Session
}

// NewHTTPContext creates a new session context
func NewHTTPContext(c echo.Context, db *gorm.DB) *HTTPContext {
	return &HTTPContext{
		Context: c,
		db:      db,
	}
}

// Session returns the current session
func (s *HTTPContext) Session() *database.Session {
	return s.sess
}

// InitSession initializes or loads an existing session
func (s *HTTPContext) InitSession() error {
	sessionID := s.getOrCreateSessionID()

	// Try to load existing session
	var sess database.Session
	result := s.db.Where("id = ?", sessionID).First(&sess)
	if result.Error != nil {
		// Create new session if not found
		bn, bv := extractBrowserInfo(s.Context)
		sess = database.Session{
			ID:             sessionID,
			BrowserName:    bn,
			BrowserVersion: bv,
		}
		if err := s.db.Create(&sess).Error; err != nil {
			return err
		}
	}

	s.sess = &sess
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

func extractBrowserInfo(c echo.Context) (string, string) {
	userAgent := common.HeaderRead(c, common.UserAgent)
	if userAgent == "" {
		return "N/A", "-1"
	}

	var name, ver string
	entries := strings.Split(strings.TrimSpace(userAgent), ",")
	for _, entry := range entries {
		entry = strings.TrimSpace(entry)
		re := regexp.MustCompile(`"([^"]+)";v="([^"]+)"`)
		matches := re.FindStringSubmatch(entry)

		if len(matches) == 3 {
			browserName := matches[1]
			version := matches[2]

			if browserName != common.BrowserNameUnknown.String() &&
				browserName != common.BrowserNameChromium.String() {
				name = browserName
				ver = version
				break
			}
		}
	}
	return name, ver
}
