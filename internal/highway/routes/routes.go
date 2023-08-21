package routes

import (
	"github.com/gin-gonic/gin"
	h "github.com/sonrhq/core/internal/highway/handler"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// RegisterRoutes registers all routes for the highway service
func RegisterRoutes(r *gin.Engine) {
	r.GET(getHealthStatusEndpoint, h.GetHealth)                              // Accounts
	r.GET(kCurrentControllerEndpoint, h.CurrentUser)                         // Authentication
	r.GET(kGetCredCreationOptionsEndpoint, h.GetCredentialAttestationParams) // Authentication/Register
	r.GET(kGetCredAssertionOptionsEndpoint, h.GetCredentialAssertionParams)  // Authentication/Login
	r.POST(kRegisterCredForClaimsEndpoint, h.RegisterControllerIdentity)     // Authentication/Register
	r.POST(kVerifyCredForAccountEndpoint, h.SignInWithCredential)            // Authentication/Login
	r.POST(kVerifyWithAccountEndpoint, h.VerifyWithAccount)                  // Accounts

	r.GET(kGetMagicEmailStartEndpoint, h.GetEmailAssertionParams)
	r.POST(kRegisterMagicEmailEndpoint, h.SignInWithEmail)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
