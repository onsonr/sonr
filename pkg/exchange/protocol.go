package exchange

import (
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/fs"
	"github.com/sonr-io/core/internal/host"
	"github.com/sonr-io/core/pkg/common"
)

type ExchangeProtocol struct {
	ctx  context.Context
	host *host.SNRHost
	node api.NodeImpl
	// mail    *local.Mail
	//mailbox *local.Mailbox
	invites map[peer.ID]*InviteRequest
}

// New creates a new ExchangeProtocol
func New(ctx context.Context, host *host.SNRHost, node api.NodeImpl) (*ExchangeProtocol, error) {
	//mail := local.NewMail(cmd.NewClients(TextileClientURL, true, TextileMinerIdx), local.DefaultConfConfig())

	// Create Mailbox Protocol
	mailProtocol := &ExchangeProtocol{
		ctx:  ctx,
		host: host,
		//	mail: mail,
		node:    node,
		invites: make(map[peer.ID]*InviteRequest),
	}

	// Create new mailbox
	if fs.ThirdParty.Exists(TextileMailboxDirName) {
		// Return Existing Mailbox
		if err := mailProtocol.loadMailbox(); err != nil {
			return nil, err
		}
	} else {
		// Create New Mailbox
		if err := mailProtocol.newMailbox(); err != nil {
			return nil, err
		}
	}
	logger.Debug("✅  MailboxProtocol is Activated \n")
	// go mailProtocol.handleMailboxEvents()
	return mailProtocol, nil
}

// // // Handle Mailbox Events
// func (ts *MailboxProtocol) handleMailboxEvents() {
// 	// Initialize State
// 	connState := cmd.Online

// 	// Handle mailbox events as they arrive
// 	events := make(chan local.MailboxEvent)
// 	defer close(events)
// 	go func() {
// 		for e := range events {
// 			switch e.Type {
// 			case local.NewMessage:
// 				ts.onNewMessage(e, connState)
// 				continue
// 			}
// 		}
// 	}()

// 	// Start watching (the third param indicates we want to keep watching when offline)
// 	state, err := ts.mailbox.WatchInbox(ts.ctx, events, true)
// 	if err != nil {
// 		logger.Errorf("%s - Error watching Mailbox", err)
// 		return
// 	}

// 	// Handle Mailbox State
// 	for s := range state {
// 		// Update Connection State
// 		connState = s.State

// 		// handle connectivity state
// 		switch s.State {
// 		case cmd.Online:
// 			logger.Debug(fmt.Sprintf("Mailbox is Online: %s", s.Err))
// 		case cmd.Offline:
// 			logger.Debug(fmt.Sprintf("Mailbox is Offline: %s", s.Err))
// 		}
// 	}
// }

// // Handle New Mailbox Message
// func (ts *MailboxProtocol) onNewMessage(e local.MailboxEvent, state cmd.ConnectionState) {
// 	// List Total Inbox
// 	inbox, err := ts.mailbox.ListInboxMessages(ts.ctx)
// 	if err != nil {
// 		logger.Errorf("%s - Failed to List Inbox Messages", err)
// 		return
// 	}

// 	// Logging and Open Body
// 	logger.Debug(fmt.Sprintf("Received new message: %s", inbox[0].From))
// 	body, err := inbox[0].Open(ts.ctx, ts.mailbox.Identity())
// 	if err != nil {
// 		logger.Errorf("%s - Failed to Open Inbox Messages", err)
// 		return
// 	}

// 	// Log Valid Lobby Length
// 	logger.Debug(fmt.Sprintf("Valid Body Length: %d", len(body)))

// 	// Unmarshal InviteRequest from JSON
// 	msg := MailboxMessage{}
// 	err = protojson.Unmarshal(body, &msg)
// 	if err != nil {
// 		logger.Errorf("%s - Failed to Unmarshal Mailbox Message", err)
// 	}

// 	// Send Foreground Event
// 	if state == cmd.Online {
// 		// Create Mail Event
// 		mail := &api.MailboxEvent{
// 			To:     msg.GetTo(),
// 			From:   msg.GetFrom(),
// 			Buffer: msg.GetBuffer(),
// 			Id:     msg.GetId(),
// 		}

// 		// Send Mail Event
// 		ts.node.OnMailbox(mail)
// 	}
// }

