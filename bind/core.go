package sonr

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p-core/protocol"
	sh "github.com/sonr-io/core/pkg/host"
	pb "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// Callback returns updates from p2p
type Callback interface {
	OnRefreshed(data []byte)
	OnInvited(data []byte)
	OnResponded(data []byte)
	OnProgress(data []byte)
	OnError(data []byte)
}

// ^ Start begins the mobile host ^
func Start(data []byte, call *Callback) *Node {
	// ** Create Context and Node - Begin Setup **
	ctx := context.Background()
	node := new(Node)
	node.Callback = call
	node.files = make([]pb.Metadata, maxFileBufferSize)

	// @I. Unmarshal Connection Event
	connEvent := pb.RequestMessage{}
	err := proto.Unmarshal(data, &connEvent)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// @1. Create Host
	node.host, err = sh.NewHost(ctx)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// @2. Set Stream Handlers
	node.host.SetStreamHandler(protocol.ID("/sonr/auth"), node.HandleAuthStream)
	node.host.SetStreamHandler(protocol.ID("/sonr/transfer"), node.HandleTransferStream)

	// @3. Set Node User Information
	err = node.setUser(&connEvent)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// @4. Setup Discovery w/ Lobby
	err = node.setDiscovery(ctx, &connEvent)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// ** Callback Node User Information ** //
	return node
}

// ^ Exit Ends Communication ^
func (sn *Node) Exit() {
	sn.Lobby.End()
	sn.host.Close()
}
