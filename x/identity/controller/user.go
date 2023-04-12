package controller

import (
	"github.com/goccy/go-json"
	"github.com/sonrhq/core/x/identity/types"
)

type User struct {
	// DID of the user
	Did string `json:"_id"`

	// DID document of the primary identity
	Username string `json:"username"`

	// Map of the dids of the keyshares to the dids of the accounts
	Accounts []string `json:"accounts"`

	// DidDocument of the primary identity
	PrimaryIdentity *types.DidDocument `json:"primaryIdentity"`
}

func NewUser(c Controller) *User {
	accDids := make([]string, 0)
	accs, err := c.ListAccounts()
	if err != nil {
		return nil
	}

	for _, acc := range accs {
		accDids = append(accDids, acc.Did())
	}

	return &User{
		Did:             c.Did(),
		Accounts:        accDids,
		PrimaryIdentity: c.PrimaryIdentity(),
	}
}

func LoadUser(data []byte) (*User, error) {
	var u User
	err := json.Unmarshal(data, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (u *User) Marshal() ([]byte, error) {
	return json.Marshal(u)
}
