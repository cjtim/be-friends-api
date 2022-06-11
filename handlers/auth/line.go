package auth

import (
	"fmt"
	"net/http"
	"time"

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

	jwtToken, err := auth.Callback(code)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	clientHost, err := repository.RedisCallback.GetCallback(state)
	if err != nil || clientHost == "" {
		zap.L().Error(
			"error redis - cannot get callback by state",
			zap.String("jwt", jwtToken),
			zap.String("state", state),
			zap.Error(err),
		)
	}
	redirectURL := fmt.Sprintf("http://%s%s", clientHost, configs.Config.LINE_WEB_CALLBACK_PATH)
	authCookie := fmt.Sprintf(
		"%s=%s; Max-Age=%d; Path=/; SameSite=None; Secure;",
		configs.Config.JWTCookies,
		jwtToken,
		int64(auth.TOKEN_EXPIRE/time.Second),
	)
	c.Response().Header.Add("set-cookie", authCookie)
	return c.Redirect(redirectURL)
}
