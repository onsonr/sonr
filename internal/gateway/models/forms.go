package models

type CreatePasskeyParams struct {
	Address       string
	Handle        string
	Name          string
	Challenge     string
	CreationBlock string
}

type CreateProfileParams struct {
	TurnstileSiteKey string
	FirstNumber      int
	LastNumber       int
}

func (d CreateProfileParams) Sum() int {
	return d.FirstNumber + d.LastNumber
}
