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
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/kataras/golog"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var execServeCmd = exec.Command("sonrd", "serve")
var cliPtr bool
var latPtr float64
var lngPtr float64
var profilePtr string
var varsPtr string
var cfgFile string
var PIDFile = "/tmp/sonrd.pid"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sonrd",
	Short: "Daemon for Sonr Binary, interact with Node through this CLI.",
	Long:  `Sonr's Core Framework manages Discovery, Connection, Data-Transfer, Authorization/Authentication and Peer Lobby.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Name() != "close" {
			if err := execServeCmd.Start(); err != nil {
				golog.Child("[Sonr.daemon] ").Fatal(err)
			}
			savePID(execServeCmd.Process.Pid)
			err := execServeCmd.Process.Signal(syscall.SIGSTOP)
			if err != nil {
				golog.Child("[Sonr.daemon] ").Fatal(err)
			}
			execServeCmd.Wait()
		}
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
		viper.Set("vars", base64.StdEncoding.EncodeToString([]byte(varsPtr)))
		viper.SafeWriteConfig()
	}

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

		// Split String Values
		keyValuePairs := strings.Split(string(encodedVars), ",")
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

func savePID(pid int) {
	golog.Child("[Sonr.daemon] ").Info("Process ID is : ", pid)
	file, err := os.Create(PIDFile)
	if err != nil {
		golog.Child("[Sonr.daemon] ").Error("Unable to create pid file : %v\n", err)
		os.Exit(1)
	}

	defer file.Close()

	_, err = file.WriteString(strconv.Itoa(pid))

	if err != nil {
		golog.Child("[Sonr.daemon] ").Error("Unable to create pid file : %v\n", err)
		os.Exit(1)
	}

	file.Sync() // flush to disk
}
