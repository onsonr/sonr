package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sonr-io/sonr/internal/highway"
)

func main() {
	// Start the app
	hw, err := highway.NewHighway(context.Background())
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()

	r.HandleFunc("/register/begin/{username}", hw.StartRegisterName).Methods("GET")
	r.HandleFunc("/register/finish/{username}", hw.FinishRegisterName).Methods("POST")
	r.HandleFunc("/login/begin/{username}", hw.StartAccessName).Methods("GET")
	r.HandleFunc("/login/finish/{username}", hw.FinishAccessName).Methods("POST")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./cmd/highwayd/")))

	serverAddress := ":8081"
	log.Println("starting server at", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, r))
}
