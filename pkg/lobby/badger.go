package lobby

import (
	"encoding/json"
	"errors"
	"fmt"

	badger "github.com/dgraph-io/badger/v2"
	"github.com/libp2p/go-libp2p-core/peer"
)

// GetPeer returns ONE Peer in Datastore
func (lob *Lobby) GetPeer(queryID string) (Peer, error) {
	// Initialize Object
	var peer Peer

	// ** Create Transaction ** //
	err := lob.peerDB.View(func(txn *badger.Txn) error {
		// Set Transaction Query
		item, err := txn.Get([]byte(queryID))

		// @ Find Item
		err = item.Value(func(val []byte) error {
			// Convert Value to String Add to Slice
			cm := new(Peer)
			err := json.Unmarshal(val, cm)
			// Check for Error
			if err != nil {
				fmt.Println("JSON Error ", err)
			} else {
				// @ Add Item value to Object
				peer = *cm
			}
			return nil
		})

		// Check for Error
		if err != nil {
			return err
		}
		return nil
	})

	// Check for Error
	if err != nil {
		fmt.Println("Search Error ", err)
	}
	return peer, nil
}

// GetPeer returns ONE Peer in Datastore
func (lob *Lobby) GetPeerID(idStr string) (peer.ID, error) {
	// Get Lobby PeerID Slice
	lobbyPeers := lob.ps.ListPeers(lob.Code)

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
func (lob *Lobby) GetAllPeers() string {
	// ** Initialize Variables ** //
	var peerSlice []Peer

	// ** Open Data Store Read Transaction ** //
	err := lob.peerDB.View(func(txn *badger.Txn) error {
		// @ Create Iterator
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		// @ Iterate over bucket
		for it.Rewind(); it.Valid(); it.Next() {
			// Get Item and Key
			item := it.Item()

			// Get Item Value
			err := item.Value(func(peer []byte) error {
				// Convert Value to String Add to Slice
				cm := new(Peer)
				err := json.Unmarshal(peer, cm)

				// Check for Error
				if err != nil {
					fmt.Println("JSON Error", err)
				} else {
					// Add Item value to Slice
					peerSlice = append(peerSlice, *cm)
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
		fmt.Println("Transaction Erro ", err)
	}

	// ** Convert slice to bytes ** //
	bytes, err := json.Marshal(peerSlice)
	if err != nil {
		println("Error converting peers to json ", err)
	}

	// Return as string
	return string(bytes)
}
