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
	active   bool
	ctxAuth  context.Context
	ctxToken context.Context

	// Parameters
	apiKeys *md.APIKeys
	host    host.HostNode
	keyPair *md.KeyPair
	options *md.ConnectionRequest_TextileOptions

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
		keyPair: sc.user.KeyPair(),
		options: sc.textileOpts,
		apiKeys: sc.apiKeys,
		host:    sc.host,
		active:  false,
	}
	sc.Textile = textile

	// Check Textile Enabled
	if textile.options.GetEnabled() {
		log.Println("Found Textile Enabled")

		// Initialize
		var err error
		creds := credentials.NewTLS(&tls.Config{})
		auth := common.Credentials{}

		// Dial GRPC
		opts := []grpc.DialOption{grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(auth)}
		textile.client, err = client.NewClient(util.TEXTILE_API_URL, opts...)
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
		textile.active = true
		log.Println("Activated Textile")

		// Initialize Threads
		if textile.options.GetThreads() {
			log.Println("Found Threads Enabled")
			textile.InitThreads()
		} else {
			log.Println("Found Threads DISABLED")
		}

		// Initialize Mailbox
		if textile.options.GetMailbox() {
			log.Println("Found Mailbox Enabled")
			textile.InitMail(sc.user.GetDevice())
		} else {
			log.Println("Found Mailbox DISABLED")
		}
	}
	return nil
}

// @ Returns Instance Host
func (tn *TextileService) PubKey() thread.PubKey {
	return tn.identity.GetPublic()
}

// @ Initializes Threads
func (tn *TextileService) InitThreads() *md.SonrError {
	// Generate a new thread ID
	threadID := thread.NewIDV1(thread.Raw, 32)

	// Create your new thread
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
	log.Println("> Success!")
	log.Println(fmt.Sprintf("ID: %s \n Maddr: %s \n Key: %s \n Name: %s \n", threadID.String(), info.Addrs, info.Key.String(), info.Name))
	return nil
}

// @ Initializes Mailbox
func (tn *TextileService) InitMail(d *md.Device) *md.SonrError {
	// Setup the mail lib
	mail := local.NewMail(cmd.NewClients(util.TEXTILE_API_URL, true, util.TEXTILE_MINER_IDX), local.DefaultConfConfig())
	tn.mail = mail

	// Create a new mailbox with config
	mailbox, err := mail.NewMailbox(context.Background(), local.Config{
		Path:      d.WorkingSupportDirectory(),
		Identity:  tn.identity,
		APIKey:    tn.apiKeys.GetTextileKey(),
		APISecret: tn.apiKeys.GetTextileSecret(),
	})

	// Check Error
	if err != nil {
		return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
	}
	tn.mailbox = mailbox
	return nil
}

// @ Method Reads Inbox and Returns List of Mail Entries
func (tn *TextileService) ReadMail() ([]*md.MailEntry, *md.SonrError) {
	// Check Mail Enabled
	if tn.active && tn.options.GetMailbox() {
		// List the recipient's inbox
		inbox, err := tn.mailbox.ListInboxMessages(context.Background())

		if err != nil {
			return nil, md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
		}

		// Initialize Entry List
		entries := make([]*md.MailEntry, len(inbox))

		// Iterate over Entries
		for i, v := range inbox {
			// Open decrypts the message body
			body, err := v.Open(context.Background(), tn.identity)
			if err != nil {
				return nil, md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
			}

			// Unmarshal Body to entry
			entry := &md.MailEntry{}
			err = proto.Unmarshal(body, entry)
			if err != nil {
				return nil, md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
			}
			entries[i] = entry
		}
		return entries, nil
	}
	return nil, nil
}

// @ Method Sends Mail Entry to Peer
func (tn *TextileService) SendMail(e *md.MailEntry) *md.SonrError {
	// Check Mail Enabled
	if tn.active && tn.options.GetMailbox() {
		// Send Message to Mailbox
		_, err := tn.mailbox.SendMessage(context.Background(), e.ToPubKey(), e.Buffer())

		// Check Error
		if err != nil {
			return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
		}
	}
	return nil
}
