package highway

import (
	"fmt"
	"log"
	"net/http"

	"github.com/duo-labs/webauthn/protocol"
	"github.com/duo-labs/webauthn/webauthn"
	"github.com/fxamacker/cbor/v2"
	"github.com/gorilla/mux"
	rtv1 "github.com/sonr-io/blockchain/x/registry/types"
	"github.com/sonr-io/core/crypto"
	"github.com/sonr-io/core/did"
	"github.com/sonr-io/core/did/ssi"
	"github.com/sonr-io/core/highway/user"
)

func (s *HighwayServer) BeginRegistration(w http.ResponseWriter, r *http.Request) {
	// get username/friendly name
	vars := mux.Vars(r)
	username, ok := vars["username"]
	if !ok {
		JsonResponse(w, fmt.Errorf("must supply a valid username i.e. foo@bar.com"), http.StatusBadRequest)
		return
	}

	// get user
	usr, err := s.userDb.GetUser(username)
	// user doesn't exist, create new user
	if err != nil {
		usr = user.NewUser(username, fmt.Sprintf("%s.snr", username))
		s.userDb.PutUser(usr)
	}
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
		credCreationOpts.CredentialExcludeList = usr.CredentialExcludeList()
	}

	// generate PublicKeyCredentialCreationOptions, session data
	options, sessionData, err := s.auth.BeginRegistration(
		usr,
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

func (s *HighwayServer) FinishRegistration(w http.ResponseWriter, r *http.Request) {

	// get username
	vars := mux.Vars(r)
	username := vars["username"]

	// get user
	user, err := s.userDb.GetUser(username)
	// user doesn't exist
	if err != nil {
		log.Println(err)
		JsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// load the session data
	sessionData, err := s.sessionStore.GetWebauthnSession("registration", r)
	if err != nil {
		log.Println(err)
		JsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	credential, err := s.auth.FinishRegistration(user, sessionData, r)
	if err != nil {
		log.Println(err)
		JsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	coseKey := crypto.COSEKey{}
	err = cbor.Unmarshal(credential.PublicKey, &coseKey)
	if err != nil {
		log.Println(err)
		JsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	pubKey, err := crypto.DecodePublicKey(&coseKey)
	if err != nil {
		log.Println(err)
		JsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// account `alice` was initialized during `starport chain serve`
	accountName := "alice"

	// get account from the keyring by account name and return a bech32 address
	acc, err := s.cosmos.Account(accountName)
	if err != nil {
		log.Println(err)
		JsonResponse(w, "Failed to find blockchain account", http.StatusNotFound)
	}

	didStr := fmt.Sprintf("did:sonr:%s", acc.Address("snr"))
	id, err := did.ParseDID(didStr)
	if err != nil {
		log.Println(err)
		JsonResponse(w, "Failed to parse DID", http.StatusNotFound)
	}

	verf, err := did.NewVerificationMethod(*id, ssi.JsonWebKey2020, *id, pubKey)
	if err != nil {
		log.Println(err)
		JsonResponse(w, "Failed to create verification method", http.StatusNotFound)
	}

	log.Println(verf)

	// define a message to create a did
	msg := &rtv1.MsgRegisterName{
		Creator:         acc.Address("snr"),
		NameToRegister:  username,
		PublicKeyBuffer: credential.PublicKey,
	}
	log.Println(msg.String())

	// broadcast a transaction from account `alice` with the message to create a did
	// store response in txResp
	txResp, err := s.cosmos.BroadcastTx(accountName, msg)
	if err != nil {
		log.Println(err)
		JsonResponse(w, "Failed to broadcast to blockchain", http.StatusBadRequest)
	}
	log.Println(txResp.String())

	// Return response from broadcasting a transaction
	user.AddCredential(*credential)
	JsonResponse(w, txResp.String(), http.StatusOK)
}

func (s *HighwayServer) BeginLogin(w http.ResponseWriter, r *http.Request) {

	// get username
	vars := mux.Vars(r)
	username := vars["username"]

	// get user
	user, err := s.userDb.GetUser(username)

	// user doesn't exist
	if err != nil {
		log.Println(err)
		JsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// generate PublicKeyCredentialRequestOptions, session data
	options, sessionData, err := s.auth.BeginLogin(user)
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

func (s *HighwayServer) FinishLogin(w http.ResponseWriter, r *http.Request) {

	// get username
	vars := mux.Vars(r)
	username := vars["username"]

	// get user
	user, err := s.userDb.GetUser(username)

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
	_, err = s.auth.FinishLogin(user, sessionData, r)
	if err != nil {
		log.Println(err)
		JsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// handle successful login
	JsonResponse(w, "Login Success", http.StatusOK)
}
