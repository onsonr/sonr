package modulesapi

import (
	"encoding/json"
	"net/http"
	"strconv"

	govv1 "cosmossdk.io/api/cosmos/gov/v1"
	"github.com/go-chi/chi/v5"

	"github.com/sonrhq/sonr/app/highway/middleware"
)

type GovHandler struct{}

func (h GovHandler) GetConstitution(w http.ResponseWriter, r *http.Request) {
	res, err := middleware.NewGovClient(middleware.GrpcClientConn(r)).Constitution(r.Context(), &govv1.QueryConstitutionRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

func (h GovHandler) GetProposal(w http.ResponseWriter, r *http.Request) {
	proposalIdStr := chi.URLParam(r, "proposalId")
	i, _ := strconv.ParseUint(proposalIdStr, 10, 64)
	res, err := middleware.NewGovClient(middleware.GrpcClientConn(r)).Proposal(r.Context(), &govv1.QueryProposalRequest{ProposalId: i})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

func (h GovHandler) GetProposals(w http.ResponseWriter, r *http.Request) {
	res, err := middleware.NewGovClient(middleware.GrpcClientConn(r)).Proposals(r.Context(), &govv1.QueryProposalsRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

func (h GovHandler) GetVote(w http.ResponseWriter, r *http.Request) {
	proposalIdStr := chi.URLParam(r, "proposalId")
	voterStr := chi.URLParam(r, "voter")
	i, _ := strconv.ParseUint(proposalIdStr, 10, 64)
	res, err := middleware.NewGovClient(middleware.GrpcClientConn(r)).Vote(r.Context(), &govv1.QueryVoteRequest{ProposalId: i, Voter: voterStr})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

func (h GovHandler) GetVotes(w http.ResponseWriter, r *http.Request) {
	proposalIdStr := chi.URLParam(r, "proposalId")
	i, _ := strconv.ParseUint(proposalIdStr, 10, 64)
	res, err := middleware.NewGovClient(middleware.GrpcClientConn(r)).Votes(r.Context(), &govv1.QueryVotesRequest{ProposalId: i})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

func (h GovHandler) GetDeposit(w http.ResponseWriter, r *http.Request) {
	proposalIdStr := chi.URLParam(r, "proposalId")
	depositorStr := chi.URLParam(r, "depositor")
	i, _ := strconv.ParseUint(proposalIdStr, 10, 64)
	res, err := middleware.NewGovClient(middleware.GrpcClientConn(r)).Deposit(r.Context(), &govv1.QueryDepositRequest{ProposalId: i, Depositor: depositorStr})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

func (h GovHandler) GetDeposits(w http.ResponseWriter, r *http.Request) {
	proposalIdStr := chi.URLParam(r, "proposalId")
	i, _ := strconv.ParseUint(proposalIdStr, 10, 64)
	res, err := middleware.NewGovClient(middleware.GrpcClientConn(r)).Deposits(r.Context(), &govv1.QueryDepositsRequest{ProposalId: i})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}

func (h GovHandler) GetTally(w http.ResponseWriter, r *http.Request) {
	proposalIdStr := chi.URLParam(r, "proposalId")
	i, _ := strconv.ParseUint(proposalIdStr, 10, 64)
	res, err := middleware.NewGovClient(middleware.GrpcClientConn(r)).TallyResult(r.Context(), &govv1.QueryTallyResultRequest{ProposalId: i})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rBz, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rBz)
}
