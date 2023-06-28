package sfs

import (
	"github.com/sonrhq/core/internal/crypto"
	identitytypes "github.com/sonrhq/core/x/identity/types"
	"github.com/sonrhq/core/x/vault/types"
)



type ClaimAccountResponse struct {
	Alias string `json:"alias"`
	UcwDid string `json:"ucw_did"`
	CoinType crypto.CoinType `json:"coin_type"`
	Address string `json:"address"`
	Account *types.AccountInfo `json:"account"`
	DIDDocument *identitytypes.DIDDocument `json:"did_document"`
	JWT string `json:"jwt"`
}

type UnlockAccountResponse struct {
	Did string `json:"ucw_did"`
	Account *types.AccountInfo `json:"account"`
	JWT string `json:"jwt"`
}

type JWTClaims struct {
	Did string `json:"did"`
	Credential []byte `json:"credential"`
	ExpiresAt int64 `json:"expires_at"`
}
