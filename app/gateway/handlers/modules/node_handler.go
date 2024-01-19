package modulesapi

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	"github.com/go-chi/chi/v5"

	"github.com/sonrhq/sonr/app/gateway/middleware"
)

type NodeHandler struct{}

func (h NodeHandler) GetLatestBlock(w http.ResponseWriter, r *http.Request) {
	res, err := middleware.NewCometClient(middleware.GrpcClientConn(r)).GetLatestBlock(r.Context(), &cmtservice.GetLatestBlockRequest{})
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

func (h NodeHandler) GetNodeInfo(w http.ResponseWriter, r *http.Request) {
	res, err := middleware.NewCometClient(middleware.GrpcClientConn(r)).GetNodeInfo(r.Context(), &cmtservice.GetNodeInfoRequest{})
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

func (h NodeHandler) GetSyncing(w http.ResponseWriter, r *http.Request) {
	res, err := middleware.NewCometClient(middleware.GrpcClientConn(r)).GetSyncing(r.Context(), &cmtservice.GetSyncingRequest{})
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

func (h NodeHandler) GetBlockByHeight(w http.ResponseWriter, r *http.Request) {
	heightStr := chi.URLParam(r, "height")
	i, _ := strconv.ParseInt(heightStr, 10, 64)
	res, err := middleware.NewCometClient(middleware.GrpcClientConn(r)).GetBlockByHeight(r.Context(), &cmtservice.GetBlockByHeightRequest{Height: i})
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

func (h NodeHandler) GetLatestValidatorSet(w http.ResponseWriter, r *http.Request) {
	res, err := middleware.NewCometClient(middleware.GrpcClientConn(r)).GetLatestValidatorSet(r.Context(), &cmtservice.GetLatestValidatorSetRequest{})
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

func (h NodeHandler) GetValidatorSetByHeight(w http.ResponseWriter, r *http.Request) {
	heightStr := chi.URLParam(r, "height")
	i, _ := strconv.ParseInt(heightStr, 10, 64)
	res, err := middleware.NewCometClient(middleware.GrpcClientConn(r)).GetValidatorSetByHeight(r.Context(), &cmtservice.GetValidatorSetByHeightRequest{Height: i})
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
