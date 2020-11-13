package user

import "encoding/json"

// Contact is user contact info
type Contact struct {
	firstname  string
	lastname   string
	profilepic string
}

// NewContact creates profile from string json
func NewContact(ctactJSON string) Contact {
	profile := new(Contact)
	json.Unmarshal([]byte(ctactJSON), &profile)
	return *profile
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
