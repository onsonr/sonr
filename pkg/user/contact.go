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

// SetContact creates profile from string json
func SetContact(jsonString string) Contact {
	// Generate Map
	byt := []byte(jsonString)
	var data map[string]interface{}
	err := json.Unmarshal(byt, &data)

	// Check error
	if err != nil {
		panic(err)
	}

	// Set Values
	return Contact{
		FirstName:  data["firstName"].(string),
		LastName:   data["lastName"].(string),
		ProfilePic: data["profilePic"].(string),
	}
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
