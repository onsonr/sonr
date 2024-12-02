package types

import (
	"encoding/json"
	fmt "fmt"
)

// DefaultParams returns default module parameters.
func DefaultParams() Params {
	return Params{}
}

// DefaultSeedMessage returns the default seed message
func DefaultSeedMessage() string {
	l1 := "The Sonr Network shall make no protocol that respects the establishment of centralized authority,"
	l2 := "or prohibits the free exercise of decentralized identity; or abridges the freedom of data sovereignty,"
	l3 := "or of encrypted communication; or the right of the users to peaceally interact and transact,"
	l4 := "and to petition the Network for the redress of vulnerabilities."
	return fmt.Sprintf("%s %s %s %s", l1, l2, l3, l4)
}

// Stringer method for Params.
func (p Params) String() string {
	bz, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	return string(bz)
}

// Validate does the sanity check on the params.
func (p Params) Validate() error {
	// TODO:
	return nil
}
