package keeper

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/x/bucket/types"
)

func RegisterInvariants(ir sdk.InvariantRegistry, keeper Keeper) {
	ir.RegisterRoute(types.ModuleName, "test",
		SingleAliasOwner(keeper))
}

func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		if res, broken := SingleAliasOwner(k)(ctx); broken {
			return res, broken
		}

		return "Every invariant condition is fulfilled correctly", true
	}
}

func SingleAliasOwner(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		allWhoIs := k.GetAllWhoIs(ctx)

		records := make(map[string][]string)
		for _, whoIs := range allWhoIs {
			owner := whoIs.Owner
			for _, alias := range whoIs.GetAlias() {
				aliasName := alias.Name
				records[aliasName] = append(records[aliasName], owner)
			}
		}
		faultyRecords := make(map[string][]string)
		count := 0
		for aliasName, owners := range records {
			if len(owners) > 1 {
				faultyRecords[aliasName] = owners
				count += 1
			}
		}
		broken := count != 0
		b, err := json.MarshalIndent(faultyRecords, "", "  ")
		msg := ""
		if err != nil {
			msg = fmt.Sprintf("%v", faultyRecords)
		} else {
			msg = fmt.Sprint(string(b))
		}

		return sdk.FormatInvariant(types.ModuleName,
				"single-alias-owner",
				fmt.Sprintf("Following Aliases have multiple owners: %s", msg)),
			broken
	}
}
