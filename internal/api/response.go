package api

import common "github.com/sonr-io/core/pkg/common"

func NewInitialzeResponse(gpf common.GetProfileFunc, success bool) *InitializeResponse {
	resp := &InitializeResponse{Success: success}
	if !success || gpf == nil {
		return resp
	}
	p, err := gpf()
	if err != nil {
		logger.Error("Failed to get profile", err)
		return resp
	}
	resp.Profile = p
	return resp
}
