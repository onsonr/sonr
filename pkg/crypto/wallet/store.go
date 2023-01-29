package wallet

type Store interface {
	// GetShare returns a *cmp.Config for the given name
	GetAccount(name string) (Account, error)

	// PutShare stores the given *cmp.Config under the given name
	PutAccount(acc Account) error

	// JWKClaims returns the JWKClaims for the store to be signed by the identity
	JWKClaims(aacc Account) (string, error)

	// VerifyJWKClaims verifies the JWKClaims for the store
	VerifyJWKClaims(claims string, acc Account) error
}
