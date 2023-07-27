package types

const (
	// ModuleName defines the module name
	ModuleName = "identity"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_identity"

	// Version defines the current version the IBC module supports
	Version = "identity-1"

	// PortID is the default port id that module binds to
	PortID = "identity"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("identity-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	ControllerAccountKeyPrefix      = "ControllerAccount/value/"
	ControllerAccountCountKeyPrefix = "ControllerAccount/count/"
)

const (
	EscrowAccountKeyPrefix      = "EscrowAccount/value/"
	EscrowAccountCountKeyPrefix = "EscrowAccount/count/"
)
