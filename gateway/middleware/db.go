package middleware

import (
	"net/http"
)



func SetDBMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    //timeoutContext, _ := context.WithTimeout(context.Background(), time.Second)
    //ctx := context.WithValue(r.Context(), "DB", db.WithContext(timeoutContext))
    //next.ServeHTTP(w, r.WithContext(ctx))
  })
}
