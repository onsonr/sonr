package stream

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	pb "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"

	gorpc "github.com/libp2p/go-libp2p-gorpc"
)

const protocolID = protocol.ID("/sonr/data/progress")

// ** Service Types ** //
type ProgressArgs struct {
	Current  int
	Total    int
	Progress float32
}
type ProgressReply struct {
	HasReceived bool
}

type ProgressService struct{}

func (t *ProgressService) Progress(ctx context.Context, argType ProgressArgs, replyType *ProgressReply) error {
	log.Println("Received a Progress call: ", argType.Progress)
	replyType.HasReceived = true
	// Send Callback
	return nil
}

// ** DataStream Handling Types ** //
type ProgressClient struct {
	progressed OnProgressed
	rpcClient  *gorpc.Client
	hostID     peer.ID
}

type ProgressServer struct {
	progressed OnProgressed
	rpcServer  *gorpc.Server
}

// ^ Set Sender as Server ^ //
func setSender(host host.Host, op OnProgressed) ProgressServer {
	log.Println("Setting Data Sender as RPC Server")
	rpcHost := gorpc.NewServer(host, protocolID)
	svc := ProgressService{}
	err := rpcHost.Register(&svc)
	if err != nil {
		panic(err)
	}
	return ProgressServer{
		rpcServer:  rpcHost,
		progressed: op,
	}
}

// ^ Set Receiver as Client ^ //
func setReceiver(client host.Host, peerID peer.ID, op OnProgressed) ProgressClient {
	log.Println("Setting Data Receiver as RPC Server")
	rpc := gorpc.NewClient(client, protocolID)
	return ProgressClient{
		rpcClient:  rpc,
		progressed: op,
		hostID:     peerID,
	}
}

// ^ Send Progress to Data Sender ^ //
func (pc *ProgressClient) SendProgress(current int32, total int32) {
	// Send Callback
	callbackProtobuf(pc.progressed, current, total)

	var reply ProgressReply
	var args ProgressArgs
	p := float32(current) / float32(total)

	args.Progress = p
	args.Current = int(current)
	args.Total = int(total)

	startTime := time.Now()
	err := pc.rpcClient.Call(pc.hostID, "ProgressService", "Progress", args, &reply)
	if err != nil {
		panic(err)
	}
	if !(reply.HasReceived) {
		panic("Server did not receive progress")
	}
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Printf("%f progress from %s: time=%s\n", p, pc.hostID.String(), diff)
}

// ^ Callback to Frontend with Protobuf ^ //
func callbackProtobuf(op OnProgressed, current int32, total int32) {
	// Create Callback for User
	progressMessage := pb.ProgressUpdate{
		Current: current,
		Total:   total,
		Percent: float32(current) / float32(total),
	}

	// Convert to bytes
	bytes, err := proto.Marshal(&progressMessage)
	if err != nil {
		fmt.Println(err, "SendProgress")
	}

	// Send Callback
	op(bytes)
}
