package util

import (
	"fmt"
	"time"

	dscl "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
)

// ** ─── Variables ────────────────────────────────────────────────────────
// Bootstrap MDNS Peer Discovery Interval
const REFRESH_INTERVAL = time.Second * 4

// RPC Binary AutoUpdate Interval
const AUTOUPDATE_INTERVAL = time.Second * 3

// Method Returns MDNS Discovery Options
func MDNS_OPTS() (time.Duration, string) {
	return REFRESH_INTERVAL, HOST_RENDEVOUZ_POINT
}

// Method Returns DHT Discovery Options
func DHT_OPTS() (string, dscl.Option) {
	return HOST_RENDEVOUZ_POINT, dscl.TTL(REFRESH_INTERVAL)
}

// Pubsub Room Max Messages
const MAX_EMITTER_CAP = 2048

// Pubsub Room Max Messages
const MAX_CHAN_DATA = 128

// Maximum Chunk Size During Transfer
const CHUNK_SIZE = 4 * 1024

// Directory Name of Private Key Folder
const PRIVATE_KEY_DIR = ".sonr_private"

// Filename for Account
const ACCOUNT_FILE = "account.snr"

// ** ─── Protocols ────────────────────────────────────────────────────────
// Auth Service Protocol
const AUTH_PROTOCOL = protocol.ID("/sonr/auth-service/0.2")

// Exchange Service Protocol
const EXCHANGE_PROTOCOL = protocol.ID("/sonr/exchange-service/0.2")

// Remote Service Protocol
const REMOTE_PROTOCOL = protocol.ID("/sonr/remote-service/0.2")

// Exchange Service Protocol
const ACCOUNT_PROTOCOL = protocol.ID("/sonr/account-protocol/0.2")

// ** ─── API ────────────────────────────────────────────────────────
// Textile Client API URL
const TEXTILE_API_URL = "api.hub.textile.io:443"

// Textile Miner Index Target
const TEXTILE_MINER_IDX = "api.minerindex.hub.textile.io:443"

// Textile Mailbox Directory
const TEXTILE_MAILBOX_DIR = ".textile"

// ** ─── Services ────────────────────────────────────────────────────────
// Local RPC Service Name
const AUTH_RPC_SERVICE = "AuthService"

// Local RPC Service Method for Invite
const AUTH_METHOD_INVITE = "InviteWith"

// Local RPC Service Method for Invite
const AUTH_METHOD_LINK = "LinkWith"

// Local RPC Service Name
const DEVICE_RPC_SERVICE = "DeviceService"

// Local RPC Service Method for Push
const DEVICE_METHOD_DIRECT = "DirectWith"

// Local RPC Service Method for Push
const DEVICE_METHOD_MODIFY = "ModifyWith"

// Local RPC Service Method for Exchange
const DEVICE_METHOD_VERIFY = "VerifyWith"

// Local RPC Service Name
const EXCHANGE_RPC_SERVICE = "ExchangeService"

// Local RPC Service Method for Exchange
const EXCHANGE_METHOD_EXCHANGE = "ExchangeWith"

// ** ─── Host ────────────────────────────────────────────────────────
// Libp2p Host Rendevouz Point
const HOST_RENDEVOUZ_POINT = "/sonr/rendevouz/0.9.2"

// RPC Server Localhost Port
const RPC_SERVER_PORT = 60214

// ─── Methods ────────────────────────────────────────────────────────
// Construct New Protocol ID given Method Name String and Value String
func NewValueProtocol(method string, value string) protocol.ID {
	return protocol.ID(fmt.Sprintf("/sonr/%s/%s", method, value))
}

// Construct New Protocol ID given Method Name String and id Peer.ID
func NewIDProtocol(method string, id peer.ID) protocol.ID {
	return protocol.ID(fmt.Sprintf("/sonr/%s/%s", method, id.Pretty()))
}
