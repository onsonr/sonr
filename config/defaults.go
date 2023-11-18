package config

import (
	"github.com/spf13/viper"
)

// Masthead is the masthead of the application.
const Masthead = `
MMMMMMMMMMMMMMMMMMMWWMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMWXkoc:;;:cokXWMMMMMMMMMMMMM
MMMMMMMMMMMWOc.          'oXMMMMMMMMMMMM
MMMMMMMMMWO:.            .oXMMMMMMMMMMMM
MMMMMMMWO:.     .lk00kl;oKWMMN0KWMMMMMMM
MMMMMWO:.     .lKWMMMMWWWMMNk,..lKWMMMMM
MMMW0:.     .lKWMWKOOKWMMMNo.    .oKWMMM
MMWx.     .lKWMW0l.  .c0WMWKl.     'kWMM
MMO'    .lKWMW0l.      .c0WMWKl.    'OMM
MWd    .dWMMMk.          .kMMMWd.    dWM
MMx.    :KWMMKl.        .lKMMWK:    .xMM
MMXc     .lKWMW0c.    .l0WMWKl.     cXMM
MMMXo.     .dNMMW0o::l0WMWKl.     .oXMMM
MMMMWKl.   ;kNMMMMMMWMMWKl.     .lKWMMMM
MMMMMMWKocxNMMWKkKWMMWKl.     .lKWMMMMMM
MMMMMMMMWMMWKl. .;cc;.     .lKWMMMMMMMMM
MMMMMMMMMMMK:            .lKWMMMMMMMMMMM
MMMMMMMMMMMWKd;..    ..;dKWMMMMMMMMMMMMM
MMMMMMMMMMMMMMWX0OkkO0XWMMMMMMMMMMMMMMMM

Sonr Node
> Sonr is an Encrypted & Private by default Identity Verification System for the IBC Protocol.
🌐 - https://sonr.io
🚀 - https://github.com/sonrhq/sonr
`

const bip44_purpose = 44

const bip44_coin_type = 703

// Init initializes the configuration parameters.
func Init() {
	// Set defaults
	initSonrConfig()
}

// initSonrConfig sets the default values for the configuration parameters.
func initSonrConfig() {
	viper.SetDefault("launch.chain-id", "sonr-localnet-1")
	viper.SetDefault("launch.environment", "development")
	viper.SetDefault("launch.moniker", "alice")
	viper.SetDefault("launch.val_address", "")
	viper.SetDefault("highway.enabled", true)
	viper.SetDefault("highway.jwt.key", "@re33lyb@dsecret")
	viper.SetDefault("highway.api.host", "localhost")
	viper.SetDefault("highway.api.timeout", 15)

	viper.SetDefault("highway.icefirekv.host", "localhost")
	viper.SetDefault("highway.icefirekv.port", 6001)

	viper.SetDefault("highway.icefiresql.host", "localhost")
	viper.SetDefault("highway.icefiresql.port", 23306)

	viper.SetDefault("node.api.host", "localhost")
	viper.SetDefault("node.api.port", 1317)
	viper.SetDefault("node.grpc.host", "localhost")
	viper.SetDefault("node.grpc.port", 9090)
	viper.SetDefault("node.rpc.host", "localhost")
	viper.SetDefault("node.rpc.port", 26657)
	viper.SetDefault("node.p2p.host", "validator")
	viper.SetDefault("node.p2p.port", 26656)

}
