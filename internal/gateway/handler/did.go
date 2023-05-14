package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sonrhq/core/internal/local"
)

func GetDID(c *fiber.Ctx) error {
	doc, err := local.Context().GetDID(c.Context(), c.Params("did"))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"did":      c.Params("did"),
		"document": doc,
	})
}

func GetDIDByAlias(c *fiber.Ctx) error {
	doc, _ := local.Context().GetDIDByAlias(c.Context(), c.Params("alias"))
	if doc == nil {
		return c.JSON(fiber.Map{
			"available": true,
			"alias":     c.Params("alias"),
		})
	}
	return c.JSON(fiber.Map{
		"available": false,
		"did":       doc.Id,
		"document":  doc,
		"alias":     c.Params("alias"),
		"address":  doc.FindPrimaryAddress(),
	})
}

func GetDIDByOwner(c *fiber.Ctx) error {
	doc, err := local.Context().GetDIDByOwner(c.Context(), c.Params("owner"))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"did":      doc.Id,
		"document": doc,
		"owner":    c.Params("owner"),
		"alias":    doc.FindUsername(),
	})
}

func ListDIDs(c *fiber.Ctx) error {
	docs, err := local.Context().GetAllDIDs(c.Context())
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"documents": docs,
		"count":     len(docs),
	})
}

func GetOldestUnclaimed(c *fiber.Ctx) error {
	uw, err := local.Context().OldestUnclaimedWallet(c.Context())
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"success": true,
		"data": uw,
	})
}

func ListAllUnclaimed(c *fiber.Ctx) error {
	uw, err := local.Context().GetUnclaimedWallets(c.Context())
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"success": true,
		"data": uw,
	})
}
