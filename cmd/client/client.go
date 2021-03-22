package client

import (
	"context"
	"log"
	"time"

	md "github.com/sonr-io/core/pkg/models"
	sn "github.com/sonr-io/core/pkg/node"
)

const interval = 500 * time.Millisecond

// @ Interface: Callback is implemented from Plugin to receive updates
type Callback interface {
	OnConnected(data bool)     // Node Host has Bootstrapped
	OnReady(data bool)         // Node Host Connection Result
	OnRefreshed(data []byte)   // Lobby Updates
	OnEvent(data []byte)       // Lobby Event
	OnInvited(data []byte)     // User Invited
	OnRemoteStart(data []byte) // User started remote
	OnDirected(data []byte)    // User Direct-Invite from another Device
	OnResponded(data []byte)   // Peer has responded
	OnProgress(data float32)   // File Progress Updated
	OnReceived(data []byte)    // User Received File
	OnTransmitted(data []byte) // User Sent File
	OnError(data []byte)       // Internal Error
}

// @ Struct: Reference for Client Info
type SysInfo struct {
	OLC           string
	Device        md.Device
	Directory     md.Directories
	TempFirstName string
	TempLastName  string
}

// @ Struct: Reference for Exposed Sonr Client
type Client struct {
	ctx             context.Context
	Info            SysInfo
	ID              string
	DeviceID        string
	UserID          uint32
	node            *sn.Node
	hasStarted      bool
	hasBootstrapped bool
	hostOpts        *sn.HostOptions
}

// ^ Create New DeskClient Node ^ //
func NewClient(ctx context.Context, req *md.ConnectionRequest, call Callback) *Client {
	// Set Default Info
	var c = new(Client)
	c.Info = SystemInfo()
	c.ctx = ctx

	// Create New Client
	c.node = sn.NewNode(req, call)
	hostOpts, err := sn.NewHostOpts(req)
	if err != nil {
		log.Println(err)
	}

	// Set Host Opts
	c.hostOpts = hostOpts
	return c
}

// @ Start Host
func (c *Client) Connect() {
	// Start Node
	result := c.node.Start(c.hostOpts)
	if result {
		// Set Peer Info
		peer := c.node.Peer()
		c.ID = peer.Id.Peer
		c.DeviceID = peer.Id.Device
		c.UserID = peer.Id.User

		// Set Started
		c.hasStarted = true

		// Bootstrap to Peers
		strapResult := c.node.Bootstrap(c.hostOpts)
		if strapResult {
			c.hasBootstrapped = true

			// Start Routine
			go c.UpdateAuto(time.NewTicker(interval))
		} else {
			log.Println("Failed to bootstrap node")
		}
	} else {
		log.Println("Failed to start host")
	}
}

// ^ Method to Periodically Update Presence ^ //
func (dc *Client) UpdateAuto(ticker *time.Ticker) {
	for {
		select {
		case <-dc.ctx.Done():
			dc.node.Stop()
			return
		case <-ticker.C:
			dc.node.Update(0, 0)
		}
	}
}

// @ Invite Processes Data and Sends Invite to Peer
func (dc *Client) Invite(req *md.InviteRequest) {
	if dc.IsReady() {
		dc.node.Invite(req)
	}
}

// @ Checks for is Ready
func (dc *Client) IsReady() bool {
	return dc.hasBootstrapped && dc.hasStarted
}
