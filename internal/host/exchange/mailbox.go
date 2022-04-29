package exchange

import (
	"os"

	"github.com/kataras/golog"
	device "github.com/sonr-io/sonr/pkg/fs"
)

// // // Handle Mailbox Events
// func (ts *ExchangeProtocol) handleMailboxEvents() {
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
// func (ts *ExchangeProtocol) onNewMessage(e local.MailboxEvent, state cmd.ConnectionState) {
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
// func (ts *ExchangeProtocol) DeleteMail(id string) error {
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
// func (ts *ExchangeProtocol) ReadMail(id string) error {
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
// func (ts *ExchangeProtocol) SendMail(to thread.PubKey, message *MailboxMessage) error {
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
func (mb *ExchangeProtocol) getMailboxPath() (string, error) {
	// Get Mailbox Path
	path, err := device.ThirdParty.GenPath(TextileMailboxDirName)
	if err != nil {
		logger.Errorf("%s - Failed to Find Existing Mailbox at Path", err)
		return "", err
	}
	return path, nil
}

// loadMailbox loads an existing mailbox instance
func (mb *ExchangeProtocol) loadMailbox() error {
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
func (mb *ExchangeProtocol) newMailbox() error {
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
