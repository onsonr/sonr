package service

import (
	"log"

	op "github.com/skratchdot/open-golang/open"
	md "github.com/sonr-io/core/internal/models"
	ui "github.com/sonr-io/core/pkg/menu"
	"google.golang.org/protobuf/proto"
)

// @ Inherited Method: On Connected ^ //
func (c *Client) OnConnected() {
	log.Println("Connected")
}

// @ Inherited Method: Handle Event ^ //
func (c *Client) OnEvent(data []byte) {
	lob := &md.LobbyEvent{}
	err := proto.Unmarshal(data, lob)
	if err != nil {
		log.Panicln("Error Unmarshalling Request")
	}
	log.Println(lob.String())
}

// @ Inherited Method: Handle Refresh ^ //
func (c *Client) OnRefreshed(data []byte) {
	lob := &md.Lobby{}
	err := proto.Unmarshal(data, lob)
	if err != nil {
		log.Panicln("Error Unmarshalling Request")
	}
	c.menu.RefreshPeers(lob, c.node)
}

// @ Inherited Method: Handle Invite ^ //
func (c *Client) OnInvited(data []byte) {
	// Unmarshal Invite
	m := &md.AuthInvite{}
	err := proto.Unmarshal(data, m)
	if err != nil {
		log.Panicln("Error Unmarshalling Request")
	}

	// Check Auth
	decs := ui.ShowAuthDialog(m)

	// Check Invite
	if m.Payload == md.Payload_MEDIA {
		c.node.Respond(decs)
	} else if m.Payload == md.Payload_URL {
		err := op.Start(m.Card.Url.Link)
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
func (c *Client) OnDirected(data []byte) {
	log.Println(data)
}

// @ Inherited Method: Handle Progress ^ //
func (c *Client) OnProgress(data float32) {
	log.Println(data)
}

// @ Inherited Method: Handle Received ^ //
func (c *Client) OnReceived(data []byte) {
	// Unmarshal Data
	meta := &md.Metadata{}
	err := proto.Unmarshal(data, meta)
	if err != nil {
		log.Panicln("Error Unmarshalling Request")
	}

	// Move File
	err = op.Start(meta.Path)
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
	// Unmarshal Data
	errMsg := &md.ErrorMessage{}
	err := proto.Unmarshal(data, errMsg)
	if err != nil {
		log.Panicln("Error Unmarshalling Request")
	}

	// Display Message
	ui.ShowErrorDialog(errMsg)
}
