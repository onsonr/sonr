package transfer

import (
	"context"
	"log"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ****************** //
// ** GRPC Service ** //
// ****************** //
// Argument is AuthMessage protobuf
type AuthArgs struct {
	Data []byte
	From string
}

// Reply is also AuthMessage protobuf
type AuthReply struct {
	Data     []byte
	Decision bool
}

// Service Struct
type Authorization struct {
	// Current Data
	currArgs  AuthArgs
	currReply *AuthReply
	invited   OnProtobuf
	responded OnProtobuf
	peerConn  *PeerConnection
}

// ^ Calls Invite on Remote Peer ^ //
func (as *Authorization) Invite(ctx context.Context, args AuthArgs, reply *AuthReply) error {
	log.Println("Received a Invite call: ", args.Data)
	// Set Current Data
	as.currArgs = args
	as.currReply = reply

	// Set Peer ID
	as.peerConn.peerID = as.peerConn.Find(args.From)

	// Set Current Message
	err := proto.Unmarshal(args.Data, as.peerConn.currMessage)
	if err != nil {
		onError(err, "process")
		return err
	}

	// Send Callback
	as.invited(args.Data)
	if err != nil {
		return err
	}
	return nil
}

// ^ Send Invite to a Peer ^ //
func (pc *PeerConnection) Invite(id peer.ID, info *md.Peer, sm *sf.SafeMetadata) {
	// Set SafeFile
	pc.safeFile = sm

	// Create Invite Message
	reqMsg := &md.AuthMessage{
		Event:    md.AuthMessage_REQUEST,
		From:     info,
		Metadata: sm.GetMetadata(),
	}

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(reqMsg)
	if err != nil {
		onError(err, "sendInvite")
		log.Println(err)
	}

	// Send GRPC Call
	go func(id peer.ID, data []byte) {
		// Initialize Vars
		var reply AuthReply
		var args AuthArgs
		args.Data = msgBytes

		// Set Data
		startTime := time.Now()

		// Call to Peer
		err = pc.rpcClient.Call(id, "Authorization", "Invite", args, &reply)

		if err != nil {
			onError(err, "sendInvite")
			log.Panicln(err)
		}

		// End Tracking
		endTime := time.Now()
		diff := endTime.Sub(startTime)
		log.Printf("Auth from %s: time=%s\n", id, diff)

		// Send Callback and Reset
		pc.respondedCall(reply.Data)

		// Handle Response
		if reply.Decision {
			// Begin Transfer
			pc.SendFile()
		}
	}(id, msgBytes)
}

// ^ Send Accept Message on Stream ^ //
func (pc *PeerConnection) SendResponse(decision bool, selfInfo *md.Peer) {
	// Initialize Message
	var respMsg *md.AuthMessage

	// Check Decision
	if decision {
		// Initialize Transfer
		savePath := "/" + pc.currMessage.Metadata.Name + "." + pc.currMessage.Metadata.Mime.Subtype
		pc.transfer = NewTransfer(savePath, pc.currMessage.Metadata, pc.currMessage.From, pc.progressCall, pc.completedCall)

		// Create Accept Response
		respMsg = &md.AuthMessage{
			From:  selfInfo,
			Event: md.AuthMessage_ACCEPT,
		}
	} else {
		// Reset Peer Info
		pc.peerID = ""
		pc.currMessage = nil

		// Create Decline Response
		respMsg = &md.AuthMessage{
			From:  selfInfo,
			Event: md.AuthMessage_DECLINE,
		}
	}

	// Convert Protobuf to bytes
	msgBytes, err := proto.Marshal(respMsg)
	if err != nil {
		onError(err, "sendInvite")
		log.Println(err)
	}

	// Send GRPC Call
	go func(d bool, msgBytes []byte) {
		// Send Reply
		pc.ascv.currReply.Data = msgBytes
		pc.ascv.currReply.Decision = d
	}(decision, msgBytes)
}
