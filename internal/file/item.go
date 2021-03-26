package file

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
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
	Path    string

	// Outgoing Properties
	inInfo  *md.InFileInfo
	outInfo *md.OutFileInfo
	request *md.InviteRequest
}

// ^ NewOutgoingFileItem Processes Outgoing File ^ //
func NewOutgoingFileItem(req *md.InviteRequest, p *md.Peer, hc chan bool) (*FileItem, error) {
	// Check Values
	if req == nil || p == nil {
		hc <- false
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
			hc <- false
			return nil, err
		}

		// @ 1a. Get File Info
		preview := thumbWriter.Bytes()
		info, err := md.GetOutFileInfoWithPreview(file.Path, preview)
		if err != nil {
			hc <- false
			return nil, err
		}

		// @ 2a. Create new SafeFile
		sm := &FileItem{
			Path:    file.Path,
			outInfo: info,
			Owner:   p,
			request: req,
		}

		// @ 3a. Callback with Preview
		hc <- true
		return sm, nil
	} else {
		// @ 1b. Get File Info
		info, err := md.GetOutFileInfo(file.Path)
		if err != nil {
			hc <- false
			return nil, err
		}

		// @ 2b. Create new SafeFile
		sm := &FileItem{
			Path:    file.Path,
			outInfo: info,
			Owner:   p,
			request: req,
		}

		// @ 3b. Callback with Preview
		hc <- true
		return sm, nil
	}
}

// ^ NewIncomingFileItem Prepares for Incoming Data ^ //
func NewIncomingFileItem(inv *md.AuthInvite) (*FileItem, error) {
	info := md.GetInFileInfo(inv)
	return &FileItem{
		Owner:  inv.From,
		inInfo: info,
	}, nil
}

func (pf *FileItem) InfoOut() (*md.OutFileInfo, error) {
	if pf.outInfo != nil {
		return pf.outInfo, nil
	}
	return nil, errors.New("No Outgoing Info")
}

func (pf *FileItem) InfoIn() (*md.InFileInfo, error) {
	if pf.inInfo != nil {
		return pf.inInfo, nil
	}
	return nil, errors.New("No Incoming Info")
}
