package routes

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/nebula/pages"
)

type Route string

const (
	HomeRoute      Route = "/home"
	LoginRoute     Route = "/login"
	RegisterRoute  Route = "/register"
	ProfileRoute   Route = "/profile"
	AuthorizeRoute Route = "/authorize"
)

func (r Route) Route() string {
	return string(r)
}

func (r Route) Component(c echo.Context) templ.Component {
	switch r {
	case HomeRoute:
		return pages.Home(c)
	case LoginRoute:
		return pages.Login(c)
	case RegisterRoute:
		return pages.Register(c)
	case ProfileRoute:
		return pages.Profile(c)
	case AuthorizeRoute:
		return pages.Authorize(c)
	}
	return nil
}
