package config

import "fmt"

var c Config

// Config is defining a struct type named `Config`. This struct is used to store the configuration values for the application. It has various fields that correspond to different configuration parameters, such as version, chain ID, launch settings, database settings, node settings, genesis settings, etc. Each field is annotated with `mapstructure` tags, which are used to map the corresponding configuration values from a YAML file to the struct fields.
type Config struct {
	Version     string        `mapstructure:"version"`
	ChainID     string        `mapstructure:"chain-id"`
	Environment string        `mapstructure:"environment"`
	Launch      LaunchConfig  `mapstructure:"launch"`
	Genesis     GenesisConfig `mapstructure:"genesis"`
}

// LaunchConfig is defining a struct type named `LaunchConfig`. This struct is used to store the configuration values for the launch settings of the application. It has various fields that correspond to different configuration parameters, such as name, address, highway settings, node settings, etc. Each field is annotated with `mapstructure` tags, which are used to map the corresponding configuration values from a YAML file to the struct fields.
type LaunchConfig struct {
	Name    string        `mapstructure:"name"`
	Address string        `mapstructure:"address"`
	Highway HighwayConfig `mapstructure:"highway"`
	Node    NodeConfig    `mapstructure:"node"`
}

// HighwayConfig is defining a struct type named `HighwayConfig`. This struct is used to store the configuration values for the highway settings of the application. It has various fields that correspond to different configuration parameters, such as API settings, database settings, JWT signing key, etc. Each field is annotated with `mapstructure` tags, which are used to map the corresponding configuration values from a YAML file to the struct fields.
type HighwayConfig struct {
	API           HighwayAPIConfig `mapstructure:"api"`
	DB            HighwayDBConfig  `mapstructure:"db"`
	JWTSigningKey string           `mapstructure:"signing-key"`
}

// HighwayAPIConfig is defining a struct type named `HighwayAPIConfig`. This struct is used to store the configuration values for the API settings of the application. It has various fields that correspond to different configuration parameters, such as host, port, timeout, etc. Each field is annotated with `mapstructure` tags, which are used to map the corresponding configuration values from a YAML file to the struct fields.
type HighwayAPIConfig struct {
	Host    string `mapstructure:"host"`
	Port    int    `mapstructure:"port"`
	Timeout int    `mapstructure:"timeout"`
}

// HighwayDBConfig is defining a struct type named `HighwayDBConfig`. This struct is used to store the configuration values for the database settings of the application. It has various fields that correspond to different configuration parameters, such as IcefireKV settings, IcefireSQL settings, etc. Each field is annotated with `mapstructure` tags, which are used to map the corresponding configuration values from a YAML file to the struct fields.
type HighwayDBConfig struct {
	IcefireKV  IcefireKVConfig  `mapstructure:"icefirekv"`
	IcefireSQL IcefireSQLConfig `mapstructure:"icefiresql"`
}

// IcefireKVConfig is defining a struct type named `IcefireKVConfig`. This struct is used to store the configuration values for the IcefireKV settings of the application. It has various fields that correspond to different configuration parameters, such as host, port, binary, etc. Each field is annotated with `mapstructure` tags, which are used to map the corresponding configuration values from a YAML file to the struct fields.
type IcefireKVConfig struct {
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	Binary string `mapstructure:"binary"`
}

// IcefireSQLConfig is defining a struct type named `IcefireSQLConfig`. This struct is used to store the configuration values for the IcefireSQL settings of the application. It has various fields that correspond to different configuration parameters, such as host, port, binary, etc. Each field is annotated with `mapstructure` tags, which are used to map the corresponding configuration values from a YAML file to the struct fields.
type IcefireSQLConfig struct {
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	Binary string `mapstructure:"binary"`
}

