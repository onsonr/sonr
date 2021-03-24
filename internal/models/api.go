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
