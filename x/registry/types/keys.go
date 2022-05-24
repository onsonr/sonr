package types

const (
	// ModuleName defines the module name
	ModuleName = "registry"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_registry"

	// Version defines the current version the IBC module supports
	Version = "registry-1"

	// PortID is the default port id that module binds to
	PortID = "registry"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("registry-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	WhoIsCountKey  = "WhoIs-count-"
	WhoIsKeyPrefix = "WhoIs/value/"
)

// WhoIsKey returns the store key to retrieve a WhoIs from the did field
func WhoIsKey(did string) []byte {
	var key []byte

	didBytes := []byte(did)
	key = append(key, didBytes...)
	key = append(key, []byte("/")...)

	return key
}
