package models

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"time"

	"github.com/denisbrodbeck/machineid"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/pkg/util"
)

// ** ─── KeyPair MANAGEMENT ────────────────────────────────────────────────────────
// Constructer that Initializes KeyPair without Buffer
func (d *Device) SetKeyPair() *SonrError {
	if d.HasKeys() {
		// Get PrivKey File
		privBuf, serr := d.ReadKey()
		if serr != nil {
			return serr
		}

		// Get Private Key from Buffer
		privKey, err := crypto.UnmarshalPrivateKey(privBuf)
		if err != nil {
			return NewError(err, ErrorMessage_KEY_INVALID)
		}

		// Get Public Key from Private and Marshal
		pubKey := privKey.GetPublic()
		pubBuf, err := crypto.MarshalPublicKey(pubKey)
		if err != nil {
			return NewError(err, ErrorMessage_KEY_SET)
		}

		// Get ID from Pub Key
		id, err := peer.IDFromPublicKey(pubKey)
		if err != nil {
			return NewError(err, ErrorMessage_KEY_ID)
		}

		// Set Key Pair
		d.KeyPair = &KeyPair{
			Type: KeyType_Ed25519,
			Public: &KeyPair_Public{
				Id:     id.String(),
				Buffer: pubBuf,
			},
			Private: &KeyPair_Private{
				Path:   d.WorkingKeyPath(),
				Buffer: privBuf,
			},
		}
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

		// Write Private Key to File
		path, werr := d.WriteKey(privBuf)
		if werr != nil {
			return NewError(err, ErrorMessage_USER_SAVE)
		}

		// Get ID from Pub Key
		id, err := peer.IDFromPublicKey(pubKey)
		if err != nil {
			return NewError(err, ErrorMessage_KEY_ID)
		}

		// Set Keys
		d.KeyPair = &KeyPair{
			Type: KeyType_Ed25519,
			Public: &KeyPair_Public{
				Id:     id.String(),
				Buffer: pubBuf,
			},
			Private: &KeyPair_Private{
				Path:   path,
				Buffer: privBuf,
			},
		}
	}
	return nil
}

// Method Returns PeerID from Public Key
func (kp *KeyPair) ID() (peer.ID, *SonrError) {
	id, err := peer.IDFromPublicKey(kp.PubKey())
	if err != nil {
		return "", NewError(err, ErrorMessage_KEY_ID)
	}
	return id, nil
}

// Method Returns Private Key
func (kp *KeyPair) PrivKey() crypto.PrivKey {
	// Get Key from Buffer
	key, err := crypto.UnmarshalPrivateKey(kp.GetPrivate().GetBuffer())
	if err != nil {
		return nil
	}
	return key
}

// Method Returns Private Key
func (kp *KeyPair) PrivBuffer() []byte {
	return kp.GetPrivate().GetBuffer()
}

// Method Returns Public Key
func (kp *KeyPair) PubKey() crypto.PubKey {
	// Get Key from Buffer
	privKey, err := crypto.UnmarshalPrivateKey(kp.GetPrivate().GetBuffer())
	if err != nil {
		return nil
	}
	return privKey.GetPublic()
}

