package highwaycmd

import (
	"context"

	"github.com/kataras/golog"
	"github.com/sonr-io/sonr/pkg/highway"
	"github.com/sonr-io/sonr/pkg/highway/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func bootstrapServeCommand(ctx context.Context) (serveCmd *cobra.Command) {
	// env
	var (
		RP_SERVER_PORT = viper.GetString("RP_SERVER_PORT")
		GRPC_PORT      = viper.GetInt("GRPC_PORT")
		HTTP_PORT      = viper.GetInt("HTTP_PORT")
		DISPLAY_NAME   = viper.GetString("DISPLAY_NAME")
		RP_ID          = viper.GetString("RP_ID")
		RP_AUTH_ORIGIN = viper.GetString("RP_AUTH_ORIGIN")
		IS_DEBUG       = viper.GetBool("IS_DEBUG")
	)

	logger := golog.Default.Child("serve")

	// ServeCmd represents the serve command
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Starts the highwayd and launches frontend in browser.",

		Run: func(cmd *cobra.Command, args []string) {
			node, err := highway.NewHighway(
				ctx,
				config.WithHighwayAPISettings("tcp", "localhost", GRPC_PORT, HTTP_PORT),
				config.WithWebAuthnConfig(DISPLAY_NAME, RP_ID, RP_AUTH_ORIGIN, IS_DEBUG))
			cobra.CheckErr(err)

			node.Serve()
			logger.Info("Server started at ", RP_SERVER_PORT)
		},
	}

	return
}
