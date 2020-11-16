package sonr

import (
	"bufio"
	"fmt"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-msgio"
)

// ^ Auth Stream Struct ^ //
type AuthStreamConn struct {
	rw     msgio.ReadWriter
	stream network.Stream
}

// ^ Handle Incoming Stream ^ //
func (sn *Node) HandleAuthStream(stream network.Stream) {
	fmt.Println("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	buffrw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	mrw := msgio.NewReadWriter(buffrw)

	// Create/Set Auth Stream
	asc := &AuthStreamConn{
		rw:     mrw,
		stream: stream,
	}
	sn.AuthStream = *asc

	// Initialize Routine
	go asc.Read()
}

// ^ Create New Stream ^ //
func (sn *Node) NewAuthStream(stream network.Stream) {
	fmt.Println("Creating New Stream")

	// Create new Buffer
	buffrw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	mrw := msgio.NewReadWriter(buffrw)

	// Create/Set Auth Stream
	asc := &AuthStreamConn{
		rw:     mrw,
		stream: stream,
	}
	sn.AuthStream = *asc

	// Initialize Routine
	go asc.Read()
}

// ^ Read Data from Msgio ^ //
func (asc *AuthStreamConn) Read() {
	for {
		msg, err := asc.rw.ReadMsg()
		if err != nil {
			fmt.Println("Error: ", err)
		}

		fmt.Println("Received: ", string(msg))
	}
}

// ^ Message on Stream ^ //
func (asc *AuthStreamConn) Write(text string) {
	err := asc.rw.WriteMsg([]byte(text))
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Sent: ", text)
	}
}
