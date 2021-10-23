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
var isDebug bool
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Path to config file (default is $HOME/.sonr-config/sonrd.yaml)")
	rootCmd.PersistentFlags().BoolVar(&isDebug, "debug", false, "Sets Logging to Debug mode (default is false)")
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
		golog.Debugf("Loading %v Enviornment variables from Config.", len(keyValuePairs))
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
