package bind

import (
	"context"

	md "github.com/sonr-io/core/internal/models"
)

// * Struct: Reference for Binded Proxy Node * //
type mobileConfig struct {
	CtxFS   context.Context
	CtxNode context.Context
	CtxUser context.Context

	HasStarted      bool
	HasBootstrapped bool
	HasJoinedLocal  bool

	Status md.Status
}

func newMobileConfig() mobileConfig {
	return mobileConfig{
		CtxFS:           context.Background(),
		CtxNode:         context.Background(),
		CtxUser:         context.Background(),
		
		HasStarted:      false,
		HasBootstrapped: false,
		HasJoinedLocal:  false,

		Status:          md.Status_NONE,
	}
}

func (mc *mobileConfig) contextFS() context.Context {
	return mc.CtxFS
}

func (mc *mobileConfig) contextNode() context.Context {
	return mc.CtxNode
}

func (mc *mobileConfig) contextUser() context.Context {
	return mc.CtxUser
}

func (mc *mobileConfig) setConnected(val bool) {
	mc.HasStarted = val
	mc.Status = md.Status_CONNECTED
}

func (mc *mobileConfig) setBootstrapped(val bool) {
	mc.HasBootstrapped = val
	mc.Status = md.Status_BOOTSTRAPPED
}

func (mc *mobileConfig) setJoinedLocal(val bool) {
	mc.HasJoinedLocal = val
	mc.Status = md.Status_AVAILABLE
}

func (mc *mobileConfig) isReady() bool {
	return mc.HasBootstrapped && mc.HasStarted
}
