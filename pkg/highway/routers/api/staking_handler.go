package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	stakingv1beta1 "cosmossdk.io/api/cosmos/staking/v1beta1"
	"github.com/go-chi/chi/v5"

	"github.com/sonrhq/sonr/pkg/highway/middleware"
)

// StakingHandler is a handler for the staking module
var StakingHandler = stakingHandler{}

// stakingHandler is a handler for the staking module
type stakingHandler struct{}

// GetDelegation returns a delegation
func (h stakingHandler) GetDelegation(w http.ResponseWriter, r *http.Request) {
	delegatorAddr := chi.URLParam(r, "delegatorAddr")
	validatorAddr := chi.URLParam(r, "validatorAddr")
	resp, err := middleware.StakingClient(r, w).Delegation(r.Context(), &stakingv1beta1.QueryDelegationRequest{
		DelegatorAddr: delegatorAddr,
		ValidatorAddr: validatorAddr,
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

// GetUnbondingDelegation returns an unbonding delegation
func (h stakingHandler) GetUnbondingDelegation(w http.ResponseWriter, r *http.Request) {
	delegatorAddr := chi.URLParam(r, "delegatorAddr")
	validatorAddr := chi.URLParam(r, "validatorAddr")
	resp, err := middleware.StakingClient(r, w).UnbondingDelegation(r.Context(), &stakingv1beta1.QueryUnbondingDelegationRequest{
		DelegatorAddr: delegatorAddr,
		ValidatorAddr: validatorAddr,
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

// GetDelegatorDelegations returns all delegations for a delegator
func (h stakingHandler) GetDelegatorDelegations(w http.ResponseWriter, r *http.Request) {
	delegatorAddr := chi.URLParam(r, "delegatorAddr")
	resp, err := middleware.StakingClient(r, w).DelegatorDelegations(r.Context(), &stakingv1beta1.QueryDelegatorDelegationsRequest{
		DelegatorAddr: delegatorAddr,
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

// GetDelegatorUnbondingDelegations returns all unbonding delegations for a delegator
func (h stakingHandler) GetDelegatorUnbondingDelegations(w http.ResponseWriter, r *http.Request) {
	delegatorAddr := chi.URLParam(r, "delegatorAddr")
	resp, err := middleware.StakingClient(r, w).DelegatorUnbondingDelegations(r.Context(), &stakingv1beta1.QueryDelegatorUnbondingDelegationsRequest{
		DelegatorAddr: delegatorAddr,
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

// GetRedelegations returns all redelegations for a delegator
func (h stakingHandler) GetRedelegations(w http.ResponseWriter, r *http.Request) {
	delegatorAddr := chi.URLParam(r, "delegatorAddr")
	srcValidatorAddr := chi.URLParam(r, "srcValidatorAddr")
	dstValidatorAddr := chi.URLParam(r, "dstValidatorAddr")
	resp, err := middleware.StakingClient(r, w).Redelegations(r.Context(), &stakingv1beta1.QueryRedelegationsRequest{
		DelegatorAddr:    delegatorAddr,
		SrcValidatorAddr: srcValidatorAddr,
		DstValidatorAddr: dstValidatorAddr,
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

// GetValidator returns a validator
func (h stakingHandler) GetValidator(w http.ResponseWriter, r *http.Request) {
	validatorAddr := chi.URLParam(r, "validatorAddr")
	resp, err := middleware.StakingClient(r, w).Validator(r.Context(), &stakingv1beta1.QueryValidatorRequest{
		ValidatorAddr: validatorAddr,
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

// GetValidators returns all validators
func (h stakingHandler) GetValidators(w http.ResponseWriter, r *http.Request) {
	resp, err := middleware.StakingClient(r, w).Validators(r.Context(), &stakingv1beta1.QueryValidatorsRequest{})
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

// GetValidatorDelegations returns all delegations for a validator
func (h stakingHandler) GetValidatorDelegations(w http.ResponseWriter, r *http.Request) {
	validatorAddr := chi.URLParam(r, "validatorAddr")
	resp, err := middleware.StakingClient(r, w).ValidatorDelegations(r.Context(), &stakingv1beta1.QueryValidatorDelegationsRequest{
		ValidatorAddr: validatorAddr,
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

// GetDelegatorValidators returns all validators for a delegator
func (h stakingHandler) GetDelegatorValidators(w http.ResponseWriter, r *http.Request) {
	delegatorAddr := chi.URLParam(r, "delegatorAddr")
	resp, err := middleware.StakingClient(r, w).DelegatorValidators(r.Context(), &stakingv1beta1.QueryDelegatorValidatorsRequest{
		DelegatorAddr: delegatorAddr,
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

// GetHistoricalInfo returns historical info
func (h stakingHandler) GetHistoricalInfo(w http.ResponseWriter, r *http.Request) {
	heightStr := chi.URLParam(r, "height")
	height, _ := strconv.ParseInt(heightStr, 10, 64)
	resp, err := middleware.StakingClient(r, w).HistoricalInfo(r.Context(), &stakingv1beta1.QueryHistoricalInfoRequest{
		Height: height,
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

// RegisterRoutes registers the staking routes
func (h stakingHandler) RegisterRoutes(r chi.Router) {
	r.Get("/delegators/{delegatorAddr}", h.GetDelegatorDelegations)
	r.Get("/delegators/{delegatorAddr}/unbonding", h.GetDelegatorUnbondingDelegations)
	r.Get("/delegators/{delegatorAddr}/validators", h.GetDelegatorValidators)
	r.Get("/delegators/{delegatorAddr}/validators/{validatorAddr}", h.GetDelegation)
	r.Get("/delegators/{delegatorAddr}/validators/{validatorAddr}/unbonding", h.GetUnbondingDelegation)
	r.Get("/delegators/{delegatorAddr}/validators/{srcValidatorAddr}/redelegate/{dstValidatorAddr}", h.GetRedelegations)
	r.Get("/history/{height}", h.GetHistoricalInfo)
	r.Get("/validators", h.GetValidators)
	r.Get("/validators/{validatorAddr}", h.GetValidator)
	r.Get("/validators/{validatorAddr}/delegations", h.GetValidatorDelegations)
}
