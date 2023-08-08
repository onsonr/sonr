package highway

import (
	"fmt"
	// swagger embed files
	// gin-swagger middleware

	"github.com/spf13/viper"
)

// StartService starts the highway service
func StartService() {
	var err error
	err = runHighway()
	if err != nil {
		fmt.Println("Cannot start the service: " + err.Error())
	}

	kvArgs := viper.GetStringSlice("highway.icefirekv.args")
	kvExecutable := viper.GetString("highway.icefirekv.executable")
	err = runIcefireKv(kvExecutable, kvArgs)
	if err != nil {
		fmt.Println("Cannot start the service: " + err.Error())
	}

	sqlArgs := viper.GetStringSlice("highway.icefiresql.args")
	sqlExecutable := viper.GetString("highway.icefiresql.executable")
	err = runIcefireSQL(sqlExecutable, sqlArgs)
	if err != nil {
		fmt.Println("Cannot start the service: " + err.Error())
	}
}
