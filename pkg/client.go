package main

import (
	"context"
	"log"
	"os"

	sonr "github.com/sonr-io/core/bind"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	ctx context.Context
	sonr.Callback
	node *sonr.Node
}

const weybridgeOLC = "87c4xfjv+"

// ^ Create New Client Node ^ //
func NewClient(ctx context.Context) *Client {
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
		Username: "",
		Device: &md.Device{
			Platform: "Mac",
			Model:    "MBP",
			Name:     name,
		},
		Directory: &md.Directories{
			Documents: docDir,
			Temporary: "local/temp",
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
	return c
}

// ^ Method To Share File ^ //
func (c *Client) ShareFile() {

}

// ^ Method To Share Text ^ //
func (c *Client) ShareText() {

}
