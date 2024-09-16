package models

func (a *Account) Table() string {
	return "accounts"
}

func (a *Asset) Table() string {
	return "assets"
}

func (a *Credential) Table() string {
	return "credentials"
}

func (a *Keyshare) Table() string {
	return "keyshares"
}

func (a *Permission) Table() string {
	return "permissions"
}

func (a *Profile) Table() string {
	return "profiles"
}

func (a *Property) Table() string {
	return "properties"
}
