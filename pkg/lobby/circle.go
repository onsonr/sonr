package lobby

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
)

func findPeerID(slice []peer.ID, val string) bool {
	for _, item := range slice {
		if item.String() == val {
			return true
		}
	}
	return false
}

// GetCircle returns available peers as string
func (lob *Lobby) GetCircle() string {
	// Create new map
	temp := lob.circle

	// Remove Peers that arent available
	for peer := range temp {
		// Not Available
		if peer.Status != "Available" {
			// Delete from temp map
			fmt.Println("Peer is not available")
			delete(temp, peer)
		}
	}

	// Get all peers from map
	keys := make([]Peer, 0, len(temp))
	for k := range temp {
		keys = append(keys, k)
	}

	// Convert map to bytes
	bytes, err := json.Marshal(keys)
	if err != nil {
		println("Error converting peers to json ", err)
	}

	// Return as string
	return string(bytes)
}

// ListPeers returns peerids in room
func (lob *Lobby) handleEvents() {
	peerRefreshTicker := time.NewTicker(time.Second)
	defer peerRefreshTicker.Stop()

	for {
		select {
		// ** when we receive a message from the lobby room **
		case m := <-lob.Messages:
			println(m.Event)

		// ** refresh the list of peers in the chat room periodically **
		case <-peerRefreshTicker.C:
			// Check if Peer is in Lobby
			for peer := range lob.circle {
				// Find Peer
				found := findPeerID(lob.ps.ListPeers(olcName(lob.Code)), peer.ID)

				// Not Found
				if !found {
					// Delete peer from circle
					fmt.Println("Peer no longer in lobby")
					delete(lob.circle, peer)

					// Send Callback of new peers
					lob.callback.OnRefresh(lob.GetCircle())
				}
			}
			continue

		case <-lob.ctx.Done():
			return

		case <-lob.doneCh:
			return
		}
	}
}

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
	peer.Status = data["Status"].(string)
	peer.Device = data["Device"].(string)
	peer.FirstName = data["FirstName"].(string)
	peer.LastName = data["LastName"].(string)
	peer.ProfilePic = data["ProfilePic"].(string)

	// Add Peer
	lob.circle[Peer{ID: peer.String()}] = -1
}

// updatePeer changes peer values in circle
func (lob *Lobby) updatePeer(jsonString string) {
	// Generate Map
	notif := new(Notification)
	err := json.Unmarshal([]byte(jsonString), notif)
	if err != nil {
		fmt.Println("Sonr P2P Error: ", err)
	}

	// Iterate Peers in circle
	for peer := range lob.circle {
		// Find Peers
		if peer.ID == notif.ID {
			// Update Values for Peer
			peer.Direction = notif.Direction
			peer.Status = notif.Status
		}
	}

	// Send Callback with updated peers
	lob.callback.OnRefresh(lob.GetCircle())
}
