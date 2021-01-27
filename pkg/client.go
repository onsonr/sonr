package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/skip2/go-qrcode"
	sonr "github.com/sonr-io/core/bind"
	md "github.com/sonr-io/core/internal/models"
	"github.com/sonr-io/core/pkg/ui"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

const weybridgeOLC = "87C4XFJV+"
const interval = 500 * time.Millisecond

type Client struct {
	ctx  context.Context
	menu ui.SystemMenu
	sonr.Callback
	node       *sonr.Node
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
			Temporary: filepath.Join(docDir, "Downloads"),
		},
		Contact: &md.Contact{
			FirstName: "MacTest",
			LastName:  "MacTest",
		},
	}

	// Convert to Bytes
	bytes, err := proto.Marshal(&request)
	if err != nil {
		log.Panicln("Error Marshalling Request")
	}

	// Create New Client
	var c = new(Client)
	c.ctx = ctx
	c.node = sonr.NewNode(bytes, c)
	// go c.HandleLinkInput()
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

// ^ Display QR Code of Peer Info ^ //
func (c *Client) DisplayCode() []byte {
	// Get Node JSON
	jsonBytes, err := protojson.Marshal(c.node.Peer())
	if err != nil {
		log.Println(err)
		return nil
	}
	json := string(jsonBytes)
	print(json)

	// Encode to QR
	png, err := qrcode.Encode(json, qrcode.Medium, 256)
	if err != nil {
		log.Println(err)
		return nil
	}
	return png
}

// // ^ Routine Handles Menu Input ^ //
// func (c *Client) HandleLinkInput() {
// 	<-c.menu.MLink.ClickedCh
// 	c.node.DisplayCode()
// }

// ^ Method To Queue File ^ //
func (c *Client) QueueFile(path string) {
	c.node.Process(path)
}

// ^ Method To Share File ^ //
func (c *Client) ShareFile(id string, path string) {
	c.node.InviteWithFile(id)
}

// ^ Method To Share Text ^ //
func (c *Client) ShareText() {

}
