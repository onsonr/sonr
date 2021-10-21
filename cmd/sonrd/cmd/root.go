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
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sonrd",
	Short: "Daemon for Sonr Binary, interact with Node through this CLI.",
	Long:  `Sonr's Core Framework manages Discovery, Connection, Data-Transfer, Authorization/Authentication and Peer Lobby.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Path to config file (default is $HOME/.sonr-config/sonrd.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
			golog.Error("Failed to read Enviornment Vars from Config")
			return
		}

		keyValuesBuf, err := base64.StdEncoding.DecodeString(encodedVars)
		if err != nil {
			golog.Error("Failed to decode Enviornment Vars from Config")
			return
		}

		// Split String Values
		keyValuePairs := strings.Split(string(keyValuesBuf), ",")
		golog.Infof("Loading %v Enviornment variables from Config.", len(keyValuePairs))
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
}
