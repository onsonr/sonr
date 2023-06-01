package types

const (
	// ModuleName defines the module name
	ModuleName = "vault"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_vault"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func KeysharePrefix(v string) string {
	return "ks/" + v
}

func AccountPrefix(v string) string {
	return "acc/" + v
}

func WebauthnPrefix(v string) string {
	return "webauthn/" + v
}

const (
	ClaimableWalletKey      = "ClaimableWallet/value/"
	ClaimableWalletCountKey = "ClaimableWallet/count/"
)
