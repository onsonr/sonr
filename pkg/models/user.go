package models

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	crypto "github.com/libp2p/go-libp2p-crypto"
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

// Saves Transfer File to Disk
func (d *Device) SaveTransfer(f *SonrFile, i int, data []byte) error {
	// Get Save Path
	path := f.Files[i].SetPath(d)

	// Write File to Disk
	if err := os.WriteFile(path, data, 0644); err != nil {
		return err
	}
	return nil
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
	return &User{
		Contact:  cr.GetContact(),
		Device:   cr.GetDevice(),
		Location: cr.GetLocation(),
		Connection: &User_Connection{
			HasConnected:    false,
			HasBootstrapped: false,
			HasJoinedLocal:  false,
			Connectivity:    cr.GetConnectivity(),
			Router: &User_Router{
				Rendevouz:    fmt.Sprintf("/sonr/%s", cr.GetLocation().MajorOLC()),
				LocalIPTopic: fmt.Sprintf("/sonr/topic/%s", cr.GetLocation().IPOLC()),
			},
			Status: Status_IDLE,
		},
	}
}

// Method Loads User Data from Disk
func (u *User) LoadUser() (*User, *SonrError) {
	// Read File
	dat, serr := u.GetDevice().ReadFile("user.snr")
	if serr != nil {
		return nil, serr
	}

	// Get User Data
	userRef := &User{}
	err := proto.Unmarshal(dat, userRef)
	if err != nil {
		return nil, NewError(err, ErrorMessage_UNMARSHAL)
	}

	// Set and Return
	u = userRef
	u.Settings = userRef.GetSettings()
	u.Devices = userRef.GetDevices()
	return u, nil
}

// Method Returns Private Key
func (u *User) PrivateKey() crypto.PrivKey {
	pk, err := u.Device.PrivateKey()
	if err != nil {
		return nil
	}
	return pk
}

// Method Updates User Contact
func (u *User) SaveContact(c *Contact) *SonrError {
	// Load User
	user, err := u.LoadUser()
	if err != nil {
		return err
	}

	// Set Contact
	user.Contact = c

	// Save User
	if err := u.SaveUser(user); err != nil {
		return err
	}
	return nil
}

// Write User Data at Path
func (u *User) SaveUser(user *User) *SonrError {
	// Convert User to Bytes
	data, err := proto.Marshal(user)
	if err != nil {
		return NewError(err, ErrorMessage_UNMARSHAL)
	}

	// Write File to Disk
	_, serr := u.GetDevice().WriteFile("user.snr", data)
	if err != nil {
		return serr
	}
	return nil
}

// Updates User Peer
func (u *User) Update(ur *UpdateRequest) {
	u.GetPeer().Update(ur)
}

// ** ─── Router MANAGEMENT ────────────────────────────────────────────────────────
// @ Local Lobby Topic Protocol ID
func (r *User) LocalIPTopic() string {
	return fmt.Sprintf("/sonr/topic/%s", r.Location.IPOLC())
}

func (r *User) LocalGeoTopic() (string, error) {
	geoOlc, err := r.Location.GeoOLC()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("/sonr/topic/%s", geoOlc), nil
}

// @ Transfer Controller Data Protocol ID
func (r *User_Router) Transfer(id peer.ID) protocol.ID {
	return protocol.ID(fmt.Sprintf("/sonr/transfer/%s", id.Pretty()))
}

// @ Lobby Topic Protocol ID
func (r *User_Router) Topic(name string) string {
	return fmt.Sprintf("/sonr/topic/%s", name)
}

// @ Major Rendevouz Advertising Point
func (u *User) Router() *User_Router {
	return u.GetConnection().GetRouter()
}

// ** ─── Status MANAGEMENT ────────────────────────────────────────────────────────
// Update Connected Connection Status
func (u *User) SetConnected(value bool) *StatusUpdate {
	// Set Value
	u.Connection.HasConnected = value

	// Update Status
	if value {
		u.Connection.Status = Status_CONNECTED
	} else {
		u.Connection.Status = Status_FAILED
	}

	// Returns Status Update
	return &StatusUpdate{Value: u.Connection.GetStatus()}
}

// Update Bootstrap Connection Status
func (u *User) SetBootstrapped(value bool) *StatusUpdate {
	// Set Value
	u.Connection.HasBootstrapped = value

	// Update Status
	if value {
		u.Connection.Status = Status_BOOTSTRAPPED
	} else {
		u.Connection.Status = Status_FAILED
	}

	// Returns Status Update
	return &StatusUpdate{Value: u.Connection.GetStatus()}
}

// Update Bootstrap Connection Status
func (u *User) SetJoinedLocal(value bool) *StatusUpdate {
	// Set Value
	u.Connection.HasJoinedLocal = value

	// Update Status
	if value {
		u.Connection.Status = Status_AVAILABLE
	} else {
		u.Connection.Status = Status_BOOTSTRAPPED
	}

	// Returns Status Update
	return &StatusUpdate{Value: u.Connection.GetStatus()}
}

// Update Node Status
func (u *User) SetStatus(ns Status) *StatusUpdate {
	// Set Value
	u.Connection.Status = ns

	// Returns Status Update
	return &StatusUpdate{Value: u.Connection.GetStatus()}
}

// Checks if Status is Given Value
func (u *User) IsStatus(gs Status) bool {
	return u.GetConnection().GetStatus() == gs
}

// Checks if Status is Not Given Value
func (u *User) IsNotStatus(gs Status) bool {
	return u.GetConnection().GetStatus() != gs
}
