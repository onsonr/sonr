package models

func (u *URLLink) GetTransfer() *Transfer {
	return &Transfer{
		Data: &Transfer_Url{
			Url: u,
		},
	}
}

func (f *SonrFile) GetTransfer() *Transfer {
	return &Transfer{
		Data: &Transfer_File{
			File: f,
		},
	}
}

func (c *Contact) GetTransfer() *Transfer {
	return &Transfer{
		Data: &Transfer_Contact{
			Contact: c,
		},
	}
}
