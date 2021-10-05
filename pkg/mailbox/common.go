package mailbox

import (
	"errors"
	"os"

	"github.com/sonr-io/core/internal/device"
	"github.com/sonr-io/core/tools/logger"
	"github.com/textileio/textile/v2/mail/local"
	"go.uber.org/zap"
)

// Textile API definitions
const (
	// Textile Client API URL
	TextileClientURL = "https://api.textile.io"

	// Textile Miner Index Target
	TextileMinerIdx = "api.minerindex.hub.textile.io:443"

	// Textile Mailbox Directory
	TextileMailboxDir = ".mailbox"
)

var (
	ErrMissingAPIKey    = errors.New("Missing Textile API Key in env")
	ErrMissingAPISecret = errors.New("Missing Textile API Secret in env")
)

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

// loadMailbox loads an existing mailbox instance
func (mb *MailboxProtocol) loadMailbox() error {
	// Get Mailbox Path
	path, err := device.Mailbox.Path()
	if err != nil {
		return logger.Error("Failed to Find Existing Mailbox at Path", err)
	}

	// Return Existing Mailbox
	mailbox, err := mb.mail.GetLocalMailbox(mb.ctx, path)
	if err != nil {
		return logger.Error("Failed to Load Existing Mailbox", err)
	}

	// Set mailbox
	mb.mailbox = mailbox
	logger.Info("Existing Mailbox has been loaded.", zap.String("path", path))
	return nil
}

// newMailbox creates a new mailbox instance
func (mb *MailboxProtocol) newMailbox() error {
	// Get Mailbox Path
	path, err := device.Mailbox.Path()
	if err != nil {
		return logger.Error("Failed to Find Existing Mailbox at Path", err)
	}

	// Get Device ThreadIdentity
	id, err := device.KeyChain.ThreadIdentity()
	if err != nil {
		return logger.Error("Failed to get thread Identity", err)
	}

	// Fetch API Keys
	key, secret, err := fetchApiKeys()
	if err != nil {
		return logger.Error("Failed to create new mailbox", err)
	}

	// Create a new mailbox with config
	mailbox, err := mb.mail.NewMailbox(mb.ctx, local.Config{
		Path:      path,
		Identity:  id,
		APIKey:    key,
		APISecret: secret,
	})

	// Check for errors
	if err != nil {
		return logger.Error("Failed to create mailbox", err)
	}

	// Set mailbox
	mb.mailbox = mailbox
	logger.Info("New Mailbox has been created.", zap.String("path", device.MailboxPath))
	return nil
}
