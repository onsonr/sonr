package sql

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
)

var ifq *IceFireMySQL

// IceFireMySQL is a wrapper around the mysql client
type IceFireMySQL struct {
	ctx context.Context
}

// EnvIceFireSQLHost returns the host and port of the IceFire KV store
func envIceFireSQLHost() string {
	return fmt.Sprintf("%s:%d", viper.GetString("highway.icefiresql.host"), viper.GetInt("highway.icefiresql.port"))
}
