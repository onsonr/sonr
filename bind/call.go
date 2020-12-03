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

// ^ Callback Method with type ^
func (sn *Node) Callback(call sonrModel.CallbackType, data proto.Message) {
	// Convert Message to bytes
	bytes, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("Cannot Marshal Error Protobuf: ", err)
	}

	// Check Call Type
	switch call {
	case sonrModel.CallbackType_REFRESHED:
		sn.call.OnRefreshed(bytes)
	case sonrModel.CallbackType_QUEUED:
		sn.call.OnQueued(bytes)
	case sonrModel.CallbackType_INVITED:
		sn.call.OnQueued(bytes)
	case sonrModel.CallbackType_RESPONDED:
		sn.call.OnResponded(bytes)
	case sonrModel.CallbackType_COMPLETED:
		sn.call.OnQueued(bytes)
	}
}

// ^ Error Callback to Plugin with error  ^
func (sn *Node) Error(err error, method string) {
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

	// Check and callback
	if sn.call != nil {
		// Reference
		sn.call.OnError(bytes)
	}

	// Log In Core
	log.Fatalln(fmt.Sprintf("[Error] At Method %s : %s", err.Error(), method))
}
