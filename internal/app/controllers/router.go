package controller

import (
	"github.com/cjtim/be-friends-api/internal/app/controllers/users"
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
	usersRouteSetup(r)
}

func usersRouteSetup(r *fiber.App) {
	usersRoute := r.Group("/users")
	usersRoute.Get("/me", middlewares.GetJWTMiddleware, users.Me)
	usersRoute.Post("/login", users.Login)
}
