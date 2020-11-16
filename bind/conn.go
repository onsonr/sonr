package sonr

import (
	"bufio"
	"encoding/binary"
	"fmt"

	"math/rand"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-msgio"
)

// ^ Auth Stream Struct ^ //
type AuthStreamConn struct {
	readWriter *bufio.ReadWriter
	writer     msgio.Writer
	reader     msgio.Reader
	stream     network.Stream
}

// ^ Handle Incoming Stream ^ //
func (sn *Node) HandleAuthStream(stream network.Stream) {
	fmt.Println("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	buffrw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	mwtr := msgio.NewWriter(buffrw)
	mrdr := msgio.NewReader(buffrw)

	// Create/Set Auth Stream
	sn.AuthStream = AuthStreamConn{
		readWriter: buffrw,
		writer:     mwtr,
		reader:     mrdr,
		stream:     stream,
	}
	// Initialize Routine
	go sn.AuthStream.Read()

	sn.AuthStream.Send("Third Message")
}

// ^ Create New Stream ^ //
func (sn *Node) NewAuthStream(stream network.Stream) {
	fmt.Println("Creating New Stream")

	// Create new Buffer
	buffrw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	mwtr := msgio.NewWriter(buffrw)
	mrdr := msgio.NewReader(buffrw)

	// Create/Set Auth Stream
	sn.AuthStream = AuthStreamConn{
		readWriter: buffrw,
		writer:     mwtr,
		reader:     mrdr,
		stream:     stream,
	}
	// Initialize Routine
	go sn.AuthStream.Read()

	sn.AuthStream.Send("First Message")
}

// ^ Read Data from Msgio ^ //
func (asc *AuthStreamConn) Read() {
	readMsg(asc.readWriter)
}

func readMsg(rw *bufio.ReadWriter) {
	for {
		// read bytes until new line
		msg, err := rw.ReadBytes('\n')
		if err != nil {
			fmt.Println("Error reading from buffer")
			continue
		}

		// get the id
		id := int64(binary.LittleEndian.Uint64(msg[0:8]))

		// get the content, last index is len(msg)-1 to remove the new line char
		content := string(msg[8 : len(msg)-1])

		if content != "" {
			// we print [message ID] content
			fmt.Printf("[%d] %s", id, content)
		}
	}
}

// ^ Message on Stream ^ //
func (asc *AuthStreamConn) Send(text string) {
	content := []byte(text)
	sendMsg(asc.readWriter, rand.Int63(), content)
}

func sendMsg(rw *bufio.ReadWriter, id int64, content []byte) error {
	// allocate our slice of bytes with the correct size 4 + size of the message + 1
	msg := make([]byte, 4+len(content)+1)

	// write id
	binary.LittleEndian.PutUint64(msg, uint64(id))

	// add content to msg
	copy(msg[13:], content)

	// add new line at the end
	msg[len(msg)-1] = '\n'

	// write msg to stream
	_, err := rw.Write(msg)
	if err != nil {
		fmt.Println("Error writing to buffer")
		return err
	}
	err = rw.Flush()
	if err != nil {
		fmt.Println("Error flushing buffer")
		return err
	}
	return nil
}
