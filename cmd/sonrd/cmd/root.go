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
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/app"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/pkg/common"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/spf13/viper"
)

var cliPtr bool
var latPtr float64
var lngPtr float64
var profilePtr string
var varsPtr string
var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sonrd",
	Short: "Daemon for Sonr Binary, interact with Node through this CLI.",
	Long:  `Sonr's Core Framework manages Discovery, Connection, Data-Transfer, Authorization/Authentication and Peer Lobby.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Set Enviornment variables
		if varsPtr != "" {
			// Decode base64 encoded string
			keyValuesBuf, err := base64.StdEncoding.DecodeString(varsPtr)
			if err != nil {
				golog.Child("[Sonr.daemon] ").Error("Failed to decode Enviornment Vars from Config")
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
					golog.Child("[Sonr.daemon] ").Fatal("Invalid Enviornment Variable Format")
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
				golog.Child("[Sonr.daemon] ").Warn("Failed to set Profile from flag")
			}
		}
		app.Start(req, cliPtr, "full")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().BoolVar(&cliPtr, "cli", false, "run in CLI Mode")
	rootCmd.Flags().Float64Var(&latPtr, "lat", 34.102920, "latitude for InitializeRequest")
	rootCmd.Flags().Float64Var(&lngPtr, "lng", -118.394190, "longitude for InitializeRequest")
	rootCmd.Flags().StringVar(&profilePtr, "profile", "", "profile JSON string")
	rootCmd.Flags().StringVar(&varsPtr, "vars", "", "enviornment variables encoded as base64")
	viper.BindPFlags(rootCmd.Flags())

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Path to config file (default is $HOME/.sonr-config/sonrd.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".sonrd" (without extension).
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
		}
		cfPath := filepath.Join(homeDir, ".sonr-config")
		os.MkdirAll(cfPath, os.ModePerm)
		viper.AddConfigPath(cfPath)
		viper.SetConfigType("yaml")
		viper.SetConfigName("sonrd")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		encodedVars := viper.GetString("vars")
		if encodedVars == "" {
			golog.Child("[Sonr.daemon] ").Error("Failed to read Enviornment Vars from Config")
			return
		}

		keyValuesBuf, err := base64.StdEncoding.DecodeString(encodedVars)
		if err != nil {
			golog.Child("[Sonr.daemon] ").Error("Failed to decode Enviornment Vars from Config")
			return
		}

		// Split String Values
		keyValuePairs := strings.Split(string(keyValuesBuf), ",")
		golog.Child("[Sonr.daemon] ").Debugf("Loading %v Enviornment variables from Config.", len(keyValuePairs))
		// Iterate over keyValuePairs
		for _, v := range keyValuePairs {
			// Trim White Space
			tv := strings.TrimSpace(v)

			// Split Key Value Pairs
			value := strings.Split(tv, "=")
			if len(value) != 2 {
				golog.Child("[Sonr.daemon] ").Fatal("Invalid Enviornment Variable Format")
			}

			// Set Env Variables
			os.Setenv(value[0], value[1])
		}
	}
}
