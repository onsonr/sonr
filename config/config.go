package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

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

// Get returns the configuration values of the application.
func Get() Config {
	return c
}

// ChainID returns the chain ID from the environment variable SONR_CHAIN_ID. (default: sonr-localnet-1)
func ChainID() string {
	return viper.GetString("SONR_CHAIN_ID")
}

// Environment returns the environment from the environment variable SONR_ENVIRONMENT. (default: development)
func Environment() string {
	return viper.GetString("SONR_ENVIRONMENT")
}

// IsProduction returns true if the environment is production
func IsProduction() bool {
	return viper.GetString("SONR_ENVIRONMENT") == "production"
}

// JWTSigningKey returns the JWT signing key
func JWTSigningKey() []byte {
	return []byte(viper.GetString("SONR_LAUNCH_HIGHWAY_SIGNING_KEY"))
}

// HighwayHostAddress returns the host and port of the Highway API
func HighwayHostAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString("SONR_LAUNCH_HIGHWAY_API_HOST"), viper.GetInt("SONR_LAUNCH_HIGHWAY_API_PORT"))
}

// HighwayRequestTimeout returns the timeout for Highway API requests
func HighwayRequestTimeout() time.Duration {
	return time.Duration(viper.GetInt("SONR_LAUNCH_HIGHWAY_API_TIMEOUT")) * time.Second
}

// IceFireKVHost returns the host and port of the IceFire KV store
func IceFireKVHost() string {
	return fmt.Sprintf("%s:%d", viper.GetString("SONR_LAUNCH_HIGHWAY_DB_ICEFIREKV_HOST"), viper.GetInt("SONR_LAUNCH_HIGHWAY_DB_ICEFIREKV_PORT"))
}

// NodeAPIHostAddress returns the host and port of the Node API
func NodeAPIHostAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString("SONR_LAUNCH_NODE_API_HOST"), viper.GetInt("SONR_LAUNCH_NODE_API_PORT"))
}

// NodeGrpcHostAddress returns the host and port of the Node P2P
func NodeGrpcHostAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString("SONR_LAUNCH_NODE_GRPC_HOST"), viper.GetInt("SONR_LAUNCH_NODE_GRPC_PORT"))
}

// NodeP2PHostAddress returns the host and port of the Node P2P
func NodeP2PHostAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString("SONR_LAUNCH_NODE_P2P_HOST"), viper.GetInt("SONR_LAUNCH_NODE_P2P_PORT"))
}

// NodeRPCHostAddress returns the host and port of the Node RPC
func NodeRPCHostAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString("SONR_LAUNCH_NODE_RPC_HOST"), viper.GetInt("SONR_LAUNCH_NODE_RPC_PORT"))
}

// ValidatorAddress returns the validator address from the environment variable SONR_VALIDATOR_ADDRESS.
func ValidatorAddress() string {
	return viper.GetString("SONR_VALIDATOR_ADDRESS")
}
