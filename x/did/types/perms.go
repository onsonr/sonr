package types

var (
	PermissionScopeStrings = [...]string{
		"profile",
		"permissions.read",
		"permissions.write",
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
		"PERMISSION_SCOPE_PROFILE_NAME":           PermissionScope_PERMISSION_SCOPE_BASIC_INFO,
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
