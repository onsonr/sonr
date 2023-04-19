package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sonrhq/core/internal/protocol/middleware"
)

func ReadInboxMessages(c *fiber.Ctx) error {
	usr, err := middleware.FetchUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	msgs, err := usr.ReadMail(c.Params("address"))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"success":  true,
		"messages": msgs,
	})
}

func SendInboxMessage(c *fiber.Ctx) error {
	usr, err := middleware.FetchUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	err = usr.SendMail(c.Params("address"), c.Params("to"), c.Query("message"))
	if err != nil {
		return c.Status(501).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"success": true,
		"to": c.Params("to"),
		"from": c.Params("from"),
		"message": c.Query("message"),
	})
}
