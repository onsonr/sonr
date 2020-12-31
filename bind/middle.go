package sonr

import (
	"log"

	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ^ complete is middleware method to handle post transfer
func (sn *Node) complete(isReceiver bool, data []byte) {
	// Send Callback
	if isReceiver {
		sn.call.OnReceived(data)
	} else {
		sn.call.OnTransmitted(data)
	}

	// Announce Available
	err := sn.lobby.Update()
	if err != nil {
		log.Println(err)
	}
}

// ^ error is middleware method with error instance, and method ^
func (sn *Node) error(err error, method string) {
	// Create Error ProtoBuf
	errorMsg := md.ErrorMessage{
		Message: err.Error(),
		Method:  method,
	}

	// Convert Message to bytes
	bytes, err := proto.Marshal(&errorMsg)
	if err != nil {
		log.Println("Cannot Marshal Error Protobuf: ", err)
	}
	// Send Callback
	sn.call.OnError(bytes)

	// Log In Core
	log.Fatalf("[Error] At Method %s : %s", err.Error(), method)
}
