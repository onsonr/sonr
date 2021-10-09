package api

import (
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
)

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

// IsDev returns true if the node is running in development mode.
func (ir *InitializeRequest) IsDev() bool {
	return ir.GetEnvironment().IsDev()
}

// ToDeviceOpts converts InitializeRequest_DeviceOptions to device.Opts.
func (ir *InitializeRequest) ToDeviceOpts() []device.FSOption {
	fsOpts := make([]device.FSOption, 0)
	dOpts := ir.GetDeviceOptions()

	// Check Device Options
	if dOpts != nil {
		// Set Device ID
		err := device.SetDeviceID(dOpts.GetId())
		if err != nil {
			logger.Error("Failed to Set Device ID", err)
			return fsOpts
		}

		// Set Temporary Path
		fsOpts = append(fsOpts, device.FSOption{
			Path: dOpts.GetCacheDir(),
			Type: device.Temporary,
		}, device.FSOption{
			Path: dOpts.GetDownloadsDir(),
			Type: device.Downloads,
		}, device.FSOption{
			Path: dOpts.GetDocumentsDir(),
			Type: device.Documents,
		}, device.FSOption{
			Path: dOpts.GetSupportDir(),
			Type: device.Support,
		}, device.FSOption{
			Path: dOpts.GetDatabaseDir(),
			Type: device.Database,
		}, device.FSOption{
			Path: dOpts.GetTextileDir(),
			Type: device.Textile,
		})
	}
	return fsOpts
}
