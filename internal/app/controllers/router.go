package controller

import (
	"github.com/cjtim/be-friends-api/internal/app/controllers/auth"
	"github.com/cjtim/be-friends-api/internal/app/middlewares"
	"github.com/gofiber/fiber/v2"
)

// Route for all api request
func Route(r *fiber.App) {
	r.Get("/api/v1", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"msg": "Hello, world"})
	})
	r.Get("/api/v1/health", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
	authRouteSetup(r)
}

func authRouteSetup(r *fiber.App) {
	authRoute := r.Group("/api/v1/auth")
	authRoute.Get("/me", middlewares.GetJWTMiddleware, auth.Me)
	authRoute.Get("/line", auth.LoginLine)
	authRoute.Get("/line/jwt", auth.LineCallback)
}
