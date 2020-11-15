package lobby

import (
	"encoding/json"
	"fmt"
)

// ^ Interface ^

// Peer is a representative in the lobby for a device
type Peer struct {
	ID         string
	Device     string
	FirstName  string
	LastName   string
	ProfilePic string
	Direction  float64
}

// Bytes converts message struct to JSON bytes
func (p *Peer) Bytes() []byte {
	// Convert to Bytes
	msgBytes, err := json.Marshal(p)
	if err != nil {
		println(err)
	}
	return msgBytes
}

// String converts message struct to JSON String
func (p *Peer) String() string {
	// Convert to JSON
	peerBytes, err := json.Marshal(p)
	if err != nil {
		println(err)
	}
	return string(peerBytes)
}

// ^ Manage Peers ^
// joinPeer adds a peer to dictionary
func (lob *Lobby) joinPeer(jsonString string) {
	// Generate Map
	byt := []byte(jsonString)
	var data map[string]interface{}
	err := json.Unmarshal(byt, &data)

	// Check error
	if err != nil {
		panic(err)
	}

	// Set Values
	peer := new(Peer)
	peer.ID = data["ID"].(string)
	peer.Device = data["Device"].(string)
	peer.FirstName = data["FirstName"].(string)
	peer.LastName = data["LastName"].(string)
	peer.ProfilePic = data["ProfilePic"].(string)
	peer.Direction = data["Direction"].(float64)

	// Add Peer to dictionary
	lob.peers[peer.ID] = peer

	// Send Callback with updated peers
	lob.callback.OnRefresh(lob.GetPeers())
}

// ^ removePeer deletes a peer from the circle ^
func (lob *Lobby) removePeer(id string) {
	// Delete peer at id
	delete(lob.peers, id)

	// Send Callback with updated peers
	lob.callback.OnRefresh(lob.GetPeers())
	println("")
}

// ^ Search for Peer in Pub/Sub Topic ^ //
func (lob *Lobby) searchPeer(queryID string) bool {
	// Get Pub/Sub Topic Peers and Iterate
	for _, id := range lob.ListPeers() {
		// If Found
		if id.String() == queryID {
			return true
		}
	}
	// If Not Found
	return false
}

// ^ updatePeer changes peer values in circle ^
func (lob *Lobby) updatePeer(jsonString string) {
	// Generate Map
	peer := new(Peer)
	err := json.Unmarshal([]byte(jsonString), peer)
	if err != nil {
		fmt.Println("Sonr P2P Error: ", err)
	}

	// Update peer in Dictionary
	lob.peers[peer.ID] = peer

	// Send Callback with updated peers
	lob.callback.OnRefresh(lob.GetPeers())
}
