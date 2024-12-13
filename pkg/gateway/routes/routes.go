package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/pkg/gateway/handlers"
)

func Register(e *echo.Echo) error {
	// Register View Handlers
	e.GET("/", handlers.HandleIndex)
	e.GET("/register", handlers.RenderProfileCreate)
	e.POST("/register/passkey", handlers.RenderPasskeyCreate)
	e.POST("/register/finish", handlers.RenderVaultLoading)

	// Register Validation Handlers
	e.POST("/register/profile/handle", handlers.CheckProfileHandle)
	e.POST("/register/profile/is_human", handlers.CheckIsHumanSum)
	e.POST("/submit/profile/handle", handlers.SubmitProfileHandle)
	e.POST("/submit/credential", handlers.SubmitPublicKeyCredential)
	return nil
}
