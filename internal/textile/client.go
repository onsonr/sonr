package textile

import (
	"context"
	"crypto/tls"
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
)

type TextileNode interface {
	InitThreads() *md.SonrError
	PubKey() thread.PubKey
	ReadMail() ([]*md.MailEntry, *md.SonrError)
	SendMail(*md.MailEntry) *md.SonrError
}

type textileNode struct {
	TextileNode
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
func NewTextile(hn host.HostNode, req *md.ConnectionRequest, keyPair *md.KeyPair) (TextileNode, *md.SonrError) {
	// Initialize
	node := &textileNode{
		keyPair: keyPair,
		options: req.GetTextileOptions(),
		apiKeys: req.GetApiKeys(),
		host:    hn,
		active:  false,
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
		node.active = true
		log.Println("Activated Textile")
	}
	return node, nil
}

// @ Returns Instance Host
func (tn *textileNode) PubKey() thread.PubKey {
	return tn.identity.GetPublic()
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
func (tn *textileNode) newTokenCtx() (context.Context, error) {
	// Generate a new token for the user
	token, err := tn.client.GetToken(tn.ctxAuth, tn.identity)
	if err != nil {
		return nil, err
	}
	return thread.NewTokenContext(tn.ctxAuth, token), nil
}