// // Method Sends Mail Entry to Peer
// func (ts *MailboxProtocol) DeleteMail(id string) error {
// 	// Check if Mailbox exists
// 	if ts.mail == nil || ts.mailbox == nil {
// 		return ErrMailboxDisabled
// 	}

// 	// Mark Item as Read
// 	err := ts.mailbox.DeleteInboxMessage(ts.ctx, id)
// 	if err != nil {
// 		logger.Errorf("%s - Failed to Delete Mailbox Message", err)
// 		return err
// 	}
// 	return nil
// }

// // Method Sends Mail Entry to Peer
// func (ts *MailboxProtocol) ReadMail(id string) error {
// 	// Check if Mailbox exists
// 	if ts.mail == nil || ts.mailbox == nil {
// 		return ErrMailboxDisabled
// 	}

// 	// Mark Item as Read
// 	err := ts.mailbox.ReadInboxMessage(ts.ctx, id)
// 	if err != nil {
// 		logger.Errorf("%s - Failed to set Mailbox Message as Read", err)
// 		return err
// 	}
// 	return nil
// }

// // Method Sends Mail Entry to Peer
// func (ts *MailboxProtocol) SendMail(to thread.PubKey, message *MailboxMessage) error {
// 	// Check if Mailbox exists
// 	if ts.mail == nil || ts.mailbox == nil {
// 		return ErrMailboxDisabled
// 	}

// 	// Marshal Data
// 	buf, err := protojson.Marshal(message)
// 	if err != nil {
// 		logger.Errorf("%s - Failed to Marshal Outbound Mailbox Message with JSON", err)
// 		return err
// 	}

// 	// 	// Send Message to Mailbox
// 	msg, err := ts.mailbox.SendMessage(ts.ctx, to, buf)
// 	if err != nil {
// 		logger.Error(fmt.Sprintf("Failed to Send Message to Peer with ThreadIdentity: %s", to.String()), err)
// 		return err
// 	}

// 	// Log Result
// 	logger.Debug("Succesfully sent mail!", golog.Fields{"ID": msg.ID, "SentAt": msg.CreatedAt, "To": msg.To.String()})
// 	return nil
// }

// Validate takes list of Requests and returns true if Request exists in List and UUID is verified.
// Method also returns the InviteRequest that points to the Response.
func (sq *ExchangeProtocol) Validate(peer peer.ID, resp *InviteResponse) (*common.Payload, error) {
	// Authenticate Message
	valid := sq.host.AuthenticateMessage(resp, resp.Metadata)
	if !valid {
		return nil, ErrFailedAuth
	}

	// Check Decision
	if !resp.GetDecision() {
		return nil, nil
	}

	// Get Next Entry
	entry, ok := sq.invites[peer]
	if !ok {
		return nil, ErrFailedEntry
	}
	return entry.GetPayload(), nil
}

// Request Method sends a request to Transfer Data to a remote peer
func (p *ExchangeProtocol) Request(to *common.Peer) error {
	// Create Request
	id, req, err := p.createRequest(to)
	if err != nil {
		logger.Errorf("%s - Failed to Create Request", err)
		return err
	}

	// Check if the response is valid
	if req == nil {
		return ErrInvalidRequest
	}

	// sign the data
	signature, err := p.host.SignMessage(req)
	if err != nil {
		logger.Errorf("%s - Failed to Sign Response Message", err)
		return err
	}

	// add the signature to the message
	req.Metadata.Signature = signature
	err = p.host.SendMessage(id, RequestPID, req)
	if err != nil {
		logger.Errorf("%s - Failed to Send Message to Peer", err)
		return err
	}

	p.invites[id] = req
	return nil
}

// Respond Method authenticates or declines a Transfer Request
func (p *ExchangeProtocol) Respond(decs bool, to *common.Peer) error {
	// Create Response
	id, resp, err := p.createResponse(decs, to)
	if err != nil {
		logger.Errorf("%s - Failed to Create Request", err)
		return err
	}

	// Check if the response is valid
	if resp == nil {
		return ErrInvalidResponse
	}

	// sign the data
	signature, err := p.host.SignMessage(resp)
	if err != nil {
		logger.Errorf("%s - Failed to Sign Response Message", err)
		return err
	}

	// add the signature to the message
	resp.Metadata.Signature = signature

	// Send Response
	err = p.host.SendMessage(id, ResponsePID, resp)
	if err != nil {
		logger.Errorf("%s - Failed to Send Message to Peer", err)
		return err
	}
	return nil
}
