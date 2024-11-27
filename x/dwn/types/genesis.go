package types

// Define capability hierarchy for smart account operations
const (
	// Root capabilities
	CAP_OWNER    = "OWNER"    // Full account control
	CAP_OPERATOR = "OPERATOR" // Can perform operations
	CAP_OBSERVER = "OBSERVER" // Can view account state

	// Operation capabilities
	CAP_EXECUTE = "EXECUTE" // Can execute transactions
	CAP_PROPOSE = "PROPOSE" // Can propose transactions
	CAP_SIGN    = "SIGN"    // Can sign transactions

	// Policy capabilities
	CAP_SET_POLICY    = "SET_POLICY"    // Can modify account policies
	CAP_SET_THRESHOLD = "SET_THRESHOLD" // Can modify signing threshold

	// Recovery capabilities
	CAP_RECOVER = "RECOVER" // Can initiate recovery
	CAP_SOCIAL  = "SOCIAL"  // Can act as social recovery
)

// Resource types for smart account operations
type ResourceType string

const (
	RES_ACCOUNT     = "account"
	RES_TRANSACTION = "tx"
	RES_POLICY      = "policy"
	RES_RECOVERY    = "recovery"
)

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
