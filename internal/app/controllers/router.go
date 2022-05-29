package controller

import (
	"github.com/cjtim/be-friends-api/internal/app/controllers/auth"
	"github.com/cjtim/be-friends-api/internal/app/middlewares"
	"github.com/gofiber/fiber/v2"
)

// Route for all api request
func Route(r *fiber.App) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"msg": "Hello, world"})
	})
	r.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
	// r.Post("/line/webhook", line_controllers.Webhook)
	authRouteSetup(r)
}

func authRouteSetup(r *fiber.App) {
	authRoute := r.Group("/auth")
	authRoute.Get("/me", middlewares.GetJWTMiddleware, auth.Me)
	authRoute.Get("/line", auth.LoginLine)
	authRoute.Get("/line/callback", auth.LineCallback)
}
