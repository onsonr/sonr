package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sonrhq/sonr/internal/crypto"
	identitytypes "github.com/sonrhq/sonr/x/identity/types"
)

// AuthenticationResult struct is defining a data structure that represents the response returned after a successful authentication process. It contains various fields such as `Account`, `Address`, `Alias`, `DID`, `DIDDocument`, and `JWT`, which hold information related to
// the authenticated user.
type AuthenticationResult struct {
	Account     *crypto.AccountData        `json:"account"`
	Address     string                     `json:"address"`
	Alias       string                     `json:"alias"`
	DID         string                     `json:"did"`
	DIDDocument *identitytypes.DIDDocument `json:"did_document"`
	JWT         string                     `json:"jwt"`
}

// ClaimsResult struct is defining a data structure used to allocate a challenge and random wallet address to a user.
type ClaimsResult struct {
	Challenge  string `json:"challenge"`
	UCWAddress string `json:"ucw_address"`
}

// TxResponse is a wrapper struct for the response returned after a transaction is broadcasted to the blockchain.
type TxResponse = sdk.TxResponse
