package modulesapi

import (
	"encoding/json"
	"net/http"
	"strconv"

	stakingv1beta1 "cosmossdk.io/api/cosmos/staking/v1beta1"
	"github.com/go-chi/chi/v5"

	"github.com/sonrhq/sonr/app/gateway/middleware"
)

type StakingHandler struct{}

func (h StakingHandler) GetDelegation(w http.ResponseWriter, r *http.Request) {
	delegatorAddr := chi.URLParam(r, "delegatorAddr")
	validatorAddr := chi.URLParam(r, "validatorAddr")
	resp, err := middleware.NewStakingClient(middleware.GrpcClientConn(r)).Delegation(r.Context(), &stakingv1beta1.QueryDelegationRequest{
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

func (h StakingHandler) GetUnbondingDelegation(w http.ResponseWriter, r *http.Request) {
	delegatorAddr := chi.URLParam(r, "delegatorAddr")
	validatorAddr := chi.URLParam(r, "validatorAddr")
	resp, err := middleware.NewStakingClient(middleware.GrpcClientConn(r)).UnbondingDelegation(r.Context(), &stakingv1beta1.QueryUnbondingDelegationRequest{
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

func (h StakingHandler) GetDelegatorDelegations(w http.ResponseWriter, r *http.Request) {
	delegatorAddr := chi.URLParam(r, "delegatorAddr")
	resp, err := middleware.NewStakingClient(middleware.GrpcClientConn(r)).DelegatorDelegations(r.Context(), &stakingv1beta1.QueryDelegatorDelegationsRequest{
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

func (h StakingHandler) GetDelegatorUnbondingDelegations(w http.ResponseWriter, r *http.Request) {
	delegatorAddr := chi.URLParam(r, "delegatorAddr")
	resp, err := middleware.NewStakingClient(middleware.GrpcClientConn(r)).DelegatorUnbondingDelegations(r.Context(), &stakingv1beta1.QueryDelegatorUnbondingDelegationsRequest{
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

func (h StakingHandler) GetRedelegations(w http.ResponseWriter, r *http.Request) {
	delegatorAddr := chi.URLParam(r, "delegatorAddr")
	srcValidatorAddr := chi.URLParam(r, "srcValidatorAddr")
	dstValidatorAddr := chi.URLParam(r, "dstValidatorAddr")
	resp, err := middleware.NewStakingClient(middleware.GrpcClientConn(r)).Redelegations(r.Context(), &stakingv1beta1.QueryRedelegationsRequest{
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

func (h StakingHandler) GetValidator(w http.ResponseWriter, r *http.Request) {
	validatorAddr := chi.URLParam(r, "validatorAddr")
	resp, err := middleware.NewStakingClient(middleware.GrpcClientConn(r)).Validator(r.Context(), &stakingv1beta1.QueryValidatorRequest{
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

func (h StakingHandler) GetValidators(w http.ResponseWriter, r *http.Request) {
	resp, err := middleware.NewStakingClient(middleware.GrpcClientConn(r)).Validators(r.Context(), &stakingv1beta1.QueryValidatorsRequest{})
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

func (h StakingHandler) GetValidatorDelegations(w http.ResponseWriter, r *http.Request) {
	validatorAddr := chi.URLParam(r, "validatorAddr")
	resp, err := middleware.NewStakingClient(middleware.GrpcClientConn(r)).ValidatorDelegations(r.Context(), &stakingv1beta1.QueryValidatorDelegationsRequest{
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

func (h StakingHandler) GetDelegatorValidators(w http.ResponseWriter, r *http.Request) {
	delegatorAddr := chi.URLParam(r, "delegatorAddr")
	resp, err := middleware.NewStakingClient(middleware.GrpcClientConn(r)).DelegatorValidators(r.Context(), &stakingv1beta1.QueryDelegatorValidatorsRequest{
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

func (h StakingHandler) GetHistoricalInfo(w http.ResponseWriter, r *http.Request) {
	heightStr := chi.URLParam(r, "height")
	height, _ := strconv.ParseInt(heightStr, 10, 64)
	resp, err := middleware.NewStakingClient(middleware.GrpcClientConn(r)).HistoricalInfo(r.Context(), &stakingv1beta1.QueryHistoricalInfoRequest{
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
