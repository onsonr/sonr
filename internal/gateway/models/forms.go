package models

type CreatePasskeyData struct {
	Address       string
	Handle        string
	Name          string
	Challenge     string
	CreationBlock string
}

type CreateProfileData struct {
	TurnstileSiteKey string
	FirstNumber      int
	LastNumber       int
}

func (d CreateProfileData) Sum() int {
	return d.FirstNumber + d.LastNumber
}
