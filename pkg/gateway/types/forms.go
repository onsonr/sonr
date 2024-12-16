package types

// ╭───────────────────────────────────────────────────────────╮
// │            Create Passkey (/register/passkey)             │
// ╰───────────────────────────────────────────────────────────╯

// DefaultCreatePasskeyParams returns a default CreatePasskeyParams
func DefaultCreatePasskeyParams() CreatePasskeyParams {
	return CreatePasskeyParams{
		Address:       "",
		Handle:        "",
		Name:          "",
		Challenge:     "",
		CreationBlock: "",
	}
}

// CreatePasskeyParams represents the parameters for creating a passkey
type CreatePasskeyParams struct {
	Address       string
	Handle        string
	Name          string
	Challenge     string
	CreationBlock string
}

// ╭───────────────────────────────────────────────────────────╮
// │            Create Profile (/register/profile)             │
// ╰───────────────────────────────────────────────────────────╯

// DefaultCreateProfileParams returns a default CreateProfileParams
func DefaultCreateProfileParams() CreateProfileParams {
	return CreateProfileParams{
		TurnstileSiteKey: "",
		FirstNumber:      0,
		LastNumber:       0,
	}
}

// CreateProfileParams represents the parameters for creating a profile
type CreateProfileParams struct {
	TurnstileSiteKey string
	FirstNumber      int
	LastNumber       int
}

// Sum returns the sum of the first and last number
func (d CreateProfileParams) Sum() int {
	return d.FirstNumber + d.LastNumber
}
