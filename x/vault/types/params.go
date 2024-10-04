package types

import (
	"encoding/json"

	"github.com/onsonr/sonr/pkg/orm"
)

// DefaultParams returns default module parameters.
func DefaultParams() Params {
	// TODO:
	return Params{
		IpfsActive: true,
		Schema:     DefaultSchema(),
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
		Version:    orm.SchemaVersion,
		Account:    orm.GetSchema(&orm.Account{}),
		Asset:      orm.GetSchema(&orm.Asset{}),
		Chain:      orm.GetSchema(&orm.Chain{}),
		Credential: orm.GetSchema(&orm.Credential{}),
		Grant:      orm.GetSchema(&orm.Grant{}),
		Keyshare:   orm.GetSchema(&orm.Keyshare{}),
		Profile:    orm.GetSchema(&orm.Profile{}),
	}
}
