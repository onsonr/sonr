package cmd

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sonr-hq/sonr/app"
)

func initSDKConfig() {
	// Set and seal config
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(app.AccountAddressPrefix, "")
	config.SetBech32PrefixForValidator(app.AccountAddressPrefix, "")
	config.SetBech32PrefixForConsensusNode(app.AccountAddressPrefix, "")
	config.Seal()
}
