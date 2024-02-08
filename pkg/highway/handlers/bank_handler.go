package handlers

import (
	"encoding/json"
	"net/http"

	types "cosmossdk.io/api/cosmos/bank/v1beta1"
	"github.com/go-chi/chi/v5"

	"github.com/sonrhq/sonr/pkg/highway/middleware"
)

// ! ||--------------------------------------------------------------------------------||
// ! ||                                  API Endpoints                                 ||
// ! ||--------------------------------------------------------------------------------||

// BankAPI is a handler for the bank module
var BankAPI = bankAPI{}

// bankAPI is a handler for the bank module
type bankAPI struct{}

// GetAllBalances returns all balances for an address
func (h bankAPI) GetAllBalances(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")
	resp, err := middleware.BankClient(r, w).AllBalances(r.Context(), &types.QueryAllBalancesRequest{
		Address: address,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

// GetBalance returns a balance for an address and denom
func (h bankAPI) GetBalance(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")
	denom := chi.URLParam(r, "denom")
	resp, err := middleware.BankClient(r, w).Balance(r.Context(), &types.QueryBalanceRequest{Address: address, Denom: denom})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

// GetTotalSupply returns the total supply
func (h bankAPI) GetTotalSupply(w http.ResponseWriter, r *http.Request) {
	resp, err := middleware.BankClient(r, w).TotalSupply(r.Context(), &types.QueryTotalSupplyRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

// GetSupplyOf returns the supply of a denom
func (h bankAPI) GetSupplyOf(w http.ResponseWriter, r *http.Request) {
	denom := chi.URLParam(r, "denom")
	resp, err := middleware.BankClient(r, w).SupplyOf(r.Context(), &types.QuerySupplyOfRequest{Denom: denom})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

// GetSpendableBalances returns the spendable balances for an address
func (h bankAPI) GetSpendableBalances(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")
	resp, err := middleware.BankClient(r, w).SpendableBalances(r.Context(), &types.QuerySpendableBalancesRequest{Address: address})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

// GetSpendableBalancesByDenom returns the spendable balances for an address and denom
func (h bankAPI) GetSpendableBalancesByDenom(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")
	denom := chi.URLParam(r, "denom")
	resp, err := middleware.BankClient(r, w).SpendableBalanceByDenom(r.Context(), &types.QuerySpendableBalanceByDenomRequest{Address: address, Denom: denom})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

// RegisterRoutes registers the bank routes
func (h bankAPI) RegisterRoutes(r chi.Router) {
	r.Get("/balance/{address}", h.GetAllBalances)
	r.Get("/balance/{address}/spendable", h.GetSpendableBalances)
	r.Get("/balance/{address}/{denom}", h.GetBalance)
	r.Get("/balance/{address}/{denom}/spendable", h.GetSpendableBalancesByDenom)
	r.Get("/supply", h.GetTotalSupply)
	r.Get("/supply/{denom}", h.GetSupplyOf)
}
