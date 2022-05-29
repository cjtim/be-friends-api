package middlewares

import (
	"github.com/cjtim/be-friends-api/configs"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

var GetJWTMiddleware = func(c *fiber.Ctx) error {
	jwt := c.Cookies(configs.Config.JWTCookies)
	c.Request().Header.Add("Authorization", "Bearer "+jwt)
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(configs.Config.JWTSecret),
	})(c)
}
