package render

import (
	"net/http"

	"github.com/a-h/templ"
)

func TemplComponent(comp templ.Component) (handlerFn http.HandlerFunc) {
    return func(w http.ResponseWriter, r *http.Request) {
        err := comp.Render(r.Context(), w)
        if err!= nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
}
