package types

// Capability hierarchy for smart account operations
// ----------------------------------------------
// OWNER
//
//	└─ OPERATOR
//	     ├─ EXECUTE
//	     ├─ PROPOSE
//	     └─ SIGN
//	└─ SET_POLICY
//	     └─ SET_THRESHOLD
//	└─ RECOVER
//	     └─ SOCIAL
//
// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

func (s *Schema) Equal(that *Schema) bool {
	if that == nil {
		return false
	}
	if s.Version != that.Version {
		return false
	}
	if s.Account != that.Account {
		return false
	}
	if s.Asset != that.Asset {
		return false
	}
	if s.Chain != that.Chain {
		return false
	}
	if s.Credential != that.Credential {
		return false
	}
	if s.Jwk != that.Jwk {
		return false
	}
	if s.Grant != that.Grant {
		return false
	}
	if s.Keyshare != that.Keyshare {
		return false
	}
	if s.Profile != that.Profile {
		return false
	}
	return true
}
