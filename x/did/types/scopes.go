package types

import "gopkg.in/macaroon-bakery.v2/bakery/checkers"

var (
	PermissionScopeStrings = [...]string{
		"profile.name",
		"identifiers.email",
		"identifiers.phone",
		"transactions.read",
		"transactions.write",
		"wallets.read",
		"wallets.create",
		"wallets.subscribe",
		"wallets.update",
		"transactions.verify",
		"transactions.broadcast",
		"admin.user",
		"admin.validator",
	}

	StringToPermissionScope = map[string]PermissionScope{
		"PERMISSION_SCOPE_UNSPECIFIED":            PermissionScope_PERMISSION_SCOPE_UNSPECIFIED,
		"PERMISSION_SCOPE_BASIC_INFO":             PermissionScope_PERMISSION_SCOPE_BASIC_INFO,
		"PERMISSION_SCOPE_IDENTIFIERS_EMAIL":      PermissionScope_PERMISSION_SCOPE_PERMISSIONS_READ,
		"PERMISSION_SCOPE_IDENTIFIERS_PHONE":      PermissionScope_PERMISSION_SCOPE_PERMISSIONS_WRITE,
		"PERMISSION_SCOPE_TRANSACTIONS_READ":      PermissionScope_PERMISSION_SCOPE_TRANSACTIONS_READ,
		"PERMISSION_SCOPE_TRANSACTIONS_WRITE":     PermissionScope_PERMISSION_SCOPE_TRANSACTIONS_WRITE,
		"PERMISSION_SCOPE_WALLETS_READ":           PermissionScope_PERMISSION_SCOPE_WALLETS_READ,
		"PERMISSION_SCOPE_WALLETS_CREATE":         PermissionScope_PERMISSION_SCOPE_WALLETS_CREATE,
		"PERMISSION_SCOPE_WALLETS_SUBSCRIBE":      PermissionScope_PERMISSION_SCOPE_WALLETS_SUBSCRIBE,
		"PERMISSION_SCOPE_WALLETS_UPDATE":         PermissionScope_PERMISSION_SCOPE_WALLETS_UPDATE,
		"PERMISSION_SCOPE_TRANSACTIONS_VERIFY":    PermissionScope_PERMISSION_SCOPE_TRANSACTIONS_VERIFY,
		"PERMISSION_SCOPE_TRANSACTIONS_BROADCAST": PermissionScope_PERMISSION_SCOPE_TRANSACTIONS_BROADCAST,
		"PERMISSION_SCOPE_ADMIN_USER":             PermissionScope_PERMISSION_SCOPE_ADMIN_USER,
		"PERMISSION_SCOPE_ADMIN_VALIDATOR":        PermissionScope_PERMISSION_SCOPE_ADMIN_VALIDATOR,
	}
)

func ResolvePermissionScope(scope string) (PermissionScope, bool) {
	uriToPrefix := make(map[string]string)
	for _, scope := range PermissionScopeStrings {
		uriToPrefix["https://example.com/auth/"+scope] = scope
	}
	PermissionNamespace := checkers.NewNamespace(uriToPrefix)

	prefix, ok := PermissionNamespace.Resolve("https://example.com/auth/" + scope)
	if !ok {
		return 0, false
	}
	permScope, ok := StringToPermissionScope[prefix]
	return permScope, ok
}
