package middleware

import (
	"encoding/json"
	"fmt"

	"github.com/sonr-io/sonr/internal/local"
	domaintypes "github.com/sonr-io/sonr/x/domain/types"
)

// CheckAliasAvailable checks if an alias is already registered
func CheckAliasAvailable(email string) (bool, error) {
	endpoint := fmt.Sprintf("%s/core/domain/username/%s", baseAPIUrl, domaintypes.EmailIndex(email))
	resp := new(domaintypes.QueryGetUsernameRecordsResponse)
	bz, err := local.GetJSON(endpoint)
	if err != nil {
		return true, nil
	}
	err = json.Unmarshal(bz, resp)
	if err != nil {
		return true, err
	}
	return resp.UsernameRecords.Address == "", nil
}

// CheckAliasUnavailable checks if an alias is already registered
func CheckAliasUnavailable(email string) (bool, error) {
	endpoint := fmt.Sprintf("%s/core/domain/username/%s", baseAPIUrl, domaintypes.EmailIndex(email))
	resp := new(domaintypes.QueryGetUsernameRecordsResponse)
	bz, err := local.GetJSON(endpoint)
	if err != nil {
		return false, nil
	}
	err = json.Unmarshal(bz, resp)
	if err != nil {
		return false, err
	}
	return resp.UsernameRecords.Address != "", nil
}
