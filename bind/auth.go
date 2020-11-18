package sonr

import (
	"bufio"
	"encoding/json"
	"fmt"

	"github.com/libp2p/go-libp2p-core/network"
)

// ^ authStreamMessage is for Auth Stream Request ^
type authStreamMessage struct {
	subject  string
	decision bool
	peerInfo string
	metadata string
}

// ^ Auth Stream Struct ^ //
type authStreamConn struct {
	readWriter *bufio.ReadWriter
	stream     network.Stream
	callback   Callback
}

// ^ Handle Incoming Stream ^ //
func (sn *Node) HandleAuthStream(stream network.Stream) {
	// Create a buffer stream for non blocking read and write.
	buffrw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	// Create/Set Auth Stream
	sn.AuthStream = authStreamConn{
		readWriter: buffrw,
		stream:     stream,
		callback:   sn.Callback,
	}
	// Initialize Routine
	go sn.AuthStream.Read()
}

// ^ Create New Stream ^ //
func (sn *Node) NewAuthStream(stream network.Stream) {
	// Create new Buffer
	buffrw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	// Create/Set Auth Stream
	sn.AuthStream = authStreamConn{
		readWriter: buffrw,
		stream:     stream,
		callback:   sn.Callback,
	}
	// Initialize Routine
	go sn.AuthStream.Read()
}

// ^ Write Message on Stream ^ //
func (asc *authStreamConn) Write(authMsg authStreamMessage) error {
	// Check Message
	fmt.Println("Auth Message being passed: ", authMsg)

	// Convert Request to JSON String
	bytes, err := json.Marshal(authMsg)
	if err != nil {
		println("Error Converting Meta to JSON", err)
		return err
	}
	msg := string(bytes)
	fmt.Println("Auth Request: ", msg)

	// Write Message with "Delimiter"=(Seperator for Message Values)
	_, err = asc.readWriter.WriteString(fmt.Sprintf("%s\n", msg))
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

// ^ Read Data from Msgio ^ //
func (asc *authStreamConn) Read() {
	for {
		// ** Read the Buffer **
		str, err := asc.readWriter.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from buffer")
			panic(err)
		}

		// ** Empty String **
		if str == "" {
			return
		}

		// ** Contains Data **
		if str != "\n" {
			// Construct message
			asm := new(authStreamMessage)
			err := json.Unmarshal([]byte(str), asm)
			if err != nil {
				fmt.Println("Error Unmarshalling Auth Stream Message")
			}

			// Check Message Subject
			switch asm.subject {
			// @ Request to Invite
			case "Request":
				fmt.Println("Auth Invited: ", str)
				// Callback the Invitation
				asc.callback.OnInvited(asm.peerInfo, asm.metadata)

			// @ Response to Invite
			case "Response":
				// Check peer decision
				if asm.decision {
					fmt.Println("Auth Accepted: ", str)
					// User Accepted
					asc.callback.OnAccepted("Great")
				} else {
					fmt.Println("Auth Declined: ", str)
					// User Declined
					asc.callback.OnDenied("Unlucky")
				}

			// ! Invalid Subject
			default:
				fmt.Printf("%s.\n", asm.subject)
			}
		}
	}
}
