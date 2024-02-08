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

// QueryOrigin returns the service for the origin host
func (h serviceHandler) QueryOrigin(w http.ResponseWriter, r *http.Request) {
	origin := chi.URLParam(r, "origin")
	resp, err := middleware.ServiceClient(r, w).Service(r.Context(), &servicev1.QueryServiceRequest{Origin: origin})
	if err != nil {
		middleware.NotFound(w, err)
		return
	}
	middleware.JSONResponse(w, resp.Service)
}

// StartRegistration returns credential creation options for the origin host
func (h serviceHandler) StartRegistration(w http.ResponseWriter, r *http.Request) {
	handleStr := chi.URLParam(r, "handle")
	origin := chi.URLParam(r, "origin")
	resp, err := middleware.ServiceClient(r, w).Service(r.Context(), &servicev1.QueryServiceRequest{Origin: origin})
	if err != nil {
		middleware.NotFound(w, err)
		return
	}
	c, err := vault.Create(r.Context())
	if err != nil {
		middleware.InternalServerError(w, err)
		return
	}
	opts := service.GetPublicKeyCredentialCreationOptions(resp.Service, protocol.UserEntity{
		DisplayName: handleStr,
		ID:          []byte(c.Address),
	})
	middleware.JSONResponse(w, opts)
}

// FinishRegistration returns the result of the credential creation
func (h serviceHandler) FinishRegistration(w http.ResponseWriter, r *http.Request) {
	origin := chi.URLParam(r, "origin")
	resp, err := middleware.ServiceClient(r, w).Service(r.Context(), &servicev1.QueryServiceRequest{Origin: origin})
	if err != nil {
		middleware.NotFound(w, err)
		return
	}
	var credential protocol.PublicKeyCredential
	if err := json.NewDecoder(r.Body).Decode(&credential); err != nil {
		middleware.BadRequest(w, err)
		return
	}
	_, err = service.FinishRegistration(r.Context(), resp.Service, credential)
	if err != nil {
		middleware.BadRequest(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// StartLogin returns credential request options for the origin host
func (h serviceHandler) StartLogin(w http.ResponseWriter, r *http.Request) {
	origin := chi.URLParam(r, "origin")
	resp, err := middleware.ServiceClient(r, w).Service(r.Context(), &servicev1.QueryServiceRequest{Origin: origin})
	if err != nil {
		middleware.NotFound(w, err)
		return
	}
	opts := service.GetPublicKeyCredentialRequestOptions(resp.Service, []protocol.CredentialDescriptor{})
	middleware.JSONResponse(w, opts)
}

// FinishLogin returns the result of the credential request
func (h serviceHandler) FinishLogin(w http.ResponseWriter, r *http.Request) {
	origin := chi.URLParam(r, "origin")
	resp, err := middleware.ServiceClient(r, w).Service(r.Context(), &servicev1.QueryServiceRequest{Origin: origin})
	if err != nil {
		middleware.NotFound(w, err)
		return
	}
	var credential protocol.PublicKeyCredential
	if err := json.NewDecoder(r.Body).Decode(&credential); err != nil {
		return
	}
	_, err = service.FinishLogin(r.Context(), resp.Service, credential)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}

// RegisterRoutes registers the node routes
func (h serviceHandler) RegisterRoutes(r chi.Router) {
	r.Get("/service/{origin}", h.StartLogin)
	r.Get("/service/{origin}/login/{handle}/start", h.StartLogin)
	r.Post("/service/{origin}/login/{handle}/finish", h.FinishLogin)
	r.Get("/service/{origin}/register/{handle}/start", h.StartRegistration)
	r.Post("/service/{origin}/register/{handle}/finish", h.FinishRegistration)
}
