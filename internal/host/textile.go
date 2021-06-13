package host

import (
	"context"
	"crypto/tls"
	"time"

	crypto "github.com/libp2p/go-libp2p-core/crypto"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/textileio/go-threads/api/client"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/textile/v2/api/common"
	"github.com/textileio/textile/v2/cmd"
	"github.com/textileio/textile/v2/mail/local"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/proto"
)

// @ Initializes New Textile Instance
func (hn *hostNode) StartTextile(d *md.Device) *md.SonrError {
	// Initialize
	var err error
	creds := credentials.NewTLS(&tls.Config{})
	auth := common.Credentials{}

	// Dial GRPC
	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(auth)}
	hn.tileClient, err = client.NewClient(textileApiUrl, opts...)
	if err != nil {
		return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
	}

	// Get Identity
	hn.tileIdentity = getIdentity(hn.keyPair.PrivKey())

	// Create Auth Context
	hn.ctxTileAuth, err = newUserAuthCtx(context.Background(), hn.apiKeys)
	if err != nil {
		return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
	}

	// Create Token Context
	hn.ctxTileToken, err = hn.newTokenCtx()
	if err != nil {
		return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
	}
	// Setup the mail lib
	hn.tileMail = local.NewMail(cmd.NewClients(textileApiUrl, true, textileMinerIdx), local.DefaultConfConfig())

	// Create a new mailbox with config
	hn.tileMailbox, err = hn.tileMail.NewMailbox(context.Background(), local.Config{
		Path:      d.WorkingSupportPath(".mailbox"),
		Identity:  hn.tileIdentity,
		APIKey:    hn.apiKeys.TextileKey,
		APISecret: hn.apiKeys.TextileSecret,
	})
	if err != nil {
		return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
	}
	return nil
}

// @ Method Reads Inbox and Returns List of Mail Entries
func (hn *hostNode) ReadMail() ([]*md.MailEntry, *md.SonrError) {
	// List the recipient's inbox
	inbox, err := hn.tileMailbox.ListInboxMessages(context.Background())

	if err != nil {
		return nil, md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
	}

	// Initialize Entry List
	entries := make([]*md.MailEntry, len(inbox))

	// Iterate over Entries
	for i, v := range inbox {
		// Open decrypts the message body
		body, err := v.Open(context.Background(), hn.tileIdentity)
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

// @ Method Sends Mail Entry to Peer
func (hn *hostNode) SendMail(e *md.MailEntry) *md.SonrError {
	// Send Message to Mailbox
	_, err := hn.tileMailbox.SendMessage(context.Background(), e.ToPubKey(), e.Buffer())

	// Check Error
	if err != nil {
		return md.NewError(err, md.ErrorMessage_HOST_TEXTILE)
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
	return common.CreateAPISigContext(ctx, time.Now().Add(time.Minute), keys.TextileSecret)
}

// # Helper: Creates Auth Token Context from AuthContext, Client, Identity
func (hn *hostNode) newTokenCtx() (context.Context, error) {
	// Generate a new token for the user
	token, err := hn.tileClient.GetToken(hn.ctxTileAuth, hn.tileIdentity)
	if err != nil {
		return nil, err
	}
	return thread.NewTokenContext(hn.ctxTileAuth, token), nil
}
