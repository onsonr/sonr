package sonr

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/denisbrodbeck/machineid"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ^ Method Adds Device to User ^ //
func addDevice(device *md.Device, docPath string) error {
	// Load User
	user, err := loadUser(docPath)
	if err != nil {
		return err
	}

	// Append Devices List
	user.Devices = append(user.Devices, device)

	// Save User
	err = saveUser(user, docPath)
	if err != nil {
		return err
	}
	return nil
}

// ^ Method Creates User and Saves Data to Disk ^ //
func createUser(connEvent *md.ConnectionRequest) error {
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
		Profile: connEvent.Profile,
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
func loadUser(docPath string) (*md.User, error) {
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
func saveUser(user *md.User, docPath string) error {
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
func updateContact(contact *md.Contact, docPath string) error {
	// Load User
	user, err := loadUser(docPath)
	if err != nil {
		return err
	}

	// Set Contact
	user.Contact = contact

	// Save User
	err = saveUser(user, docPath)
	if err != nil {
		return err
	}
	return nil
}

// ^ getDeviceID sets node device ID from path if Exists ^ //
func getDeviceID(connEvent *md.ConnectionRequest) error {
	// Check if ID already provided
	if connEvent.Device.Id != "" {
		return nil
	}

	// Create Device ID Path
	path := filepath.Join(connEvent.Directories.Documents, ".sonr-device-id")

	// @ Check for Path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Generate ID
		id, err := machineid.ProtectedID("Sonr")
		if err != nil {
			return err
		}

		// Write ID To File
		f, err := os.Create(path)
		if err != nil {
			return err
		}

		// Defer Close
		defer f.Close()

		// Write to File
		_, err = f.WriteString(id)
		if err != nil {
			return err
		}

		// Update Device
		connEvent.Device.Id = id
		return nil
	} else {
		// @ Read Device ID Data
		dat, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		// Convert to String
		id := string(dat)

		// Update Device
		connEvent.Device.Id = id

		return nil
	}
}
