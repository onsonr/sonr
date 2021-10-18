package mailbox

import (
	"context"

	"github.com/sonr-io/core/internal/api"
	"github.com/sonr-io/core/internal/fs"
	"github.com/sonr-io/core/internal/host"
)

type MailboxProtocol struct {
	ctx  context.Context
	host *host.SNRHost
	node api.NodeImpl
	// mail    *local.Mail
	//mailbox *local.Mailbox
}

// NewProtocol creates a new lobby protocol instance.
func NewProtocol(ctx context.Context, host *host.SNRHost, node api.NodeImpl) (*MailboxProtocol, error) {
	//mail := local.NewMail(cmd.NewClients(TextileClientURL, true, TextileMinerIdx), local.DefaultConfConfig())

	// Create Mailbox Protocol
	mailProtocol := &MailboxProtocol{
		ctx:  ctx,
		host: host,
		//	mail: mail,
		node: node,
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
	logger.Info("✅  MailboxProtocol is Activated \n")
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
// 	state, err := ts.mailbox.WatchInbox(context.Background(), events, true)
// 	if err != nil {
// 		logger.Error("Error watching Mailbox", err)
// 		return
// 	}

// 	// Handle Mailbox State
// 	for s := range state {
// 		// Update Connection State
// 		connState = s.State

// 		// handle connectivity state
// 		switch s.State {
// 		case cmd.Online:
// 			logger.Info(fmt.Sprintf("Mailbox is Online: %s", s.Err))
// 		case cmd.Offline:
// 			logger.Info(fmt.Sprintf("Mailbox is Offline: %s", s.Err))
// 		}
// 	}
// }

// // Handle New Mailbox Message
// func (ts *MailboxProtocol) onNewMessage(e local.MailboxEvent, state cmd.ConnectionState) {
// 	// List Total Inbox
// 	inbox, err := ts.mailbox.ListInboxMessages(context.Background())
// 	if err != nil {
// 		logger.Error("Failed to List Inbox Messages", err)
// 		return
// 	}

// 	// Logging and Open Body
// 	logger.Info(fmt.Sprintf("Received new message: %s", inbox[0].From))
// 	body, err := inbox[0].Open(context.Background(), ts.mailbox.Identity())
// 	if err != nil {
// 		logger.Error("Failed to Open Inbox Messages", err)
// 		return
// 	}

// 	// Log Valid Lobby Length
// 	logger.Info(fmt.Sprintf("Valid Body Length: %d", len(body)))

// 	// Unmarshal InviteRequest from JSON
// 	msg := MailboxMessage{}
// 	err = protojson.Unmarshal(body, &msg)
// 	if err != nil {
// 		logger.Error("Failed to Unmarshal Mailbox Message", err)
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
// 	err := ts.mailbox.DeleteInboxMessage(context.Background(), id)
// 	if err != nil {
// 		logger.Error("Failed to Delete Mailbox Message", err)
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
// 	err := ts.mailbox.ReadInboxMessage(context.Background(), id)
// 	if err != nil {
// 		logger.Error("Failed to set Mailbox Message as Read", err)
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
// 		logger.Error("Failed to Marshal Outbound Mailbox Message with JSON", err)
// 		return err
// 	}

// 	// 	// Send Message to Mailbox
// 	msg, err := ts.mailbox.SendMessage(context.Background(), to, buf)
// 	if err != nil {
// 		logger.Error(fmt.Sprintf("Failed to Send Message to Peer with ThreadIdentity: %s", to.String()), err)
// 		return err
// 	}

// 	// Log Result
// 	logger.Info("Succesfully sent mail!", golog.Fields{"ID": msg.ID, "SentAt": msg.CreatedAt, "To": msg.To.String()})
// 	return nil
// }
