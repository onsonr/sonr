package landing

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"

	"github.com/sonrhq/sonr/pkg/context"
	landing_ui "github.com/sonrhq/sonr/web/htmx/home/ui"
)

type Handler struct{}

func (b Handler) IndexPage(w http.ResponseWriter, r *http.Request) {
	err := landing_ui.Home().Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (b Handler) GetLatestBlock(w http.ResponseWriter, r *http.Request) {
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

func (b Handler) GetServicesCount(w http.ResponseWriter, r *http.Request) {}

func (b Handler) GetIdentitiesCount(w http.ResponseWriter, r *http.Request) {}

func (b Handler) GetValidatorSet(w http.ResponseWriter, r *http.Request) {}

func (b Handler) GetBondedRatio(w http.ResponseWriter, r *http.Request) {}
