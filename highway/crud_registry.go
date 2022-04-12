package highway

import (
	context "context"
	"fmt"
	"log"
	"net/http"

	"github.com/duo-labs/webauthn/protocol"
	"github.com/duo-labs/webauthn/webauthn"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	rtv1 "github.com/sonr-io/blockchain/x/registry/types"
	rt "go.buf.build/grpc/go/sonr-io/sonr/registry"
)

// StartRegisterName starts the registration process for webauthn on http
func (s *HighwayServer) StartRegisterName(w http.ResponseWriter, r *http.Request) {
	// get username/friendly name
	vars := mux.Vars(r)
	username, ok := vars["username"]
	if !ok {
		JsonResponse(w, fmt.Errorf("must supply a valid username i.e. foo.snr"), http.StatusBadRequest)
		return
	}

	// Check if user exists and return error if it does
	_, err := s.cosmos.QueryName(username)
	if err == nil {
		log.Println(err)
		JsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	whois := &rtv1.WhoIs{
		Name:        username,
		Did:         "",
		Document:    nil,
		Creator:     s.cosmos.AccountName(),
		Credentials: make([]*rtv1.Credential, 0),
	}

	// Want performance? Store pointers!
	s.cache.Set(username, whois, cache.DefaultExpiration)

	// Updating the AuthenticatorSelection options.
	// See the struct declarations for values
	authSelect := protocol.AuthenticatorSelection{
		AuthenticatorAttachment: protocol.AuthenticatorAttachment("platform"),
		RequireResidentKey:      protocol.ResidentKeyUnrequired(),
		UserVerification:        protocol.VerificationRequired,
	}

	// Updating the ConveyencePreference options.
	// See the struct declarations for values
	conveyencePref := protocol.ConveyancePreference(protocol.PreferNoAttestation)
	registerOptions := func(credCreationOpts *protocol.PublicKeyCredentialCreationOptions) {
		credCreationOpts.CredentialExcludeList = whois.CredentialExcludeList()
	}

	// generate PublicKeyCredentialCreationOptions, session data
	options, sessionData, err := s.auth.BeginRegistration(
		whois,
		registerOptions,
		webauthn.WithAuthenticatorSelection(authSelect),
		webauthn.WithConveyancePreference(conveyencePref),
	)

	if err != nil {
		log.Println(err)
		JsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// store session data as marshaled JSON
	err = s.sessionStore.SaveWebauthnSession("registration", sessionData, r, w)
	if err != nil {
		log.Println(err)
		JsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JsonResponse(w, options, http.StatusOK)
}

// FinishRegisterName handles the registration of a new credential
func (s *HighwayServer) FinishRegisterName(w http.ResponseWriter, r *http.Request) {
	// get username
	vars := mux.Vars(r)
	username := vars["username"]

	// get user
	x, found := s.cache.Get(username)
	if !found {
		log.Println("Cache expired. user not found")
		JsonResponse(w, "Cache expired. User not found", http.StatusBadRequest)
		return
	}
	whois := x.(*rtv1.WhoIs)

	// load the session data
	sessionData, err := s.sessionStore.GetWebauthnSession("registration", r)
	if err != nil {
		log.Println(err)
		JsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	credential, err := s.auth.FinishRegistration(whois, sessionData, r)
	if err != nil {
		log.Println(err)
		JsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// define a message to create a did
	msg := rtv1.NewMsgRegisterName(s.cosmos.Address(), username, *credential)

	// broadcast a transaction from account `alice` with the message to create a did
	// store response in txResp
	txResp, err := s.cosmos.BroadcastRegisterName(msg)
	if err != nil {
		log.Println(err)
		JsonResponse(w, "Failed to broadcast to blockchain", http.StatusBadRequest)
		return
	}

	s.cache.Set("session", txResp.GetSession(), -1)
	JsonResponse(w, txResp.String(), http.StatusOK)
}

// StartAccessName accesses the user's existing credentials and returns a PublicKeyCredentialRequestOptions
func (s *HighwayServer) StartAccessName(w http.ResponseWriter, r *http.Request) {
	// get username
	vars := mux.Vars(r)
	username := vars["username"]

	// get user
	whois, err := s.cosmos.QueryName(username)
	if err != nil {
		log.Println(err)
		JsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// user doesn't exist
	if err != nil {
		log.Println(err)
		JsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// generate PublicKeyCredentialRequestOptions, session data
	options, sessionData, err := s.auth.BeginLogin(whois)
	if err != nil {
		log.Println(err)
		JsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// store session data as marshaled JSON
	err = s.sessionStore.SaveWebauthnSession("authentication", sessionData, r, w)
	if err != nil {
		log.Println(err)
		JsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JsonResponse(w, options, http.StatusOK)
}

// FinishAccessName handles the login of a credential and returns a PublicKeyCredentialResponse
func (s *HighwayServer) FinishAccessName(w http.ResponseWriter, r *http.Request) {
	// get username
	vars := mux.Vars(r)
	username := vars["username"]

	// get user
	whois, err := s.cosmos.QueryName(username)
	if err != nil {
		log.Println(err)
		JsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// user doesn't exist
	if err != nil {
		log.Println(err)
		JsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// load the session data
	sessionData, err := s.sessionStore.GetWebauthnSession("authentication", r)
	if err != nil {
		log.Println(err)
		JsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// in an actual implementation, we should perform additional checks on
	// the returned 'credential', i.e. check 'credential.Authenticator.CloneWarning'
	// and then increment the credentials counter
	_, err = s.auth.FinishLogin(whois, sessionData, r)
	if err != nil {
		log.Println(err)
		JsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// handle successful login
	JsonResponse(w, "Login Success", http.StatusOK)
}

// UpdateName updates a name.
func (s *HighwayServer) UpdateName(ctx context.Context, req *rt.MsgUpdateName) (*rt.MsgUpdateNameResponse, error) {
	return nil, ErrMethodUnimplemented
}
