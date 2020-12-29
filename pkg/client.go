package main

import (
	"context"
	"log"
	"os"
	"time"

	sonr "github.com/sonr-io/core/bind"
	md "github.com/sonr-io/core/internal/models"
	"github.com/sonr-io/core/pkg/ui"
	"google.golang.org/protobuf/proto"
)

const weybridgeOLC = "87C4XFJV+"
const interval = 500 * time.Millisecond

type Client struct {
	ctx  context.Context
	menu ui.SystemMenu
	sonr.Callback
	node *sonr.Node
}

// ^ Create New Client Node ^ //
func NewClient(ctx context.Context, m ui.SystemMenu) *Client {
	// Get Info
	name, err := os.Hostname()
	if err != nil {
		log.Println(err)
		name = "Undefined"
	}

	docDir, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
		docDir = "local/temp"
	}

	// Create Request Message
	request := md.ConnectionRequest{
		Olc:      weybridgeOLC,
		Username: "@TestUser",
		Device: &md.Device{
			Platform: "Mac",
			Model:    "MBP",
			Name:     name,
		},
		Directory: &md.Directories{
			Documents: docDir,
			Temporary: "local/temp",
		},
		Contact: &md.Contact{
			FirstName: "MacTest",
			LastName:  "MacTest",
		},
	}

	bytes, err := proto.Marshal(&request)
	if err != nil {
		log.Panicln("Error Marshalling Request")
	}

	// Create New Client
	var c = new(Client)
	c.ctx = ctx
	c.node = sonr.NewNode(bytes, c)
	go c.UpdatePeriodically(time.NewTicker(interval))
	return c
}

// ^ Method to Periodically Update Presence ^ //
func (c *Client) UpdatePeriodically(ticker *time.Ticker) {
	for {
		select {
		case <-c.ctx.Done():
			return
		case <-ticker.C:
			c.node.Update(0)
		}
	}
}

// ^ Method To Share File ^ //
func (c *Client) ShareFile() {

}

// ^ Method To Share Text ^ //
func (c *Client) ShareText() {

}
