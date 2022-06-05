package middlewares

import (
	"net/http"

	"github.com/cjtim/be-friends-api/configs"
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

var GetJWTMiddleware = func(c *fiber.Ctx) error {
	jwt := c.Cookies(configs.Config.JWTCookies)
	c.Request().Header.Add(configs.AuthorizationHeader, "Bearer "+jwt)
	if jwt != "" && !repository.RedisRepo.IsJwtValid(jwt) {
		return c.Status(http.StatusBadRequest).SendString("Invalid JWT")
	}
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(configs.Config.JWTSecret),
	})(c)
}
