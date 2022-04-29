package highwaycmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func bootstrapRegisterCommand(ctx context.Context) (registerCmd *cobra.Command) {
	// env
	var (
		RP_ORIGIN = viper.GetString("RP_ORIGIN")
		HTTP_PORT = viper.GetInt("HTTP_PORT")
	)

	registerCmd = &cobra.Command{
		Use:   "register <name>",
		Short: "Register a domain on the Sonr network",

		Run: func(cmd *cobra.Command, args []string) {
      
		},
	}
	return
}

