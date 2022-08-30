package middlewares

import "github.com/gofiber/fiber/v2"

func NoCache() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
		return c.Next()
	}
}
