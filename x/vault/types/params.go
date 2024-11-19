package types

import (
	"encoding/json"

	"github.com/onsonr/sonr/pkg/common/types"
	"github.com/onsonr/sonr/pkg/common/types/orm"
)

// DefaultParams returns default module parameters.
func DefaultParams() Params {
	// TODO:
	return Params{
		IpfsActive:               true,
		LocalRegistrationEnabled: true,
		Schema:                   DefaultSchema(),
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
		Version:    types.SchemaVersion,
		Account:    types.GetSchema(&orm.Account{}),
		Asset:      types.GetSchema(&orm.Asset{}),
		Chain:      types.GetSchema(&orm.Chain{}),
		Credential: types.GetSchema(&orm.Credential{}),
		Grant:      types.GetSchema(&orm.Grant{}),
		Keyshare:   types.GetSchema(&orm.Keyshare{}),
		Profile:    types.GetSchema(&orm.Profile{}),
	}
}
