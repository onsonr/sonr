package sonr

import (
	"log"

	"github.com/skip2/go-qrcode"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/encoding/protojson"
)

// ^ Link with a QR Code ^ //
func (sn *Node) LinkDevice(peerString string) error {
	// Convert String to Bytes
	b := []byte(peerString)
	peer := md.Peer{}

	// Convert to Peer Protobuf
	err := protojson.Unmarshal(b, &peer)
	if err != nil {
		return err
	}

	// TODO: Save Peer to Disk

	return nil
}

// ^ Display QR Code of Peer Info ^ //
func (sn *Node) DisplayCode() []byte {
	// Get Node JSON
	jsonBytes, err := protojson.Marshal(sn.peer)
	if err != nil {
		log.Println(err)
		return nil
	}
	json := string(jsonBytes)
	print(json)

	// Encode to QR
	png, err := qrcode.Encode(json, qrcode.Medium, 256)
	if err != nil {
		log.Println(err)
		return nil
	}
	return png
}
