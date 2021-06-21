package textile

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	crypto "github.com/libp2p/go-libp2p-crypto"
	"github.com/sonr-io/core/internal/host"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/util"
	"github.com/textileio/go-threads/api/client"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/textile/v2/api/common"
	"github.com/textileio/textile/v2/mail/local"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/proto"
)

type TextileNode interface {
	PubKey() thread.PubKey
	SendMail(*md.MailEntry) *md.SonrError
	ReadMail() ([]*md.MailEntry, *md.SonrError)
}

type textileNode struct {
	TextileNode
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
func NewTextile(hn host.HostNode, req *md.ConnectionRequest, keyPair *md.KeyPair) (TextileNode, *md.SonrError) {
	// Initialize
	node := &textileNode{
		keyPair: keyPair,
		options: req.GetTextileOptions(),
		apiKeys: req.GetApiKeys(),
		host:    hn,
	}

	// Check Textile Enabled
	if node.options.GetEnabled() {
		// Initialize
		var err error
		creds := credentials.NewTLS(&tls.Config{})
		auth := common.Credentials{}

		// Dial GRPC
		opts := []grpc.DialOption{grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(auth)}
		node.client, err = client.NewClient(util.TEXTILE_API_URL, opts...)
		if err != nil {
			return nil, md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
		}

		// Get Identity
		node.identity = getIdentity(node.keyPair.PrivKey())

		// Create Auth Context
		node.ctxAuth, err = newUserAuthCtx(context.Background(), node.apiKeys)
		if err != nil {
			return nil, md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
		}

		// Create Token Context
		node.ctxToken, err = node.newTokenCtx()
		if err != nil {
			return nil, md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
		}

		// Check Thread Enabled
		if node.options.GetThreads() {
			// Generate a new thread ID
			threadID := thread.NewIDV1(thread.Raw, 32)

			// Create your new thread
			err = node.client.NewDB(node.ctxToken, threadID)
			if err != nil {
				return nil, md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
			}

			// Get DB Info
			info, err := node.client.GetDBInfo(node.ctxToken, threadID)
			if err != nil {
				return nil, md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
			}

			// Log DB Info
			log.Println("> Success!")
			log.Println(fmt.Sprintf("ID: %s \n Maddr: %s \n Key: %s \n Name: %s \n", threadID.String(), info.Addrs, info.Key.String(), info.Name))
		}
	}
	return node, nil
}

// @ Returns Instance Host
func (hn *textileNode) PubKey() thread.PubKey {
	return hn.identity.GetPublic()
}

// @ Method Reads Inbox and Returns List of Mail Entries
func (hn *textileNode) ReadMail() ([]*md.MailEntry, *md.SonrError) {
	// Check Mail Enabled
	if hn.options.GetMailbox() {
		// List the recipient's inbox
		inbox, err := hn.mailbox.ListInboxMessages(context.Background())

		if err != nil {
			return nil, md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
		}

		// Initialize Entry List
		entries := make([]*md.MailEntry, len(inbox))

		// Iterate over Entries
		for i, v := range inbox {
			// Open decrypts the message body
			body, err := v.Open(context.Background(), hn.identity)
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
func (hn *textileNode) SendMail(e *md.MailEntry) *md.SonrError {
	// Check Mail Enabled
	if hn.options.GetMailbox() {
		// Send Message to Mailbox
		_, err := hn.mailbox.SendMessage(context.Background(), e.ToPubKey(), e.Buffer())

		// Check Error
		if err != nil {
			return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
		}
		return nil
	}
	return nil
}

// # Helper: Gets Thread Identity from Private Key
func getIdentity(privateKey crypto.PrivKey) thread.Identity {
	myIdentity := thread.NewLibp2pIdentity(privateKey)
	return myIdentity
}

// # Helper: Creates User Auth Context from API Keys
func newUserAuthCtx(ctx context.Context, keys *md.APIKeys) (context.Context, error) {
	// Add our user group key to the context
	ctx = common.NewAPIKeyContext(ctx, keys.TextileKey)

	// Add a signature using our user group secret
	return common.CreateAPISigContext(ctx, time.Now().Add(time.Hour), keys.TextileSecret)
}

// # Helper: Creates Auth Token Context from AuthContext, Client, Identity
func (hn *textileNode) newTokenCtx() (context.Context, error) {
	// Generate a new token for the user
	token, err := hn.client.GetToken(hn.ctxAuth, hn.identity)
	if err != nil {
		return nil, err
	}
	return thread.NewTokenContext(hn.ctxAuth, token), nil
}
