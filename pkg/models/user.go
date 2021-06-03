package models

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"time"

	crypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
	"google.golang.org/protobuf/proto"
)

// ** ─── DEVICE MANAGEMENT ────────────────────────────────────────────────────────
func (d *Device) IsDesktop() bool {
	return d.Platform == Platform_MacOS || d.Platform == Platform_Linux || d.Platform == Platform_Windows
}

func (d *Device) IsMobile() bool {
	return d.Platform == Platform_IOS || d.Platform == Platform_Android
}

func (d *Device) IsIOS() bool {
	return d.Platform == Platform_IOS
}

func (d *Device) IsAndroid() bool {
	return d.Platform == Platform_Android
}

func (d *Device) IsMacOS() bool {
	return d.Platform == Platform_MacOS
}

func (d *Device) IsLinux() bool {
	return d.Platform == Platform_Linux
}

func (d *Device) IsWeb() bool {
	return d.Platform == Platform_Web
}

func (d *Device) IsWindows() bool {
	return d.Platform == Platform_Windows
}

// Returns Path for Application/User Data
func (d *Device) DataSavePath(fileName string, IsDesktop bool) string {
	// Check for Desktop
	if IsDesktop {
		return filepath.Join(d.FileSystem.GetLibrary(), fileName)
	} else {
		return filepath.Join(d.FileSystem.GetSupport(), fileName)
	}
}

