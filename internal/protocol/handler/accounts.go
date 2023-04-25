package handler

import (
	"encoding/base64"

	"github.com/gofiber/fiber/v2"
	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/internal/protocol/middleware"
)

func GetAccount(c *fiber.Ctx) error {
	usr, err := middleware.FetchUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	acc, err := usr.GetAccount(c.Params("address", ""))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"success":   true,
		"account":   acc.ToProto(),
		"coin_type": acc.CoinType().Ticker(),
		"address":   c.Params("address"),
	})
}

func ListAccounts(c *fiber.Ctx) error {
	usr, err := middleware.FetchUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	accs, err := usr.ListAccounts()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"success":  true,
		"accounts": accs,
	})
}

func CreateAccount(c *fiber.Ctx) error {
	usr, err := middleware.FetchUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	acc, err := usr.CreateAccount(c.Params("name"), crypto.CoinTypeFromName(c.Params("coin_type")))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"success":   true,
		"account":   acc.ToProto(),
		"coin_type": acc.CoinType().Ticker(),
		"address":   acc.Address(),
	})
}

func SignWithAccount(c *fiber.Ctx) error {
	usr, err := middleware.FetchUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	bz, err := base64.RawStdEncoding.DecodeString(c.Query("message"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	sig, err := usr.Sign(c.Params("address"), bz)
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

func VerifyWithAccount(c *fiber.Ctx) error {
	usr, err := middleware.FetchUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	bz, err := base64.RawStdEncoding.DecodeString(c.Query("message"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	sig, err := base64.RawStdEncoding.DecodeString(c.Query("signature"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	ok, err := usr.Verify(c.Params("address"), bz, sig)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"success":   true,
		"verified":  ok,
		"message":   c.Query("message"),
		"signature": c.Query("signature"),
		"address":   c.Params("address"),
	})
}

func SendTransaction(c *fiber.Ctx) error {
	return nil
}
