package user

import (
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ^ Method Loads User Data from Disk ^ //
func (u *UserConfig) LoadUser() (*md.User, *md.SonrError) {
	// Read File
	dat, serr := u.fileSystem.ReadFile(K_SONR_USER_PATH)
	if serr != nil {
		return nil, serr
	}

	// Get User Data
	user := &md.User{}
	err := proto.Unmarshal(dat, user)
	if err != nil {
		return nil, md.NewError(err, md.ErrorMessage_UNMARSHAL)
	}

	// Set and Return
	u.user = user
	u.settings = user.GetSettings()
	return user, nil
}

// ^ Method Updates User Contact ^ //
func (u *UserConfig) SaveContact(contact *md.Contact) *md.SonrError {
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
func (u *UserConfig) SaveDevice(device *md.Device) *md.SonrError {

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
func (u *UserConfig) SaveUser(user *md.User) *md.SonrError {
	// Convert User to Bytes
	data, err := proto.Marshal(user)
	if err != nil {
		return md.NewError(err, md.ErrorMessage_UNMARSHAL)
	}

	// Write File to Disk
	_, serr := u.fileSystem.WriteFile(K_SONR_USER_PATH, data)
	if err != nil {
		return serr
	}
	return nil
}
