package user

import (
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// ^ Method Loads User Data from Disk ^ //
func (u *User) LoadUser() (*md.User, error) {
	// Read File
	dat, err := u.FS.ReadFile(K_SONR_USER_PATH)
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
	u.protoRef = user
	return user, nil
}

// ^ Method Updates User Contact ^ //
func (u *User) SaveContact(contact *md.Contact) error {
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
func (u *User) SaveDevice(device *md.Device) error {

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
func (u *User) SaveUser(user *md.User) error {
	// Convert User to Bytes
	data, err := proto.Marshal(user)
	if err != nil {
		return err
	}

	// Write File to Disk
	_, err = u.FS.WriteFile(K_SONR_USER_PATH, data)
	if err != nil {
		return err
	}
	return nil
}
