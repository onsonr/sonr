package sonr

import (
	"bufio"
	"encoding/json"
	"fmt"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/sonr-io/core/pkg/file"
	"github.com/sonr-io/core/pkg/lobby"
)

// ^ AuthRequestMessage is for Auth Stream Request ^
type AuthStreamMessage struct {
	subject  string
	decision bool
	peerInfo lobby.Peer
	metadata file.Metadata
}

// ^ Auth Stream Struct ^ //
type AuthStreamConn struct {
	readWriter *bufio.ReadWriter
	stream     network.Stream
	callback   Callback
}

// ^ Handle Incoming Stream ^ //
func (sn *Node) HandleAuthStream(stream network.Stream) {
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

// ^ Write Message on Stream ^ //
func (asc *AuthStreamConn) Write(authMsg AuthStreamMessage) error {
	// Convert Request to JSON String
	msgBytes, err := json.Marshal(authMsg)
	if err != nil {
		println("Error Converting Meta to JSON", err)
		return err
	}

	// Write Message with "Delimiter"=(Seperator for Message Values)
	_, err = asc.readWriter.WriteString(fmt.Sprintf("%s\n", string(msgBytes)))
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
func (asc *AuthStreamConn) Read() {
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
			asm := new(AuthStreamMessage)
			err := json.Unmarshal([]byte(str), asm)
			if err != nil {
				fmt.Println("Error Unmarshalling Auth Stream Message")
			}

			// Check Message Subject
			switch asm.subject {
			// @ Request to Invite
			case "Request":
				// Callback the Invitation
				asc.callback.OnInvited(str)

			// @ Response to Invite
			case "Response":
				// Check peer decision
				if asm.decision {
					// User Accepted
					asc.callback.OnAccepted("Greate")
				} else {
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
