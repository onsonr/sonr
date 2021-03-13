package data

import (
	"errors"
	"hash/fnv"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/denisbrodbeck/machineid"
	"github.com/getsentry/sentry-go"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// @ Constant Variables
const K_SONR_USER_PATH = "user.snr"

// ^ Method Adds Device to User ^ //
func (fs *SonrFS) AddDevice(device *md.Device) error {
	// Load User
	user, err := fs.LoadUser()
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	// Append Devices List
	user.Devices = append(user.Devices, device)

	// Save User
	err = fs.WriteUser(user)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}
	return nil
}

// ^ GetID returns ID Reference ^ //
func (sfs *SonrFS) GetID(connEvent *md.ConnectionRequest, profile *md.Profile, peerID string) *md.Peer_ID {
	// Initialize
	deviceID := connEvent.Device.Id

	// Get User ID
	userID := fnv.New32a()
	_, err := userID.Write([]byte(profile.Username))
	if err != nil {
		return nil
	}

	// Check if ID not provided
	if deviceID == "" {
		// Generate ID
		if id, err := machineid.ProtectedID("Sonr"); err != nil {
			sentry.CaptureException(err)
			deviceID = ""
		} else {
			deviceID = id
		}
	}

	return &md.Peer_ID{
		Peer:   peerID,
		Device: deviceID,
		User:   userID.Sum32(),
	}
}

// ^ Method Loads User Data from Disk ^ //
func (fs *SonrFS) LoadUser() (*md.User, error) {
	path := filepath.Join(fs.Root, K_SONR_USER_PATH)

	// @ Check for Path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, errors.New("User File Does Not Exist")
	} else {
		// @ Read User Data File
		dat, err := ioutil.ReadFile(path)
		if err != nil {
			sentry.CaptureException(err)
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

// ^ Method Updates User Contact ^ //
func (fs *SonrFS) UpdateContact(contact *md.Contact) error {
	// Load User
	user, err := fs.LoadUser()
	if err != nil {
		return err
	}

	// Set Contact
	user.Contact = contact

	// Save User
	err = fs.WriteUser(user)
	if err != nil {
		return err
	}
	return nil
}
