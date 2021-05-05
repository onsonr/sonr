package models

import (
	"hash/fnv"
	"log"
	"math"
	"os"
	"path/filepath"
	"time"

	mid "github.com/denisbrodbeck/machineid"
	"github.com/libp2p/go-libp2p-core/peer"
	crypto "github.com/libp2p/go-libp2p-crypto"
	"github.com/multiformats/go-multiaddr"
	"google.golang.org/protobuf/proto"
)

// ** ─── PEER MANAGEMENT ────────────────────────────────────────────────────────

// ^ Create New Peer from Connection Request and Host ID ^ //
func NewPeer(cr *ConnectionRequest, id peer.ID, maddr multiaddr.Multiaddr) (*Peer, *SonrError) {
	// Initialize
	deviceID := cr.Device.GetId()
	c := cr.GetContact()
	profile := c.GetProfile()

	// Get User ID
	userID := fnv.New32a()
	_, err := userID.Write([]byte(profile.GetUsername()))
	if err != nil {
		return nil, NewError(err, ErrorMessage_HOST_KEY)
	}

	// Check if ID not provided
	if deviceID == "" {
		// Generate ID
		if id, err := mid.ProtectedID("Sonr"); err != nil {
			log.Println(err)
			deviceID = ""
		} else {
			deviceID = id
		}
	}

	// Set Peer
	return &Peer{
		Id: &Peer_ID{
			Peer:   id.String(),
			Device: deviceID,
			User:   userID.Sum32(),
		},
		Profile:  profile,
		Platform: cr.Device.Platform,
		Model:    cr.Device.Model,
	}, nil
}