// Method Signs given data and returns response
func (kp *KeyPair) Sign(value string) string {
	h := hmac.New(sha256.New, kp.PrivBuffer())
	h.Write([]byte(value))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

// Method verifies 'sig' is the signed hash of 'data'
func (kp *KeyPair) Verify(data []byte, sig []byte) (bool, error) {
	// Check for Public Key
	if pubKey := kp.PubKey(); pubKey != nil {
		result, err := pubKey.Verify(data, sig)
		if err != nil {
			return false, err
		}
		return result, nil
	}
	// Return Error
	return false, errors.New("Public Key Doesnt Exist")
}

// ** ─── DEVICE MANAGEMENT ────────────────────────────────────────────────────────
// Method Initializes Device
func (d *Device) Initialize() {
	if d.GetId() == "" {
		id, err := machineid.ID()
		if err != nil {
			log.Println(err)
			return
		}

		// Set ID
		d.Id = id
	}
}

// Method Checks if Device has Keys
func (d *Device) HasKeys() bool {
	if _, err := os.Stat(d.WorkingFilePath(util.KEY_FILE_NAME)); os.IsNotExist(err) {
		return false
	}
	return true
}

// Method Checks for Desktop
func (d *Device) IsDesktop() bool {
	return d.Platform == Platform_MacOS || d.Platform == Platform_Linux || d.Platform == Platform_Windows
}

// Method Checks for Mobile
func (d *Device) IsMobile() bool {
	return d.Platform == Platform_IOS || d.Platform == Platform_Android
}

// Method Checks for IOS
func (d *Device) IsIOS() bool {
	return d.Platform == Platform_IOS
}

// Method Checks for Android
func (d *Device) IsAndroid() bool {
	return d.Platform == Platform_Android
}

// Method Checks for MacOS
func (d *Device) IsMacOS() bool {
	return d.Platform == Platform_MacOS
}

// Method Checks for Linux
func (d *Device) IsLinux() bool {
	return d.Platform == Platform_Linux
}

// Method Checks for Web
func (d *Device) IsWeb() bool {
	return d.Platform == Platform_Web
}

// Method Checks for Windows
func (d *Device) IsWindows() bool {
	return d.Platform == Platform_Windows
}

// Checks if File Exists
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

// Loads Private Key Buf from Device FS Directory
func (d *Device) ReadKey() ([]byte, *SonrError) {
	dat, err := os.ReadFile(d.WorkingKeyPath())
	if err != nil {
		return nil, NewError(err, ErrorMessage_USER_LOAD)
	}
	return dat, nil
}

// Loads File from Disk as Buffer
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

// Returns Path for Private Key File
func (d *Device) WorkingKeyPath() string {
	// Check for Desktop
	if d.IsDesktop() {
		return filepath.Join(d.FileSystem.GetLibrary(), util.KEY_FILE_NAME)
	} else {
		return filepath.Join(d.FileSystem.GetSupport(), util.KEY_FILE_NAME)
	}
}

// Returns Path for Application/User Data
func (d *Device) WorkingFilePath(fileName string) string {
	// Check for Desktop
	if d.IsDesktop() {
		return filepath.Join(d.FileSystem.GetDownloads(), fileName)
	} else {
		return filepath.Join(d.FileSystem.GetDocuments(), fileName)
	}
}

// Returns Path for Application/User Data
func (d *Device) WorkingSupportPath(fileName string) string {
	// Check for Desktop
	if d.IsDesktop() {
		return filepath.Join(d.FileSystem.GetLibrary(), fileName)
	} else {
		return filepath.Join(d.FileSystem.GetSupport(), fileName)
	}
}

// Returns Directory for Device Working Support Folder
func (d *Device) WorkingSupportDir() string {
	if d.IsDesktop() {
		return d.GetFileSystem().GetLibrary()
	} else {
		if d.IsAndroid() {
			return fmt.Sprintf(d.GetFileSystem().GetSupport()+"%s", "/.textile/")
		}
		return d.GetFileSystem().GetSupport()
	}
}

// Writes a File to Disk and Returns Path
func (d *Device) WriteKey(data []byte) (string, *SonrError) {
	// Create File Path
	path := d.WorkingKeyPath()

	// Write File to Disk
	if err := os.WriteFile(path, data, 0644); err != nil {
		return "", NewError(err, ErrorMessage_USER_FS)
	}
	return path, nil
}

// Writes a File to Disk and Returns Path for Downloads/Documents
func (d *Device) WriteFile(name string, data []byte) (string, *SonrError) {
	// Create File Path
	path := d.WorkingFilePath(name)

	// Write File to Disk
	if err := os.WriteFile(path, data, 0644); err != nil {
		return "", NewError(err, ErrorMessage_USER_FS)
	}
	return path, nil
}

// ** ─── User MANAGEMENT ────────────────────────────────────────────────────────
// ^ Method Initializes User Info Struct ^ //
func NewUser(ir *InitializeRequest) (*User, *SonrError) {
	// Initialize Device
	d := ir.GetDevice()
	d.Initialize()

	// Fetch Key Pair
	err := d.SetKeyPair()
	if err != nil {
		return nil, err
	}

	// Return User
	u := &User{
		Device:  d,
		ApiKeys: ir.GetApiKeys(),
		Status:  Status_DEFAULT,
	}
	return u, nil
}

// Set the User with ConnectionRequest
func (u *User) InitConnection(cr *ConnectionRequest) {
	u.Contact = cr.GetContact()
	u.SName = u.Contact.Profile.SName
	u.Router = &User_Router{
		Rendevouz:  "/sonr/rendevouz/0.9.2",
		LocalTopic: fmt.Sprintf("/sonr/topic/%s", cr.GetLocation().OLC()),
		Location:   cr.GetLocation(),
	}
	u.Status = Status_IDLE
}

// Checks Whether User is Ready to Communicate
func (u *User) IsReady() bool {
	return u.Contact != nil && u.SName != "" && u.Router != nil && u.Status != Status_DEFAULT
}

// Return Client API Keys
func (u *User) APIKeys() *APIKeys {
	return u.GetApiKeys()
}

// Method Returns DeviceID
func (u *User) DeviceID() string {
	return u.Device.GetId()
}

// Method Returns Profile First Name
func (u *User) FirstName() string {
	return u.GetPeer().GetProfile().GetFirstName()
}

// Method Returns Peer_ID
func (u *User) ID() *Peer_ID {
	return u.GetPeer().GetId()
}

// Method Returns KeyPair
func (u *User) KeyPair() *KeyPair {
	return u.GetDevice().GetKeyPair()
}

// Method Returns Profile Last Name
func (u *User) LastName() string {
	return u.GetPeer().GetProfile().GetLastName()
}

// Method Returns Profile
func (u *User) Profile() *Profile {
	return u.GetContact().GetProfile()
}

// Method Signs Data with KeyPair
func (u *User) Sign(req *AuthRequest) *AuthResponse {
	// Create Prefix
	prefixResult := u.KeyPair().Sign(fmt.Sprintf("%s%s", req.GetSName(), u.DeviceID()))

	// Get Prefix Appended and Place
	prefix := util.Substring(prefixResult, 0, 16)
	// Get FingerPrint from Mnemonic and Place
	fingerprint := u.KeyPair().Sign(req.GetMnemonic())

	// Get ID from Public Key
	identity := u.KeyPair().GetPublic().GetId()

	// Return Response
	return &AuthResponse{
		SignedPrefix:      prefix,
		SignedFingerprint: fingerprint,
		PublicIdentity:    identity,
	}
}

// Method Returns SName
func (u *User) PrettySName() string {
	return fmt.Sprintf("%s.snr/", u.Profile().GetSName())
}

// Method Updates User Contact
func (u *User) UpdateContact(c *Contact) {
	u.Peer.Profile = &Profile{
		SName:     c.GetProfile().GetSName(),
		FirstName: c.GetProfile().GetFirstName(),
		LastName:  c.GetProfile().GetLastName(),
		Picture:   c.GetProfile().GetPicture(),
		Platform:  c.GetProfile().GetPlatform(),
	}
}

// Method Updates User Position
func (u *User) UpdatePosition(pos *Position) {
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
}

// Method Updates User Contact
func (u *User) UpdateProperties(props *Peer_Properties) {
	u.Peer.Properties = props
}

// ^ Signs InviteResponse with Flat Contact
func (u *User) ReplyToFlat(from *Peer) *InviteResponse {
	return &InviteResponse{
		Type:     InviteResponse_Contact,
		FlatMode: true,
		From:     u.GetPeer(),
		Transfer: &Transfer{
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

// ^ NewUpdateEvent Creates Lobby Event with Peer Data ^
func (p *Peer) NewUpdateEvent(topic *Topic) *LobbyEvent {
	return &LobbyEvent{
		Subject: LobbyEvent_UPDATE,
		Peer:    p,
		Id:      p.Id.Peer,
		Topic:   topic,
	}
}

// ^ NewUpdateEvent Creates Lobby Event with Peer Data ^
func (p *Peer) NewExitEvent(topic *Topic) *LobbyEvent {
	return &LobbyEvent{
		Subject: LobbyEvent_EXIT,
		Peer:    p,
		Id:      p.Id.Peer,
		Topic:   topic,
	}
}
