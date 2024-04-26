package types

import (
	"encoding/json"
)

// DefaultParams returns default module parameters.
func DefaultParams() Params {
	// TODO:
	return Params{
		PropertyAllowlist: []string{
			"email",
			"phone",
		},
		DefaultCurve:             "P-256",
		WhitelistedVerifications: []string{},
		AssertionRewardRate:      0.35,
		EncryptionRewardRate:     0.5,
		ReferralRewardRate:       0.15,
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
