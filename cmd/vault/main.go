//go:build js && wasm
// +build js,wasm

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	wasmhttp "github.com/nlepage/go-wasm-http-server"
)

var db *sql.DB

func main() {
	// Initialize the database
	initDB()

	// Define your handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/items", itemsHandler)

	// Use wasmhttp.Serve to start the server
	wasmhttp.Serve(nil)
}

func initDB() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS items (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL
		)
	`)
	if err != nil {
		panic(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the WASM SQLite Server!")
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getItems(w, r)
	case "POST":
		addItem(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getItems(w http.ResponseWriter, _ *http.Request) {
	rows, err := db.Query("SELECT id, name FROM items")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []map[string]interface{}
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, map[string]interface{}{"id": id, "name": name})
	}

	json.NewEncoder(w).Encode(items)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	var item struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO items (name) VALUES (?)", item.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	json.NewEncoder(w).Encode(map[string]interface{}{"id": id, "name": item.Name})
}
