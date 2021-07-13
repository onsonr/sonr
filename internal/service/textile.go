package service

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/internal/host"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/util"
	"github.com/textileio/go-threads/api/client"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/go-threads/db"
	"github.com/textileio/textile/v2/api/common"
	"github.com/textileio/textile/v2/cmd"
	"github.com/textileio/textile/v2/mail/local"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/proto"
)

type TextileService struct {
	ctxAuth  context.Context
	ctxToken context.Context

	// Parameters
	apiKeys     *md.APIKeys
	host        host.HostNode
	keyPair     *md.KeyPair
	options     *md.ConnectionRequest_TextileOptions
	onConnected md.OnConnected

	// Properties
	identity       thread.Identity
	client         *client.Client
	mail           *local.Mail
	mailbox        *local.Mailbox
	isMailReady    bool
	isThreadsReady bool
	isBucketsReady bool
}

// @ Set Service Status for Buckets
func (sc *TextileService) SetBucketsStatus(val bool) {
	sc.isBucketsReady = val
}

// @ Set Service Status for Mailbox
func (sc *TextileService) SetMailboxStatus(val bool) {
	sc.isMailReady = val
}

// @ Set Service Status for Threads
func (sc *TextileService) SetThreadsStatus(val bool) {
	sc.isThreadsReady = val
}

// @ Initializes New Textile Instance
func (sc *serviceClient) StartTextile() *md.SonrError {
	// Initialize
	textile := &TextileService{
		keyPair:     sc.user.KeyPair(),
		options:     sc.request.GetTextileOptions(),
		apiKeys:     sc.apiKeys,
		host:        sc.host,
		onConnected: sc.handler.OnConnected,
	}
	sc.Textile = textile

	// Check Textile Enabled
	if textile.options.GetEnabled() {
		// Initialize
		var err error
		creds := credentials.NewTLS(&tls.Config{})
		auth := common.Credentials{}

		// Dial GRPC
		textile.client, err = client.NewClient(util.TEXTILE_API_URL, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(auth))
		if err != nil {
			return md.NewError(err, md.ErrorMessage_TEXTILE_START_CLIENT)
		}

		// Get Identity
		textile.identity = getIdentity(textile.keyPair.PrivKey())

		// Create Auth Context
		textile.ctxAuth, err = newUserAuthCtx(context.Background(), textile.apiKeys)
		if err != nil {
			return md.NewError(err, md.ErrorMessage_TEXTILE_USER_CTX)
		}

		// Create Token Context
		textile.ctxToken, err = textile.newTokenCtx()
		if err != nil {
			return md.NewError(err, md.ErrorMessage_TEXTILE_TOKEN_CTX)
		}

		// Initialize Threads
		serr := textile.InitThreads(sc)
		if err != nil {
			return serr
		}

		// Initialize Mailbox
		serr = textile.InitMail(sc.user.GetDevice(), sc.request.GetStatus(), sc)
		if err != nil {
			return serr
		}
	}
	return nil
}

// @ Returns Instance Host
func (tn *TextileService) PubKey() thread.PubKey {
	return tn.identity.GetPublic()
}

// @ Initializes Threads
func (tn *TextileService) InitThreads(sc *serviceClient) *md.SonrError {
	// Verify Ready to Init
	if tn.ctxToken != nil {
		// Generate a new thread ID
		threadID := thread.NewIDV1(thread.Raw, 32)
		err := tn.client.NewDB(tn.ctxToken, threadID, db.WithNewManagedName("Sonr-Users"))
		if err != nil {
			return md.NewError(err, md.ErrorMessage_THREADS_START_NEW)
		}

		// Get DB Info
		allInfo, err := tn.client.ListDBs(tn.ctxToken)
		if err != nil {
			log.Println(err)
			return md.NewError(err, md.ErrorMessage_THREADS_LIST_ALL)
		}

		// Log DB Info
		log.Println("> Success!: Textile Threads Enabled -- > ALL DBs")
		for k, v := range allInfo {
			log.Println(fmt.Sprintf("Key: %s", k.String()))
			log.Println(fmt.Sprintf("ID: %s \n Maddr: %s \n Key: %s \n Name: %s \n", threadID.String(), v.Addrs, v.Key.String(), v.Name))
		}
		tn.SetThreadsStatus(true)
	}
	return nil
}

