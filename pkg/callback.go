package main

import (
	"log"

	op "github.com/skratchdot/open-golang/open"
	md "github.com/sonr-io/core/internal/models"
	"github.com/sonr-io/core/pkg/ui"
	"google.golang.org/protobuf/proto"
)

// @ Inherited Method: On Connected ^ //
func (c *Client) OnConnected() {
	log.Println("Connected")
}

// @ Inherited Method: Handle Event ^ //
func (c *Client) OnEvent(data []byte) {
	m := &md.LobbyEvent{}
	err := proto.Unmarshal(data, m)
	if err != nil {
		log.Panicln("Error Unmarshalling Request")
	}
	log.Println(m.String())
}

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
	// Unmarshal Invite
	m := &md.AuthInvite{}
	err := proto.Unmarshal(data, m)
	if err != nil {
		log.Panicln("Error Unmarshalling Request")
	}
	ui.PushInvited(m)

	// Check Invite
	if m.Payload == md.Payload_FILE {
		c.Node.Respond(true)
	} else if m.Payload == md.Payload_URL {
		err := op.Start(m.Url)
		if err != nil {
			log.Println(err)
		}
	}

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
	// Unmarshal Data
	m := &md.Metadata{}
	err := proto.Unmarshal(data, m)
	if err != nil {
		log.Panicln("Error Unmarshalling Request")
	}
	print(m.String())

	// Move File
	err = op.Start(m.Path)
	if err != nil {
		return
	}

	// Notify
	ui.BeepCompleted()
}

// @ Inherited Method: Handle Sent ^ //
func (c *Client) OnTransmitted(data []byte) {
	ui.BeepCompleted()
}

// @ Inherited Method: Handle Error ^ //
func (c *Client) OnError(data []byte) {

}
