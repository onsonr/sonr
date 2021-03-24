package models

func (cr *ConnectionRequest) HasUser() bool {
	return cr.User != nil
}

func (cr *ConnectionRequest) IsDesktop() bool {
	return cr.Device.IsDesktop
}
