package sonr

import (
	"bufio"
	"fmt"

	"github.com/libp2p/go-libp2p-core/network"
)

// ^ Auth Stream Struct ^ //
type AuthStreamConn struct {
	readWriter *bufio.ReadWriter
	stream     network.Stream
	callback   Callback
}

// ^ Handle Incoming Stream ^ //
func (sn *Node) HandleAuthStream(stream network.Stream) {
	fmt.Println("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	buffrw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	// Create/Set Auth Stream
	sn.AuthStream = AuthStreamConn{
		readWriter: buffrw,
		stream:     stream,
		callback:   sn.Callback,
	}
	// Initialize Routine
	go sn.AuthStream.Read()
}

// ^ Create New Stream ^ //
func (sn *Node) NewAuthStream(stream network.Stream) {
	fmt.Println("Creating New Stream")

	// Create new Buffer
	buffrw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	// Create/Set Auth Stream
	sn.AuthStream = AuthStreamConn{
		readWriter: buffrw,
		stream:     stream,
		callback:   sn.Callback,
	}
	// Initialize Routine
	go sn.AuthStream.Read()
}

// ^ Read Data from Msgio ^ //
func (asc *AuthStreamConn) Read() {
	for {
		// Read the Buffer
		str, err := asc.readWriter.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from buffer")
			panic(err)
		}

		// Empty String
		if str == "" {
			return
		}

		// Contains Data
		if str != "\n" {
			//fmt.Println("Received Message: ", str)
			// Callback the Message
			asc.callback.OnInvited(str)
		}
	}
}

// ^ Message on Stream ^ //
func (asc *AuthStreamConn) Send(text string) error {
	// Write Message with "Delimiter"=(Seperator for Message Values)
	_, err := asc.readWriter.WriteString(fmt.Sprintf("%s\n", text))
	if err != nil {
		fmt.Println("Error writing to buffer")
		return err
	}

	// Write buffered data
	err = asc.readWriter.Flush()
	if err != nil {
		fmt.Println("Error flushing buffer")
		return err
	}
	return nil
}
