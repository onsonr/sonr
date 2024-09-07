package front

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

func NewServeFrontendCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "serve-web",
		Short: "TUI for managing the local Sonr validator node",
		Run: func(cmd *cobra.Command, args []string) {
			e := echo.New()
			if err := e.Start(":42069"); err != http.ErrServerClosed {
				log.Fatal(err)
			}
		},
	}
}
