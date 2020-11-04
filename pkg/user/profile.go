package user

import "encoding/json"

// Profile is user contact info
type Profile struct {
	firstname  string
	lastname   string
	profilepic string
}

// NewProfile creates profile from string json
func NewProfile(profileJSON string) Profile {
	profile := new(Profile)
	json.Unmarshal([]byte(profileJSON), &profile)
	return *profile
}
