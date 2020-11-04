package user

import (
	"encoding/json"

	"github.com/sonr-io/p2p/pkg/lobby"
)

// User is Model with device, location, profile information
type User struct {
	// Management
	id      string
	olc     string
	device  string
	status  Status
	profile Profile

	// Sensory Variables
	direction float64
	distance  float64
}

// NewUser returns user object
func NewUser(peerID string, olc string, device string, profileJSON string) User {
	// Create User
	return User{
		id:      peerID,
		olc:     olc,
		device:  device,
		status:  Available,
		profile: NewProfile(profileJSON),
	}
}

// Update takes json and updates status/direction
func (u *User) Update(dataJSON string) {
	// Get Data
	result := new(lobby.UpdateMessage)
	json.Unmarshal([]byte(dataJSON), &result)

	// Set New Data
	u.direction = float64(result.Direction)
	u.status = GetStatus(result.Status)
}
