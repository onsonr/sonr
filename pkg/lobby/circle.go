package lobby

import (
	"encoding/json"
	"fmt"
	"math"
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

// Returns Difference Between two peers
func getDifference(sendDir float64, recDir float64) float64 {
	// Get Receiver Antipodal Degrees
	receiverAntipodal := getAntipodal(recDir)

	// Difference between angles
	if receiverAntipodal > sendDir {
		theta := receiverAntipodal - sendDir
		return theta * (math.Pi / 180)
	}
	theta := sendDir - receiverAntipodal
	return theta * (math.Pi / 180)
}

// Gets antipodal version of degrees
func getAntipodal(degrees float64) float64 {
	if degrees > 180 {
		return degrees - 180
	}
	return degrees + 180
}

// ListPeers returns peerids in room
func (lob *Lobby) handleEvents() {
	peerRefreshTicker := time.NewTicker(time.Second * 3)
	defer peerRefreshTicker.Stop()

	for {
		select {
		// ** when we receive a message from the lobby room **
		case m := <-lob.Messages:
			// Update Circle by event
			if m.Event == "Join" {
				lob.joinPeer(m.Value)
			} else if m.Event == "Update" {
				lob.updatePeer(m.Value)
			} else if m.Event == "Leave" {
				lob.removePeer(m.Value)
			}

		// ** refresh the list of peers in the chat room periodically **
		case <-peerRefreshTicker.C:
			// Check if Peer is in Lobby
			for _, peer := range lob.peers {
				// Find Peer in range
				inRange := findPeerID(lob.ps.ListPeers(olcName(lob.Code)), peer.ID)

				// No longer available
				if !inRange {
					// Verify not user
					if peer.ID != lob.Self.ID {
						// Delete peer from circle
						fmt.Println("Peer no longer in lobby")

						// Send Callback of new peers
						// Todo: lob.callback.OnRefresh(lob.GetCircle())
					}
					continue
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
	peer.Device = data["Device"].(string)
	peer.FirstName = data["FirstName"].(string)
	peer.LastName = data["LastName"].(string)
	peer.ProfilePic = data["ProfilePic"].(string)

	// Add Peer
	lob.peers = append(lob.peers, *peer)
}

// updatePeer changes peer values in circle
func (lob *Lobby) removePeer(jsonString string) {

}

// updatePeer changes peer values in circle
func (lob *Lobby) updatePeer(jsonString string) {
	// Generate Map
	notif := new(Notification)
	err := json.Unmarshal([]byte(jsonString), notif)
	if err != nil {
		fmt.Println("Sonr P2P Error: ", err)
	}

	// Update peer in Slice
	for _, peer := range lob.peers {
		// Find Peers
		if peer.ID == notif.ID {
			// Update Values for Peer
			peer.Direction = notif.Direction
		}
	}

	// Send Callback with updated peers
	// Todo: lob.callback.OnRefresh(lob.GetCircle())
}
