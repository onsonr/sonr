package lobby

import (
	"encoding/json"
	"fmt"
	"math"
)

// ******************************** //
// ********** Interface  ********** //
// ******************************** //

// Peer is a representative in the lobby for a device
type Peer struct {
	ID         string
	Device     string
	FirstName  string
	LastName   string
	ProfilePic string
	Direction  float64
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

// ******************************** //
// ********** Manage Peers ******** //
// ******************************** //
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
	lob.peers[peer.ID] = *peer
}

// removePeer deletes a peer from the circle
func (lob *Lobby) removePeer(id string) {
	// Delete peer at id
	delete(lob.peers, id)

	// Send Callback with updated peers
	lob.callback.OnRefresh(lob.GetPeers())
}

// updatePeer changes peer values in circle
func (lob *Lobby) updatePeer(jsonString string) {
	// Generate Map
	notif := new(Notification)
	err := json.Unmarshal([]byte(jsonString), notif)
	if err != nil {
		fmt.Println("Sonr P2P Error: ", err)
	}

	println("Update Message: ", jsonString)

	// Update peer in Dictionary
	peerRef := lob.peers[notif.ID]
	peerRef.Direction = notif.Direction
	lob.peers[notif.ID] = peerRef

	println("Lobby Count: ", len(lob.peers))

	// Send Callback with updated peers
	lob.callback.OnRefresh(lob.GetPeers())
}
