package highway

import (
	"fmt"

	"github.com/sonrhq/core/internal/highway/types"
)

// StartAPI starts the highway api service
func StartAPI() {
	if types.EnvEnabled() {

		err := runHighway()
		if err != nil {
			fmt.Println("Cannot start the service: " + err.Error())
		}
	}
}
