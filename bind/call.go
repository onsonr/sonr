package sonr

import (
	"fmt"
	"log"

	sonrModel "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ** Summary ** //
// Class to handle Callbacks uses callback in Node
// Marshals data and sends Callback

// ^ callback Method with type ^
func (sn *Node) callback(call sonrModel.CallbackType, data proto.Message) {
	// ** Convert Message to bytes **
	bytes, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("Cannot Marshal Error Protobuf: ", err)
	}

	// ** Check Call Type **
	switch call {
	// @ Lobby Refreshed
	case sonrModel.CallbackType_REFRESHED:
		sn.callbackRef.OnRefreshed(bytes)

	// @ File has Queued
	case sonrModel.CallbackType_QUEUED:
		sn.callbackRef.OnQueued(bytes)

	// @ Peer has been Invited
	case sonrModel.CallbackType_INVITED:
		sn.callbackRef.OnInvited(bytes)

	// @ Peer has Responded
	case sonrModel.CallbackType_RESPONDED:
		sn.callbackRef.OnResponded(bytes)

	// @ Transfer has Completed
	case sonrModel.CallbackType_COMPLETED:
		sn.callbackRef.OnCompleted(bytes)
	}
}

// ^ error Callback with error instance, and method ^
func (sn *Node) error(err error, method string) {
	// Create Error ProtoBuf
	errorMsg := sonrModel.ErrorMessage{
		Message: err.Error(),
		Method:  method,
	}

	// Convert Message to bytes
	bytes, err := proto.Marshal(&errorMsg)
	if err != nil {
		fmt.Println("Cannot Marshal Error Protobuf: ", err)
	}
	// Send Callback
	sn.callbackRef.OnError(bytes)

	// Log In Core
	log.Fatalln(fmt.Sprintf("[Error] At Method %s : %s", err.Error(), method))
}
