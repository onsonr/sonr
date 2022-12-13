// TODO: Update this Package to utlize: https://github.com/go-webauthn/example

package webauthn

import (
	"fmt"
	"log"
	"net/http"
)

func BeginLogin(w http.ResponseWriter, r *http.Request) {
	// get username/friendly name
	vals := r.URL.Query()
	username := vals.Get("username")
	if username == "" {
		JsonResponse(w, fmt.Errorf("must supply a valid username i.e. foo@bar.com"), http.StatusBadRequest)
		return
	}

	// get user
	user, err := userDB.GetUser(username)

	// user doesn't exist
	if err != nil {
		log.Println(err)
		JsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// generate PublicKeyCredentialRequestOptions, session data
	options, _, err := webAuthn.BeginLogin(user)
	if err != nil {
		log.Println(err)
		JsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// The above code is saving the session data to the session store.
	// // store session data as marshaled JSON
	// err = sessionStore.SaveWebauthnSession("authentication", sessionData, r, w)
	// if err != nil {
	// 	log.Println(err)
	// 	JsonResponse(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	JsonResponse(w, options, http.StatusOK)
}
