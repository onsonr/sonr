package vfs

import (
	modulev1 "github.com/sonrhq/sonr/api/identity/module/v1"
)

//  Sonr Vault File Structure
// ───────┐
//        ├──────▶ /credentials.json - 'Serialized list of accumulators mapped to did methods'
//        ├──────▶ /wallets.json - 'Serialized list of wallets'
//        ├──────▶ /index.html - 'Self-recovery and user portal'
//        └──────┐  /.keyshares/* - 'Serialized bytes of the validator/user encrypted mpc shares'
//               ├────────▶ /{did}:{coin_method}:{coin_network}:{wallet_address}/public.share
//               └────────▶ /{did}:{coin_method}:{coin_network}:{wallet_address}/private.{idxval123abc}.share

// Wallets struct contains account details structured in this HD Wallet
type Wallets struct {
	Accounts []Account `json:"accounts"`
}

// Account struct represents an individual account in the wallet
type Account struct {
	Address     string            `json:"address"`
	CoinType    modulev1.CoinType `json:"coinType"`
	CoinNetwork string            `json:"coinNetwork"`
	PublicKey   []byte            `json:"publicKey"`
}

// Credentials keeps the JSON file containing a list of accumulated credentials
type Credentials struct {
	Accumulators map[string]string `json:"accumulators"`
	Witnesses    map[string]string `json:"witnesses"`
}
