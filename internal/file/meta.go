package file

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/h2non/filetype"
	md "github.com/sonr-io/core/internal/models"
)

func GetMetadata(path string, owner *md.Peer) *md.Metadata {
	// @ 1. Get File Information
	// Open File at Path
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// Get Info
	info, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	// Read File to required bytes
	head := make([]byte, 261)
	_, err = file.Read(head)
	if err != nil {
		log.Fatalln(err)
	}

	// Get File Type
	kind, err := filetype.Match(head)
	if err != nil {
		log.Fatalln(err)
	}

	// @ 2. Set Metadata Protobuf Values
	// Set Metadata
	base := filepath.Base(path)
	return &md.Metadata{
		Name: strings.TrimSuffix(base, filepath.Ext(path)),
		Path: path,
		Size: int32(info.Size()),
		Mime: &md.MIME{
			Type:    md.MIME_Type(md.MIME_Type_value[kind.MIME.Type]),
			Subtype: kind.MIME.Subtype,
			Value:   kind.MIME.Value,
		},
		Received: int32(time.Now().Unix()),
		Owner:    owner,
	}
}
