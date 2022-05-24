package highwaycmd

import (
	"context"

	"github.com/kataras/golog"
	"github.com/sonr-io/sonr/internal/highway"
	"github.com/sonr-io/sonr/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func bootstrapServeCommand(ctx context.Context) (serveCmd *cobra.Command) {
	// env
	var (
		RP_SERVER_PORT = viper.GetString("RP_SERVER_PORT")
		GRPC_PORT      = viper.GetInt("GRPC_PORT")
		HTTP_PORT      = viper.GetInt("HTTP_PORT")
	)

	logger := golog.Default.Child("serve")

	// ServeCmd represents the serve command
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Starts the highwayd and launches frontend in browser.",

		Run: func(cmd *cobra.Command, args []string) {
			logger.Infof("Serving new highway instance")
			node, err := highway.NewHighway(
				ctx,
				config.WithHighwayAPISettings("tcp", "localhost", GRPC_PORT, HTTP_PORT),
				nil)
			cobra.CheckErr(err)

			node.Serve()
			logger.Info("Server started at ", RP_SERVER_PORT)
		},
	}

	return
}
