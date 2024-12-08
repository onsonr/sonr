package session

import (
	"regexp"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common"
	"gorm.io/gorm"
)

// Middleware creates a new session middleware
func Middleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := NewHTTPContext(c, db)
			if err := cc.InitSession(); err != nil {
				return err
			}
			return next(cc)
		}
	}
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
