package main

import (
	"fmt"
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/kataras/golog"
	"github.com/sonr-io/sonr/app"
	"github.com/spf13/viper"
	cmd "github.com/tendermint/spm/cosmoscmd"
)

const (
	CONFIG_PATH_ROOT = "../../"
)

var (
	logger = golog.Default.Child("sonrd")
)

// load environment variables
func loadEnv() error {
	viper.AddConfigPath(CONFIG_PATH_ROOT)

	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	return viper.ReadInConfig()
}

func main() {
	// err := loadEnv()
	// if err != nil {
	// 	logger.Errorf("Error while loading config for enviorment %s", err)
	// }

	rootCmd, _ := cmd.NewRootCmd(
		app.Name,
		app.AccountAddressPrefix,
		app.DefaultNodeHome,
		app.Name,
		app.ModuleBasics,
		app.New,
		// this line is used by starport scaffolding # root/arguments

	)
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// Execute the root command.
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
