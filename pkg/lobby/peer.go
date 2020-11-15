package lobby

import (
	"encoding/json"
	"fmt"
	"math"
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
	lob.peers[peer.ID] = *peer
}

// ^ removePeer deletes a peer from the circle ^
func (lob *Lobby) removePeer(id string) {
	// Delete peer at id
	delete(lob.peers, id)

	// Send Callback with updated peers
	lob.callback.OnRefresh(lob.GetPeers())
	println("")
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
	lob.peers[peer.ID] = *peer

	// Send Callback with updated peers
	lob.callback.OnRefresh(lob.GetPeers())
}

// ^ validatePeers checks if all peers in dictionary are still in lobby ^
func (lob *Lobby) validatePeers() {
	// Get Pub/Sub Topic Peers
	inLobbyPeers := lob.ListPeers()
	inDictNotLobbyPeers := lob.peers

	// Temp Logging
	fmt.Println("In Lobby Count: ", len(inLobbyPeers))
	fmt.Println("In Dict Count: ", len(inDictNotLobbyPeers))

	// Iterate through Slice:inLobbyPeers and remove from inDictNotLobbyPeers
	for _, id := range inLobbyPeers {
		// Remove Peers that are still in lobby
		delete(inDictNotLobbyPeers, id.String())
	}

	// Temp Logging
	fmt.Println("In Dict Not Lobby Count: ", len(inDictNotLobbyPeers))

	// Check if Peers need to be disposed
	if len(inDictNotLobbyPeers) > 0 {
		// Iterate through Dict:inDictNotLobbyPeers and delete from actual dictionary
		for id := range inDictNotLobbyPeers {
			delete(lob.peers, id)
		}

		// Send Callback with updated peers after disposal
		lob.callback.OnRefresh(lob.GetPeers())
	}
}
