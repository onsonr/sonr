package ctx

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type SonrContext struct {
	sdk.Context
}

func GetSonrCTX(ctx sdk.Context) *SonrContext {
	return &SonrContext{ctx}
}

func (s *SonrContext) GetBlockExpiration(duration time.Duration) int64 {
	blockTime := s.BlockTime()
	avgBlockTime := float64(blockTime.Sub(blockTime).Seconds())
	return int64(duration.Seconds() / avgBlockTime)
}