// NodeConfig is defining a struct type named `NodeConfig`. This struct is used to store the configuration values for the node settings of the application. It has various fields that correspond to different configuration parameters, such as API settings, P2P settings, RPC settings, GRPC settings, etc. Each field is annotated with `mapstructure` tags, which are used to map the corresponding configuration values from a YAML file to the struct fields.
type NodeConfig struct {
	API  NodeAPIConfig  `mapstructure:"api"`
	P2P  NodeP2PConfig  `mapstructure:"p2p"`
	RPC  NodeRPCConfig  `mapstructure:"rpc"`
	GRPC NodeGRPCConfig `mapstructure:"grpc"`
}

// NodeAPIConfig is defining a struct type named `NodeAPIConfig`. This struct is used to store the configuration values for the API settings of the application. It has various fields that correspond to different configuration parameters, such as host, port, etc. Each field is annotated with `mapstructure` tags, which are used to map the corresponding configuration values from a YAML file to the struct fields.
type NodeAPIConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

// NodeP2PConfig is defining a struct type named `NodeP2PConfig`. This struct is used to store the configuration values for the P2P settings of the application. It has various fields that correspond to different configuration parameters, such as host, port, etc. Each field is annotated with `mapstructure` tags, which are used to map the corresponding configuration values from a YAML file to the struct fields.
type NodeP2PConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

// NodeRPCConfig is defining a struct type named `NodeRPCConfig`. This struct is used to store the configuration values for the RPC settings of the application. It has various fields that correspond to different configuration parameters, such as host, port, etc. Each field is annotated with `mapstructure` tags, which are used to map the corresponding configuration values from a YAML file to the struct fields.
type NodeRPCConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

// NodeGRPCConfig is defining a struct type named `NodeGRPCConfig`. This struct is used to store the configuration values for the GRPC settings of the application. It has various fields that correspond to different configuration parameters, such as host, port, etc. Each field is annotated with `mapstructure` tags, which are used to map the corresponding configuration values from a YAML file to the struct fields.
type NodeGRPCConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

// GenesisConfig is defining a struct type named `GenesisConfig`. This struct is used to store the configuration values for the genesis settings of the application. It has various fields that correspond to different configuration parameters, such as accounts, faucet, validators, etc. Each field is annotated with `mapstructure` tags, which are used to map the corresponding configuration values from a YAML file to the struct fields.
type GenesisConfig struct {
	Accounts   []GenesisAccountConfig   `mapstructure:"accounts"`
	Faucet     GenesisFaucetConfig      `mapstructure:"faucet"`
	Validators []GenesisValidatorConfig `mapstructure:"validators"`
}

// GenesisAccountConfig is defining a struct type named `GenesisAccountConfig`. This struct is used to store the configuration values for the account settings of the application. It has various fields that correspond to different configuration parameters, such as name, coins, etc. Each field is annotated with `mapstructure` tags, which are used to map the corresponding configuration values from a YAML file to the struct fields.
type GenesisAccountConfig struct {
	Name  string   `mapstructure:"name"`
	Coins []string `mapstructure:"coins"`
}

// GenesisFaucetConfig is defining a struct type named `GenesisFaucetConfig`. This struct is used to store the configuration values for the faucet settings of the application. It has various fields that correspond to different configuration parameters, such as name, coins, etc. Each field is annotated with `mapstructure` tags, which are used to map the corresponding configuration values from a YAML file to the struct fields.
type GenesisFaucetConfig struct {
	Name  string   `mapstructure:"name"`
	Coins []string `mapstructure:"coins"`
}

// GenesisValidatorConfig is defining a struct type named `GenesisValidatorConfig`. This struct is used to store the configuration values for the validator settings of the application. It has various fields that correspond to different configuration parameters, such as name, bonded, etc. Each field is annotated with `mapstructure` tags, which are used to map the corresponding configuration values from a YAML file to the struct fields.
type GenesisValidatorConfig struct {
	Name   string `mapstructure:"name"`
	Bonded string `mapstructure:"bonded"`
}

// Print is a function that prints the configuration values of the application.
func Print() {
	fmt.Print(c)
}
