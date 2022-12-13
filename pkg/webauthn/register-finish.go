// TODO: Update this Package to utlize: https://github.com/go-webauthn/example

package webauthn

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type userdb struct {
	users map[string]*User
	mu    sync.RWMutex
}

var db *userdb

// DB returns a userdb singleton
func DB() *userdb {

	if db == nil {
		db = &userdb{
			users: make(map[string]*User),
		}
	}

	return db
}

// GetUser returns a *User by the user's username
func (db *userdb) GetUser(name string) (*User, error) {

	db.mu.Lock()
	defer db.mu.Unlock()
	user, ok := db.users[name]
	if !ok {
		return &User{}, fmt.Errorf("error getting user '%s': does not exist", name)
	}

	return user, nil
}

// PutUser stores a new user by the user's username
func (db *userdb) PutUser(user *User) {

	db.mu.Lock()
	defer db.mu.Unlock()
	db.users[user.name] = user
}

func FinishRegistration(w http.ResponseWriter, r *http.Request) {
	// get username/friendly name
	vals := r.URL.Query()
	username := vals.Get("username")
	if username == "" {
		JsonResponse(w, fmt.Errorf("must supply a valid username i.e. foo@bar.com"), http.StatusBadRequest)
		return
	}
	// get user
	_, err := userDB.GetUser(username)
	// user doesn't exist
	if err != nil {
		log.Println(err)
		JsonResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// // load the session data
	// sessionData, err := sessionStore.GetWebauthnSession("registration", r)
	// if err != nil {
	// 	log.Println(err)
	// 	JsonResponse(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// credential, err := webAuthn.FinishRegistration(user, sessionData, r)
	// if err != nil {
	// 	log.Println(err)
	// 	JsonResponse(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// user.AddCredential(*credential)

	JsonResponse(w, "Registration Success", http.StatusOK)
}
