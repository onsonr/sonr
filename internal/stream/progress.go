package stream

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"time"

	"log"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"

	gorpc "github.com/libp2p/go-libp2p-gorpc"

	multiaddr "github.com/multiformats/go-multiaddr"
)

type PingArgs struct {
	Data []byte
}
type PingReply struct {
	Data []byte
}
type PingService struct{}

func (t *PingService) Ping(ctx context.Context, argType PingArgs, replyType *PingReply) error {
	log.Println("Received a Ping call")
	replyType.Data = argType.Data
	return nil
}

var protocolID = protocol.ID("/p2p/rpc/ping")

func startServer(host host.Host) {
	log.Println("Launching host")
	log.Printf("Hello World, my hosts ID is %s\n", host.ID().Pretty())
	for _, addr := range host.Addrs() {
		ipfsAddr, err := multiaddr.NewMultiaddr("/ipfs/" + host.ID().Pretty())
		if err != nil {
			panic(err)
		}
		peerAddr := addr.Encapsulate(ipfsAddr)
		log.Printf("I'm listening on %s\n", peerAddr)
	}

	rpcHost := gorpc.NewServer(host, protocolID)

	svc := PingService{}
	err := rpcHost.Register(&svc)
	if err != nil {
		panic(err)
	}

	fmt.Println("Done")

	for {
		time.Sleep(time.Second * 1)
	}
}

func startClient(client host.Host, hostID peer.ID, pingCount, randomDataSize int) {
	fmt.Println("Launching client")
	rpcClient := gorpc.NewClient(client, protocolID)
	numCalls := 0
	durations := []time.Duration{}
	betweenPingsSleep := time.Second * 1

	for numCalls < pingCount {
		var reply PingReply
		var args PingArgs

		c := randomDataSize
		b := make([]byte, c)
		_, err := rand.Read(b)
		if err != nil {
			panic(err)
		}

		args.Data = b

		time.Sleep(betweenPingsSleep)
		startTime := time.Now()
		err = rpcClient.Call(hostID, "PingService", "Ping", args, &reply)
		if err != nil {
			panic(err)
		}
		if !bytes.Equal(reply.Data, b) {
			panic("Received wrong amount of bytes back!")
		}
		endTime := time.Now()
		diff := endTime.Sub(startTime)
		fmt.Printf("%d bytes from %s: seq=%d time=%s\n", c, hostID.String(), numCalls+1, diff)
		numCalls += 1
		durations = append(durations, diff)
	}

	totalDuration := int64(0)
	for _, dur := range durations {
		totalDuration = totalDuration + dur.Nanoseconds()
	}
	averageDuration := totalDuration / int64(len(durations))
	fmt.Printf("Average duration for ping reply: %s\n", time.Duration(averageDuration))

}
