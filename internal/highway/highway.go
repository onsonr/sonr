package highway

import (
	"fmt"
	// swagger embed files
	// gin-swagger middleware
)

// StartAPI starts the highway api service
func StartAPI() {
	err := runHighway()
	if err != nil {
		fmt.Println("Cannot start the service: " + err.Error())
	}
}