// @ Checks if File Exists
func (d *Device) IsFile(name string) bool {
	// Initialize
	var path string

	// Create File Path
	if d.IsDesktop() {
		path = filepath.Join(d.FileSystem.GetLibrary(), name)
	} else {
		path = filepath.Join(d.FileSystem.GetDocuments(), name)
	}

	// Check Path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

// @ Returns Private key from disk if found
func (d *Device) NewPrivateKey() *SonrError {
	K_SONR_PRIV_KEY := "snr-priv-key"
	K_SONR_PUB_KEY := "snr-pub-key"

	// Get Private Key
	if ok := d.IsFile(K_SONR_PRIV_KEY); ok {
		// Get Key File
		buf, serr := d.ReadFile(K_SONR_PRIV_KEY)
		if serr != nil {
			return serr
		}

		// Set Buffer for Key
		d.PrivateKey = buf

		// Set Key Ref
		return nil
	} else {
		// Create New Key
		privKey, pubKey, err := crypto.GenerateKeyPair(crypto.Ed25519, -1)
		if err != nil {
			return NewError(err, ErrorMessage_HOST_KEY)
		}

		// Marshal Data
		privBuf, err := crypto.MarshalPrivateKey(privKey)
		if err != nil {
			return NewError(err, ErrorMessage_MARSHAL)
		}

		// Marshal Data
		pubBuf, err := crypto.MarshalPublicKey(pubKey)
		if err != nil {
			return NewError(err, ErrorMessage_MARSHAL)
		}

		// Set Buffer for Key
		d.PrivateKey = privBuf
		d.PublicKey = pubBuf

		// Write Private Key to File
		_, werr := d.WriteFile(K_SONR_PRIV_KEY, privBuf)
		if werr != nil {
			return NewError(err, ErrorMessage_USER_SAVE)
		}

		// Write Public Key to File
		_, wpuberr := d.WriteFile(K_SONR_PUB_KEY, pubBuf)
		if wpuberr != nil {
			return NewError(err, ErrorMessage_USER_SAVE)
		}
		return nil
	}
}

// Loads User File
func (d *Device) ReadFile(name string) ([]byte, *SonrError) {
	// Initialize
	var path string

	// Create File Path
	if d.IsDesktop() {
		path = filepath.Join(d.FileSystem.GetLibrary(), name)
	} else {
		path = filepath.Join(d.FileSystem.GetDocuments(), name)
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

// Writes a File to Disk
func (d *Device) WriteFile(name string, data []byte) (string, *SonrError) {
	// Create File Path
	path := d.DataSavePath(name, d.IsDesktop())

	// Write File to Disk
	if err := os.WriteFile(path, data, 0644); err != nil {
		return "", NewError(err, ErrorMessage_USER_FS)
	}
	return path, nil
}

// ** ─── User MANAGEMENT ────────────────────────────────────────────────────────
// ^ Method Initializes User Info Struct ^ //
func NewUser(cr *ConnectionRequest) *User {
	// Initialize Device
	d := cr.GetDevice()
	d.NewPrivateKey()

	// Get Crypto
	crypto := cr.GetCrypto()

	// Return User
	return &User{
		Id:       crypto.GetPrefix(),
		Device:   d,
		Contact:  cr.GetContact(),
		Location: cr.GetLocation(),
		Crypto:   crypto,
		Connection: &User_Connection{
			ApiKeys: cr.GetClientKeys(),
			Router: &User_Router{
				Rendevouz:  "/sonr/rendevouz/0.9.2",
				LocalTopic: fmt.Sprintf("/sonr/topic/%s", cr.GetLocation().OLC()),
				Location:   cr.GetLocation(),
			},
			Status: Status_IDLE,
		},
		Devices:  cr.GetDevices(),
		Settings: cr.GetSettings(),
	}
}

func (u *User) DeviceID() string {
	return u.Device.GetId()
}

func (u *User) ID() *Peer_ID {
	return u.GetPeer().GetId()
}

// Method Returns Crypto Prefix With Signature
func (u *User) Prefix() string {
	return u.GetCrypto().Prefix
}

func (u *User) Profile() *Profile {
	return u.GetContact().GetProfile()
}

// Method Returns Private Key
func (u *User) PrivateKey() crypto.PrivKey {
	// Get Key from Buffer
	key, err := crypto.UnmarshalPrivateKey(u.GetDevice().GetPrivateKey())
	if err != nil {
		return nil
	}
	return key
}

// Method Returns SName
func (u *User) SName() string {
	return fmt.Sprintf("%s.snr/", u.Profile().GetSName())
}

// Updates User Peer
func (u *User) Update(ur *UpdateRequest) {
	switch ur.Data.(type) {
	case *UpdateRequest_Position:
		// Extract Data
		pos := ur.GetPosition()
		facing := pos.GetFacing()
		heading := pos.GetHeading()

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
		u.Peer.Position = &Position{
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
			Orientation: pos.GetOrientation(),
		}

	case *UpdateRequest_Contact:
		u.Contact = ur.GetContact()
		u.Peer.Profile = &Profile{
			FirstName: u.Contact.GetProfile().GetFirstName(),
			LastName:  u.Contact.GetProfile().GetLastName(),
			Picture:   u.Contact.GetProfile().GetPicture(),
		}
	case *UpdateRequest_Properties:
		props := ur.GetProperties()
		u.Peer.Properties = props
	default:
		return
	}
}

// ** ─── Peer MANAGEMENT ────────────────────────────────────────────────────────
// ^ Create New Peer from Connection Request and Host ID ^ //
func (u *User) NewPeer(id peer.ID, maddr multiaddr.Multiaddr) *SonrError {
	u.Peer = &Peer{
		Id: &Peer_ID{
			Peer:   id.String(),
			Device: u.DeviceID(),
			SName:  u.SName(),
		},
		Profile:  u.Profile(),
		Platform: u.Device.Platform,
		Model:    u.Device.Model,
	}
	// Set Device Topic
	u.Connection.Router.DeviceTopic = fmt.Sprintf("/sonr/topic/%s", u.Peer.UserID())
	return nil
}

// ^ Returns Peer as Buffer ^ //
func (p *Peer) Buffer() ([]byte, error) {
	buf, err := proto.Marshal(p)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// ^ Returns Peer User ID ^ //
func (p *Peer) DeviceID() string {
	return string(p.Id.GetDevice())
}

// ^ Returns Peer ID String Value
func (p *Peer) PeerID() string {
	return p.Id.Peer
}

// ^ Returns Peer User ID ^ //
func (p *Peer) UserID() string {
	return p.Id.GetSName()
}

// ^ Checks if Two Peers are the Same by Device ID and Peer ID
func (p *Peer) IsSame(other *Peer) bool {
	return p.PeerID() == other.PeerID() && p.DeviceID() == other.DeviceID() && p.UserID() == other.UserID()
}

// ^ Checks if PeerDeviceIDID is the Same
func (p *Peer) IsSameDeviceID(other *Peer) bool {
	return p.DeviceID() == other.DeviceID()
}

// ^ Checks if PeerID is the Same
func (p *Peer) IsSamePeerID(pid peer.ID) bool {
	return p.PeerID() == pid.String()
}

// ^ Checks if Two Peers are NOT the Same by Device ID and Peer ID
func (p *Peer) IsNotSame(other *Peer) bool {
	return p.PeerID() != other.PeerID() && p.DeviceID() != other.DeviceID() && p.UserID() != other.UserID()
}

// ^ Checks if DeviceID is NOT the Same
func (p *Peer) IsNotSameDeviceID(other *Peer) bool {
	return p.DeviceID() == other.DeviceID()
}

// ^ Checks if PeerID is NOT the Same
func (p *Peer) IsNotSamePeerID(pid peer.ID) bool {
	return p.PeerID() != pid.String()
}

// ^ Signs AuthReply with Flat Contact
func (u *User) SignFlatReply(from *Peer) *AuthReply {
	return &AuthReply{
		Type: AuthReply_FlatContact,
		From: u.GetPeer(),
		Data: &Transfer{
			// SQL Properties
			Payload:  Payload_CONTACT,
			Received: int32(time.Now().Unix()),

			// Owner Properties
			Owner:    u.GetPeer().Profile,
			Receiver: from.GetProfile(),

			// Data Properties
			Data: u.GetContact().ToData(),
		},
	}
}

// ^ SignUpdate Creates Lobby Event with Peer Data ^
func (p *Peer) SignUpdate() *LocalEvent {
	return &LocalEvent{
		Subject: LocalEvent_UPDATE,
		From:    p,
		Id:      p.Id.Peer,
	}
}
