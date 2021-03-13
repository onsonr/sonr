package data

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/denisbrodbeck/machineid"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

type SonrFS struct {
	Devices   []*md.Device
	User      *md.User
	Cache     string
	Documents string
	Downloads string
	Home      string
	Sonr      string
	Temporary string
}

// ^ Method Initializes Root Sonr Directory ^ //
func InitFS(connEvent *md.ConnectionRequest, profile *md.Profile) *SonrFS {
	return &SonrFS{
		Cache:     connEvent.Directories.Cache,
		Documents: connEvent.Directories.Documents,
		Downloads: connEvent.Directories.Downloads,
		Home:      connEvent.Directories.Home,
		Sonr:      connEvent.Directories.Sonr,
		Temporary: connEvent.Directories.Temporary,
	}
}

// ^ Method Adds Device to User ^ //
func (fs *SonrFS) AddDevice(device *md.Device, docPath string) error {
	// Load User
	user, err := fs.LoadUser(docPath)
	if err != nil {
		return err
	}

	// Append Devices List
	user.Devices = append(user.Devices, device)

	// Save User
	err = fs.SaveUser(user, docPath)
	if err != nil {
		return err
	}
	return nil
}

// ^ Method Creates User and Saves Data to Disk ^ //
func (fs *SonrFS) CreateUser(connEvent *md.ConnectionRequest, profile *md.Profile) error {
	// Initialize Path
	path := filepath.Join(connEvent.Directories.Documents, "user.snr")

	// Set Device Directories
	connEvent.Device.Directories = connEvent.Directories

	// Create Devices
	devices := make([]*md.Device, 32)
	devices = append(devices, connEvent.Device)

	// Create User
	user := &md.User{
		Contact: connEvent.Contact,
		Profile: profile,
		Devices: devices,
	}

	// Convert User to Bytes
	userData, err := proto.Marshal(user)
	if err != nil {
		return err
	}

	// @ Check for Path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Write ID To File
		f, err := os.Create(path)
		if err != nil {
			return err
		}

		// Defer Close
		defer f.Close()

		// Write to File
		_, err = f.Write(userData)
		if err != nil {
			return err
		}
		return nil
	} else {
		// @ Over write file if Exists
		f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			return err
		}

		// Defer Close
		defer f.Close()

		// Write to File
		_, err = f.Write(userData)
		if err != nil {
			return err
		}

		return nil
	}
}

// ^ Method Loads User Data from Disk ^ //
func (fs *SonrFS) LoadUser(docPath string) (*md.User, error) {
	path := filepath.Join(docPath, "user.snr")

	// @ Check for Path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, errors.New("User File Does Not Exist")
	} else {
		// @ Read User Data File
		dat, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}

		// Get User Data
		user := &md.User{}
		err = proto.Unmarshal(dat, user)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
}

// ^ Method Creates User and Saves Data to Disk ^ //
func (fs *SonrFS) SaveUser(user *md.User, docPath string) error {
	// Initialize Path
	path := filepath.Join(docPath, "user.snr")

	// Convert User to Bytes
	userData, err := proto.Marshal(user)
	if err != nil {
		return err
	}

	// @ Check for Path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Write ID To File
		f, err := os.Create(path)
		if err != nil {
			return err
		}

		// Defer Close
		defer f.Close()

		// Write to File
		_, err = f.Write(userData)
		if err != nil {
			return err
		}
		return nil
	} else {
		// @ Over write file if Exists
		f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			return err
		}

		// Defer Close
		defer f.Close()

		// Write to File
		_, err = f.Write(userData)
		if err != nil {
			return err
		}

		return nil
	}
}

// ^ Method Updates User Contact ^ //
func (fs *SonrFS) UpdateContact(contact *md.Contact, docPath string) error {
	// Load User
	user, err := fs.LoadUser(docPath)
	if err != nil {
		return err
	}

	// Set Contact
	user.Contact = contact

	// Save User
	err = fs.SaveUser(user, docPath)
	if err != nil {
		return err
	}
	return nil
}

// ^ getDeviceID sets node device ID from path if Exists ^ //
func (fs *SonrFS) GetDeviceID(connEvent *md.ConnectionRequest) (string, error) {
	// Check if ID already provided
	if connEvent.Device.Id != "" {
		return connEvent.Device.Id, nil
	}

	// Create Device ID Path
	path := filepath.Join(connEvent.Directories.Documents, ".sonr-device-id")

	// @ Check for Path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Generate ID
		id, err := machineid.ProtectedID("Sonr")
		if err != nil {
			return "", err
		}

		// Write ID To File
		f, err := os.Create(path)
		if err != nil {
			return "", err
		}

		// Defer Close
		defer f.Close()

		// Write to File
		_, err = f.WriteString(id)
		if err != nil {
			return "", err
		}

		// Update Device
		connEvent.Device.Id = id
		return id, nil
	} else {
		// @ Read Device ID Data
		dat, err := ioutil.ReadFile(path)
		if err != nil {
			return "", err
		}

		// Convert to String
		id := string(dat)

		// Update Device
		connEvent.Device.Id = id

		return id, nil
	}
}
