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
	"github.com/kataras/golog"
	"github.com/sonr-io/core/app"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/pkg/common"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
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
				golog.Child("[Sonr.daemon] ").Warn("Failed to set Profile from flag")
			}
		}
		app.Start(req, cliPtr, "full")
		app.Persist()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
