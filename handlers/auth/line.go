package auth

import (
	"net/http"

	"github.com/cjtim/be-friends-api/internal/auth"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// LoginLine - GET line login url
func LoginLine(c *fiber.Ctx) error {
	url := auth.GetLoginURL()
	return c.Status(http.StatusOK).SendString(url)
}

// LineCallback - Users being redirect here to register with us
func LineCallback(c *fiber.Ctx) error {
	code := c.Query("code", "")
	if code == "" {
		return c.SendStatus(http.StatusBadRequest)
	}
	state := c.Query("state", "")
	if state == "" {
		return c.SendStatus(http.StatusBadRequest)
	}

	jwtToken, err := auth.Callback(code, state)
	if err != nil {
		zap.L().Error("error line callback", zap.Error(err))
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).SendString(jwtToken)
}
