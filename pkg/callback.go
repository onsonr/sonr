package main

import (
	"log"

	md "github.com/sonr-io/core/internal/models"
	"github.com/sonr-io/core/pkg/ui"
	"google.golang.org/protobuf/proto"
)

// @ Inherited Method: Handle Refresh ^ //
func (c *Client) OnRefreshed(data []byte) {
	m := &md.Lobby{}
	err := proto.Unmarshal(data, m)
	if err != nil {
		log.Panicln("Error Unmarshalling Request")
	}
	c.menu.UpdatePeers(m)
}

// @ Inherited Method: Handle Invite ^ //
func (c *Client) OnInvited(data []byte) {
	m := &md.AuthInvite{}
	err := proto.Unmarshal(data, m)
	if err != nil {
		log.Panicln("Error Unmarshalling Request")
	}
	ui.PushInvited(m)
	c.node.Respond(true)
}

// @ Inherited Method: Handle Response ^ //
func (c *Client) OnResponded(data []byte) {
	log.Println(data)
}

// @ Inherited Method: Handle Queue ^ //
func (c *Client) OnQueued(data []byte) {
	log.Println(data)
}

// @ Inherited Method: Handle Progress ^ //
func (c *Client) OnProgress(data float32) {
	log.Println(data)
}

// @ Inherited Method: Handle Received ^ //
func (c *Client) OnReceived(data []byte) {
	ui.BeepCompleted()
}

// @ Inherited Method: Handle Sent ^ //
func (c *Client) OnTransmitted(data []byte) {
	ui.BeepCompleted()
}

// @ Inherited Method: Handle Error ^ //
func (c *Client) OnError(data []byte) {

}
