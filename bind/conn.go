package sonr

import (
	"bufio"
	"fmt"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-msgio"
)

// ^ Auth Stream Struct ^ //
type AuthStreamConn struct {
	channel msgio.Chan
	rw      msgio.ReadWriter
	stream  network.Stream
}

// ^ Handle Incoming Stream ^ //
func (sn *Node) HandleAuthStream(stream network.Stream) {
	fmt.Println("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	buffrw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	mrw := msgio.NewReadWriter(buffrw)
	chn := msgio.NewChan(128)

	// Create/Set Auth Stream
	asc := &AuthStreamConn{
		rw:      mrw,
		stream:  stream,
		channel: *chn,
	}
	sn.AuthStream = *asc
	asc.Write("Third Message")

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
	asc.Write("First Message")

	// Initialize Routine
	go asc.Read()
}

// ^ Read Data from Msgio ^ //
func (asc *AuthStreamConn) Read() {
	for {
		select {
		case msg := <-asc.channel.MsgChan:
			fmt.Println("Received: ", string(msg))
			asc.Write("This is a Reply")
		}
	}
}

// ^ Message on Stream ^ //
func (asc *AuthStreamConn) Write(text string) {
	asc.channel.MsgChan <- []byte(text)
}
