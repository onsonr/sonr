package keeper

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/x/registry/types"
)

const SINGLE_ALIAS_OWNER = "single-alias-owner"
const NON_NEGATIVE_SELL_PRICE = "nonnegative-sell-price"

func RegisterInvariants(ir sdk.InvariantRegistry, keeper Keeper) {
	ir.RegisterRoute(types.ModuleName, SINGLE_ALIAS_OWNER,
		SingleAliasOwner(keeper))
	ir.RegisterRoute(types.ModuleName, NON_NEGATIVE_SELL_PRICE,
		NonNegativeAliasSellPrice(keeper))
}

func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		if res, broken := SingleAliasOwner(k)(ctx); broken {
			return res, broken
		}
		return NonNegativeAliasSellPrice(k)(ctx)
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
				SINGLE_ALIAS_OWNER,
				fmt.Sprintf("Following Aliases have multiple owners: %s", msg)),
			broken
	}
}

func NonNegativeAliasSellPrice(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		allWhoIs := k.GetAllWhoIs(ctx)

		faultyRecords := make(map[string]int32)
		for _, whoIs := range allWhoIs {
			for _, alias := range whoIs.GetAlias() {
				aliasAmount := alias.Amount
				aliasName := alias.Name
				if alias.IsForSale && aliasAmount < 0 {
					faultyRecords[aliasName] = aliasAmount
				}
			}
		}
		broken := len(faultyRecords) != 0
		b, err := json.MarshalIndent(faultyRecords, "", "  ")
		msg := ""
		if err != nil {
			msg = fmt.Sprintf("%v", faultyRecords)
		} else {
			msg = fmt.Sprint(string(b))
		}

		return sdk.FormatInvariant(types.ModuleName,
				NON_NEGATIVE_SELL_PRICE,
				fmt.Sprintf("Following aliases are listed for sale but have inapporipriate prices: %s", msg)),
			broken
	}
}
