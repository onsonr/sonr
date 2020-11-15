package sonr

import (
	"bufio"
	"fmt"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-msgio"
)

type AuthStream struct {
	ReadWriter msgio.ReadWriter
}

type DataStream struct {
}

func handleStream(stream network.Stream) {
	fmt.Println("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	mrw := msgio.NewReadWriter(rw)

	// Send Initial Message
	m := []byte("Hello from stream")
	mrw.WriteMsg(m)

	// Loop Messages
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
	}
}
