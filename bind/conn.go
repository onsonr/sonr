package sonr

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-msgio"
)

func handleStream(stream network.Stream) {
	fmt.Println("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	mrw := msgio.NewReadWriter(rw)

	for {
		msg, err := mrw.ReadMsg()
		if err != nil {
			panic(err)
		}

		// echo it back :)
		err = mrw.WriteMsg(msg)
		if err != nil {
			panic(err)
		}
		// 'stream' will stay open until you close it (or the other side closes it).
	}
}

func readData(rw *bufio.ReadWriter) {
	for {
		str, err := rw.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from buffer")
			panic(err)
		}

		if str == "" {
			return
		}
		if str != "\n" {
			// Green console colour: 	\x1b[32m
			// Reset console colour: 	\x1b[0m
			fmt.Printf("\x1b[32m%s\x1b[0m> ", str)
		}

	}
}

func writeData(rw *bufio.ReadWriter) {
	stdReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from stdin")
			panic(err)
		}

		_, err = rw.WriteString(fmt.Sprintf("%s\n", sendData))
		if err != nil {
			fmt.Println("Error writing to buffer")
			panic(err)
		}
		err = rw.Flush()
		if err != nil {
			fmt.Println("Error flushing buffer")
			panic(err)
		}
	}
}

// SetStreamHandler sets the protocol handler on the Host's Mux.
func (sn *Node) setStreamHandler(pid protocol.ID, handler network.StreamHandler) {
	sn.Host.SetStreamHandler(pid, handler)
}

// SetStreamHandlerMatch sets the protocol handler on the Host's Mux
// using a matching function for protocol selection.
func (sn *Node) setStreamHandlerMatch(pid protocol.ID, m func(string) bool, h network.StreamHandler) {
	sn.Host.SetStreamHandlerMatch(pid, m, h)
}

// RemoveStreamHandler removes a handler on the mux that was set by
// SetStreamHandler
func (sn *Node) removeStreamHandler(pid protocol.ID) {
	sn.Host.RemoveStreamHandler(pid)
}

// NewStream opens a new stream to given peer p, and writes a p2p/protocol
func (sn *Node) newStream(ctx context.Context, p peer.ID, pids ...protocol.ID) (network.Stream, error) {
	return sn.Host.NewStream(ctx, p, pids...)
}
