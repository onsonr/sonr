package user

import (
	"encoding/json"
	"fmt"
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

// State returns user State information as string
func (u *Profile) State() string {
	slice := [2]string{fmt.Sprintf("%f", u.Direction), u.Device}
	bytes, err := json.Marshal(slice)

	// Check for Error
	if err != nil {
		println("Error creating update message")
	}

	return string(bytes)
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
