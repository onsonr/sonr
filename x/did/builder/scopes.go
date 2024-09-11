package builder

import (
	"github.com/onsonr/sonr/x/did/types"
	"gopkg.in/macaroon-bakery.v2/bakery/checkers"
)

var (
	GenericPermissionScopeStrings = [...]string{
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

	StringToModulePermissionScope = map[string]types.PermissionScope{
		"PERMISSION_SCOPE_UNSPECIFIED":            types.PermissionScope_PERMISSION_SCOPE_UNSPECIFIED,
		"PERMISSION_SCOPE_BASIC_INFO":             types.PermissionScope_PERMISSION_SCOPE_BASIC_INFO,
		"PERMISSION_SCOPE_IDENTIFIERS_EMAIL":      types.PermissionScope_PERMISSION_SCOPE_PERMISSIONS_READ,
		"PERMISSION_SCOPE_IDENTIFIERS_PHONE":      types.PermissionScope_PERMISSION_SCOPE_PERMISSIONS_WRITE,
		"PERMISSION_SCOPE_TRANSACTIONS_READ":      types.PermissionScope_PERMISSION_SCOPE_TRANSACTIONS_READ,
		"PERMISSION_SCOPE_TRANSACTIONS_WRITE":     types.PermissionScope_PERMISSION_SCOPE_TRANSACTIONS_WRITE,
		"PERMISSION_SCOPE_WALLETS_READ":           types.PermissionScope_PERMISSION_SCOPE_WALLETS_READ,
		"PERMISSION_SCOPE_WALLETS_CREATE":         types.PermissionScope_PERMISSION_SCOPE_WALLETS_CREATE,
		"PERMISSION_SCOPE_WALLETS_SUBSCRIBE":      types.PermissionScope_PERMISSION_SCOPE_WALLETS_SUBSCRIBE,
		"PERMISSION_SCOPE_WALLETS_UPDATE":         types.PermissionScope_PERMISSION_SCOPE_WALLETS_UPDATE,
		"PERMISSION_SCOPE_TRANSACTIONS_VERIFY":    types.PermissionScope_PERMISSION_SCOPE_TRANSACTIONS_VERIFY,
		"PERMISSION_SCOPE_TRANSACTIONS_BROADCAST": types.PermissionScope_PERMISSION_SCOPE_TRANSACTIONS_BROADCAST,
		"PERMISSION_SCOPE_ADMIN_USER":             types.PermissionScope_PERMISSION_SCOPE_ADMIN_USER,
		"PERMISSION_SCOPE_ADMIN_VALIDATOR":        types.PermissionScope_PERMISSION_SCOPE_ADMIN_VALIDATOR,
	}
)

func ResolvePermissionScope(scope string) (types.PermissionScope, bool) {
	uriToPrefix := make(map[string]string)
	for _, scope := range GenericPermissionScopeStrings {
		uriToPrefix["https://example.com/auth/"+scope] = scope
	}
	PermissionNamespace := checkers.NewNamespace(uriToPrefix)

	prefix, ok := PermissionNamespace.Resolve("https://example.com/auth/" + scope)
	if !ok {
		return 0, false
	}
	permScope, ok := StringToModulePermissionScope[prefix]
	return permScope, ok
}
