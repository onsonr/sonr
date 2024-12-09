package session

import (
	"regexp"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/app/gateway/internal/database"
	"github.com/onsonr/sonr/pkg/common"
	"github.com/segmentio/ksuid"
)

// InitSession initializes or loads an existing session
func (s *HTTPContext) InitSession() error {
	sessionID := s.getOrCreateSessionID()

	// Try to load existing session
	var sess database.Session
	result := s.db.Where("id = ?", sessionID).First(&sess)
	if result.Error != nil {
		// Create new session if not found
		bn, bv, arch, plat, platVer, model := extractBrowserInfo(s.Context)
		sess = database.Session{
			ID:               sessionID,
			BrowserName:      bn,
			BrowserVersion:   bv,
			UserArchitecture: arch,
			Platform:         plat,
			PlatformVersion:  platVer,
			DeviceModel:      model,
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

func extractBrowserInfo(c echo.Context) (string, string, string, string, string, string) {
	// Extract all relevant headers
	browserName := common.HeaderRead(c, common.UserAgent)
	arch := common.HeaderRead(c, common.Architecture)
	platform := common.HeaderRead(c, common.Platform)
	platformVer := common.HeaderRead(c, common.PlatformVersion)
	model := common.HeaderRead(c, common.Model)
	fullVersionList := common.HeaderRead(c, common.FullVersionList)

	// Default values if headers are empty
	if browserName == "" {
		browserName = "N/A"
	}
	if arch == "" {
		arch = "unknown"
	}
	if platform == "" {
		platform = "unknown"
	}
	if platformVer == "" {
		platformVer = "unknown"
	}
	if model == "" {
		model = "unknown"
	}

	// Extract browser version from full version list
	version := "-1"
	if fullVersionList != "" {
		entries := strings.Split(strings.TrimSpace(fullVersionList), ",")
		for _, entry := range entries {
			entry = strings.TrimSpace(entry)
			re := regexp.MustCompile(`"([^"]+)";v="([^"]+)"`)
			matches := re.FindStringSubmatch(entry)

			if len(matches) == 3 {
				browserName = matches[1]
				version = matches[2]
				if browserName != "Not.A/Brand" &&
					browserName != "Chromium" {
					break
				}
			}
		}
	}

	return browserName, version, arch, platform, platformVer, model
}
