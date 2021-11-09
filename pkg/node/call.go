package node

import (
	context "context"
	"strings"

	api "github.com/sonr-io/core/pkg/api"
)

// Edit method edits the node's properties in the Key/Value Store
func (s *NodeMotorStub) Edit(ctx context.Context, req *api.EditRequest) (*api.EditResponse, error) {
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
func (s *NodeMotorStub) Fetch(ctx context.Context, req *api.FetchRequest) (*api.FetchResponse, error) {
	// Call Lobby Update
	if err := s.Update(); err != nil {
		logger.Warnf("%s - Failed to Update Lobby", err)
	}

	// Call Internal Fetch4
	profile, err := s.node.Profile()
	if err != nil {
		return &api.FetchResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	recents, err := s.node.identity.GetRecents()
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
	}, nil
}

// Share method sends supplied files/urls with a peer
func (s *NodeMotorStub) Share(ctx context.Context, req *api.ShareRequest) (*api.ShareResponse, error) {
	// Call Lobby Update
	if err := s.Update(); err != nil {
		logger.Warnf("%s - Failed to Update Lobby", err)
	}

	// Request Peer to Transmit File
	if s.TransmitProtocol != nil {
		err := s.ExchangeProtocol.Request(req)
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
func (s *NodeMotorStub) Search(ctx context.Context, req *api.SearchRequest) (*api.SearchResponse, error) {
	// Call Lobby Update
	if err := s.Update(); err != nil {
		logger.Warnf("%s - Failed to Update Lobby", err)
	}

	// Call Internal Ping
	if s.DiscoverProtocol != nil {
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
func (s *NodeMotorStub) Respond(ctx context.Context, req *api.RespondRequest) (*api.RespondResponse, error) {
	// Call Lobby Update
	if err := s.Update(); err != nil {
		logger.Warnf("%s - Failed to Update Lobby", err)
	}

	// buildRespFunc is a function that builds a response to a received InviteRequest.
	buildRespFunc := func(err error) *api.RespondResponse {
		if err == nil {
			return &api.RespondResponse{
				Success: true,
			}
		}
		return &api.RespondResponse{
			Success: false,
			Error:   err.Error(),
		}
	}

	// Get Request Parameters
	decs := req.GetDecision()
	peer := req.GetPeer()

	// Respond on ExchangeProtocol
	payload, err := s.ExchangeProtocol.Respond(decs, peer)
	if err != nil {
		return buildRespFunc(err), nil
	}

	// Check decision
	if decs {
		// Prepare on TransmitProtocol
		if err := s.TransmitProtocol.Incoming(payload, peer); err != nil {
			return buildRespFunc(err), nil
		}
	}

	// Send Response
	return buildRespFunc(nil), nil

}

// Authorize Signing Method Request for Data
func (s *NodeHighwayStub) Authorize(ctx context.Context, req *api.AuthenticateRequest) (*api.AuthenticateResponse, error) {
	logger.Debug("HighwayStub.Authorize() is Unimplemented")
	return nil, nil
}

// Link a new Device to the Node
func (s *NodeHighwayStub) Link(ctx context.Context, req *api.LinkRequest) (*api.LinkResponse, error) {
	logger.Debug("HighwayStub.Link() is Unimplemented")
	return nil, nil
}

// Register a new domain with the Node on the highway
func (s *NodeHighwayStub) Register(ctx context.Context, req *api.RegisterRequest) (*api.RegisterResponse, error) {
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

	// Create Record
	resp, err := s.RegistryProtocol.Register(req)
	if err != nil {
		return &api.RegisterResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	// Return Response
	return &api.RegisterResponse{
		Success: true,
		Records: resp,
	}, nil
}
