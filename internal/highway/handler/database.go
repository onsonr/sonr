package handler

import (
	"context"

	databasepb "github.com/sonrhq/core/types/highway/database/v1"
)

// DatabaseHandler is the handler for the authentication service
type DatabaseHandler struct {}

// Health returns the health of the service.
//
// @Summary Health of the service
// @Description Returns the health of the service.
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string "Status message"
// @Router /getHealth [get]
func (a *DatabaseHandler) Health(ctx context.Context, req *databasepb.HealthRequest) (*databasepb.HealthResponse, error) {
	return &databasepb.HealthResponse{
		Ok: true,
		Message: "Sonr node is active and running.",
		Details: map[string]string{
			"api.address": "",
			"grpc.address": "",
			"highway.address": "",
			"rpc.address": "",
		},
	}, nil
}
