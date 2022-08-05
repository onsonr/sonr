package types

const (
	// ModuleName defines the module name
	ModuleName = "bucket"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_bucket"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	WhereIsCountKey   = "WhereIs-count-"
	WhereIsKeyPrefix  = "Bucket/value/"
	DoucmnetKeyPrefix = "Document/value/"
)

func WhereIsKey(creator string) []byte {
	var key []byte
	creatorBytes := []byte(creator)
	key = append(key, creatorBytes...)
	key = append(key, []byte("/")...)

	return key
}
