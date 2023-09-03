package routes

import (
	"github.com/gin-gonic/gin"
	h "github.com/sonrhq/core/internal/highway/handler"
)

// RegisterRoutes registers all routes for the highway service
func RegisterRoutes(r *gin.Engine) {
	r.GET(kCurrentControllerEndpoint, h.CurrentUser)                         // Authentication
	r.GET(kGetCredCreationOptionsEndpoint, h.GetCredentialAttestationParams) // Authentication/Register
	r.GET(kGetCredAssertionOptionsEndpoint, h.GetCredentialAssertionParams)  // Authentication/Login
	r.POST(kRegisterCredForClaimsEndpoint, h.RegisterControllerIdentity)     // Authentication/Register
	r.POST(kVerifyCredForAccountEndpoint, h.SignInWithCredential)            // Authentication/Login
	r.POST(kVerifyWithAccountEndpoint, h.VerifyWithAccount)                  // Accounts

	r.GET(kGetMagicEmailStartEndpoint, h.GetEmailAssertionParams)
	r.POST(kRegisterMagicEmailEndpoint, h.SignInWithEmail)
}
