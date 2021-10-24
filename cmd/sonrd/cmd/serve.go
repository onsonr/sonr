package cmd

import (
	"encoding/base64"
	"os"
	"strings"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/app"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/node"
	"github.com/sonr-io/core/pkg/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/encoding/protojson"
)

var hostPtr string
var portPtr int
var profilePtr string
var varsPtr string

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves the Sonr Daemon on localhost",
	Run: func(cmd *cobra.Command, args []string) {
		// Set Enviornment variables
		if varsPtr != "" {
			// Decode base64 encoded string
			keyValuesBuf, err := base64.StdEncoding.DecodeString(varsPtr)
			if err != nil {
				golog.Error("Failed to decode Enviornment Vars from Config")
				return
			}

			// Split String Values
			keyValuePairs := strings.Split(string(keyValuesBuf), ",")
			golog.Debugf("Updating %v Enviornment variables.", len(keyValuePairs))
			// Iterate over keyValuePairs
			for _, v := range keyValuePairs {
				// Trim White Space
				tv := strings.TrimSpace(v)

				// Split Key Value Pairs
				value := strings.Split(tv, "=")
				if len(value) != 2 {
					golog.Fatal("Invalid Enviornment Variable Format")
				}

				// Set Env Variables
				os.Setenv(value[0], value[1])
			}
		}

		// Set Location
		req := api.DefaultInitializeRequest()
		// Set Profile
		if profilePtr != "" {
			golog.Debug("Setting Profile from JSON.")
			p := &common.Profile{}

			// Unmarshal JSON String
			err := protojson.Unmarshal([]byte(profilePtr), p)
			if err == nil {
				req.Profile = p
			} else {
				golog.Warn("Failed to set Profile from flag")
			}
		}
		// Set Log Level
		logLevel := app.InfoLevel
		if isDebug {
			logLevel = app.DebugLevel
		}

		// Initialize App
		app.Start(req, app.WithMode(node.StubMode_BIN),
			app.WithHost(hostPtr),
			app.WithPort(portPtr),
			app.WithLogLevel(logLevel),
		)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVar(&hostPtr, "host", ":", "host address for Node -- Default: localhost")
	serveCmd.Flags().IntVar(&portPtr, "port", 26225, "port for RPC NodeStub Service -- Default: 26225")
	serveCmd.Flags().StringVar(&profilePtr, "profile", "", "profile JSON string")
	serveCmd.Flags().StringVar(&varsPtr, "vars", "", "enviornment variables encoded as base64")
	viper.BindPFlags(serveCmd.Flags())
}
