package device

var (
	// deviceID is the device ID. Either provided or found
	deviceID string
)

// SetDeviceID sets the device ID.
func SetDeviceID(id string) error {
	if id != "" {
		deviceID = id
		return nil
	}
	logger.Error("Failed to Set Device ID", ErrEmptyDeviceID)
	return ErrEmptyDeviceID
}
