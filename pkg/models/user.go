package models

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"time"

	"github.com/denisbrodbeck/machineid"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/sonr-io/core/pkg/util"
)

// ** ─── DEVICE MANAGEMENT ────────────────────────────────────────────────────────
// Method Initializes Device
func (d *Device) Initialize(r *InitializeRequest) (*KeyChain, *SonrError) {
	// Init FileSystem
	d.Status = Status_DEFAULT
	serr := d.GetFileSystem().Initialize()
	if serr != nil {
		return nil, serr
	}

	// Get Machine ID of Device
	if d.GetId() == "" {
		id, err := machineid.ID()
		if err != nil {
			return nil, NewError(err, ErrorEvent_DEVICE_ID)
		}

		// Set ID
		d.Id = id
	}

	// Get Hostname of Device
	if d.GetHostName() == "" {
		name, err := os.Hostname()
		if err != nil {
			return nil, NewError(err, ErrorEvent_DEVICE_ID)
		}
		d.HostName = name
	}

	// Check Initialize Options
	if r.ShouldLoadKeychain() {
		return d.loadKeyChain()
	} else if r.ShouldCreateTempKeys() {
		return d.tempKeyChain()
	} else {
		return d.newKeyChain()
	}
}

// Method Sets Connection for Device
func (d *Device) SetConnection(cr *ConnectionRequest) {
	// Initialize Params
	d.Location = cr.GetLocation()
	d.Status = Status_IDLE
}

// Method Checks for Desktop
func (d *Device) IsDesktop() bool {
	return d.Platform == Platform_MACOS || d.Platform == Platform_LINUX || d.Platform == Platform_WINDOWS
}

// Method Checks for Mobile
func (d *Device) IsMobile() bool {
	return d.Platform == Platform_IOS || d.Platform == Platform_ANDROID
}

// Method Checks for IOS
func (d *Device) IsIOS() bool {
	return d.Platform == Platform_IOS
}

// Method Checks for Android
func (d *Device) IsAndroid() bool {
	return d.Platform == Platform_ANDROID
}

// Method Checks for MacOS
func (d *Device) IsMacOS() bool {
	return d.Platform == Platform_MACOS
}

// Method Checks for Linux
func (d *Device) IsLinux() bool {
	return d.Platform == Platform_LINUX
}

// Method Checks for Web
func (d *Device) IsWeb() bool {
	return d.Platform == Platform_WEB
}

// Method Checks for Windows
func (d *Device) IsWindows() bool {
	return d.Platform == Platform_WINDOWS
}


// Method Updates User Position
func (u *Device) UpdatePosition(faceDir float64, headDir float64, orientation *Position_Orientation) {
	// Update User Values
	var faceAnpd float64
	var headAnpd float64
	faceDir = math.Round(faceDir*100) / 100
	headDir = math.Round(headDir*100) / 100
	faceDesg := int((faceDir / 11.25) + 0.25)
	headDesg := int((headDir / 11.25) + 0.25)

	// Find Antipodal
	if faceDir > 180 {
		faceAnpd = math.Round((faceDir-180)*100) / 100
	} else {
		faceAnpd = math.Round((faceDir+180)*100) / 100
	}

	// Find Antipodal
	if headDir > 180 {
		headAnpd = math.Round((headDir-180)*100) / 100
	} else {
		headAnpd = math.Round((headDir+180)*100) / 100
	}

	// Set Position
	pos := &Position{
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
		Orientation: orientation,
	}

	// Update Position
	u.GetPeer().Position = pos
}

// Method Updates User Contact
func (u *Device) UpdateProperties(props *Peer_Properties) {
	u.GetPeer().Properties = props
}

// Method initializes FileSystem Private Directory
func (f *FileSystem) Initialize() *SonrError {
	// Init Default Private Dir Path
	path := filepath.Join(f.GetSupport().GetPath(), util.PRIVATE_KEY_DIR)

	// Set Directory Reference
	if IsExisting(path) {
		f.Private = &FileSystem_Directory{
			Path: path,
			Type: FileSystem_Directory_PRIVATE,
		}
		return nil
	}

	// Create Private Dir
	err := os.MkdirAll(path, 0700)
	if err != nil {
		return NewError(err, ErrorEvent_USER_FS)
	}

	// Set Reference
	f.Private = &FileSystem_Directory{
		Path: path,
		Type: FileSystem_Directory_PRIVATE,
	}
	return nil
}

