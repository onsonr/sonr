package user

// Status is enum for Node State
type Status int

// Offline is default Node State
const (
	Offline   Status = 0
	Available Status = 1
	Searching Status = 2
	Busy      Status = 3
	Unknown   Status = 4
)

// const list of status names
var statusNames = [...]string{
	"Offline",
	"Available",
	"Searching",
	"Busy"}

// String converts a status to a string
func (st Status) String() string {
	// Prevent out of range
	if st < Offline || st > Busy {
		return "Unknown"
	}

	// return the name of a Status
	return statusNames[st]
}

// GetStatus converts a string to a status
func GetStatus(s string) Status {
	for i, v := range statusNames {
		if v == s {
			return Status(i)
		}
	}
	return Status(4)
}

// IsStatus checks if given status is legitimate
func (st Status) IsStatus() bool {
	switch st {
	case Offline, Available, Searching, Busy:
		return true
	}
	return false
}
