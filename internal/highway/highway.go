package highway

import (
	"fmt"
	// swagger embed files
	// gin-swagger middleware

	"github.com/spf13/viper"
)

// StartAPI starts the highway api service
func StartAPI() {
	err := runHighway()
	if err != nil {
		fmt.Println("Cannot start the service: " + err.Error())
	}
}

// StartDB starts the highway db service
func StartDB() {
	kvArgs := viper.GetStringSlice("highway.icefirekv.args")
	kvExecutable := viper.GetString("highway.icefirekv.executable")
	err := runIcefireKv(kvExecutable, kvArgs)
	if err != nil {
		fmt.Println("Cannot start the service: " + err.Error())
	}
}

// StartSQL starts the highway sql service
func StartSQL() {
	sqlArgs := viper.GetStringSlice("highway.icefiresql.args")
	sqlExecutable := viper.GetString("highway.icefiresql.executable")
	err := runIcefireSQL(sqlExecutable, sqlArgs)
	if err != nil {
		fmt.Println("Cannot start the service: " + err.Error())
	}
}
