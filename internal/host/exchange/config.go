package exchange

import (
	"errors"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/protocol"
	device "github.com/sonr-io/sonr/pkg/fs"
)

// Textile API definitions
const (
	// Textile Client API URL
	TextileClientURL = "https://api.textile.io"

	// Textile Miner Index Target
	TextileMinerIdx = "api.minerindex.hub.textile.io:443"

	// Textile Mailbox Directory
	TextileMailboxDirName = "mailbox"

	RequestPID protocol.ID = "/exchange/request/0.0.1"

	ResponsePID protocol.ID = "/exchange/response/0.0.1"
)

var (
	logger              = golog.Default.Child("protocols/exchange")
	ErrMailboxDisabled  = errors.New("Mailbox not enabled, cannot perform request.")
	ErrMissingAPIKey    = errors.New("Missing Textile API Key in env")
	ErrMissingAPISecret = errors.New("Missing Textile API Secret in env")
	ErrRequestNotFound  = errors.New("Request not found in protocol cache")
	ErrNotSupported     = errors.New("Action not supported for StubMode")
)

// Option is a function that can be applied to ExchangeProtocol config
type Option func(*options)

// options for ExchangeProtocol config
type options struct {
	enableMailbox bool
}

// defaultOptions for ExchangeProtocol config
func defaultOptions() *options {
	return &options{
		enableMailbox: false,
	}
}

// EnableMailbox enables the mailbox
func EnableMailbox() Option {
	return func(o *options) {
		o.enableMailbox = true
	}
}

// Apply applies the options to the ExchangeProtocol
func (o *options) Apply(p *ExchangeProtocol) error {
	// Apply options
	p.mode = p.node.Role()

	// Set enableMailbox
	if o.enableMailbox {
		//mail := local.NewMail(cmd.NewClients(TextileClientURL, true, TextileMinerIdx), local.DefaultConfConfig())

		// Create new mailbox
		if device.ThirdParty.Exists(TextileMailboxDirName) {
			// Return Existing Mailbox
			if err := p.loadMailbox(); err != nil {
				return err
			}
		} else {
			// Create New Mailbox
			if err := p.newMailbox(); err != nil {
				return err
			}
		}
		// go p.handleMailboxEvents()
	}
	return nil
}
