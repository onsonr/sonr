package types

import (
	"encoding/json"

	"github.com/onsonr/sonr/pkg/common"
	"github.com/onsonr/sonr/x/dwn/types/models"
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
		Account:    common.GetSchema(&models.Account{}),
		Asset:      common.GetSchema(&models.Asset{}),
		Chain:      common.GetSchema(&models.Chain{}),
		Credential: common.GetSchema(&models.Credential{}),
		Grant:      common.GetSchema(&models.Grant{}),
		Keyshare:   common.GetSchema(&models.Keyshare{}),
		Profile:    common.GetSchema(&models.Profile{}),
	}
}
