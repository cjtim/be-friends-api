package auth

import (
	"github.com/cjtim/be-friends-api/internal/auth"
	"github.com/gofiber/fiber/v2"
)

// Me - JWT claim infomation
func Me(c *fiber.Ctx) error {
	claims, err := auth.GetUserExtendedFromFiberCtx(c)
	if err != nil {
		return err
	}
	return c.JSON(claims)
}
