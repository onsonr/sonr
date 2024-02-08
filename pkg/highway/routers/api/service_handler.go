package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-webauthn/webauthn/protocol"

	servicev1 "github.com/sonrhq/sonr/api/sonr/service/v1"
	"github.com/sonrhq/sonr/pkg/highway/middleware"
	"github.com/sonrhq/sonr/pkg/vault"
	"github.com/sonrhq/sonr/x/service"
)

// ServiceHandler is a handler for the staking module
var ServiceHandler = serviceHandler{}

// serviceHandler is a handler for the staking module
type serviceHandler struct{}

// StartRegistration returns credential creation options for the origin host
func StartRegistration(w http.ResponseWriter, r *http.Request) {
	heightStr := chi.URLParam(r, "handle")
	hc := middleware.NewServiceClient(middleware.GrpcClientConn(r))
	resp, err := hc.Service(r.Context(), &servicev1.QueryServiceRequest{})
	if err != nil {
		return
	}
	c, err := vault.Create(r.Context())
	if err != nil {
		return
	}
	opts := service.GetPublicKeyCredentialCreationOptions(resp.Service, protocol.UserEntity{
		DisplayName: heightStr,
		ID:          []byte(c.Address),
	})
	optsBz, err := json.Marshal(opts)
	if err != nil {
		return
	}
	w.Write(optsBz)
}

// StartLogin returns credential request options for the origin host
func StartLogin(w http.ResponseWriter, r *http.Request) {
	hc := middleware.NewServiceClient(middleware.GrpcClientConn(r))
	resp, err := hc.Service(r.Context(), &servicev1.QueryServiceRequest{})
	if err != nil {
		return
	}
	opts := service.GetPublicKeyCredentialRequestOptions(resp.Service, []protocol.CredentialDescriptor{})
	optsBz, err := json.Marshal(opts)
	if err != nil {
		return
	}
	w.Write(optsBz)
}
