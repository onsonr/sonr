package service

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"

	"github.com/sonr-io/core/internal/host"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/util"
	"github.com/textileio/go-threads/api/client"
	"github.com/textileio/go-threads/core/thread"
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
	identity thread.Identity
	client   *client.Client
	mail     *local.Mail
	mailbox  *local.Mailbox
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
		log.Println("ENABLE: Textile Service")

		// Initialize
		var err error
		creds := credentials.NewTLS(&tls.Config{})
		auth := common.Credentials{}

		// Dial GRPC
		textile.client, err = client.NewClient(util.TEXTILE_API_URL, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(auth))
		if err != nil {
			return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
		}

		// Get Identity
		textile.identity = getIdentity(textile.keyPair.PrivKey())

		// Create Auth Context
		textile.ctxAuth, err = newUserAuthCtx(context.Background(), textile.apiKeys)
		if err != nil {
			return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
		}

		// Create Token Context
		textile.ctxToken, err = textile.newTokenCtx()
		if err != nil {
			return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
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
		// Log
		log.Println("ACTIVATE: Textile Threads")

		// Generate a new thread ID
		threadID := thread.NewIDV1(thread.Raw, 32)
		err := tn.client.NewDB(tn.ctxToken, threadID)
		if err != nil {
			return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
		}

		// Get DB Info
		info, err := tn.client.GetDBInfo(tn.ctxToken, threadID)
		if err != nil {
			return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
		}

		// Log DB Info
		log.Println("> Success!: Textile Threads Enabled")
		log.Println(fmt.Sprintf("ID: %s \n Maddr: %s \n Key: %s \n Name: %s \n", threadID.String(), info.Addrs, info.Key.String(), info.Name))
	}

	// Update Status
	if !tn.options.GetMailbox() {
		sc.status.EnableTextile(true, false)
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
				sc.status.EnableTextile(true, false)
				return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
			}

			// Set Mailbox and Update Status
			tn.mailbox = mailbox
			log.Println("> Success!: Textile Mailbox Enabled, New Mailbox")
			sc.status.EnableTextile(true, true)
		} else {
			// Return Existing Mailbox
			mailbox, err := tn.mail.GetLocalMailbox(context.Background(), d.WorkingSupportDir())
			if err != nil {
				sc.status.EnableTextile(true, false)
				return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
			}

			// Set Mailbox and Update Status
			tn.mailbox = mailbox
			log.Println("> Success!: Textile Mailbox Enabled, Existing Mailbox")
			sc.status.EnableTextile(true, true)
		}

		// Read Existing Mai.
		mailevent, err := sc.ReadMail()
		if err != nil {
			log.Println(err)
		}
		sc.handler.OnMail(mailevent)
	}
	return nil
}

// @ Method Reads Inbox and Returns List of Mail Entries
func (sc *serviceClient) ReadMail() (*md.MailEvent, *md.SonrError) {
	// Check Mail Enabled
	if sc.HasMailbox() {
		// List the recipient's inbox
		inbox, err := sc.Textile.mailbox.ListInboxMessages(context.Background())

		if err != nil {
			return nil, md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
		}

		// Initialize Entry List
		entries := make([]*md.InviteRequest, len(inbox))

		// Iterate over Entries
		for i, v := range inbox {
			// Open decrypts the message body
			body, err := v.Open(context.Background(), sc.Textile.identity)
			if err != nil {
				return nil, md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
			}

			// Unmarshal Body to entry
			entry := &md.InviteRequest{}
			err = proto.Unmarshal(body, entry)
			if err != nil {
				return nil, md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
			}
			entries[i] = entry
		}

		return &md.MailEvent{
			Invites: entries,
		}, nil
	}
	return nil, nil
}

// @ Method Sends Mail Entry to Peer
func (sc *serviceClient) SendMail(e *md.InviteRequest) *md.SonrError {
	// Check Mail Enabled
	if sc.HasMailbox() {
		pubKey := e.GetTo().ThreadKey()

		buf, err := proto.Marshal(e)
		if err != nil {
			return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
		}
		// Send Message to Mailbox
		_, err = sc.Textile.mailbox.SendMessage(context.Background(), pubKey, buf)

		// Check Error
		if err != nil {
			return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
		}
	}
	return nil
}

// @ Helper: Checks if Mailbox Enabled
func (sc *serviceClient) HasMailbox() bool {
	if sc.Textile.options.GetMailbox() && sc.Textile.options.GetEnabled() {
		return sc.status.Textile == md.ServiceStatus_FULL
	}
	return false
}
