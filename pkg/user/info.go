package user

import "encoding/json"

// Contact is user contact info
type Info struct {
	ID         string
	FirstName  string
	LastName   string
	ProfilePic string
	Device     string
}

// Return User Info Given Profile/Contact
func GetInfo(p Profile, c Contact) Info {
	return Info{
		ID:         p.ID,
		FirstName:  c.FirstName,
		LastName:   c.LastName,
		ProfilePic: c.ProfilePic,
		Device:     p.Device,
	}
}

// String converts message struct to JSON String
func (msg *Info) String() string {
	// Convert to JSON
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		println(err)
	}
	return string(msgBytes)
}
