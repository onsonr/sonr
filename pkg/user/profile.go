package user

import (
	"encoding/json"
	"fmt"
	"math"
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

// NewProfile returns user object
func NewProfile(peerID string, olc string, device string) Profile {
	// Create User
	return Profile{
		ID:     peerID,
		OLC:    olc,
		Device: device,
	}
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

// Update takes json and updates status/direction
func (u *Profile) Update(direction float64) {

	// Set New Data
	u.Direction = Round(direction, .5, 2)
}

// Round converts a number to be rounded
func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
