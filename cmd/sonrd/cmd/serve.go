/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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

var cliPtr bool
var latPtr float64
var lngPtr float64
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
			golog.Infof("Updating %v Enviornment variables.", len(keyValuePairs))
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
		req := &api.InitializeRequest{
			Location: &common.Location{
				Latitude:  latPtr,
				Longitude: lngPtr,
			},
			Profile: common.NewDefaultProfile(),
		}

		// Set Profile
		if profilePtr != "" {
			golog.Info("Setting Profile from JSON.")
			p := &common.Profile{}

			// Unmarshal JSON String
			err := protojson.Unmarshal([]byte(profilePtr), p)
			if err == nil {
				req.Profile = p
			} else {
				golog.Warn("Failed to set Profile from flag")
			}
		}
		if cliPtr {
			app.Start(req, node.StubMode_CLI)
		} else {
			app.Start(req, node.StubMode_BIN)
		}

	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().BoolVar(&cliPtr, "cli", false, "run in CLI Mode")
	serveCmd.Flags().Float64Var(&latPtr, "lat", 34.102920, "latitude for InitializeRequest")
	serveCmd.Flags().Float64Var(&lngPtr, "lng", -118.394190, "longitude for InitializeRequest")
	serveCmd.Flags().StringVar(&profilePtr, "profile", "", "profile JSON string")
	serveCmd.Flags().StringVar(&varsPtr, "vars", "", "enviornment variables encoded as base64")
	viper.BindPFlags(serveCmd.Flags())
}
