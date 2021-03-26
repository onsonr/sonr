package file

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"strings"
	"sync"

	md "github.com/sonr-io/core/internal/models"
)

const K_BUF_CHUNK = 32000
const K_B64_CHUNK = 31998 // Adjusted for Base64 -- has to be divisible by 3

// @ File that safely sets metadata and thumbnail in routine
type FileItem struct {
	mutex sync.Mutex

	// References
	Payload md.Payload
	Owner   *md.Peer
	Name    string
	Path    string
	Ext     string

	// Outgoing Properties
	hasPreview bool
	preview    []byte
	request    *md.InviteRequest

	// Incoming Properties
	bytesBuilder   *bytes.Buffer
	inInfo         *md.InFileInfo
	invite         *md.AuthInvite
	stringsBuilder *strings.Builder
}

// ^ NewOutgoingFileItem Processes Outgoing File ^ //
func NewOutgoingFileItem(req *md.InviteRequest, p *md.Peer) (*FileItem, error) {
	// Check Values
	if req == nil || p == nil {
		return nil, errors.New("Request or Profile not Provided")
	}

	// Get File Information
	file := req.Files[len(req.Files)-1]

	// Check Thumbnail
	if len(file.Thumbnail) > 0 {
		// Initialize
		thumbWriter := new(bytes.Buffer)
		thumbReader := bytes.NewReader(file.Thumbnail)

		// Convert to Image Object
		img, _, err := image.Decode(thumbReader)
		if err != nil {
			return nil, err
		}

		// @ Encode as Jpeg into buffer w/o scaling
		err = jpeg.Encode(thumbWriter, img, nil)
		if err != nil {
			return nil, err
		}

		// @ 1a. Get File Info
		preview := thumbWriter.Bytes()
		info, err := md.GetOutFileInfoWithPreview(file.Path, preview)
		if err != nil {
			return nil, err
		}

		// @ 3a. Callback with Preview
		return &FileItem{
			hasPreview: true,
			preview:    preview,
			Name:       info.Name,
			Path:       file.Path,
			Ext:        info.Ext(),
			Owner:      p,
			request:    req,
		}, nil
	} else {
		// @ 1b. Get File Info
		info, err := md.GetOutFileInfo(file.Path)
		if err != nil {
			return nil, err
		}

		// @ 3b. Callback with Preview
		return &FileItem{
			hasPreview: false,
			Path:       file.Path,
			Name:       info.Name,
			Ext:        info.Ext(),
			Owner:      p,
			request:    req,
		}, nil
	}
}

// ^ NewIncomingFileItem Prepares for Incoming Data ^ //
func NewIncomingFileItem(i *md.AuthInvite, p string) (*FileItem, error) {
	// Calculate Tracking Data
	totalChunks := int(i.Card.Properties.Size) / K_B64_CHUNK
	interval := totalChunks / 100

	// Get Info
	info := md.GetInFileInfo(i, interval)
	fileName := i.Card.Properties.Name + "." + i.Card.Properties.Mime.Subtype

	// Return Item
	return &FileItem{
		// Inherited Properties
		Owner:   i.From,
		Payload: i.Payload,
		Name:    fileName,
		Path:    p,
		inInfo:  info,
		invite:  i,

		// Builders
		stringsBuilder: new(strings.Builder),
		bytesBuilder:   new(bytes.Buffer),
	}, nil
}

// ^ Display Outgoing File Information ^ //
func (pf *FileItem) InfoOut() (*md.OutFileInfo, error) {
	if pf.hasPreview {
		return md.GetOutFileInfoWithPreview(pf.Path, pf.preview)
	} else {
		return md.GetOutFileInfo(pf.Path)
	}

}

// ^ Display Incoming File Information ^ //
func (pf *FileItem) InfoIn() (md.InFileInfo, error) {
	if pf.inInfo != nil {
		return *pf.inInfo, nil
	}
	return md.InFileInfo{}, errors.New("No Incoming Info")
}
