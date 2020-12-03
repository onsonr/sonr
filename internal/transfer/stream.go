package transfer

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"

	"github.com/libp2p/go-libp2p-core/network"
	msgio "github.com/libp2p/go-msgio"
	sf "github.com/sonr-io/core/internal/file"
	pb "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ^ Handle Incoming Stream ^ //
func (dsc *PeerConnection) HandleTransfer(stream network.Stream) {
	// Set Stream
	dsc.stream = stream

	// Print Stream Info
	info := stream.Stat()
	fmt.Println("Stream Info: ", info)

	// Initialize Routine
	reader := msgio.NewReader(dsc.stream)
	go dsc.readTransferStream(reader)
}

// ^ read Data from Msgio ^ //
func (dsc *PeerConnection) readTransferStream(reader msgio.ReadCloser) {
	for i := 0; ; i++ {
		// @ Read Length Fixed Bytes
		buffer, err := reader.ReadMsg()
		if err != nil {
			dsc.Call.Error(err, "ReadMsg")
			log.Fatalln(err)
			break
		}

		// @ Unmarshal Bytes into Proto
		chunk := pb.Chunk{}
		err = proto.Unmarshal(buffer, &chunk)
		if err != nil {
			dsc.Call.Error(err, "readTransferStream")
			log.Fatalln(err)
		}
	}
}

// ^ write file to Msgio ^ //
func (dsc *PeerConnection) writeTransferStream(writer msgio.WriteCloser, safeMeta *sf.SafeFile) {
	// Get Data
	metadata := safeMeta.Metadata()
	imgBuffer := new(bytes.Buffer)

	// Check Type for image
	if metadata.Mime.Type == pb.MIME_image {
		// New File for ThumbNail
		file, err := os.Open(metadata.Path)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()

		// Convert to Image Object
		img, _, err := image.Decode(file)
		if err != nil {
			log.Fatalln(err)
		}

		// Encode as Jpeg into buffer
		err = jpeg.Encode(imgBuffer, img, &jpeg.Options{Quality: 100})
		if err != nil {
			log.Fatalln(err)
		}
	}
}
