// TODO: Update this Package to utlize: https://github.com/go-webauthn/example

package webauthn

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-webauthn/webauthn/webauthn"
)

var webAuthn *webauthn.WebAuthn
var userDB *userdb

// from: https://github.com/duo-labs/webauthn.io/blob/3f03b482d21476f6b9fb82b2bf1458ff61a61d41/server/response.go#L15
func JsonResponse(w http.ResponseWriter, d interface{}, c int) {
	dj, err := json.Marshal(d)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	fmt.Fprintf(w, "%s", dj)
}

func FinishLogin(w http.ResponseWriter, r *http.Request) {
	// get username/friendly name
	vals := r.URL.Query()
	username := vals.Get("username")
	if username == "" {
		JsonResponse(w, fmt.Errorf("must supply a valid username i.e. foo@bar.com"), http.StatusBadRequest)
		return
	}

	// get user
	//user, err := userDB.GetUser(username)

	// user doesn't exist
	// if err != nil {
	// 	log.Println(err)
	// 	JsonResponse(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// // load the session data
	// sessionData, err := sessionStore.GetWebauthnSession("authentication", r)
	// if err != nil {
	// 	log.Println(err)
	// 	JsonResponse(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// in an actual implementation, we should perform additional checks on
	// the returned 'credential', i.e. check 'credential.Authenticator.CloneWarning'
	// and then increment the credentials counter
	// _, err = webAuthn.FinishLogin(user, sessionData, r)
	// if err != nil {
	// 	log.Println(err)
	// 	JsonResponse(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// handle successful login
	JsonResponse(w, "Login Success", http.StatusOK)
}
