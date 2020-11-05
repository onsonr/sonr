package user

import (
	"encoding/json"
)

// Contact is user contact info
type Contact struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	ProfilePic string `json:"profilePic"`
}

// Basic returns user Basic information as string
func (c *Contact) Basic() string {
	slice := [3]string{c.FirstName, c.LastName, c.ProfilePic}
	bytes, err := json.Marshal(slice)

	// Check for Error
	if err != nil {
		println("Error creating update message")
	}

	return string(bytes)
}

// String returns user as json string
func (c *Contact) String() string {
	// Convert to JSON
	msgBytes, err := json.Marshal(c)
	if err != nil {
		println(err)
	}
	return string(msgBytes)
}
