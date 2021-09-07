package host

import (
	"time"

	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/textile/v2/api/common"
)

// Textile Client API URL
const TEXTILE_API_URL = "api.hub.textile.io:443"

// Textile Miner Index Target
const TEXTILE_MINER_IDX = "api.minerindex.hub.textile.io:443"

// Textile Mailbox Directory
const TEXTILE_MAILBOX_DIR = ".textile"

func (h *SHost) startTextileClient(apiKey string, apiSecret string) error {
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

func (h *SHost) threadIdentity() thread.Identity {
	return thread.NewLibp2pIdentity(h.privKey)
}