// Checks if File Exists
func (d *FileSystem_Directory) IsFile(name string) bool {
	// Check Path
	if _, err := os.Stat(filepath.Join(d.GetPath(), name)); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

// Method Checks if any given path exists
func IsExisting(path string) bool {
	// Check Path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

// Checks if File/Directory Exists
func (d *FileSystem) IsDirectory(rootDir *FileSystem_Directory, subDir string) bool {
	// Check Path
	if _, err := os.Stat(filepath.Join(rootDir.GetPath(), subDir)); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

// Loads Private Key Buf from Device FS Directory
func (d *Device) ReadKey(t KeyPair_Type) ([]byte, *SonrError) {
	dat, err := os.ReadFile(d.WorkingKeyPath(t))
	if err != nil {
		return nil, NewError(err, ErrorEvent_USER_LOAD)
	}
	return dat, nil
}

// Loads File from Disk as Buffer
func (d *FileSystem_Directory) ReadFile(name string) ([]byte, *SonrError) {
	// Initialize
	path := filepath.Join(d.GetPath(), name)

	// Check for Path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, NewError(err, ErrorEvent_USER_LOAD)
	} else {
		// Read User Data File
		dat, err := os.ReadFile(path)
		if err != nil {
			return nil, NewError(err, ErrorEvent_USER_LOAD)
		}
		return dat, nil
	}
}

// Signs InviteResponse with Flat Contact
func (u *Device) ReplyToFlat(from *Peer) *InviteResponse {
	return &InviteResponse{
		Type:    InviteResponse_FLAT,
		To:      from,
		Payload: Payload_CONTACT,
		From:    u.GetPeer(),
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


// Returns Path for Private Key File
func (d *Device) WorkingKeyPath(t KeyPair_Type) string {
	// Check for Desktop
	return filepath.Join(d.GetFileSystem().GetPrivate().GetPath(), t.FileName())
}

// Returns Path for Application/User Data
func (d *Device) WorkingFilePath(fileName string) string {
	// Check for Desktop
	return filepath.Join(d.GetFileSystem().GetDownloads().GetPath(), fileName)
}

// Returns Path for Application/User Data
func (d *Device) WorkingSupportPath(fileName string) string {
	// Check for Desktop
	return filepath.Join(d.GetFileSystem().GetSupport().GetPath(), fileName)
}

// Returns Directory for Device Working Support Folder
func (d *Device) WorkingSupportDir() string {
	return d.GetFileSystem().GetSupport().GetPath()
}

// Writes a File to Disk and Returns Path
func (d *Device) WriteKey(data []byte, t KeyPair_Type) (string, *SonrError) {
	// Create File Path
	path := d.WorkingKeyPath(t)

	// Write File to Disk
	if err := os.WriteFile(path, data, 0644); err != nil {
		return "", NewError(err, ErrorEvent_USER_FS)
	}
	return path, nil
}

// Writes a File to Disk and Returns Path for Downloads/Documents
func (d *Device) WriteFile(name string, data []byte) (string, *SonrError) {
	// Create File Path
	path := d.WorkingFilePath(name)

	// Write File to Disk
	if err := os.WriteFile(path, data, 0644); err != nil {
		return "", NewError(err, ErrorEvent_USER_FS)
	}
	return path, nil
}

// ** ─── User MANAGEMENT ────────────────────────────────────────────────────────
// Method Initializes User Info Struct ^ //
func NewUser(ir *InitializeRequest) (*User, *SonrError) {
	// Initialize Device
	d := ir.GetDevice()

	// Fetch Key Pair
	err := d.Initialize(ir)
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
	// Initialize Params
	u.PushToken = cr.GetPushToken()
	u.SName = cr.GetContact().GetProfile().GetSName()
	u.Location = cr.GetLocation()
	u.Contact = cr.GetContact()
	u.Status = Status_IDLE
}

// Checks Whether User is Ready to Communicate
func (u *User) IsReady() bool {
	return u.Contact != nil && u.SName != "" && u.Location != nil && u.Status != Status_DEFAULT
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
	return u.GetPrimary().GetProfile().GetFirstName()
}

// Method Returns Peer_ID
func (u *User) ID() *Peer_ID {
	return u.GetPrimary().GetId()
}

// Method Returns KeyPair
func (u *User) KeyPair() *KeyPair {
	return u.GetDevice().GetAccountKeys()
}

// Method Returns Profile Last Name
func (u *User) LastName() string {
	return u.GetPrimary().GetProfile().GetLastName()
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
	pubKey := u.KeyPair().PubKeyBase64()

	// Return Response
	return &AuthResponse{
		SignedPrefix:      prefix,
		SignedFingerprint: fingerprint,
		PublicKey:         pubKey,
		GivenSName:        req.GetSName(),
		GivenMnemonic:     req.GetMnemonic(),
	}
}

// Method Returns SName
func (u *User) PrettySName() string {
	return fmt.Sprintf("%s.snr/", u.Profile().GetSName())
}

// Method Updates User Contact
func (u *User) UpdateContact(c *Contact) {
	u.Contact = c
	u.GetMember().UpdateProfile(c)
}

// Method Updates User Position
func (u *User) UpdatePosition(faceDir float64, headDir float64, orientation *Position_Orientation) {
	// Update User Values
	var faceAnpd float64
	var headAnpd float64
	faceDir = math.Round(faceDir*100) / 100
	headDir = math.Round(headDir*100) / 100
	faceDesg := int((faceDir / 11.25) + 0.25)
	headDesg := int((headDir / 11.25) + 0.25)

	// Find Antipodal
	if faceDir > 180 {
		faceAnpd = math.Round((faceDir-180)*100) / 100
	} else {
		faceAnpd = math.Round((faceDir+180)*100) / 100
	}

	// Find Antipodal
	if headDir > 180 {
		headAnpd = math.Round((headDir-180)*100) / 100
	} else {
		headAnpd = math.Round((headDir+180)*100) / 100
	}

	// Set Position
	pos := &Position{
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
		Orientation: orientation,
	}

	// Update Position
	u.GetPrimary().Position = pos
}

// Method Updates User Contact
func (u *User) UpdateProperties(props *Peer_Properties) {
	u.GetPrimary().Properties = props
}

// Method Updates User Contact
func (u *User) VerifyRead() *VerifyResponse {
	kp := u.KeyPair()
	return &VerifyResponse{
		PublicKey: kp.PubKeyBase64(),
		ShortID:   u.GetDevice().ShortID(),
	}
}

// Signs InviteResponse with Flat Contact
func (u *User) ReplyToFlat(from *Peer) *InviteResponse {
	return &InviteResponse{
		Type:    InviteResponse_FLAT,
		To:      from,
		Payload: Payload_CONTACT,
		From:    u.GetPrimary(),
		Transfer: &Transfer{
			// SQL Properties
			Payload:  Payload_CONTACT,
			Received: int32(time.Now().Unix()),

			// Owner Properties
			Owner:    u.GetPrimary().Profile,
			Receiver: from.GetProfile(),

			// Data Properties
			Data: u.GetContact().ToData(),
		},
	}
}

// NewUpdateEvent Creates Lobby Event with Peer Data ^
func (u *User) NewUpdateEvent(room *Room, id peer.ID) *RoomEvent {
	return &RoomEvent{
		Subject: RoomEvent_UPDATE,
		Peer:    u.GetPrimary(),
		Id:      id.String(),
		Room:    room,
	}
}

// NewDefaultUpdateEvent Updates Peer with Default Position and Returns Lobby Event with Peer Data ^
func (u *User) NewDefaultUpdateEvent(room *Room, id peer.ID) *RoomEvent {
	// Update Peer
	u.UpdatePosition(DefaultPosition().Parameters())

	// Check if User is Linker
	if u.IsLinker() {
		// Return Event
		return &RoomEvent{
			Subject: RoomEvent_LINKER,
			Peer:    u.GetPrimary(),
			Id:      id.String(),
			Room:    room,
		}
	} else {
		// Return Event
		return &RoomEvent{
			Subject: RoomEvent_UPDATE,
			Peer:    u.GetPrimary(),
			Id:      id.String(),
			Room:    room,
		}
	}

}

// NewUpdateEvent Creates Lobby Event with Peer Data ^
func (u *User) NewExitEvent(room *Room, id peer.ID) *RoomEvent {
	return &RoomEvent{
		Subject: RoomEvent_EXIT,
		Peer:    u.GetPrimary(),
		Id:      id.String(),
		Room:    room,
	}
}
