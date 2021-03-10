package service

import (
	"context"
	"log"
	"time"

	sonr "github.com/sonr-io/core/bind"
	md "github.com/sonr-io/core/internal/models"
	ui "github.com/sonr-io/core/pkg/interface"
	"google.golang.org/protobuf/proto"
)

const interval = 500 * time.Millisecond

type Client struct {
	sonr.Callback
	ctx       context.Context
	face      ui.Interface
	node      *sonr.Node
	info      SysInfo
	peerCount int32
	lobbySize int32
}

// ^ Create New DeskClient Node ^ //
func NewClient(ctx context.Context, m ui.Interface) *Client {
	// Set Default Info
	var c = new(Client)
	c.info = SystemInfo()
	c.ctx = ctx
	c.face = m
	c.peerCount = 0
	c.lobbySize = 1

	// Create Request Message
	request := md.ConnectionRequest{
		Latitude:    38.980620,
		Longitude:   -77.505890,
		Device:      &c.info.Device,
		Directories: &c.info.Directory,
		Contact: &md.Contact{
			FirstName: c.info.TempFirstName,
			LastName:  c.info.TempLastName,
		},
		Username: "@Prad",
	}

	// Convert to Bytes
	bytes, err := proto.Marshal(&request)
	if err != nil {
		log.Panicln("Error Marshalling Request")
	}

	// Create New Client
	c.node = sonr.NewNode(bytes, c)
	c.face.Initialize(c.node)
	m.Initialize(c.node)

	// Start Routine
	go c.UpdateAuto(time.NewTicker(interval))
	return c
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
