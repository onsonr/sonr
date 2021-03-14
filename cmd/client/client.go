package client

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	md "github.com/sonr-io/core/pkg/models"

	sn "github.com/sonr-io/core/pkg/node"
)

const interval = 500 * time.Millisecond

// @ Interface: Callback is implemented from Plugin to receive updates
type Callback interface {
	OnRefreshed(data []byte)   // Lobby Updates
	OnEvent(data []byte)       // Lobby Event
	OnInvited(data []byte)     // User Invited
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
	ctx  context.Context
	Info SysInfo
	Node *sn.Node
}

// ^ Create New DeskClient Node ^ //
func NewClient(ctx context.Context, call Callback) *Client {
	// Set Default Info
	var c = new(Client)
	c.Info = SystemInfo()
	c.ctx = ctx

	// Create Request Message
	request := &md.ConnectionRequest{
		Latitude:    38.980620,
		Longitude:   -77.505890,
		Device:      &c.Info.Device,
		Directories: &c.Info.Directory,
		Contact: &md.Contact{
			FirstName: c.Info.TempFirstName,
			LastName:  c.Info.TempLastName,
		},
		Username: "@Prad",
	}

	// Create New Client
	c.Node = sn.NewNode(request, call)

	// Start Routine
	go c.UpdateAuto(time.NewTicker(interval))
	return c
}

// ^ Method to Periodically Update Presence ^ //
func (dc *Client) UpdateAuto(ticker *time.Ticker) {
	for {
		select {
		case <-dc.ctx.Done():
			dc.Node.Stop()
			return
		case <-ticker.C:
			dc.Node.Update(0, 0)
		}
	}
}

// ^ Returns System Info ^ //
func SystemInfo() SysInfo {
	// Initialize Vars
	var platform md.Platform
	var model string
	var name string
	var homeDir string
	var libDir string
	var last string
	var err error

	// Get Operating System
	runOs := runtime.GOOS

	// Check Runtime OS
	switch runOs {
	// @ Windows
	case "windows":
		platform = md.Platform_Windows
		last = "PC"

		// @ Mac
	case "darwin":
		platform = md.Platform_MacOS
		last = "Mac"

		// @ Linux
	case "linux":
		platform = md.Platform_Linux

		// @ Unknown
	default:
		platform = md.Platform_Unknown
	}

	// Get Hostname
	if name, err = os.Hostname(); err != nil {
		log.Println(err)
		name = "Undefined"
	}

	// Get Directories
	if homeDir, err = os.UserHomeDir(); err != nil {
		log.Println(err)
		homeDir = "local/temp"
	}

	if libDir, err = os.UserConfigDir(); err != nil {
		log.Println(err)
		libDir = "local/temp"
	}

	// Return SysInfo Object
	return SysInfo{
		// Current Hard Code OLC
		OLC:           "87C4XFJV+",
		TempFirstName: "Prad's",
		TempLastName:  last,

		// Retreived Device Info
		Device: md.Device{
			Platform: platform,
			Model:    model,
			Name:     name,
			Desktop:  true,
		},

		// Current Directories
		Directory: md.Directories{
			Documents: libDir,
			Temporary: filepath.Join(homeDir, "Downloads"),
			Downloads: filepath.Join(homeDir, "Downloads"),
		},
	}
}
