package sonr

import (
	"bufio"
	"fmt"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-msgio"
)

type AuthStreamConn struct {
	rw     msgio.ReadWriter
	stream network.Stream
}

func (sn *Node) handleStream(stream network.Stream) {
	fmt.Println("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	buffrw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	mrw := msgio.NewReadWriter(buffrw)

	// Create AuthStream Type
	asc := &AuthStreamConn{
		rw:     mrw,
		stream: stream,
	}

	// Initialize Routine
	go sn.readData()

	// Set Auth Stream
	sn.AuthStream = *asc
}

func (sn *Node) readData() {
	for {
		msg, err := sn.AuthStream.rw.ReadMsg()
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		println("Received: ", string(msg))
	}
}

func (sn *Node) NewAuthStream(stream network.Stream) {
	fmt.Println("Creating New Stream")

	// Create new Buffer
	buffrw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	mrw := msgio.NewReadWriter(buffrw)

	// Create AuthStream Type
	asc := &AuthStreamConn{
		rw:     mrw,
		stream: stream,
	}

	// Initialize Routine
	go sn.readData()

	// Set Auth Stream
	sn.AuthStream = *asc
}

func (sn *Node) AuthStreamSend(text string) {
	err := sn.AuthStream.rw.WriteMsg([]byte(text))
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		println("Sent: ", text)
	}
}
