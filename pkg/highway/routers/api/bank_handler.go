package api

import (
	"encoding/json"
	"net/http"

	bankv1beta1 "cosmossdk.io/api/cosmos/bank/v1beta1"
	"github.com/go-chi/chi/v5"

	"github.com/sonrhq/sonr/pkg/highway/middleware"
)

// BankHandler is a handler for the bank module
var BankHandler = bankHandler{}

// bankHandler is a handler for the bank module
type bankHandler struct{}

// GetAllBalances returns all balances for an address
func (h bankHandler) GetAllBalances(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")
	resp, err := middleware.NewBankClient(middleware.GrpcClientConn(r)).AllBalances(r.Context(), &bankv1beta1.QueryAllBalancesRequest{
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
func (h bankHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")
	denom := chi.URLParam(r, "denom")
	resp, err := middleware.NewBankClient(middleware.GrpcClientConn(r)).Balance(r.Context(), &bankv1beta1.QueryBalanceRequest{Address: address, Denom: denom})
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
func (h bankHandler) GetTotalSupply(w http.ResponseWriter, r *http.Request) {
	resp, err := middleware.NewBankClient(middleware.GrpcClientConn(r)).TotalSupply(r.Context(), &bankv1beta1.QueryTotalSupplyRequest{})
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
func (h bankHandler) GetSupplyOf(w http.ResponseWriter, r *http.Request) {
	denom := chi.URLParam(r, "denom")
	resp, err := middleware.NewBankClient(middleware.GrpcClientConn(r)).SupplyOf(r.Context(), &bankv1beta1.QuerySupplyOfRequest{Denom: denom})
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
func (h bankHandler) GetSpendableBalances(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")
	resp, err := middleware.NewBankClient(middleware.GrpcClientConn(r)).SpendableBalances(r.Context(), &bankv1beta1.QuerySpendableBalancesRequest{Address: address})
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
func (h bankHandler) GetSpendableBalancesByDenom(w http.ResponseWriter, r *http.Request) {
	address := chi.URLParam(r, "address")
	denom := chi.URLParam(r, "denom")
	resp, err := middleware.NewBankClient(middleware.GrpcClientConn(r)).SpendableBalanceByDenom(r.Context(), &bankv1beta1.QuerySpendableBalanceByDenomRequest{Address: address, Denom: denom})
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
func (h bankHandler) RegisterRoutes(r chi.Router) {
	r.Get("/balance/{address}", h.GetAllBalances)
	r.Get("/balance/{address}/spendable", h.GetSpendableBalances)
	r.Get("/balance/{address}/{denom}", h.GetBalance)
	r.Get("/balance/{address}/{denom}/spendable", h.GetSpendableBalancesByDenom)
	r.Get("/supply", h.GetTotalSupply)
	r.Get("/supply/{denom}", h.GetSupplyOf)
}
