package types

import (
	"encoding/json"

	"github.com/onsonr/sonr/pkg/common"
	orm "github.com/onsonr/sonr/pkg/common/models"
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
		Version:    common.SchemaVersion,
		Account:    common.GetSchema(&orm.Account{}),
		Asset:      common.GetSchema(&orm.Asset{}),
		Chain:      common.GetSchema(&orm.Chain{}),
		Credential: common.GetSchema(&orm.Credential{}),
		Grant:      common.GetSchema(&orm.Grant{}),
		Keyshare:   common.GetSchema(&orm.Keyshare{}),
		Profile:    common.GetSchema(&orm.Profile{}),
	}
}
