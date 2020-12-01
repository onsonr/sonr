package file

import (
	"os"
	"time"

	"fmt"
	_ "image/gif"
	_ "image/jpeg"

	"github.com/google/uuid"
	"github.com/h2non/filetype"
	pb "github.com/sonr-io/core/internal/models"
)

// ^ Creates Metadata for File at Path ^ //
func GetMetadata(path string, owner *pb.Peer) (*pb.Metadata, error) {
	// @ 1. Get File Information
	// Open File at Path
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening File", err)
		return nil, err
	}

	// Get Info
	info, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}

	// Get File Type
	head := make([]byte, 261)
	file.Read(head)
	kind, err := filetype.Match(head)
	if err != nil {
		fmt.Println(err)
	}
	file.Close()

	// Get Mime Type
	mime := &pb.MIME{
		Type:    kind.MIME.Type,
		Subtype: kind.MIME.Subtype,
		Value:   kind.MIME.Value,
	}

	// @ 3. Set Metadata Protobuf Values
	return &pb.Metadata{
		Uuid:       uuid.New().String(),
		Name:       fileNameWithoutExtension(path),
		Path:       path,
		Size:       int32(info.Size()),
		Chunks:     int32(info.Size()) / BlockSize,
		Mime:       mime,
		Owner:      owner,
		LastOpened: int32(time.Now().Unix()),
	}, nil
}
