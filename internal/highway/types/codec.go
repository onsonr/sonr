package types

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

const (
	// HighwayServiceName is the name of the highway service
	HighwayServiceName = "Highway Protocol Service"

	// HighwayServiceDisplayName is the display name of the highway service
	HighwayServiceDisplayName = "highway"

	// HighwayServiceDescription is the description of the highway service
	HighwayServiceDescription = "Proxy for underlying blockchain protocol"

	// IceFireKVServiceName is the name of the IceFire KV service
	IceFireKVServiceName = "IceFire KV Service"

	// IceFireKVServiceDisplayName is the display name of the IceFire KV service
	IceFireKVServiceDisplayName = "icefirekv"

	// IceFireKVServiceDescription is the description of the IceFire KV service
	IceFireKVServiceDescription = "Key-value store for the Highway Protocol"

	// IceFireSQLServiceName is the name of the IceFire SQL service
	IceFireSQLServiceName = "IceFire SQL Service"

	// IceFireSQLServiceDisplayName is the display name of the IceFire SQL service
	IceFireSQLServiceDisplayName = "icefiresql"

	// IceFireSQLServiceDescription is the description of the IceFire SQL service
	IceFireSQLServiceDescription = "SQL database for the Highway Protocol"
)

// EnvEnabled returns true if the Highway API is enabled
func EnvEnabled() bool {
	return viper.GetBool("highway.enabled")
}

// EnvHostAddress returns the host and port of the Highway API
func EnvHostAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString("highway.api.host"), viper.GetInt("highway.api.port"))
}

// EnvRequestTimeout returns the timeout for Highway API requests
func EnvRequestTimeout() time.Duration {
	return time.Duration(viper.GetInt("highway.api.timeout")) * time.Second
}

// EnvNodeAPIHostAddress returns the host and port of the Node API
func EnvNodeAPIHostAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString("node.api.host"), viper.GetInt("node.api.port"))
}

// EnvNodeGrpcHostAddress returns the host and port of the Node P2P
func EnvNodeGrpcHostAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString("node.grpc.host"), viper.GetInt("node.grpc.port"))
}
