package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/h2non/filetype"
	"github.com/libp2p/go-libp2p-core/peer"
	"google.golang.org/protobuf/proto"
)

// ***************************** //
// ** ConnectionRequest Mgnmt ** //
// ***************************** //

func (req *ConnectionRequest) AttachGeoToRequest(geo *GeoIP) *ConnectionRequest {
	req.Latitude = geo.Geo.Latitude
	req.Longitude = geo.Geo.Longitude
	return req
}

// **************************** //
// ** Remote Info Management ** //
// **************************** //

// ^ Get Remote Point Info ^
func GetRemoteInfo(list []string) RemoteInfo {
	return RemoteInfo{
		Display: fmt.Sprintf("%s %s %s", list[0], list[1], list[2]),
		Topic:   fmt.Sprintf("%s-%s-%s", list[0], list[1], list[2]),
		Count:   int32(len(list)),
		IsJoin:  false,
		Words:   list,
	}
}

// *********************************** //
// ** Outgoing File Info Management ** //
// *********************************** //
// ^ Struct returned on GetInfo() Generate Preview/Metadata
type OutFile struct {
	Mime    *MIME
	Payload Payload
	Name    string
	Path    string
	Size    int32
	IsImage bool
}

// ^ Method Returns File Info at Path ^ //
func GetFileInfo(path string) (*OutFile, error) {
	// Initialize
	var mime *MIME
	var payload Payload

	// @ 1. Get File Information
	// Open File at Path
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Get Info
	info, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// Read File to required bytes
	head := make([]byte, 261)
	_, err = file.Read(head)
	if err != nil {
		return nil, err
	}

	// Get File Type
	kind, err := filetype.Match(head)
	if err != nil {
		return nil, err
	}

	// @ 2. Create Mime Protobuf
	mime = &MIME{
		Type:    MIME_Type(MIME_Type_value[kind.MIME.Type]),
		Subtype: kind.MIME.Subtype,
		Value:   kind.MIME.Value,
	}

	// @ 3. Find Payload
	if mime.Type == MIME_image || mime.Type == MIME_video || mime.Type == MIME_audio {
		payload = Payload_MEDIA
	} else {
		// Get Extension
		ext := filepath.Ext(path)

		// Cross Check Extension
		if ext == ".pdf" {
			payload = Payload_PDF
		} else if ext == ".ppt" || ext == ".pptx" {
			payload = Payload_PRESENTATION
		} else if ext == ".xls" || ext == ".xlsm" || ext == ".xlsx" || ext == ".csv" {
			payload = Payload_SPREADSHEET
		} else if ext == ".txt" || ext == ".doc" || ext == ".docx" || ext == ".ttf" {
			payload = Payload_TEXT
		} else {
			payload = Payload_UNDEFINED
		}
	}

	// Return Object
	return &OutFile{
		Mime:    mime,
		Payload: payload,
		Name:    strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)),
		Path:    path,
		Size:    int32(info.Size()),
		IsImage: filetype.IsImage(head),
	}, nil
}

// *********************************** //
// ** Incoming File Info Management ** //
// *********************************** //
type InFile struct {
	Payload       Payload
	Properties    *TransferCard_Properties
	ChunkBaseChan chan Chunk64
	ChunkBufChan  chan ChunkBuffer
}

// ********************** //
// ** Lobby Management ** //
// ********************** //
// ^ Returns as Lobby Buffer ^
func (l *Lobby) Buffer() ([]byte, error) {
	bytes, err := proto.Marshal(l)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return bytes, nil
}

// ^ Add/Update Peer in Lobby ^
func (l *Lobby) Add(peer *Peer) {
	// Update Peer with new data
	l.Peers[peer.Id.Peer] = peer
	l.Count = int32(len(l.Peers))
	l.Size = int32(len(l.Peers)) + 1 // Account for User
}

// ^ Remove Peer from Lobby ^
func (l *Lobby) Delete(id peer.ID) {
	// Update Peer with new data
	delete(l.Peers, id.String())
	l.Count = int32(len(l.Peers))
	l.Size = int32(len(l.Peers)) + 1 // Account for User
}

// ^ Sync Between Remote Peers Lobby ^
func (l *Lobby) Sync(ref *Lobby, remotePeer *Peer) {
	// Validate Lobbies are Different
	if l.Count != ref.Count {
		// Iterate Over List
		for id, peer := range ref.Peers {
			if l.User.IsNotPeerIDString(id) {
				l.Add(peer)
			}
		}
	}
	l.Add(remotePeer)
}
