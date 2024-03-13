package plugins

import (
	"github.com/labstack/echo/v4"

	"github.com/sonrhq/sonr/internal/prisma"
)

func UsePrisma(ctx echo.Context) *prisma.Client {
	return prisma.NewClient()
}
