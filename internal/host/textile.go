package host

import (
	"time"

	"github.com/textileio/textile/v2/api/common"
)

// Textile Client API URL
const TEXTILE_API_URL = "api.hub.textile.io:443"

// Textile Miner Index Target
const TEXTILE_MINER_IDX = "api.minerindex.hub.textile.io:443"

// Textile Mailbox Directory
const TEXTILE_MAILBOX_DIR = ".textile"

func (h *SNRHost) startTextileClient(apiKey string, apiSecret string) error {
	// Add our device group key to the context
	ctx := common.NewAPIKeyContext(h.ctxHost, apiKey)

	// Add a signature using our device group secret
	authCtx, err := common.CreateAPISigContext(ctx, time.Now().Add(time.Hour*2), apiSecret)
	if err != nil {
		return err
	}
	h.ctxTileAuth = authCtx

	// Start the Textile client
	return nil
}

// func (h *SNRHost) threadIdentity() thread.Identity {
// 	return thread.NewLibp2pIdentity(h.privKey)
// }

// var (
// 	isMailReady    = false
// 	isThreadsReady = false
// 	isBucketsReady = false
// )

// type TextileService struct {
// 	ctxAuth  context.Context
// 	ctxToken context.Context

// 	// Parameters
// 	options *data.ConnectionRequest_ServiceOptions
// 	emitter *emitter.Emitter

// 	// Properties
// 	client  *client.Client
// 	mail    *local.Mail
// 	mailbox *local.Mailbox
// }

// // Starts New Textile Instance
// func (sc *serviceClient) StartTextile() *data.SonrError {
// 	// Initialize
// 	textile := &TextileService{
// 		options: sc.request.GetServiceOptions(),
// 		apiKeys: sc.apiKeys,
// 		host:    sc.host,
// 		device:  sc.device,
// 		emitter: sc.emitter,
// 	}
// 	sc.Textile = textile

// 	// Check Textile Enabled
// 	if textile.options.GetTextile() {
// 		// Initialize
// 		var err error
// 		creds := credentials.NewTLS(&tls.Config{})
// 		auth := common.Credentials{}

// 		// Dial GRPC
// 		textile.client, err = client.NewClient(util.TEXTILE_API_URL, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(auth))
// 		if err != nil {
// 			return data.NewError(err, data.ErrorEvent_TEXTILE_START_CLIENT)
// 		}

// 		// Create Auth Context
// 		textile.ctxAuth, err = newUserAuthCtx(context.Background(), textile.apiKeys)
// 		if err != nil {
// 			return data.NewError(err, data.ErrorEvent_TEXTILE_USER_CTX)
// 		}

// 		// Create Token Context
// 		textile.ctxToken, err = textile.newTokenCtx()
// 		if err != nil {
// 			return data.NewError(err, data.ErrorEvent_TEXTILE_TOKEN_CTX)
// 		}

// 		// Initialize Threads
// 		serr := textile.InitThreads(sc)
// 		if err != nil {
// 			return serr
// 		}

// 		// Initialize Mailbox
// 		serr = textile.InitMail()
// 		if err != nil {
// 			return serr
// 		}
// 	}
// 	return nil
// }

// // Returns Instance Host
// func (ts *TextileService) PubKey() thread.PubKey {
// 	return ts.device.ThreadIdentity().GetPublic()
// }

// // Initializes Threads
// func (ts *TextileService) InitThreads(sc *serviceClient) *data.SonrError {
// 	// Verify Ready to Init
// 	if ts.ctxToken != nil {
// 		// Generate a new thread ID
// 		threadID := thread.NewIDV1(thread.Raw, 32)
// 		err := ts.client.NewDB(ts.ctxToken, threadID, db.WithNewManagedName("Sonr-Users"))
// 		if err != nil {
// 			return data.NewError(err, data.ErrorEvent_THREADS_START_NEW)
// 		}
// 		isThreadsReady = true
// 	}
// 	return nil
// }

// // Initializes Mailbox
// func (ts *TextileService) InitMail() *data.SonrError {
// 	// Verify Ready to Initialize
// 	if ts.options.GetMailbox() {
// 		// Setup the mail lib
// 		ts.mail = local.NewMail(cmd.NewClients(util.TEXTILE_API_URL, true, util.TEXTILE_MINER_IDX), local.DefaultConfConfig())

// 		// Create New Mailbox
// 		if ts.hasMailboxDirectory() {
// 			// Return Existing Mailbox
// 			mailbox, err := ts.mail.GetLocalMailbox(context.Background(), ts.device.WorkingSupportDir())
// 			if err != nil {
// 				return data.NewError(err, data.ErrorEvent_MAILBOX_START_EXISTING)
// 			}

// 			// Set Mailbox and Update Status
// 			ts.mailbox = mailbox
// 			isMailReady = true

// 			// Handle Mailbox Events
// 			ts.handleMailboxEvents()
// 		} else {
// 			// Create a new mailbox with config
// 			mailbox, err := ts.mail.NewMailbox(context.Background(), ts.defaultMailConfig())
// 			if err != nil {
// 				return data.NewError(err, data.ErrorEvent_MAILBOX_START_NEW)
// 			}

