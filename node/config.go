package node

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/sonr-io/core/common"
	"github.com/sonr-io/core/device"
	"github.com/sonr-io/core/types/go/node/motor/v1"
)

// Error Definitions
var (
	ErrEmptyQueue       = errors.New("No items in Transfer Queue.")
	ErrInvalidQuery     = errors.New("No SName or PeerID provided.")
	ErrMissingParam     = errors.New("Paramater is missing.")
	ErrProtocolsNotSet  = errors.New("Node Protocol has not been initialized.")
	ErrRoutingNotSet    = errors.New("DHT and Host have not been set by Routing Function")
	ErrListenerRequired = errors.New("Listener was not Provided")
	ErrMDNSInvalidConn  = errors.New("Invalid Connection, cannot begin MDNS Service")
)

var (
	bootstrapAddrStrs = []string{
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",
		"/ip4/104.131.131.82/tcp/4001/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
		"/ip4/104.131.131.82/udp/4001/quic/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
	}
	addrStoreTTL = time.Minute * 5
)

// CallbackImpl is the implementation of Callback interface
type CallbackImpl interface {
	// OnRefresh is called when the LobbyProtocol is refreshed and pushes a RefreshEvent
	OnRefresh(event *motor.OnLobbyRefreshResponse)

	// OnMailbox is called when the MailboxProtocol receives a MailboxEvent
	OnMailbox(event *motor.OnMailboxMessageResponse)

	// OnInvite is called when the TransferProtocol receives InviteEvent
	OnInvite(event *motor.OnTransmitInviteResponse)

	// OnDecision is called when the TransferProtocol receives a DecisionEvent
	OnDecision(event *motor.OnTransmitDecisionResponse, invite *motor.OnTransmitInviteResponse)

	// OnProgress is called when the TransferProtocol sends or receives a ProgressEvent
	OnProgress(event *motor.OnTransmitProgressResponse)

	// OnTransfer is called when the TransferProtocol completes a transfer and pushes a CompleteEvent
	OnComplete(event *motor.OnTransmitCompleteResponse)
}

// Role is the type of the node (Client, Highway)
type Role int

const (
	// StubMode_LIB is the Node utilized by Mobile and Web Clients
	Role_UNSPECIFIED Role = iota

	// StubMode_CLI is the Node utilized by CLI Clients
	Role_TEST

	// Role_MOTOR is for a Motor Node
	Role_MOTOR

	// Role_HIGHWAY is for a Highway Node
	Role_HIGHWAY
)

// Motor returns true if the node has a client.
func (m Role) IsMotor() bool {
	return m == Role_MOTOR
}

// Highway returns true if the node has a highway stub.
func (m Role) IsHighway() bool {
	return m == Role_HIGHWAY
}

// Prefix returns golog prefix for the node.
func (m Role) Prefix() string {
	var name string
	switch m {
	case Role_HIGHWAY:
		name = "highway"
	case Role_MOTOR:
		name = "motor"
	case Role_TEST:
		name = "test"
	default:
		name = "unknown"
	}
	return fmt.Sprintf("[SONR.%s] ", name)
}

type Configuration struct {
	connection       common.Connection
	deviceId         string
	location         *common.Location
	profile          *common.Profile
	homeDirectory    string
	supportDirectory string
	tempDirectory    string
}

func defaultConfiguration() *Configuration {
	// Default configuration
	c := &Configuration{
		connection: common.Connection_WIFI,
		location:   common.NewDefaultLocation(),
		profile:    common.NewDefaultProfile(),
	}

	// Check for non-mobile device
	if !device.IsMobile() {

		// Set Device ID
		if fid, err := device.ID(); err == nil {
			c.deviceId = fid
		}

		// Set Home Directory
		if hdir, err := os.UserHomeDir(); err == nil {
			c.homeDirectory = hdir
		}

		// Set Support Directory
		if sdir, err := os.UserConfigDir(); err == nil {
			c.supportDirectory = sdir
		}

		// Set Temp Directory
		if tdir, err := os.UserCacheDir(); err == nil {
			c.tempDirectory = tdir
		}
	}
	return c
}

func (c *Configuration) Apply(n *node) {
	// Initialize device
	device.Init(
		device.WithHomePath(c.homeDirectory),
		device.WithSupportPath(c.supportDirectory),
		device.SetDeviceID(c.deviceId),
	)
}
