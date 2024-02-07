package api

import (
	"net/http"

	servicev1 "github.com/sonrhq/sonr/api/sonr/service/v1"
	"github.com/sonrhq/sonr/pkg/highway/middleware"
)

// ServiceHandler is a handler for the staking module
var ServiceHandler = serviceHandler{}

// serviceHandler is a handler for the staking module
type serviceHandler struct{}

func StartRegistration(w http.ResponseWriter, r *http.Request) {
	middleware.NewServiceClient(middleware.GrpcClientConn(r)).Credentials(r.Context(), &servicev1.QueryCredentialsRequest{})
}
