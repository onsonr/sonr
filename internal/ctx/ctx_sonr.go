package ctx

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type SonrContext struct {
	sdk.Context
}

func GetSonrCTX(ctx sdk.Context) *SonrContext {
	return &SonrContext{ctx}
}
