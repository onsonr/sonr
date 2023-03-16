package keeper

import (
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sonrhq/core/x/identity/types"
	"github.com/sonrhq/core/x/identity/types/typesconnect"
)

var (
	connServer *msgServerConnectWrapper
)

type msgServerConnectWrapper struct {
	Keeper
}

type msgServer struct {
	Keeper
	Connect *msgServerConnectWrapper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func RegisterConnectServer(keeper Keeper, mux *runtime.ServeMux) {
	connSrv := &msgServerConnectWrapper{Keeper: keeper}
	api := http.NewServeMux()
	ep, handler := typesconnect.NewMsgHandler(connSrv)
	typesconnect.NewMsgHandler(connSrv)
	api.Handle(ep, handler)
	api.Handle("/", handler)
}
