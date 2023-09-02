package types

import (
	"fmt"

	"github.com/spf13/viper"
)

// EnvIceFireKVHost returns the host and port of the IceFire KV store
func EnvIceFireKVHost() string {
	return fmt.Sprintf("%s:%d", viper.GetString("highway.icefirekv.host"), viper.GetInt("highway.icefirekv.port"))
}

// EnvIceFireSQLHost returns the host and port of the IceFire KV store
func EnvIceFireSQLHost() string {
	return fmt.Sprintf("%s:%d", viper.GetString("highway.icefiresql.host"), viper.GetInt("highway.icefiresql.port"))
}
