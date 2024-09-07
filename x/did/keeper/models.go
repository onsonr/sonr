package keeper

import (
	didv1 "github.com/onsonr/sonr/api/did/v1"
	"github.com/onsonr/sonr/x/did/types"
)

func convertServiceRecord(rec *didv1.ServiceRecord) *types.Service {
	return &types.Service{
		Origin: rec.OriginUri,
	}
}
