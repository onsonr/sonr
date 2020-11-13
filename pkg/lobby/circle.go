package lobby

import (
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	"gonum.org/v1/gonum/graph"
)

func findPeerID(slice []peer.ID, val string) bool {
	for _, item := range slice {
		if item.String() == val {
			return true
		}
	}
	return false
}

func findPeerFromGraph(slice []Peer, val int64) (int, bool) {
	for i, item := range slice {
		if item.GraphID == val {
			return i, true
		}
	}
	return -1, false
}

// GetCircle returns available peers as string
func (lob *Lobby) GetCircle() string {
	// Create new map
	var peerSlice []Peer
	nodes := lob.circle.To(lob.Self.GraphID)

	for i := 0; i < nodes.Len(); i++ {
		index, found := findPeerFromGraph(lob.peers, nodes.Node().ID())
		if found {
			// Add Peer at Index
			peerSlice = append(peerSlice, lob.peers[index])
		}
	}

	// Convert map to bytes
	bytes, err := json.Marshal(peerSlice)
	if err != nil {
		println("Error converting peers to json ", err)
	}

	// Return as string
	return string(bytes)
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
	peerRefreshTicker := time.NewTicker(time.Second)
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
			}

		// ** refresh the list of peers in the chat room periodically **
		case <-peerRefreshTicker.C:
			// Check if Peer is in Lobby
			for _, peer := range lob.peers {
				// Find Peer in range and in graph
				inRange := findPeerID(lob.ps.ListPeers(olcName(lob.Code)), peer.ID)
				_, inGraph := findPeerFromGraph(lob.peers, peer.GraphID)

				// No longer available
				if !inRange && inGraph {
					// Verify not user
					if peer.ID != lob.Self.ID {
						// Delete peer from circle
						fmt.Println("Peer no longer in lobby")
						lob.circle.RemoveNode(peer.GraphID)

						// Send Callback of new peers
						lob.callback.OnRefresh(lob.GetCircle())
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
	peer.Status = data["Status"].(string)
	peer.Device = data["Device"].(string)
	peer.FirstName = data["FirstName"].(string)
	peer.LastName = data["LastName"].(string)
	peer.ProfilePic = data["ProfilePic"].(string)

	// Add Peer
	graphID := lob.circle.NewNode()
	peer.GraphID = graphID.ID()
	lob.circle.AddNode(graphID)
	lob.peers = append(lob.peers, *peer)
}

// Update Edge with new degrees difference
func (lob *Lobby) updateEdge(sender graph.Node, receiver graph.Node, difference float64) {
	// Check if edge exists
	result := lob.circle.WeightedEdge(sender.ID(), receiver.ID())

	// Remove if it exists
	if result != nil {
		lob.circle.RemoveEdge(sender.ID(), receiver.ID())
	}

	// Create new edge
	lob.circle.NewWeightedEdge(sender, receiver, difference)
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
			peer.Status = notif.Status
		}
	}

	// Check user status
	if lob.Self.Status == "Searching" {
		// Get User Node
		sender := lob.circle.Node(lob.Self.GraphID)

		// Get Receiver Node
		receiver := lob.circle.Node(notif.GraphID)

		// Update Edge
		difference := getDifference(lob.Self.Direction, notif.Direction)
		lob.updateEdge(sender, receiver, difference)
	}

	// Send Callback with updated peers
	lob.callback.OnRefresh(lob.GetCircle())
}
