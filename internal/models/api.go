package models

func (cr *ConnectionRequest) GetProfile() *Profile {
	return &Profile{
		Username:  cr.GetUsername(),
		FirstName: cr.Contact.GetFirstName(),
		LastName:  cr.Contact.GetLastName(),
		Picture:   cr.Contact.GetPicture(),
		Platform:  cr.Device.GetPlatform(),
	}
}

func (cr *ConnectionRequest) HasUser() bool {
	return cr.User != nil
}

func (cr *ConnectionRequest) IsDesktop() bool {
	return cr.Device.IsDesktop
}
