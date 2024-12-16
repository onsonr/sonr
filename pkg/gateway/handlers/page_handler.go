package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/gateway/middleware"
	"github.com/onsonr/sonr/internal/models"
	"github.com/onsonr/sonr/pkg/gateway/views"
)

func RenderIndex(c echo.Context) error {
	// Initialize the session
	err := middleware.NewSession(c)
	if err != nil {
		return middleware.RenderError(c, err)
	}
	// Render the initial view
	return middleware.RenderInitial(c)
}

func RenderProfileCreate(c echo.Context) error {
	numF, numL := middleware.GetHumanVerificationNumbers(c)
	params := models.CreateProfileParams{
		FirstNumber: int(numF),
		LastNumber:  int(numL),
	}
	return middleware.Render(c, views.RegisterProfileView(params))
}

func RenderPasskeyCreate(c echo.Context) error {
	return middleware.Render(c, views.RegisterPasskeyView(models.CreatePasskeyParams{}))
}

func RenderVaultLoading(c echo.Context) error {
	return middleware.Render(c, views.LoadingView())
}
