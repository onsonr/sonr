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
