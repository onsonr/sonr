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
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	ClaimableWalletKey      = "ClaimableWallet/value/"
	ClaimableWalletCountKey = "ClaimableWallet/count/"
)

var DidMethodKeyMap = map[string]string{
	"sonr":     SonrIdentityPrefix,
	"ethr":     EthereumIdentityPrefix,
	"btcr":     BitcoinIdentityPrefix,
	"webauthn": AuthenticationKeyPrefix,
	"key":      KeyAgreementKeyPrefix,
	"sov":      CapabilityDelegationKeyPrefix,
}
