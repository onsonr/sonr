package routes

import (
	"github.com/gin-gonic/gin"
	h "github.com/sonrhq/core/pkg/highway/handler"
)

func RegisterRoutes(r *gin.Engine) {
	// Public routes
	r.GET(kGetCredCreationOptionsEndpoint, h.GetCredentialCreationOptions)   // Services/Register
	r.GET(kGetCredAssertionOptionsEndpoint, h.GetCredentialAssertionOptions) // Services/Login
	r.POST(kRegisterCredForClaimsEndpoint, h.RegisterCredentialForClaims)    // Services/Register
	r.POST(kVerifyCredForAccountEndpoint, h.VerifyCredentialForAccount)      // Services/Login
	r.POST(kBroadcastSonrTxEndpoint, h.BroadcastSonrTx)                      // Transactions
	r.POST(kVerifyWithAccountEndpoint, h.VerifyWithAccount)                  // Accounts
}
