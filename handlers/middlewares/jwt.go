package middlewares

import (
	"net/http"
	"strings"

	"github.com/cjtim/be-friends-api/configs"
	"github.com/cjtim/be-friends-api/internal/utils"
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

var GetJWTMiddleware = func(c *fiber.Ctx) error {
	headers := utils.HeadersToMapStr(c)
	authorization := headers[configs.AuthorizationHeader]
	token := strings.Replace(authorization, "Bearer ", "", 1)
	// If Authorization header is not set
	// Use cookie instead
	if token == "" {
		jwt := c.Cookies(configs.Config.JWTCookies)
		c.Request().Header.Add(configs.AuthorizationHeader, "Bearer "+jwt)
		token = jwt
	}
	// Check token in Redis
	notValid := !repository.RedisJwt.IsJwtValid(token)
	if notValid {
		return c.Status(http.StatusBadRequest).SendString("Invalid JWT")
	}
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(configs.Config.JWTSecret),
	})(c)
}
