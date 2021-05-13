package models

import (
	"fmt"
	"math"
	"os"
	"path/filepath"

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
	pk, err := u.GetDevice().PrivateKey()
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
	if ur.Type == UpdateRequest_Position {
		// Extract Data
		facing := ur.Position.GetFacing()
		heading := ur.Position.GetHeading()

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
			Orientation: ur.Position.GetOrientation(),
		}
	}

	// Set Properties
	if ur.Type == UpdateRequest_Properties {
		u.Peer.Properties = ur.Properties
	}

	// Check for New Contact, Update Peer Profile
	if ur.Type == UpdateRequest_Contact {
		u.SaveContact(ur.Contact)
		profile := ur.Contact.GetProfile()
		u.Peer.Profile = &Profile{
			FirstName: profile.GetFirstName(),
			LastName:  profile.GetLastName(),
			Picture:   profile.GetPicture(),
		}
	}
}
