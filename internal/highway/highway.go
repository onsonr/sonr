package highway

import (
	"fmt"

	"github.com/sonrhq/core/internal/highway/types"
	// swagger embed files
	// gin-swagger middleware
)

// StartAPI starts the highway api service
// @title           Sonr Highway Protocol API
// @version         1.0
// @description     API for the Sonr Highway Protocol, a peer-to-peer identity and asset management system.
// @termsOfService  <URL_to_your_terms_of_service>
// @contact.name    Sonr API Support
// @contact.url     <URL_to_your_support>
// @contact.email   <your_support_email>
// @license.name    <Your_License_Name>
// @license.url     <URL_to_license>
// @host            <host_address>:<port>
// @BasePath        /api/v1
func StartAPI() {
	if types.EnvEnabled() {
		err := runHighway()
		if err != nil {
			fmt.Println("Cannot start the service: " + err.Error())
		}
	}
}
