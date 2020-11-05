package sonr

import (
	"encoding/json"
	"fmt"

	"github.com/sonr-io/core/pkg/lobby"
	"github.com/sonr-io/core/pkg/user"
)

// Send publishes a message to the SonrNode lobby
func (sn *Node) Send(data string) bool {
	// Log Send
	fmt.Println("Sonr P2P General-Send: ", data)

	// Create Message Type
	cm := new(lobby.Message)
	err := json.Unmarshal([]byte(data), cm)
	if err != nil {
		fmt.Println("Sonr P2P Error: ", err)
		return false
	}

	// Publish to Lobby
	err = sn.Lobby.Publish(*cm)
	if err != nil {
		fmt.Println("Sonr P2P Error: ", err)
		return false
	}

	// Return Success
	return true
}

// GetUser returns profile and contact in a map as string
func (sn *Node) GetUser() string {
	// Initialize Map
	m := make(map[string]string)
	m["profile"] = sn.Profile.String()
	m["contact"] = sn.Contact.String()
	m["id"] = sn.PeerID

	// Convert to JSON
	msgBytes, err := json.Marshal(m)
	if err != nil {
		println(err)
	}

	// Return String
	return string(msgBytes)
}

// SetUser from connection request
func (sn *Node) SetUser(cm lobby.ConnectRequest) error {
	// Set Profile
	profile := user.NewProfile(sn.Host.ID().String(), cm.OLC, cm.Device)
	sn.Profile = profile

	// Set Contact
	contact := user.NewContact(cm.Contact)
	sn.Contact = contact

	return nil
}

// Update occurs when status or direction changes
func (sn *Node) Update(data string) bool {
	// Update User Values
	err := sn.Profile.Update(data)
	if err != nil {
		fmt.Println("Sonr P2P Error: ", err)
		return false
	}

	// Log New Values
	println("Direction: ", sn.Profile.Direction)
	println("Status: ", sn.Profile.Status.String())

	// Create Update Map
	v := make(map[string]string)
	v["state"] = sn.Profile.State()
	v["info"] = sn.Contact.Basic()

	// Convert to JSON
	msgBytes, err := json.Marshal(v)
	if err != nil {
		println(err)
	}

	// Create Message
	cm := new(lobby.Message)
	cm.FirstName = sn.Contact.FirstName
	cm.Event = "Update"
	cm.SenderID = sn.PeerID
	cm.Value = string(msgBytes)

	// Inform Lobby
	err = sn.Lobby.Publish(*cm)
	if err != nil {
		fmt.Println("Sonr P2P Error: ", err)
		return false
	}

	// Return Success
	return true
}
