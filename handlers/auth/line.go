package auth

import (
	"fmt"
	"net/http"

	"github.com/cjtim/be-friends-api/configs"
	"github.com/cjtim/be-friends-api/internal/auth"
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// LoginLine - GET line login url
func LoginLine(c *fiber.Ctx) error {
	host := c.Query("host")
	url := auth.GetLoginURL(host)
	if url == "" {
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.Status(http.StatusOK).SendString(url)
}

// LineCallback - Redirect user back to the website they're coming from
func LineCallback(c *fiber.Ctx) error {
	state := c.Query("state", "")
	if state == "" {
		return c.SendStatus(http.StatusBadRequest)
	}

	clientHost, err := repository.RedisCallback.GetCallback(state)
	if err != nil || clientHost == "" {
		zap.L().Error(
			"error redis - cannot get callback by state",
			zap.String("state", state),
			zap.Error(err),
		)
	}
	redirectURL := fmt.Sprintf("http://%s%s?%s", clientHost, configs.Config.LINE_WEB_CALLBACK_PATH, string(c.Request().URI().QueryString()))
	return c.Redirect(redirectURL)
}

// LineGetJwt - After user back to website.
// Exchange code from line to jwt
func LineGetJwt(c *fiber.Ctx) error {
	code := c.Query("code", "")
	if code == "" {
		return c.SendStatus(http.StatusBadRequest)
	}
	state := c.Query("state", "")
	if state == "" {
		return c.SendStatus(http.StatusBadRequest)
	}

	j, jwtToken, err := auth.Callback(code)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	auth.SetCookie(c, jwtToken, j.Claims)
	return c.SendString(jwtToken)
}