// @ Initializes Mailbox
func (tn *TextileService) InitMail(d *md.Device, us md.ConnectionRequest_UserStatus, sc *serviceClient) *md.SonrError {
	// Verify Ready to Initialize
	if tn.options.GetMailbox() {
		// Log
		log.Println("ACTIVATE: Textile Mailbox")

		// Setup the mail lib
		tn.mail = local.NewMail(cmd.NewClients(util.TEXTILE_API_URL, true, util.TEXTILE_MINER_IDX), local.DefaultConfConfig())

		// Create New Mailbox
		if us == md.ConnectionRequest_NEW {
			// Create a new mailbox with config
			mailbox, err := tn.mail.NewMailbox(context.Background(), local.Config{
				Path:      d.WorkingSupportDir(),
				Identity:  tn.identity,
				APIKey:    tn.apiKeys.GetTextileKey(),
				APISecret: tn.apiKeys.GetTextileSecret(),
			})

			// Check Error
			if err != nil {
				tn.SetMailboxStatus(false)
				return md.NewError(err, md.ErrorMessage_MAILBOX_START_NEW)
			}

			// Set Mailbox and Update Status
			tn.mailbox = mailbox
			log.Println("> Success!: Textile Mailbox Enabled, New Mailbox")
			tn.SetMailboxStatus(true)
		} else {
			// Return Existing Mailbox
			mailbox, err := tn.mail.GetLocalMailbox(context.Background(), d.WorkingSupportDir())
			if err != nil {
				tn.SetMailboxStatus(false)
				return md.NewError(err, md.ErrorMessage_MAILBOX_START_EXISTING)
			}

			// Set Mailbox and Update Status
			tn.mailbox = mailbox
			log.Println("> Success!: Textile Mailbox Enabled, Existing Mailbox")
			tn.SetMailboxStatus(true)
		}

		// Read Existing Mail
		err := sc.ReadMail()
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

// @ Read Mailbox Mail
func (tn *TextileService) readMail() (*md.MailEvent, *md.SonrError) {
	// List the recipient's inbox
	inbox, err := tn.mailbox.ListInboxMessages(context.Background())
	if err != nil {
		return nil, md.NewError(err, md.ErrorMessage_MAILBOX_LIST_ALL)
	}

	// Initialize Entry List
	hasNewMail := false
	count := len(inbox)

	// Set Vars
	entries := make([][]byte, count)
	if count > 0 {
		hasNewMail = true
	}

	// Iterate over Entries
	for i, v := range inbox {
		// Open decrypts the message body
		body, err := v.Open(context.Background(), tn.identity)
		if err != nil {
			return nil, md.NewError(err, md.ErrorMessage_MAILBOX_MESSAGE_OPEN)
		}
		entries[i] = body
	}
	return &md.MailEvent{
		Invites:    entries,
		HasNewMail: hasNewMail,
	}, nil
}

// @ Method Reads Inbox and Returns List of Mail Entries
func (sc *serviceClient) ReadMail() *md.SonrError {
	// Check Mail Enabled
	if sc.Textile.options.GetMailbox() {
		// Fetch Mail Event
		event, serr := sc.Textile.readMail()
		if serr != nil {
			serr.Print()
			return serr
		}
		// Create Mail and Marshal Data
		buf, err := proto.Marshal(event)
		if err != nil {
			return md.NewMarshalError(err)
		}

		// Callback Event
		sc.handler.OnMail(buf)
	}
	return nil
}

func (ts *TextileService) sendMail(to thread.PubKey, buf []byte) ([]byte, *md.SonrError) {
	// Send Message to Mailbox
	msg, err := ts.mailbox.SendMessage(context.Background(), to, buf)
	if err != nil {
		log.Println(err)
		return nil, md.NewError(err, md.ErrorMessage_MAILBOX_MESSAGE_SEND)
	}

	// Marshal Response
	buf, err = proto.Marshal(&md.InviteResponse{
		Type: md.InviteResponse_Mailbox,
		MailInfo: &md.InviteResponse_MailInfo{
			To:        msg.To.String(),
			From:      msg.From.String(),
			CreatedAt: int32(msg.CreatedAt.Unix()),
			ReadAt:    int32(msg.ReadAt.Unix()),
			Body:      msg.Body,
			Signature: msg.Signature,
		},
	})
	if err != nil {
		return nil, md.NewMarshalError(err)
	}

	// Return Message Info
	return buf, nil
}

// @ Method Sends Mail Entry to Peer
func (sc *serviceClient) SendMail(e *md.InviteRequest) *md.SonrError {
	// Check Mail Enabled
	if sc.Textile.options.GetMailbox() {
		// Fetch Peer Thread Key
		pubKey := e.ToThreadKey()
		log.Println(pubKey)

		// Marshal Data
		buf, err := proto.Marshal(e)
		if err != nil {
			log.Println(err)
			return md.NewMarshalError(err)
		}

		// Send to Mailbox
		resp, serr := sc.Textile.sendMail(pubKey, buf)
		if serr != nil {
			log.Println(err)
			return serr
		}
		sc.handler.OnReply(peer.ID(""), resp)
		log.Println("SUCCESS: Mail has been sent")
		return nil
	} else {
		log.Println("Mail is not Ready")
	}
	return nil
}
