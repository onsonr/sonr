package middleware

import (
	"encoding/json"
	"net/http"
)

// JSONResponse writes a JSON response to the http.ResponseWriter
func JSONResponse(w http.ResponseWriter, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bz, err := json.Marshal(body)
	if err != nil {
		InternalServerError(w, err)
		return
	}
	w.Write(bz)
}

// InternalServerError writes an internal server error to the http.ResponseWriter
func InternalServerError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

// BadRequest writes a bad request error to the http.ResponseWriter
func BadRequest(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}

// NotFound writes a not found error to the http.ResponseWriter
func NotFound(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusNotFound)
}

// Unauthorized writes an unauthorized error to the http.ResponseWriter
func Unauthorized(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusUnauthorized)
}

// Forbidden writes a forbidden error to the http.ResponseWriter
func Forbidden(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusForbidden)
}
