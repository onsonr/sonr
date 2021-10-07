package node

import (
	common "github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
)

// newInitResponse creates a response for the initialize request.
func (n *Node) newInitResponse(err error) *InitializeResponse {
	// Get Profile from Store
	profile, err := n.Profile()
	if err != nil {
		logger.Error("Failed to create initialize Response", err)
		return &InitializeResponse{
			Success: false,
			Error:   err.Error(),
		}
	}

	// Check for provided error
	if err != nil {
		return &InitializeResponse{
			Success: false,
			Error:   err.Error(),
		}
	}

	// Return Response
	return &InitializeResponse{
		Success: true,
		Profile: profile,
		//	Recents: r,
	}
}

// ToFindResponse converts PeerInfo to a FindResponse.
func ToFindResponse(p *common.PeerInfo) *SearchResponse {
	return &SearchResponse{
		Success: true,
		Peer:    p.Peer,
		PeerId:  p.PeerID.String(),
		SName:   p.SName,
	}
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
