package user

import (
	"encoding/json"
)

// Profile is Model with device, location, profile information
type Profile struct {
	// Management
	ID     string
	OLC    string
	Device string

	// Sensory Variables
	Direction float64
	Distance  float64
}

// String returns user as json string
func (u *Profile) String() string {
	// Create user map
	m := make(map[string]string)
	m["id"] = u.ID
	m["olc"] = u.OLC
	m["device"] = u.Device

	// Convert to JSON
	msgBytes, err := json.Marshal(m)
	if err != nil {
		println(err)
	}

	// Return String
	return string(msgBytes)
}
