package auth

import (
	"net/http"

	"github.com/cjtim/be-friends-api/configs"
	"github.com/cjtim/be-friends-api/internal/auth"
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {
	err := repository.RedisJwt.RemoveJwt(c.Cookies(configs.Config.JWTCookies))
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	auth.RemoveCookie(c)
	return c.SendStatus(http.StatusOK)
}
