package db

// DIDNamespace defines the different namespaces of DID
type DIDNamespace int

const (
	DIDNamespaceUnspecified DIDNamespace = iota
	DIDNamespaceIPFS
	DIDNamespaceSonr
	DIDNamespaceBitcoin
	DIDNamespaceEthereum
	DIDNamespaceIBC
	DIDNamespaceWebauthn
	DIDNamespaceDWN
	DIDNamespaceService
)

// PermissionScope defines the Capabilities Controllers can grant for Services
type PermissionScope int

const (
	PermissionScopeUnspecified PermissionScope = iota
	PermissionScopeBasicInfo
	PermissionScopeRecordsRead
	PermissionScopeRecordsWrite
	PermissionScopeTransactionsRead
	PermissionScopeTransactionsWrite
	PermissionScopeWalletsRead
	PermissionScopeWalletsCreate
	PermissionScopeWalletsSubscribe
	PermissionScopeWalletsUpdate
	PermissionScopeTransactionsVerify
	PermissionScopeTransactionsBroadcast
	PermissionScopeAdminUser
	PermissionScopeAdminValidator
)
