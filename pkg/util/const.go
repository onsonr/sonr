package util

import (
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
)

// ** ─── Variables ────────────────────────────────────────────────────────
// Bootstrap Peer Discovery Interval
const REFRESH_INTERVAL = time.Second * 5

// Pubsub Topic Max Messages
const TOPIC_MAX_MESSAGES = 128

// Maximum Chunk Size During Transfer
const TRANSFER_CHUNK_SIZE = 4 * 1024

// Private Key File Name
const KEY_FILE_NAME = ".sonr_private_key"

// ** ─── Protocols ────────────────────────────────────────────────────────
// Global Service Protocol ID
const GLOBAL_PROTOCOL = protocol.ID("/sonr/global-service/0.2")

// Local Service Protocol
const LOCAL_PROTOCOL = protocol.ID("/sonr/local-service/0.2")

// Remote Service Protocol
const REMOTE_PROTOCOL = protocol.ID("/sonr/remote-service/0.2")

// ** ─── API ────────────────────────────────────────────────────────
// Textile Client API URL
const TEXTILE_API_URL = "api.hub.textile.io:443"

// Textile Miner Index Target
const TEXTILE_MINER_IDX = "f022352"

// ** ─── Services ────────────────────────────────────────────────────────
// Local RPC Service Name
const LOCAL_RPC_SERVICE = "LocalService"

// Local RPC Service Method for Exchange
const LOCAL_METHOD_EXCHANGE = "ExchangeWith"

// Local RPC Service Method for Invite
const LOCAL_METHOD_INVITE = "InviteWith"

// ^ ─── Methods ────────────────────────────────────────────────────────
// Construct New Protocol ID given Method Name String and Value String
func NewValueProtocol(method string, value string) protocol.ID {
	return protocol.ID(fmt.Sprintf("/sonr/%s/%s", method, value))
}

// Construct New Protocol ID given Method Name String and id Peer.ID
func NewIDProtocol(method string, id peer.ID) protocol.ID {
	return protocol.ID(fmt.Sprintf("/sonr/%s/%s", method, id.Pretty()))
}
