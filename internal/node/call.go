package node

import (
	context "context"
	"strings"

	api "github.com/sonr-io/core/internal/api"
)

// RPC_SERVER_PORT is the port the RPC service listens on.
const RPC_SERVER_PORT = 52006

// Supply supplies the node with the given amount of resources.
func (s *ClientNodeStub) Supply(ctx context.Context, req *api.SupplyRequest) (*api.SupplyResponse, error) {
	// Call Internal Supply
	err := s.TransferProtocol.Supply(req)
	if err != nil {
		return &api.SupplyResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	// Send Response
	return &api.SupplyResponse{
		Success: true,
	}, nil
}

// Edit method edits the node's properties in the Key/Value Store
func (s *ClientNodeStub) Edit(ctx context.Context, req *api.EditRequest) (*api.EditResponse, error) {
	// Call Internal Update
	if err := s.Update(); err != nil {
		return &api.EditResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	// Send Response
	return &api.EditResponse{
		Success: true,
	}, nil
}

// Fetch method retreives Node properties from Key/Value Store
func (s *ClientNodeStub) Fetch(ctx context.Context, req *api.FetchRequest) (*api.FetchResponse, error) {
	// Call Internal Fetch4
	profile, err := s.node.Profile()
	if err != nil {
		return &api.FetchResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	recents, err := s.node.GetRecents()
	if err != nil {
		return &api.FetchResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	history, err := s.node.GetHistory()
	if err != nil {
		return &api.FetchResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	// Send Response
	return &api.FetchResponse{
		Success: true,
		Profile: profile,
		Recents: recents,
		History: history,
	}, nil
}

// Share method sends supplied files/urls with a peer
func (s *ClientNodeStub) Share(ctx context.Context, req *api.ShareRequest) (*api.ShareResponse, error) {
	// Request Peer to Transfer File
	if s.TransferProtocol != nil {
		err := s.TransferProtocol.Request(req.GetPeer())
		if err != nil {
			return &api.ShareResponse{
				Success: false,
				Error:   err.Error(),
			}, nil
		}
	} else {
		return &api.ShareResponse{
			Success: false,
			Error:   ErrProtocolsNotSet.Error(),
		}, nil
	}

	// Send Response
	return &api.ShareResponse{
		Success: true,
	}, nil
}

// Search Method to find a Peer by SName
func (s *ClientNodeStub) Search(ctx context.Context, req *api.SearchRequest) (*api.SearchResponse, error) {
	// Call Internal Ping
	if s.ExchangeProtocol != nil {
		// Call Internal Search
		entry, err := s.Get(strings.ToLower(req.GetSName()))
		if err != nil {
			return &api.SearchResponse{
				Success: false,
				Error:   err.Error(),
			}, nil
		}

		// Send Response
		return &api.SearchResponse{
			Success: true,
			Peer:    entry,
		}, nil
	} else {
		return &api.SearchResponse{
			Success: false,
			Error:   ErrProtocolsNotSet.Error(),
		}, nil
	}

}

// Respond method responds to a received InviteRequest.
func (s *ClientNodeStub) Respond(ctx context.Context, req *api.RespondRequest) (*api.RespondResponse, error) {
	// Call Internal Respond
	if s.TransferProtocol != nil {
		// Respond on TransferProtocol
		err := s.TransferProtocol.Respond(req.GetDecision(), req.GetPeer())
		if err != nil {
			return &api.RespondResponse{
				Success: false,
				Error:   err.Error(),
			}, nil
		}

		// Send Response
		return &api.RespondResponse{
			Success: true,
		}, nil
	} else {
		return &api.RespondResponse{
			Success: false,
			Error:   ErrProtocolsNotSet.Error(),
		}, nil
	}

}

// Authorize Signing Method Request for Data
func (hrc *HighwayNodeStub) Authorize(ctx context.Context, req *api.AuthorizeRequest) (*api.AuthorizeResponse, error) {
	logger.Debug("HighwayService.Authorize() is Unimplemented")
	return nil, nil
}

// Link a new Device to the Node
func (hrc *HighwayNodeStub) Link(ctx context.Context, req *api.LinkRequest) (*api.LinkResponse, error) {
	logger.Debug("HighwayService.Link() is Unimplemented")
	return nil, nil
}

// Register a new domain with the Node on the highway
func (hrc *HighwayNodeStub) Register(ctx context.Context, req *api.RegisterRequest) (*api.RegisterResponse, error) {
	// Get Values
	pfix := req.GetPrefix()
	name := req.GetSName()
	fprint := req.GetFingerprint()

	// Check Values
	if pfix == "" || name == "" || fprint == "" {
		return &api.RegisterResponse{
			Success: false,
			Error:   "Invalid request. One or more of the required fields are empty.",
		}, nil
	}

	// // Create Record
	// resp, err := hrc.DomainProtocol.Register(name, exchange.NewNBAuthRecord(pfix, name, fprint))
	// if err != nil {
	// 	return &api.RegisterResponse{
	// 		Success: false,
	// 		Error:   err.Error(),
	// 	}, nil
	// }

	// Return Response
	return &api.RegisterResponse{
		Success: true,
		// Records: resp,
	}, nil
}
