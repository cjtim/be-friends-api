package middlewares

import (
	"github.com/cjtim/be-friends-api/configs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Cors() fiber.Handler {
	return cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     configs.Config.CORS_ALLOW_ORIGINS,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "*",
		AllowCredentials: true,
		ExposeHeaders:    "*",
		MaxAge:           0,
	})
}
