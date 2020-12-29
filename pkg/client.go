package main

import (
	"context"
	"log"

	sonr "github.com/sonr-io/core/bind"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	ctx context.Context
	sonr.Callback
	node *sonr.Node
}

// ^ Create New Client Node ^ //
func NewClient(ctx context.Context) *Client {
	// Get Request Message
	request := md.ConnectionRequest{
		Olc: "",
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

// @ Inherited Method: Handle Refresh ^ //
func (c *Client) OnRefreshed(data []byte) {
	m := &md.Lobby{}
	err := proto.Unmarshal(data, m)
	if err != nil {
		log.Panicln("Error Unmarshalling Request")
	}
}

// @ Inherited Method: Handle Invite ^ //
func (c *Client) OnInvited(data []byte) {

}

// @ Inherited Method: Handle Response ^ //
func (c *Client) OnResponded(data []byte) {

}

// @ Inherited Method: Handle Queue ^ //
func (c *Client) OnQueued(data []byte) {

}

// @ Inherited Method: Handle Progress ^ //
func (c *Client) OnProgress(data float32) {

}

// @ Inherited Method: Handle Received ^ //
func (c *Client) OnReceived(data []byte) {

}

// @ Inherited Method: Handle Sent ^ //
func (c *Client) OnTransmitted(data []byte) {

}

// @ Inherited Method: Handle Error ^ //
func (c *Client) OnError(data []byte) {

}

// ^ Method To Share File ^ //
func (c *Client) ShareFile() {

}

// ^ Method To Share Text ^ //
func (c *Client) ShareText() {

}
