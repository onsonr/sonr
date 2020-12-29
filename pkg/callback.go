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
	ui.RefreshLobby(m)
}

// @ Inherited Method: Handle Invite ^ //
func (c *Client) OnInvited(data []byte) {
	m := &md.AuthInvite{}
	err := proto.Unmarshal(data, m)
	if err != nil {
		log.Panicln("Error Unmarshalling Request")
	}
	ui.PushInvited(m)
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
