package types

import (
	"encoding/json"

	"github.com/onsonr/sonr/x/dwn/types/models"
)

// DefaultParams returns default module parameters.
func DefaultParams() Params {
	return Params{
		IpfsActive:               true,
		LocalRegistrationEnabled: true,
		Schema:                   DefaultSchema(),
		AllowedOperators: []string{ // TODO:
			"localhost",
			"didao.xyz",
			"sonr.id",
		},
	}
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

// DefaultSchema returns the default schema
func DefaultSchema() *Schema {
	return &Schema{
		Version:    SchemaVersion,
		Account:    GetSchema(&models.Account{}),
		Asset:      GetSchema(&models.Asset{}),
		Chain:      GetSchema(&models.Chain{}),
		Credential: GetSchema(&models.Credential{}),
		Grant:      GetSchema(&models.Grant{}),
		Keyshare:   GetSchema(&models.Keyshare{}),
		Profile:    GetSchema(&models.Profile{}),
	}
}
