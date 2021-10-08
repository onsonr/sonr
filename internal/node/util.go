package node

import (
	"errors"

	"github.com/libp2p/go-libp2p-core/peer"
	common "github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/pkg/transfer"
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

// Share a peer to have a transfer
func (n *Node) NewRequest(to *common.Peer) (peer.ID, *transfer.InviteRequest, error) {
	// Fetch Element from Queue
	elem := n.queue.Front()
	if elem != nil {
		// Get Payload
		payload := n.queue.Remove(elem).(*common.Payload)

		// Create New ID for Invite
		id, err := device.KeyChain.CreateUUID()
		if err != nil {
			logger.Error("Failed to create new id for Shared Invite", err)
			return "", nil, err
		}

		// Create new Metadata
		meta, err := device.KeyChain.CreateMetadata(n.host.ID())
		if err != nil {
			logger.Error("Failed to create new metadata for Shared Invite", err)
			return "", nil, err
		}

		// Fetch User Peer
		from, err := n.Peer()
		if err != nil {
			logger.Error("Failed to get Node Peer Object", err)
			return "", nil, err
		}

		// Create Invite Request
		req := &transfer.InviteRequest{
			Payload:  payload,
			Metadata: common.SignedMetadataToProto(meta),
			To:       to,
			From:     from,
			Uuid:     common.SignedUUIDToProto(id),
		}

		// Fetch Peer ID from Public Key
		toId, err := to.PeerID()
		if err != nil {
			logger.Error("Failed to fetch peer id from public key", err)
			return "", nil, err
		}
		return toId, req, nil
	}
	return "", nil, errors.New("No items in Transfer Queue.")
}

// Respond to an invite request
func (n *Node) NewResponse(decs bool, to *common.Peer) (peer.ID, *transfer.InviteResponse, error) {
	// Create new Metadata
	meta, err := device.KeyChain.CreateMetadata(n.host.ID())
	if err != nil {
		logger.Error("Failed to create new metadata for Shared Invite", err)
		return "", nil, err
	}

	// Fetch User Peer
	from, err := n.Peer()
	if err != nil {
		return "", nil, err
	}

	// Create Invite Response
	resp := &transfer.InviteResponse{
		Decision: decs,
		Metadata: common.SignedMetadataToProto(meta),
		From:     from,
		To:       to,
	}

	// Fetch Peer ID from Public Key
	toId, err := to.PeerID()
	if err != nil {
		logger.Error("Failed to fetch peer id from public key", err)
		return "", nil, err
	}
	return toId, resp, nil
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
