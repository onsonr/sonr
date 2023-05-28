package keeper

import (
	"encoding/base64"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gofiber/fiber/v2"
	"github.com/sonrhq/core/pkg/crypto"
)

func (k Keeper) GatewayCreateAccount(c *fiber.Ctx) error {
	sess, err := k.authenticator.GetSession(c.Context(), c.Cookies("session"))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	ctx := sdk.UnwrapSDKContext(c.Context())
	acc, err := k.CreateAccountForIdentity(ctx, sess.Did, c.Params("name"), crypto.CoinTypeFromName(c.Params("coin_type")))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"success":   true,
		"account":   acc,
	})
}

func (k Keeper) GatewaySignWithAccount(c *fiber.Ctx) error {
	sess, err := k.authenticator.GetSession(c.Context(), c.Cookies("session"))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	ctx := sdk.UnwrapSDKContext(c.Context())
		bz, err := base64.RawStdEncoding.DecodeString(c.Query("message"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	sig, err := k.SignWithIdentity(ctx, sess.Did, c.Params("name"), bz)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"success":   true,
		"signature": base64.RawStdEncoding.EncodeToString(sig),
		"message":   c.Query("message"),
		"address":   c.Params("address"),
	})
}
