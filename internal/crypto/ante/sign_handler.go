package ante

import (
	// cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
)

// CmpSignModeHandler Is a struct that implements the `SignModeHandler` interface.
type CmpSignModeHandler struct {
	defaultConfig client.TxConfig
}

// NewCmpSignModeHandler returns a new SignModeHandler which supports CMP based signing
func NewCmpSignModeHandler(defaultConf client.TxConfig) authsigning.SignModeHandler {
	return CmpSignModeHandler{
		defaultConfig: defaultConf,
	}
}

// DefaultMode is the default mode that is to be used with this handler if no
func (h CmpSignModeHandler) DefaultMode() signingtypes.SignMode {
	return signingtypes.SignMode_SIGN_MODE_DIRECT
}

// Modes is the list of modes supporting by this handler
func (h CmpSignModeHandler) Modes() []signingtypes.SignMode {
	return []signingtypes.SignMode{
		signingtypes.SignMode_SIGN_MODE_DIRECT,
		signingtypes.SignMode_SIGN_MODE_DIRECT_AUX,
		signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON,
	}
}

// GetSignBytes returns the sign bytes for the provided SignMode, SignerData and Tx,
func (h CmpSignModeHandler) GetSignBytes(mode signing.SignMode, data authsigning.SignerData, tx sdk.Tx) ([]byte, error) {
	switch mode {
	case signingtypes.SignMode_SIGN_MODE_DIRECT_AUX:

		return nil, fmt.Errorf("not implemented")
	default:
		return h.defaultConfig.SignModeHandler().GetSignBytes(mode, data, tx)
	}
}
