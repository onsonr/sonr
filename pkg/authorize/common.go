package authorize

import (
	"errors"
	"os"
	"time"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/fs"
	"github.com/sonr-io/core/internal/wallet"
	common "github.com/sonr-io/core/pkg/common"
)

// Textile API definitions
const (
	// Textile Client API URL
	TextileClientURL = "https://api.textile.io"

	// Textile Miner Index Target
	TextileMinerIdx = "api.minerindex.hub.textile.io:443"

	// Textile Mailbox Directory
	TextileMailboxDirName = "mailbox"
)

var (
	logger              = golog.Default.Child("protocols/mailbox")
	ErrMailboxDisabled  = errors.New("Mailbox not enabled, cannot perform request.")
	ErrMissingAPIKey    = errors.New("Missing Textile API Key in env")
	ErrMissingAPISecret = errors.New("Missing Textile API Secret in env")
	ErrFailedEntry      = errors.New("Failed to get Topmost entry from Queue")
	ErrFailedAuth       = errors.New("Failed to Authenticate message")
	ErrEmptyRequests    = errors.New("Empty Request list provided")
	ErrRequestNotFound  = errors.New("Request not found in list")
	ErrInvalidResponse  = errors.New("Invalid InviteResponse provided to TransmitProtocol")
	ErrInvalidRequest   = errors.New("Invalid InviteRequest provided to TransmitProtocol")
)

// fetchApiKeys fetches the Textile Api/Secrect keys from the environment
func fetchApiKeys() (string, string, error) {
	// Get API Key
	key, ok := os.LookupEnv("TEXTILE_HUB_KEY")
	if !ok {
		return "", "", ErrMissingAPIKey
	}

	// Get API Secret
	secret, ok := os.LookupEnv("TEXTILE_HUB_SECRET")
	if !ok {
		return "", "", ErrMissingAPISecret
	}
	return key, secret, nil
}

// getMailboxPath returns the mailbox path from the device
func (mb *AuthorizeProtocol) getMailboxPath() (string, error) {
	// Get Mailbox Path
	path, err := fs.ThirdParty.GenPath(TextileMailboxDirName)
	if err != nil {
		logger.Errorf("%s - Failed to Find Existing Mailbox at Path", err)
		return "", err
	}
	return path, nil
}

// loadMailbox loads an existing mailbox instance
func (mb *AuthorizeProtocol) loadMailbox() error {
	logger.Debug("Loading Mailbox...")

	// Get Mailbox Path
	path, err := mb.getMailboxPath()
	if err != nil {
		logger.Errorf("%s - Failed to Create New Mailbox at Path", err)
		return err
	}

	// // Return Existing Mailbox
	// mailbox, err := mb.mail.GetLocalMailbox(mb.ctx, path)
	// if err != nil {
	// 	logger.Errorf("%s - Failed to Load Existing Mailbox", err)
	// 	return err
	// }

	// // Set mailbox
	// mb.mailbox = mailbox
	logger.Debug("Existing Mailbox has been loaded.", golog.Fields{"path": path})
	return nil
}

// newMailbox creates a new mailbox instance
func (mb *AuthorizeProtocol) newMailbox() error {
	logger.Debug("Creating new Mailbox...")

	// Get Mailbox Path
	path, err := mb.getMailboxPath()
	if err != nil {
		logger.Errorf("%s - Failed to Create New Mailbox at Path", err)
		return err
	}

	// Fetch API Keys
	// key, secret, err := fetchApiKeys()
	// if err != nil {
	// 	logger.Errorf("%s - Failed to create new mailbox", err)
	// 	return err
	// }

	// // Create a new mailbox with config
	// mailbox, err := mb.mail.NewMailbox(mb.ctx, local.Config{
	// 	Path: path,
	// 	// Identity:  privKey.ThreadIdentity(),
	// 	APIKey:    key,
	// 	APISecret: secret,
	// })

	// // Check if Err is for ErrMailboxExists
	// if err == local.ErrMailboxExists {
	// 	logger.Debug("Mailbox already exists no need to create a new one")
	// 	// Load Existing Mailbox
	// 	return mb.loadMailbox()
	// }

	// // Check for errors
	// if err != nil {
	// 	logger.Errorf("%s - Failed to create mailbox", err)
	// 	return err
	// }

	// // Set mailbox
	// mb.mailbox = mailbox
	logger.Debug("New Mailbox has been created.", golog.Fields{"path": path})
	return nil
}

// ToEvent method on InviteResponse converts InviteResponse to DecisionEvent.
func (ir *InviteResponse) ToEvent() *api.DecisionEvent {
	return &api.DecisionEvent{
		From:     ir.GetFrom(),
		Received: int64(time.Now().Unix()),
		Decision: ir.GetDecision(),
	}
}

// ToEvent method on InviteRequest converts InviteRequest to InviteEvent.
func (ir *InviteRequest) ToEvent() *api.InviteEvent {
	return &api.InviteEvent{
		Received: int64(time.Now().Unix()),
		From:     ir.GetFrom(),
		Payload:  ir.GetPayload(),
	}
}

// createRequest creates a new InviteRequest
func (p *AuthorizeProtocol) createRequest(to *common.Peer) (peer.ID, *InviteRequest, error) {
	// Call Peer from Node
	from, err := p.node.Peer()
	if err != nil {
		logger.Errorf("%s - Failed to Get Peer from Node", err)
		return "", nil, err
	}

	// Fetch Peer ID from Public Key
	toId, err := to.Libp2pID()
	if err != nil {
		logger.Errorf("%s - Failed to fetch peer id from public key", err)
		return "", nil, err
	}

	// Fetch Element from Queue
	// Get Next Entry
	entry, ok := p.invites[toId]
	if ok {
		// Create new Metadata
		meta, err := wallet.Sonr.CreateMetadata(p.host.ID())
		if err != nil {
			logger.Errorf("%s - Failed to create new metadata for Shared Invite", err)
			return "", nil, err
		}

		// Create Invite Request
		req := &InviteRequest{
			Payload:  entry.GetPayload(),
			Metadata: api.SignedMetadataToProto(meta),
			To:       to,
			From:     from,
		}
		return toId, req, nil
	}
	logger.Errorf("%s - Failed to get item from Supply Queue.")
	return "", nil, errors.New("No items in Supply Queue.")
}

// createResponse creates a new InviteResponse
func (p *AuthorizeProtocol) createResponse(decs bool, to *common.Peer) (peer.ID, *InviteResponse, error) {

	// Call Peer from Node
	from, err := p.node.Peer()
	if err != nil {
		logger.Errorf("%s - Failed to Get Peer from Node", err)
		return "", nil, err
	}

	// Create new Metadata
	meta, err := wallet.Sonr.CreateMetadata(p.host.ID())
	if err != nil {
		logger.Errorf("%s - Failed to create new metadata for Shared Invite", err)
		return "", nil, err
	}

	// Create Invite Response
	resp := &InviteResponse{
		Decision: decs,
		Metadata: api.SignedMetadataToProto(meta),
		From:     from,
		To:       to,
	}

	// Fetch Peer ID from Public Key
	toId, err := to.Libp2pID()
	if err != nil {
		logger.Errorf("%s - Failed to fetch peer id from public key", err)
		return "", nil, err
	}
	return toId, resp, nil
}
