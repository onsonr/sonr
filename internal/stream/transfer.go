package stream


// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/libp2p/go-libp2p-core/host"
// 	"github.com/libp2p/go-libp2p-core/peer"
// 	"github.com/libp2p/go-libp2p-core/protocol"
// 	pb "github.com/sonr-io/core/internal/models"
// 	"google.golang.org/protobuf/proto"

// 	gorpc "github.com/libp2p/go-libp2p-gorpc"
// )

// const transferProtocolID = protocol.ID("/sonr/data/Transfer")

// // ** Service Types ** //
// type TransferArgs struct {
// 	Current  int
// 	Total    int
// 	Transfer float32
// }

// type TransferReply struct {
// 	HasReceived bool
// }

// type TransferService struct{}

// func (t *TransferService) Transfer(ctx context.Context, argType TransferArgs, replyType *TransferReply) error {
// 	log.Println("Received a Transfer call: ", argType.Transfer)
// 	replyType.HasReceived = true
// 	// Send Callback
// 	return nil
// }

// // ** DataStream Handling Types ** //
// type TransferClient struct {
// 	Transfered OnTransfered
// 	rpcClient  *gorpc.Client
// 	hostID     peer.ID
// }

// type TransferServer struct {
// 	Transfered OnTransfered
// 	rpcServer  *gorpc.Server
// }

// // ^ Set Sender as Server ^ //
// func setSender(host host.Host, op OnTransfered) TransferServer {
// 	log.Println("Setting Data Sender as RPC Server")
// 	rpcHost := gorpc.NewServer(host, protocolID)
// 	svc := TransferService{}
// 	err := rpcHost.Register(&svc)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return TransferServer{
// 		rpcServer:  rpcHost,
// 		Transfered: op,
// 	}
// }

// // ^ Set Receiver as Client ^ //
// func setReceiver(client host.Host, peerID peer.ID, op OnTransfered) TransferClient {
// 	log.Println("Setting Data Receiver as RPC Server")
// 	rpc := gorpc.NewClient(client, protocolID)
// 	return TransferClient{
// 		rpcClient:  rpc,
// 		Transfered: op,
// 		hostID:     peerID,
// 	}
// }

// // ^ Send Transfer to Data Sender ^ //
// func (pc *TransferClient) SendTransfer(current int32, total int32) {
// 	// Send Callback
// 	callbackProtobuf(pc.Transfered, current, total)

// 	var reply TransferReply
// 	var args TransferArgs
// 	p := float32(current) / float32(total)

// 	args.Transfer = p
// 	args.Current = int(current)
// 	args.Total = int(total)

// 	startTime := time.Now()
// 	err := pc.rpcClient.Call(pc.hostID, "TransferService", "Transfer", args, &reply)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if !(reply.HasReceived) {
// 		panic("Server did not receive Transfer")
// 	}
// 	endTime := time.Now()
// 	diff := endTime.Sub(startTime)
// 	fmt.Printf("%f Transfer from %s: time=%s\n", p, pc.hostID.String(), diff)
// }

// // ^ Callback to Frontend with Protobuf ^ //
// func callbackProtobuf(op OnTransfered, current int32, total int32) {
// 	// Create Callback for User
// 	TransferMessage := pb.TransferUpdate{
// 		Current: current,
// 		Total:   total,
// 		Percent: float32(current) / float32(total),
// 	}

// 	// Convert to bytes
// 	bytes, err := proto.Marshal(&TransferMessage)
// 	if err != nil {
// 		fmt.Println(err, "SendTransfer")
// 	}

// 	// Send Callback
// 	op(bytes)
// }