// ^ Returns Peer as Buffer ^ //
func (p *Peer) Buffer() ([]byte, error) {
	buf, err := proto.Marshal(p)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// ^ Checks for Host Peer ID is Same ^ //
func (p *Peer) IsPeerID(pid peer.ID) bool {
	return p.Id.Peer == pid.String()
}

// ^ Checks for Host Peer ID String is Same ^ //
func (p *Peer) IsPeerIDString(pid string) bool {
	return p.Id.Peer == pid
}

// ^ Checks for Host Peer ID String is not Same ^ //
func (p *Peer) IsNotPeerIDString(pid string) bool {
	return p.Id.Peer != pid
}

// ^ Checks for Host Peer ID String is not Same ^ //
func (p *Peer) PeerID() string {
	return p.Id.Peer
}

// ^ SignMessage Creates Lobby Event with Message ^
func (p *Peer) SignMessage(m string, to *Peer) *LobbyEvent {
	return &LobbyEvent{
		Event:   LobbyEvent_MESSAGE,
		From:    p,
		Id:      p.Id.Peer,
		Message: m,
		To:      to.Id.Peer,
	}
}

// ^ Generate AuthInvite with Contact Payload from Request, User Peer Data and User Contact ^ //
func (p *Peer) SignInviteWithContact(c *Contact, flat bool, req *InviteRequest) AuthInvite {
	// Create Invite
	return AuthInvite{
		From:    p,
		IsFlat:  flat,
		Contact: c,
		Payload: req.GetPayload(),
		Remote:  req.GetRemote(),
		To:      req.GetTo(),
	}
}

// ^ Generate AuthInvite with Contact Payload from Request, User Peer Data and User Contact ^ //
func (p *Peer) SignInviteWithFile(req *InviteRequest) AuthInvite {
	// Create Invite
	return AuthInvite{
		From:    p,
		To:      req.GetTo(),
		Payload: req.GetPayload(),
		File:    req.GetFile(),
		Remote:  req.GetRemote(),
	}
}

// ^ Generate AuthInvite with URL Payload from Request and User Peer Data ^ //
func (p *Peer) SignInviteWithLink(req *InviteRequest) AuthInvite {
	// Get URL Data
	urlInfo, err := GetPageInfoFromUrl(req.Url)
	if err != nil {
		urlInfo = &URLLink{
			Link: req.Url,
		}
	}

	// Create Invite
	return AuthInvite{
		From:    p,
		Url:     urlInfo,
		Payload: req.GetPayload(),
		Remote:  req.GetRemote(),
		To:      req.GetTo(),
	}
}

// ^ SignReply Creates AuthReply ^
func (p *Peer) SignReply(d bool, req *RespondRequest, to *Peer) *AuthReply {
	return &AuthReply{
		From:     p,
		Type:     AuthReply_Transfer,
		Decision: d,
		Remote:   req.GetRemote(),
		Card: &TransferCard{
			// SQL Properties
			Payload:  Payload_NONE,
			Received: int32(time.Now().Unix()),

			// Owner Properties
			Owner:    p.Profile,
			Receiver: to.GetProfile(),
		},
	}
}

// ^ SignReply Creates AuthReply with Contact  ^
func (p *Peer) SignReplyWithContact(c *Contact, flat bool, req *RespondRequest, to *Peer) *AuthReply {
	// Set Reply Type
	var kind AuthReply_Type
	if flat {
		kind = AuthReply_FlatContact
	} else {
		kind = AuthReply_Contact
	}

	// Check if Request Provided
	if req != nil {
		// Build Reply
		return &AuthReply{
			From:   p,
			Type:   kind,
			Remote: req.GetRemote(),
			Card: &TransferCard{
				// SQL Properties
				Payload:  Payload_CONTACT,
				Received: int32(time.Now().Unix()),

				// Owner Properties
				Owner:    p.Profile,
				Receiver: to.GetProfile(),

				// Data Properties
				Contact: c,
			},
		}
	} else {
		// Build Reply
		return &AuthReply{
			From: p,
			Type: kind,
			Card: &TransferCard{
				// SQL Properties
				Payload:  Payload_CONTACT,
				Received: int32(time.Now().Unix()),

				// Owner Properties
				Owner:    p.Profile,
				Receiver: to.GetProfile(),

				// Data Properties
				Contact: c,
			},
		}
	}

}

// ^ SignUpdate Creates Lobby Event with Peer Data ^
func (p *Peer) SignUpdate() *LobbyEvent {
	return &LobbyEvent{
		Event: LobbyEvent_UPDATE,
		From:  p,
		Id:    p.Id.Peer,
	}
}

// ^ Processes Update Request ^ //
func (p *Peer) Update(u *UpdateRequest) {
	if u.Type == UpdateRequest_Position {
		// Extract Data
		facing := u.Position.GetFacing()
		heading := u.Position.GetHeading()

		// Update User Values
		var faceDir float64
		var faceAnpd float64
		var headDir float64
		var headAnpd float64
		faceDir = math.Round(facing.Direction*100) / 100
		headDir = math.Round(heading.Direction*100) / 100
		faceDesg := int((facing.Direction / 11.25) + 0.25)
		headDesg := int((heading.Direction / 11.25) + 0.25)

		// Find Antipodal
		if facing.Direction > 180 {
			faceAnpd = math.Round((facing.Direction-180)*100) / 100
		} else {
			faceAnpd = math.Round((facing.Direction+180)*100) / 100
		}

		// Find Antipodal
		if heading.Direction > 180 {
			headAnpd = math.Round((heading.Direction-180)*100) / 100
		} else {
			headAnpd = math.Round((heading.Direction+180)*100) / 100
		}

		// Set Position
		p.Position = &Position{
			Facing: &Position_Compass{
				Direction: faceDir,
				Antipodal: faceAnpd,
				Cardinal:  Cardinal(faceDesg % 32),
			},
			Heading: &Position_Compass{
				Direction: headDir,
				Antipodal: headAnpd,
				Cardinal:  Cardinal(headDesg % 32),
			},
			Accelerometer: u.Position.GetAccelerometer(),
			Gyroscope:     u.Position.GetGyroscope(),
			Magnometer:    u.Position.GetMagnometer(),
			Orientation:   u.Position.GetOrientation(),
		}
	}

	// Set Properties
	if u.Type == UpdateRequest_Properties {
		p.Properties = u.Properties
	}

	// Check for New Contact, Update Peer Profile
	if u.Type == UpdateRequest_Contact {
		profile := u.Contact.GetProfile()
		p.Profile = &Profile{
			FirstName: profile.GetFirstName(),
			LastName:  profile.GetLastName(),
			Picture:   profile.GetPicture(),
		}
	}
}

// ** ─── DEVICE MANAGEMENT ────────────────────────────────────────────────────────
// @ Checks if File Exists
func (d *Device) IsFile(name string) bool {
	// Initialize
	var path string

	// Create File Path
	if d.IsDesktop() {
		path = filepath.Join(d.Directories.GetLibrary(), name)
	} else {
		path = filepath.Join(d.Directories.GetDocuments(), name)
	}

	// Check Path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

// @ Returns Private key from disk if found
func (d *Device) PrivateKey() (crypto.PrivKey, *SonrError) {
	K_SONR_PRIV_KEY := "snr-peer.privkey"

	// Get Private Key
	if ok := d.IsFile(K_SONR_PRIV_KEY); ok {
		// Get Key File
		buf, serr := d.ReadFile(K_SONR_PRIV_KEY)
		if serr != nil {
			return nil, serr
		}

		// Get Key from Buffer
		key, err := crypto.UnmarshalPrivateKey(buf)
		if err != nil {
			return nil, NewError(err, ErrorMessage_HOST_KEY)
		}

		// Set Key Ref
		return key, nil
	} else {
		// Create New Key
		privKey, _, err := crypto.GenerateKeyPair(crypto.Ed25519, -1)
		if err != nil {
			return nil, NewError(err, ErrorMessage_HOST_KEY)
		}

		// Marshal Data
		buf, err := crypto.MarshalPrivateKey(privKey)
		if err != nil {
			return nil, NewError(err, ErrorMessage_MARSHAL)
		}

		// Write Key to File
		_, werr := d.WriteFile(K_SONR_PRIV_KEY, buf)
		if werr != nil {
			return nil, NewError(err, ErrorMessage_USER_SAVE)
		}

		// Set Key Ref
		return privKey, nil
	}
}

// Loads User File
func (d *Device) ReadFile(name string) ([]byte, *SonrError) {
	// Initialize
	var path string

	// Create File Path
	if d.IsDesktop() {
		path = filepath.Join(d.Directories.GetLibrary(), name)
	} else {
		path = filepath.Join(d.Directories.GetDocuments(), name)
	}

	// @ Check for Path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, NewError(err, ErrorMessage_USER_LOAD)
	} else {
		// @ Read User Data File
		dat, err := os.ReadFile(path)
		if err != nil {
			return nil, NewError(err, ErrorMessage_USER_LOAD)
		}
		return dat, nil
	}
}

// Saves Transfer File to Disk
func (d *Device) SaveTransfer(f *SonrFile, i int, data []byte) error {
	path := d.Directories.TransferSavePath(f.Files[i].Name, f.Files[i].Mime, d.IsDesktop())

	// Write File to Disk
	if err := os.WriteFile(path, data, 0644); err != nil {
		return err
	}

	// Set Path for Item
	f.Files[i].Path = path
	f.Files[i].Size = int32(len(data))
	return nil
}

// Writes a File to Disk
func (d *Device) WriteFile(name string, data []byte) (string, *SonrError) {
	// Create File Path
	path := d.Directories.DataSavePath(name, d.IsDesktop())

	// Write File to Disk
	if err := os.WriteFile(path, data, 0644); err != nil {
		return "", NewError(err, ErrorMessage_USER_FS)
	}
	return path, nil
}
