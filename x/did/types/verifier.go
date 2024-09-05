package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func (t *Token) Validate(permissions string) error {
	return nil
}

func (t *Token) Verify(msg sdk.Msg) error {
	return nil
}
