package webapp

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/common/middleware/render"
	"github.com/onsonr/sonr/pkg/common/middleware/session"
	"github.com/onsonr/sonr/pkg/webapp/auth"
)

// AuthorizeModal returns the Authorize Modal route.
func AuthorizeModal(c echo.Context) error {
	cc, err := session.Get(c)
	if err != nil {
		return err
	}
	return render.Templ(c, auth.AuthorizeModal(cc))
}

// LoginModal returns the Login Modal route.
func LoginModal(c echo.Context) error {
	cc, err := session.Get(c)
	if err != nil {
		return err
	}
	return render.Templ(c, auth.LoginModal(cc))
}

// RegisterModal returns the Register Modal route.
func RegisterModal(c echo.Context) error {
	cc, err := session.Get(c)
	if err != nil {
		return err
	}
	return render.Templ(c, auth.RegisterModal(cc))
}
