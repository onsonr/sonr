package lobby

import (
	"errors"
	"fmt"

	badger "github.com/dgraph-io/badger/v2"
	"github.com/libp2p/go-libp2p-core/peer"
	pb "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// GetPeer returns ONE Peer in Datastore
func (lob *Lobby) GetPeer(queryID string) *pb.Peer {
	// @ 1. Get all Peers
	availablePeers := lob.Info.Peers

	// @ 2. Iterate Through Peers, Return Matched Peer
	for _, peer := range availablePeers {
		// Check if peer matches query
		if peer.Id == queryID {
			return peer
		}
	}
	// ! Return Nil if no matches
	return nil
}

// GetPeer returns ONE Peer in Datastore
func (lob *Lobby) GetPubSubID(idStr string) (peer.ID, error) {
	// Get Lobby PeerID Slice
	lobbyPeers := lob.ps.ListPeers(lob.Info.Code)

	// Get Pub/Sub Topic Peers and Iterate
	for _, id := range lobbyPeers {
		// If Found
		if id.String() == idStr {
			return id, nil
		}
	}
	// Create New Error
	err := errors.New("Peer ID for given query not found")
	return "", err
}

// GetAllPeers returns ALL Peers in Datastore
func (lob *Lobby) GetAllPeers() []byte {
	// ** Initialize Variables ** //
	var peers pb.RefreshMessage
	var peerCount int32
	peerCount = 0

	// ** Open Data Store Read Transaction ** //
	err := lob.peerDB.View(func(txn *badger.Txn) error {
		// @ Create Iterator
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		// @ Iterate over bucket
		for it.Rewind(); it.Valid(); it.Next() {
			// Get Item and Key
			item := it.Item()
			peerCount += 1

			// Get Item Value
			err := item.Value(func(data []byte) error {
				// Convert Value to String Add to Slice
				peer := pb.Peer{}
				err := proto.Unmarshal(data, &peer)
				if err != nil {
					fmt.Println("unmarshaling error: ", err)
					return err
				} else {
					// Add Item value to Slice
					peers.AvailablePeers = append(peers.AvailablePeers, &peer)
				}
				return nil
			})

			// Check error
			if err != nil {
				return err
			}
		}
		return nil
	})

	// Check for Error
	if err != nil {
		fmt.Println("Transaction Error ", err)
	}

	// Add additional data
	peers.Count = peerCount
	peers.Olc = lob.Code

	// Convert to bytes
	data, err := proto.Marshal(&peers)
	if err != nil {
		fmt.Println("Error Marshaling RefreshMessage ", err)
	}

	// Return as JSON String
	return data
}

// ^ Checks for Peer in Pub/Sub Topic ^ //
func (lob *Lobby) isPeerInLobby(queryID string) bool {
	// Get Lobby PeerID Slice
	lobbyPeers := lob.ps.ListPeers(lob.Code)

	// Get Pub/Sub Topic Peers and Iterate
	for _, id := range lobbyPeers {
		// If Found
		if id.String() == queryID {
			return true
		}
	}
	// If Not Found
	return false
}

// ^ removePeer deletes a peer from the circle ^
func (lob *Lobby) removePeer(msg *pb.LobbyMessage) {
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
	lob.call.Refreshed(lob.GetAllPeers())
}

// ^ updatePeer changes peer values in circle ^
func (lob *Lobby) updatePeer(msg *pb.LobbyMessage) {
	// Create Key/Value as Bytes
	key := []byte(id)

	// Update peer in DataStore
	err := lob.peerDB.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry(key, value)
		err := txn.SetEntry(e)
		return err
	})

	// Check Error
	if err != nil {
		fmt.Println("Error Updating Peer in Badger", err)
	}

	// Send Callback with updated peers
	lob.call.Refreshed(lob.GetAllPeers())
}
