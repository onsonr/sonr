package session

import (
	"net/http"
	"time"

	"github.com/donseba/go-htmx"
	"github.com/labstack/echo/v4"
	"github.com/your-project/internal/headers"
)

type Session struct {
	echo.Context
	htmx *htmx.HTMX
	RequestHeaders  headers.RequestHeaders
	ResponseHeaders headers.ResponseHeaders
}

func (s *Session) Htmx() *htmx.HTMX {
	return s.htmx
}

func (s *Session) ID() string {
	return ReadCookie(s, "session")
}

func (s *Session) UpdateHtmxContext() {
	s.htmx.Request.Boosted = s.RequestHeaders.HXBoosted != nil && *s.RequestHeaders.HXBoosted == "true"
	s.htmx.Request.CurrentURL = s.RequestHeaders.HXCurrentURL
	s.htmx.Request.HistoryRestoreRequest = s.RequestHeaders.HXHistoryRestoreRequest != nil && *s.RequestHeaders.HXHistoryRestoreRequest == "true"
	s.htmx.Request.Prompt = s.RequestHeaders.HXPrompt
	s.htmx.Request.Target = s.RequestHeaders.HXTarget
	s.htmx.Request.TriggerName = s.RequestHeaders.HXTriggerName
	s.htmx.Request.Trigger = s.RequestHeaders.HXTrigger
}

func (s *Session) SetHtmxResponseHeaders() {
	if s.ResponseHeaders.HXLocation != nil {
		s.Response().Header().Set("HX-Location", *s.ResponseHeaders.HXLocation)
	}
	if s.ResponseHeaders.HXPushURL != nil {
		s.Response().Header().Set("HX-Push-Url", *s.ResponseHeaders.HXPushURL)
	}
	if s.ResponseHeaders.HXRedirect != nil {
		s.Response().Header().Set("HX-Redirect", *s.ResponseHeaders.HXRedirect)
	}
	if s.ResponseHeaders.HXRefresh != nil {
		s.Response().Header().Set("HX-Refresh", *s.ResponseHeaders.HXRefresh)
	}
	if s.ResponseHeaders.HXReplaceURL != nil {
		s.Response().Header().Set("HX-Replace-Url", *s.ResponseHeaders.HXReplaceURL)
	}
	if s.ResponseHeaders.HXReswap != nil {
		s.Response().Header().Set("HX-Reswap", *s.ResponseHeaders.HXReswap)
	}
	if s.ResponseHeaders.HXRetarget != nil {
		s.Response().Header().Set("HX-Retarget", *s.ResponseHeaders.HXRetarget)
	}
	if s.ResponseHeaders.HXReselect != nil {
		s.Response().Header().Set("HX-Reselect", *s.ResponseHeaders.HXReselect)
	}
	if s.ResponseHeaders.HXTrigger != nil {
		s.Response().Header().Set("HX-Trigger", *s.ResponseHeaders.HXTrigger)
	}
	if s.ResponseHeaders.HXTriggerAfterSettle != nil {
		s.Response().Header().Set("HX-Trigger-After-Settle", *s.ResponseHeaders.HXTriggerAfterSettle)
	}
	if s.ResponseHeaders.HXTriggerAfterSwap != nil {
		s.Response().Header().Set("HX-Trigger-After-Swap", *s.ResponseHeaders.HXTriggerAfterSwap)
	}
}

func ReadCookie(c echo.Context, key string) string {
	cookie, err := c.Cookie(key)
	if err != nil {
		return ""
	}
	if cookie == nil {
		return ""
	}
	return cookie.Value
}

func WriteCookie(c echo.Context, key string, value string) {
	cookie := new(http.Cookie)
	cookie.Name = key
	cookie.Value = value
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
}

func UseSession() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			session := &Session{
				Context: c,
				htmx:    htmx.New(req, res),
			}

			// Bind request headers
			if err := c.Bind(&session.RequestHeaders); err != nil {
				return err
			}

			// Update HTMX context
			session.UpdateHtmxContext()

			// Set custom context
			c.Set("session", session)

			// Call the next handler
			if err := next(session); err != nil {
				c.Error(err)
			}

			// Set HTMX response headers
			session.SetHtmxResponseHeaders()

			return nil
		}
	}
}
