package api

import (
	"encoding/json"

	"github.com/sonr-io/sonr/x/registry/types"
)

func RegisterWhoIs(registerRequest []byte) error {
	var msgCreateWhoIs types.MsgCreateWhoIs
	if err := json.Unmarshal(registerRequest, &msgCreateWhoIs); err != nil {
		return err
	}

}
