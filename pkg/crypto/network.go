// It converts a `WebauthnCredential` to a `webauthn.Credential`
package crypto

import (
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
)

// A Network is a channel that sends messages to parties and receives messages from parties.
type Network interface {
	// Ls returns a list of peers that are connected to the network.
	Ls() []party.ID

	// A function that takes in a party ID and returns a channel of protocol messages.
	Next(id party.ID) <-chan *protocol.Message

	// Sending a message to the network.
	Send(msg *protocol.Message)

	// A channel that is closed when the party is done with the protocol.
	Done(id party.ID) chan struct{}

	// A function that is called when a party is done with the protocol.
	Quit(id party.ID)

	// IsOnlineNetwork returns true if the network is an online network.
	IsOnlineNetwork() bool
}
