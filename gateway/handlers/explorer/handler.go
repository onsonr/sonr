package explorer

import (
	"net/http"
)
type Handler struct {
}

func (b Handler) IndexPage(w http.ResponseWriter, r *http.Request) {
	err := Home().Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (b Handler) GetLatestBlock(w http.ResponseWriter, r *http.Request)   {}

func (b Handler) GetServicesCount(w http.ResponseWriter, r *http.Request) {}

func (b Handler) GetIdentitiesCount(w http.ResponseWriter, r *http.Request) {}

func (b Handler) GetValidatorSet(w http.ResponseWriter, r *http.Request) {}

func (b Handler) GetBondedRatio(w http.ResponseWriter, r *http.Request) {}
