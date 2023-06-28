package vault

import "github.com/sonrhq/core/x/vault/types"

// AccountInfo is an alias for types.AccountInfo which is a struct which contains information about an account
type AccountInfo = types.AccountInfo

// KeyShare is an alias for types.VaultKeyshare which is a struct which contains information about a key share
type KeyShare = *types.VaultKeyshare
