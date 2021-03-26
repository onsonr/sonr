package node

import (
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	msg "github.com/libp2p/go-msgio"
	msgio "github.com/libp2p/go-msgio"
	sf "github.com/sonr-io/core/internal/file"
	md "github.com/sonr-io/core/internal/models"
	dt "github.com/sonr-io/core/pkg/data"
	tr "github.com/sonr-io/core/pkg/transfer"
	"google.golang.org/protobuf/proto"
)

// ^ OnReceiveTransfer: Prepares for Incoming File Transfer when Accepted ^
func (n *Node) OnReceiveTransfer(inv *md.AuthInvite, fs *sf.FileSystem) {
	n.incoming = tr.CreateIncomingFile(inv, fs, n.call)
	n.host.SetStreamHandler(n.router.Transfer(), n.handleTransferIncoming)
}

// ^ OnReply: Begins File Transfer when Accepted ^
func (n *Node) OnReply(id peer.ID, p *md.Peer, cf *sf.FileItem, reply []byte) {
	// Call Responded
	n.call.Responded(reply)

	// AuthReply Message
	resp := md.AuthReply{}
	err := proto.Unmarshal(reply, &resp)
	if err != nil {
		n.call.Error(err, "handleReply")
	}

	// Check for File Transfer
	if resp.Decision && resp.Type == md.AuthReply_Transfer {
		// Create New Auth Stream
		stream, err := n.host.NewStream(n.ctx, id, n.router.Transfer())
		if err != nil {
			n.call.Error(err, "StartOutgoing")
		}

		// Initialize Writer
		writer := msgio.NewWriter(stream)

		// Start Routine
		hasCompleted := make(chan bool)
		go cf.WriteToStream(writer, p, hasCompleted)

		// Wait For Done
		done := <-hasCompleted
		if done {
			n.call.Transmitted(p)
		}
	}
}

// ^ handleTransferIncoming: Processes Incoming Data ^ //
func (n *Node) handleTransferIncoming(stream network.Stream) {
	// Route Data from Stream
	go func(reader msg.ReadCloser, t *tr.IncomingFile) {
		for i := 0; ; i++ {
			// @ Read Length Fixed Bytes
			buffer, err := reader.ReadMsg()
			if err != nil {
				n.call.Error(err, "HandleIncoming:ReadMsg")
				break
			}

			// @ Unmarshal Bytes into Proto
			hasCompleted, err := t.AddBuffer(i, buffer)
			if err != nil {
				n.call.Error(err, "HandleIncoming:AddBuffer")
				break
			}

			// @ Check if All Buffer Received to Save
			if hasCompleted {
				// Sync file
				if err := n.incoming.Save(); err != nil {
					n.call.Error(err, "HandleIncoming:Save")
				}
				n.host.RemoveStreamHandler(n.router.Transfer())
				break
			}
			dt.GetState().NeedsWait()
		}
	}(msg.NewReader(stream), n.incoming)
}

// // ^ HandleIncomingStream Writes to Current Incoming File ^ //
// func (fs *FileSystem) HandleIncomingStream(stream network.Stream) {
// 	// Get current incoming file
// 	inFile, err := fs.DequeueIn()
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	// Process Stream Events
// 	go func(reader msg.ReadCloser, f *FileItem) {
// 		for i := 0; ; i++ {
// 			// @ Read Length Fixed Bytes
// 			buffer, err := reader.ReadMsg()
// 			if err != nil {
// 				fs.Call.Error(err, "HandleIncoming:ReadMsg")
// 				break
// 			}

// 			// @ Unmarshal Bytes into Proto
// 			res, err := f.WriteFromStream(i, buffer)
// 			if err != nil {
// 				fs.Call.Error(err, "HandleIncoming:AddBuffer")
// 				break
// 			}

// 			// @ Callback with Progress
// 			if res.MetInterval {
// 				fs.Call.Progressed(res.Progress)
// 			}

// 			// @ Check if All Buffer Received to Save
// 			if res.HasCompleted {
// 				// Save File
// 				if err := fs.SaveFile(f); err != nil {
// 					fs.Call.Error(err, "HandleIncoming:Save")
// 				}
// 				break
// 			}
// 			dt.GetState().NeedsWait()
// 		}
// 	}(msg.NewReader(stream), inFile)
// }
