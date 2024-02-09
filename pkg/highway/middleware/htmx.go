package middleware

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/donseba/go-htmx"
)

func HTMX(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		hxh := htmx.HxRequestHeaderFromRequest(r)

		ctx = context.WithValue(ctx, htmx.ContextRequestHeader, hxh)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

// HTMXResponse writes a Templ component to the http.ResponseWriter
func HTMXResponse(view templ.Component) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := view.Render(r.Context(), w); err != nil {
			RenderError(w, err)
		}
	}
}
