package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	govv1 "cosmossdk.io/api/cosmos/gov/v1"
	"github.com/go-chi/chi/v5"

	"github.com/sonrhq/sonr/pkg/highway/middleware"
)

// GovHandler is a handler for the gov module
var GovHandler = govHandler{}

// govHandler is a handler for the gov module
type govHandler struct{}

// GetConstitution returns the constitution
func (h govHandler) GetConstitution(w http.ResponseWriter, r *http.Request) {
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

// GetProposal returns a proposal
func (h govHandler) GetProposal(w http.ResponseWriter, r *http.Request) {
	proposalIDStr := chi.URLParam(r, "proposalId")
	i, _ := strconv.ParseUint(proposalIDStr, 10, 64)
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

// GetProposals returns all proposals
func (h govHandler) GetProposals(w http.ResponseWriter, r *http.Request) {
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

// GetVote returns a vote
func (h govHandler) GetVote(w http.ResponseWriter, r *http.Request) {
	proposalIDStr := chi.URLParam(r, "proposalId")
	voterStr := chi.URLParam(r, "voter")
	i, _ := strconv.ParseUint(proposalIDStr, 10, 64)
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

// GetVotes returns all votes for a proposal
func (h govHandler) GetVotes(w http.ResponseWriter, r *http.Request) {
	proposalIDStr := chi.URLParam(r, "proposalId")
	i, _ := strconv.ParseUint(proposalIDStr, 10, 64)
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

// GetDeposit returns a deposit
func (h govHandler) GetDeposit(w http.ResponseWriter, r *http.Request) {
	proposalIDStr := chi.URLParam(r, "proposalId")
	depositorStr := chi.URLParam(r, "depositor")
	i, _ := strconv.ParseUint(proposalIDStr, 10, 64)
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

// GetDeposits returns all deposits for a proposal
func (h govHandler) GetDeposits(w http.ResponseWriter, r *http.Request) {
	proposalIDStr := chi.URLParam(r, "proposalId")
	i, _ := strconv.ParseUint(proposalIDStr, 10, 64)
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

// GetTally returns the tally for a proposal
func (h govHandler) GetTally(w http.ResponseWriter, r *http.Request) {
	proposalIDStr := chi.URLParam(r, "proposalId")
	i, _ := strconv.ParseUint(proposalIDStr, 10, 64)
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

// RegisterRoutes registers the gov routes
func (h govHandler) RegisterRoutes(r chi.Router) {
	r.Get("/constitution", h.GetConstitution)
	r.Get("/proposals", h.GetProposals)
	r.Get("/proposals/{proposalId}", h.GetProposal)
	r.Get("/proposals/{proposalId}/deposits", h.GetDeposits)
	r.Get("/proposals/{proposalId}/deposits/{depositor}", h.GetDeposit)
	r.Get("/proposals/{proposalId}/tally", h.GetTally)
	r.Get("/proposals/{proposalId}/votes", h.GetVotes)
	r.Get("/proposals/{proposalId}/votes/{voter}", h.GetVote)
}
