package main

import (
	_ "embed"

	"github.com/labstack/echo/v4"

	"github.com/labstack/echo/v4/middleware"
	"github.com/onsonr/sonr/pkg/common/middleware/response"
	"github.com/onsonr/sonr/pkg/common/middleware/session"
	"github.com/onsonr/sonr/web/landing/pages/home"
	"github.com/onsonr/sonr/web/vault/pages/login"
	"github.com/onsonr/sonr/web/vault/pages/register"
)

type (
	Host struct {
		Echo *echo.Echo
	}
)

func main() {
	// Setup Echo
	hosts := map[string]*Host{}

	//---------
	// Gateway
	//---------
	gateway := echo.New()
	gateway.Use(middleware.Logger())
	gateway.Use(middleware.Recover())
	gateway.Use(session.MotrMiddleware(nil))
	hosts["*.localhost:3000"] = &Host{Echo: gateway}
	gateway.GET("/*", response.Templ(register.Page()))

	//---------
	// Website
	//---------
	site := echo.New()
	site.Use(middleware.Logger())
	site.Use(middleware.Recover())
	site.Use(session.HwayMiddleware())
	hosts["localhost:3000"] = &Host{Echo: site}
	site.GET("/", response.Templ(home.Page()))
	site.GET("/register", response.Templ(register.Page()))
	site.GET("/login", response.Templ(login.Page()))

	// Server
	e := echo.New()
	e.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()
		host := hosts[req.Host]

		if host == nil {
			err = echo.ErrNotFound
		} else {
			host.Echo.ServeHTTP(res, req)
		}

		return
	})
	e.Logger.Fatal(e.Start(":3000"))
}
