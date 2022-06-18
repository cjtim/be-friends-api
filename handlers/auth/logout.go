package auth

import (
	"net/http"
	"time"

	"github.com/cjtim/be-friends-api/configs"
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {
	err := repository.RedisJwt.RemoveJwt(c.Cookies(configs.Config.JWTCookies))
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	c.Cookie(&fiber.Cookie{
		Name:    configs.Config.JWTCookies,
		Value:   "",
		Path:    "/",
		Expires: time.Now(),
	})
	return c.SendStatus(http.StatusOK)
}
