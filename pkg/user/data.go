package user

import (
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ^ Method Loads User Data from Disk ^ //
func (u *UserConfig) LoadUser() (*md.User, error) {
	// Read File
	dat, err := u.fileSystem.ReadFile(K_SONR_USER_PATH)
	if err != nil {
		return nil, err
	}

	// Get User Data
	user := &md.User{}
	err = proto.Unmarshal(dat, user)
	if err != nil {
		return nil, err
	}

	// Set and Return
	u.user = user
	u.settings = user.GetSettings()
	return user, nil
}

// ^ Method Updates User Contact ^ //
func (u *UserConfig) SaveContact(contact *md.Contact) error {
	// Load User
	user, err := u.LoadUser()
	if err != nil {
		return err
	}

	// Set Contact
	user.Contact = contact

	// Save User
	if err := u.SaveUser(user); err != nil {
		return err
	}
	return nil
}

// ^ Method Adds Device to User ^ //
func (u *UserConfig) SaveDevice(device *md.Device) error {

	// Load User
	user, err := u.LoadUser()
	if err != nil {
		return err
	}

	// Append Devices List
	user.Devices = append(user.Devices, device)

	// Save User
	err = u.SaveUser(user)
	if err != nil {
		return err
	}
	return nil
}

// ^ Write User Data at Path ^
func (u *UserConfig) SaveUser(user *md.User) error {
	// Convert User to Bytes
	data, err := proto.Marshal(user)
	if err != nil {
		return err
	}

	// Write File to Disk
	_, err = u.fileSystem.WriteFile(K_SONR_USER_PATH, data)
	if err != nil {
		return err
	}
	return nil
}
