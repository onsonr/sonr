package context

import (
	gocontext "context"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/labstack/echo/v4"
	"github.com/medama-io/go-useragent"
	hwayorm "github.com/onsonr/sonr/internal/database/hwayorm"
	"github.com/onsonr/sonr/pkg/common"
	"github.com/segmentio/ksuid"
)

func NewSession(c echo.Context) error {
	cc, ok := c.(*GatewayContext)
	if !ok {
		return nil
	}
	baseSessionCreateParams := BaseSessionCreateParams(cc)
	cc.id = baseSessionCreateParams.ID
	if _, err := cc.CreateSession(bgCtx(), baseSessionCreateParams); err != nil {
		return err
	}
	// Set Cookie
	if err := common.WriteCookie(c, common.SessionID, cc.id); err != nil {
		return err
	}
	return nil
}

// ForbiddenDevice returns true if the device is unavailable
func ForbiddenDevice(c echo.Context) bool {
	cc, ok := c.(*GatewayContext)
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
	cc, ok := c.(*GatewayContext)
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
	cc, ok := c.(*GatewayContext)
	if !ok {
		return ""
	}
	s, err := cc.GetChallengeBySessionID(bgCtx(), cc.id)
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
	cc, ok := c.(*GatewayContext)
	if !ok {
		return ""
	}
	s, err := cc.GetSessionByID(bgCtx(), cc.id)
	if err != nil {
		return ""
	}
	profile, err := cc.GetProfileByID(bgCtx(), s.ProfileID)
	if err != nil {
		return ""
	}
	return profile.Handle
}

//
// func GetHumanVerificationNumbers(c echo.Context) (int64, int64) {
// 	cc, ok := c.(*GatewayContext)
// 	if !ok {
// 		return 0, 0
// 	}
// 	s, err := cc.dbq.GetHumanVerificationNumbers(bgCtx(), cc.id)
// 	if err != nil {
// 		return 0, 0
// 	}
// 	return s.IsHumanFirst, s.IsHumanLast
// }

// utility function to get a context
func bgCtx() gocontext.Context {
	ctx := gocontext.Background()
	return ctx
}

func BaseSessionCreateParams(e echo.Context) hwayorm.CreateSessionParams {
	// f := rand.Intn(5) + 1
	// l := rand.Intn(4) + 1
	challenge, _ := protocol.CreateChallenge()
	id := getOrCreateSessionID(e)
	ua := useragent.NewParser()
	s := ua.Parse(e.Request().UserAgent())

	return hwayorm.CreateSessionParams{
		ID:             id,
		BrowserName:    s.GetBrowser(),
		BrowserVersion: s.GetMajorVersion(),
		ClientIpaddr:   e.RealIP(),
		Platform:       s.GetOS(),
		IsMobile:       s.IsMobile(),
		IsTablet:       s.IsTablet(),
		IsDesktop:      s.IsDesktop(),
		IsBot:          s.IsBot(),
		IsTv:           s.IsTV(),
		// IsHumanFirst:   int64(f),
		// IsHumanLast:    int64(l),
		Challenge: challenge.String(),
	}
}

func getOrCreateSessionID(c echo.Context) string {
	if ok := common.CookieExists(c, common.SessionID); !ok {
		sessionID := ksuid.New().String()
		common.WriteCookie(c, common.SessionID, sessionID)
		return sessionID
	}

	sessionID, err := common.ReadCookie(c, common.SessionID)
	if err != nil {
		sessionID = ksuid.New().String()
		common.WriteCookie(c, common.SessionID, sessionID)
	}
	return sessionID
}

func boolToInt64(b bool) int64 {
	if b {
		return 1
	}
	return 0
}
