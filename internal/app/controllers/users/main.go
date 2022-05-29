package users

import (
	"net/http"

	"github.com/cjtim/be-friends-api/internal/pkg/users"
	"github.com/gofiber/fiber/v2"
)

func Me(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

func Login(c *fiber.Ctx) error {
	user := c.FormValue("user")
	pass := c.FormValue("pass")

	// Throws Unauthorized error
	if user != "john" || pass != "doe" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	userInfo := users.User{
		ID:            "1",
		Name:          "john",
		LoginMethodID: 1,
	}
	_, t, err := userInfo.GetNewToken()
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"token": t})
}