// 			// Set Mailbox and Update Status
// 			ts.mailbox = mailbox
// 			isMailReady = true

// 			// Handle Mailbox Events
// 			ts.handleMailboxEvents()
// 		}
// 	}
// 	return nil
// }

// // Handle Mailbox Events
// func (ts *TextileService) handleMailboxEvents() {
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
// 		data.NewError(err, data.ErrorEvent_MAILBOX_EVENT_STATE)
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
// func (ts *TextileService) onNewMessage(e local.MailboxEvent, state cmd.ConnectionState) {
// 	// List Total Inbox
// 	inbox, err := ts.mailbox.ListInboxMessages(context.Background())
// 	if err != nil {
// 		data.NewError(err, data.ErrorEvent_MAILBOX_MESSAGE_OPEN)
// 		return
// 	}

// 	// Logging and Open Body
// 	logger.Info(fmt.Sprintf("Received new message: %s", inbox[0].From))
// 	body, err := inbox[0].Open(context.Background(), ts.mailbox.Identity())
// 	if err != nil {
// 		data.NewError(err, data.ErrorEvent_MAILBOX_MESSAGE_OPEN)
// 		return
// 	}

// 	// Log Valid Lobby Length
// 	logger.Info(fmt.Sprintf("Valid Body Length: %d", len(body)))

// 	// Unmarshal InviteRequest from JSON
// 	invite := data.InviteRequest{}
// 	err = protojson.Unmarshal(body, &invite)
// 	if err != nil {
// 		data.NewError(err, data.ErrorEvent_MAILBOX_MESSAGE_UNMARSHAL)
// 	}

// 	// Send Foreground Event
// 	if state == cmd.Online {
// 		// Create Mail Event
// 		mail := &data.MailEvent{
// 			To:        inbox[0].To.String(),
// 			From:      inbox[0].From.String(),
// 			CreatedAt: int32(inbox[0].CreatedAt.Unix()),
// 			ReadAt:    int32(inbox[0].ReadAt.Unix()),
// 			Invite:    &invite,
// 			Signature: inbox[0].Signature,
// 			ID:        inbox[0].ID,
// 		}
// 		// Callback Mail Event
// 		ts.emitter.Emit(emitter.EMIT_MAIL_EVENT, mail)
// 	}

// }

// // Send Mail to Recipient
// func (ts *TextileService) sendMail(to thread.PubKey, buf []byte) *data.SonrError {
// 	// Send Message to Mailbox
// 	_, err := ts.mailbox.SendMessage(context.Background(), to, buf)
// 	if err != nil {
// 		return data.NewError(err, data.ErrorEvent_MAILBOX_MESSAGE_SEND)
// 	}

// 	// Return Message Info
// 	return nil
// }

// // Method sets message in inbox as read
// func (ts *TextileService) deleteMessage(id string) *data.SonrError {
// 	// Mark Item as Read
// 	err := ts.mailbox.DeleteInboxMessage(context.Background(), id)
// 	if err != nil {
// 		return data.NewError(err, data.ErrorEvent_MAILBOX_MESSAGE_DELETE)
// 	}
// 	return nil
// }

// // Method sets message in inbox as read
// func (ts *TextileService) readMessage(id string) *data.SonrError {
// 	// Mark Item as Read
// 	err := ts.mailbox.ReadInboxMessage(context.Background(), id)
// 	if err != nil {
// 		return data.NewError(err, data.ErrorEvent_MAILBOX_MESSAGE_READ)
// 	}
// 	return nil
// }

// // Helper: Creates New Textile Mailbox Config
// func (ts *TextileService) defaultMailConfig() local.Config {
// 	return local.Config{
// 		Path:      ts.device.WorkingSupportDir(),
// 		Identity:  ts.device.ThreadIdentity(),
// 		APIKey:    ts.apiKeys.GetTextileKey(),
// 		APISecret: ts.apiKeys.GetTextileSecret(),
// 	}
// }

// // Helper: Checks if Device Has Mailbox Directory
// func (ts *TextileService) hasMailboxDirectory() bool {
// 	return ts.device.GetFileSystem().IsDirectory(ts.device.FileSystem.Support, util.TEXTILE_MAILBOX_DIR)
// }

// // Helper: Creates User Auth Context from API Keys
// func newUserAuthCtx(ctx context.Context, keys *data.APIKeys) (context.Context, error) {
// 	// Add our device group key to the context
// 	ctx = common.NewAPIKeyContext(ctx, keys.TextileKey)

// 	// Add a signature using our device group secret
// 	return common.CreateAPISigContext(ctx, time.Now().Add(time.Hour), keys.TextileSecret)
// }

// // Helper: Creates Auth Token Context from AuthContext, Client, Identity
// func (ts *TextileService) newTokenCtx() (context.Context, error) {
// 	// Generate a new token for the device
// 	token, err := ts.client.GetToken(ts.ctxAuth, ts.device.ThreadIdentity())
// 	if err != nil {
// 		return nil, err
// 	}
// 	return thread.NewTokenContext(ts.ctxAuth, token), nil
// }
