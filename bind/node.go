package sonr

import (
	"context"
	"fmt"
	"math"
	"sync"

	badger "github.com/dgraph-io/badger/v2"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/protocol"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/sonr-io/core/pkg/file"
	"github.com/sonr-io/core/pkg/lobby"
	pb "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ^ Struct Management ^ //
// Node contains all values for user
type Node struct {
	CTX        context.Context
	Host       host.Host
	PubSub     *pubsub.PubSub
	Lobby      lobby.Lobby
	FileQueue  *badger.DB
	Profile    pb.Profile
	Contact    pb.Contact
	AuthStream authStreamConn
	Call       Callback
}

// ^ Sends new proximity/direction update ^ //
// Update occurs when status or direction changes
func (sn *Node) Update(data []byte) bool {
	// ** Initialize ** //
	updateEvent := pb.UpdateEvent{}
	err := proto.Unmarshal(data, &updateEvent)
	if err != nil {
		fmt.Println("unmarshaling error: ", err)
		return false
	}

	// Update User Values
	sn.Profile.Direction = math.Round(updateEvent.NewDirection*100) / 100

	// Create Message with Updated Info
	notif := &pb.LobbyMessage{
		Event:  "Update",
		Sender: sn.Profile.HostId,
		Data:   sn.getPeerInfo(),
	}

	// Inform Lobby
	err = sn.Lobby.Publish(notif)
	if err != nil {
		fmt.Println("Error Posting NotifUpdate: ", err)
		return false
	}
	return true
}

// ^ Queue adds a file to Process for Transfer, returns key ^ //
// TODO: Implement an Error Schema with proto
func (sn *Node) Queue(data []byte) bool {
	// ** Initialize ** //
	queuedFile := pb.QueueEvent{}
	err := proto.Unmarshal(data, &queuedFile)
	if err != nil {
		fmt.Println("unmarshaling error: ", err)
		sn.Callback(pb.Callback_PROCESSED, nil)
		return false
	}

	// ** Create Metadata ** //
	meta := file.GetMetadata(queuedFile.FilePath)
	if err != nil {
		fmt.Println("Error Getting Metadata", err)
		sn.Callback(pb.Callback_PROCESSED, nil)
		return false
	}

	// ** Create Thumbnail ** //
	var wg sync.WaitGroup
	wg.Add(2)

	// Run Routine
	var thumb []byte
	go func(tn []byte) {
		tn = file.GetThumbnail(meta)
		fmt.Println("Raw Thumbnail: ", thumb)
		wg.Done()
	}(thumb)

	// ** Create Processed File ** //
	go func() {
		// Create Type
		processedFile := &pb.ProcessedMessage{
			Metadata:  meta,
			Thumbnail: thumb,
		}

		// Convert to bytes
		raw, err := proto.Marshal(processedFile)
		if err != nil {
			fmt.Println("Error Marshalling Processed File", err)
			sn.Callback(pb.Callback_PROCESSED, nil)
		}

		// ** Add to Badger Store ** //
		// Update peer in DataStore
		err = sn.FileQueue.Update(func(txn *badger.Txn) error {
			e := badger.NewEntry([]byte(meta.FileId), raw)
			err := txn.SetEntry(e)
			return err
		})

		// Check Error
		if err != nil {
			fmt.Println("Error Updating Peer in Badger", err)
			sn.Callback(pb.Callback_PROCESSED, nil)
		}
		wg.Done()
	}()

	// Send Callback with file ID after both tasks finish
	wg.Wait()

	// Convert to bytes
	metaRaw, err := proto.Marshal(meta)
	if err != nil {
		fmt.Println("Error Marshalling Processed File", err)
		sn.Callback(pb.Callback_PROCESSED, nil)
	}

	// Send bytes
	sn.Callback(pb.Callback_PROCESSED, metaRaw)
	return true
}

// ^ Invite an available peer to transfer ^ //
func (sn *Node) Invite(data []byte) bool {
	// Initialize
	invite := pb.InviteEvent{}
	err := proto.Unmarshal(data, &invite)
	if err != nil {
		fmt.Println("unmarshaling error: ", err)
		return false
	}

	// ** Get Required Data **
	peerID, err := sn.Lobby.GetPeerID(invite.Peer.Id)
	if err != nil {
		fmt.Println("Search Error", err)
		return false
	}

	// Retreive File Info from Memory Store
	var fileInfoRaw []byte
	err = sn.FileQueue.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(invite.FileId))
		err = item.Value(func(val []byte) error {
			// Accessing val here is valid.
			fmt.Printf("The Metadata is: %s\n", val)
			fileInfoRaw = append([]byte{}, val...)
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})

	// Send Error
	if err != nil {
		sn.NewError(err, 4, pb.Error_BADGER)
	}

	// Unmarshal into Protobuf
	fileInfo := pb.ProcessedMessage{}
	err = proto.Unmarshal(fileInfoRaw, &fileInfo)
	if err != nil {
		fmt.Println("Error unmarshaling msg into json: ", err)
		return false
	}

	// ** Create New Auth Stream **
	stream, err := sn.Host.NewStream(sn.CTX, peerID, protocol.ID("/sonr/auth"))
	if err != nil {
		fmt.Println("Auth Stream Failed to Open ", err)
		return false
	}
	// Establish Auth Stream
	sn.NewAuthStream(stream)

	// Create Request Message
	authPbf := &pb.AuthMessage{
		Subject:   pb.AuthMessage_REQUEST,
		Peer:      sn.getPeerInfo(),
		Metadata:  fileInfo.Metadata,
		Thumbnail: fileInfo.Thumbnail,
	}

	// ** Send Invite Message **
	err = sn.AuthStream.Write(authPbf)
	if err != nil {
		return false
	}

	// Return Success
	return true
}

// ^ Respond to an Invitation ^ //
func (sn *Node) Respond(data []byte) bool {
	// Initialize Event
	response := pb.RespondEvent{}
	err := proto.Unmarshal(data, &response)
	if err != nil {
		fmt.Println("unmarshaling error: ", err)
		return false
	}

	// Initialize Response
	authMsg := new(pb.AuthMessage)

	// Check Decision
	if response.Decision == true {
		// Set as Accept
		authMsg = &pb.AuthMessage{
			Subject: pb.AuthMessage_ACCEPT,
			Peer:    sn.getPeerInfo(),
		}
	} else {
		// Set as Decline
		authMsg = &pb.AuthMessage{
			Subject: pb.AuthMessage_DECLINE,
			Peer:    sn.getPeerInfo(),
		}
	}

	// Send Message
	err = sn.AuthStream.Write(authMsg)
	if err != nil {
		return false
	}

	// Succesful
	return true
}
