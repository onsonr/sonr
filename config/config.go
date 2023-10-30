package config

import (
	"github.com/sonrhq/core/pkg/env"
	"github.com/sonrhq/core/pkg/xfilepath"
)

var c Config

// Config is defining a struct type named `Config`. This struct is used to store the configuration values for the application. It has various fields that correspond to different configuration parameters, such as version, chain ID, launch settings, database settings, node settings, genesis settings, etc. Each field is annotated with `mapstructure` tags, which are used to map the corresponding configuration values from a YAML file to the struct fields.
type Config struct {
	// Version is the version of the application.
	Version string `mapstructure:"version"`

	// Environment is the environment of the application.
	Environment string `mapstructure:"environment"`

	// ChainID is the chain ID of the application.
	ChainID string `mapstructure:"chain-id"`

	// BinPath is the binary path of the application.
	BinPath string `mapstructure:"bin-path"`

	// Launch is the launch settings of the application.
	Launch struct {
		// Name is the name of the current validator/node.
		Name string `mapstructure:"name"`

		// Address is the address of the current validator/node.
		Address string `mapstructure:"address"`

		// Highway is the highway settings of the application.
		Highway struct {
			// SigningKey is the signing key of the highway.
			SigningKey string `mapstructure:"signing-key"`

			// API is the API settings of the highway.
			API struct {
				Host    string `mapstructure:"host"`
				Port    int    `mapstructure:"port"`
				Timeout int    `mapstructure:"timeout"`
			} `mapstructure:"api"`

			// DB is the database settings of the highway.
			DB struct {
				IcefireKV struct {
					Host   string `mapstructure:"host"`
					Port   int    `mapstructure:"port"`
					Binary string `mapstructure:"binary"`
				} `mapstructure:"icefirekv"`
				IcefireSQL struct {
					Host   string `mapstructure:"host"`
					Port   int    `mapstructure:"port"`
					Binary string `mapstructure:"binary"`
				} `mapstructure:"icefiresql"`
			} `mapstructure:"db"`
		} `mapstructure:"highway"`

		// Node is the node settings of the application.
		Node struct {
			// API is the API settings of the node.
			API struct {
				Host string `mapstructure:"host"`
				Port int    `mapstructure:"port"`
			} `mapstructure:"api"`

			// P2P is the P2P settings of the node.
			P2P struct {
				Host string `mapstructure:"host"`
				Port int    `mapstructure:"port"`
			} `mapstructure:"p2p"`

			// RPC is the RPC settings of the node.
			RPC struct {
				Host string `mapstructure:"host"`
				Port int    `mapstructure:"port"`
			} `mapstructure:"rpc"`

			// GRPC is the GRPC settings of the node.
			GRPC struct {
				Host string `mapstructure:"host"`
				Port int    `mapstructure:"port"`
			} `mapstructure:"grpc"`
		} `mapstructure:"node"`
	} `mapstructure:"launch"`

	// Genesis is the genesis settings of the application.
	Genesis struct {
		// Accounts is the accounts settings of the genesis.
		Accounts []struct {
			Name  string   `mapstructure:"name"`
			Coins []string `mapstructure:"coins"`
		} `mapstructure:"accounts"`

		// Faucet is the faucet settings of the genesis.
		Faucet struct {
			Name  string   `mapstructure:"name"`
			Coins []string `mapstructure:"coins"`
		} `mapstructure:"faucet"`

		// Validators is the validators settings of the genesis.
		Validators []struct {
			Name   string `mapstructure:"name"`
			Bonded string `mapstructure:"bonded"`
		} `mapstructure:"validators"`
	} `mapstructure:"genesis"`
}

// DirPath returns the path of configuration directory of Ignite.
var DirPath = xfilepath.Mkdir(env.ConfigDir())
