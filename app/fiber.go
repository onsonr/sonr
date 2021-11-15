package app

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func initFiber() {
	App = fiber.New()

	// GET /api/register
	App.Get("/api/*", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s", c.Params("*"))
		return c.SendString(msg) // => ✋ register
	})
}
