package user

import (
	"encoding/json"
)

// Contact is user contact info
type Contact struct {
	FirstName  string
	LastName   string
	ProfilePic string
}

// NewContact creates profile from string json
func NewContact(jsonString string) Contact {
	// Generate Map
	byt := []byte(jsonString)
	var data map[string]interface{}
	err := json.Unmarshal(byt, &data)

	// Check error
	if err != nil {
		panic(err)
	}

	// Set Values
	contact := new(Contact)
	contact.FirstName = data["firstName"].(string)
	contact.LastName = data["lastName"].(string)
	contact.ProfilePic = data["profilePic"].(string)
	return *contact
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
