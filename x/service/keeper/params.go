package keeper

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"sonr.io/core/x/service/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams()
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

const (
	// ServiceRecordKeyPrefix is the prefix to retrieve all ServiceRecord
	ServiceRecordKeyPrefix = "ServiceRecord/value/"
)

// ServiceRecordKey returns the store key to retrieve a ServiceRecord from the index fields
func ServiceRecordKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

// cleanServiceDomain removes the url scheme and path from a service origin
func cleanServiceDomain(origin string) string {
	// Remove url scheme
	r := strings.NewReplacer("https://", "", "http://", "")
	origin = r.Replace(origin)

	if strings.Contains(origin, "/") {
		return strings.Split(origin, "/")[0]
	}
	return origin
}
