package types

const (
	// ModuleName defines the module name
	ModuleName = "schema"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_schema"

	Version = "schema-1"

	PortID = "schema"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	SchemaCountKey  = "Schema-count-"
	SchemaKeyPrefix = "Schema/value/"
)

func SchemaKey(creator string) []byte {
	var key []byte
	creatorBytes := []byte(creator)
	key = append(key, creatorBytes...)
	key = append(key, []byte("/")...)

	return key
}
