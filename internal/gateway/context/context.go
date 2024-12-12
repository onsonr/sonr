package context

import (
	"github.com/medama-io/go-useragent"
	"github.com/onsonr/sonr/pkg/common"
	"github.com/onsonr/sonr/pkg/database/sessions"
	"github.com/segmentio/ksuid"
)

// InitSession initializes or loads an existing session
func (s *HTTPContext) InitSession(agent useragent.UserAgent) error {
	sessionID := s.getOrCreateSessionID()

	// Try to load existing session
	var sess sessions.Session
	result := s.db.Where("id = ?", sessionID).First(&sess)
	if result.Error != nil {
		// Create new session if not found
		sess = sessions.Session{
			ID:             sessionID,
			BrowserName:    agent.GetBrowser(),
			BrowserVersion: agent.GetMajorVersion(),
			Platform:       agent.GetOS(),
			IsMobile:       agent.IsMobile(),
			IsTablet:       agent.IsTablet(),
			IsDesktop:      agent.IsDesktop(),
			IsBot:          agent.IsBot(),
			IsTV:           agent.IsTV(),
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
