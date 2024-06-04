package vault

import (
	"time"

	"github.com/di-dao/sonr/pkg/cache"
)

var vaultCache *cache.Cache[contextKey, vaultFS]

type contextKey string

func (c contextKey) String() string {
	return "vault context key " + string(c)
}

var clientCtxKey = contextKey("vault-client-id")

func init() {
	// This is a placeholder
	vaultCache = cache.New[contextKey, vaultFS](time.Minute*5, time.Minute)
}
