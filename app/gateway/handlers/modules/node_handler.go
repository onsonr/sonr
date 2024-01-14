package modules_api

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"

	"github.com/sonrhq/sonr/pkg/context"
)

type NodeHandler struct{}

func (h NodeHandler) GetLatestBlock(w http.ResponseWriter, r *http.Request) {
	req := &cmtservice.GetLatestBlockRequest{}
	cmtClient := cmtservice.NewServiceClient(context.Get().GrpcConn())
	res, err := cmtClient.GetLatestBlock(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bz, err := res.Marshal()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bz)
}
