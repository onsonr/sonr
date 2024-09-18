package dwn

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/internal/dwn/front"
	"github.com/onsonr/sonr/internal/dwn/handlers"
	"github.com/onsonr/sonr/internal/dwn/middleware"
)

// NewServer creates a new dwn server using echo
func NewServer() *echo.Echo {
	e := echo.New()
	e.Use(middleware.UseSession)
	front.RegisterViews(e)
	handlers.RegisterState(e)
	return e
}
