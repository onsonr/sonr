package lobby

import (
	"encoding/json"
	"fmt"

	badger "github.com/dgraph-io/badger/v2"
	"github.com/libp2p/go-libp2p-core/peer"
)

// ^ Interface ^

// Peer is a representative in the lobby for a device
type Peer struct {
	ID         peer.ID
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

// ^ Checks for Peer in Pub/Sub Topic ^ //
func (lob *Lobby) isPeerInLobby(queryID string) bool {
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

// ^ Returns Peer in Data Store ^ //
func (lob *Lobby) GetPeer(queryID string) Peer {
	// Initialize Object
	var value []byte

	// Create Transaction
	err := lob.peerDB.View(func(txn *badger.Txn) error {
		// Set Transaction Query
		item, err := txn.Get([]byte(queryID))

		// Find Item
		err = item.Value(func(val []byte) error {
			// Copying or parsing val is valid.
			value = append([]byte{}, val...)
			return nil
		})

		// Check for Error
		if err != nil {
			fmt.Println(err)
		}
		return nil
	})

	// Generate Map
	peer := new(Peer)
	err = json.Unmarshal(value, peer)
	if err != nil {
		fmt.Println("Marshal Error: ", err)
	}
	return *peer
}

// ^ removePeer deletes a peer from the circle ^
func (lob *Lobby) removePeer(id string) {
	// Delete peer from datastore
	key := []byte(id)
	err := lob.peerDB.Update(func(txn *badger.Txn) error {
		err := txn.Delete(key)
		return err
	})

	// Check for Error
	if err != nil {
		fmt.Println(err)
	}

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

	// Create Key/Value as Bytes
	key := []byte(peer.ID)
	value := peer.Bytes()

	// Update peer in DataStore
	err = lob.peerDB.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry(key, value)
		err := txn.SetEntry(e)
		return err
	})

	// Send Callback with updated peers
	lob.callback.OnRefresh(lob.GetPeers())
}
