package handlers

import (
	"github.com/cjtim/be-friends-api/handlers/auth"
	"github.com/cjtim/be-friends-api/handlers/middlewares"
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
)

// Route for all api request
func Route(r *fiber.App) {
	v1 := r.Group("/api/v1")
	v1.Get("", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"msg": "Hello, world"})
	})
	v1.Get("/ready", func(c *fiber.Ctx) error {
		err := repository.IsRedisReady()
		if err != nil {
			return err
		}
		err = repository.DB.Ping()
		if err != nil {
			return err
		}
		return c.SendStatus(fiber.StatusOK)
	})
	v1.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	authRoute := v1.Group("/auth")
	authRoute.Get("/me", middlewares.GetJWTMiddleware, auth.Me)
	authRoute.Get("/logout", auth.Logout)
	authRoute.Post("/login", auth.AuthLogin)
	authRoute.Post("/register", auth.AuthRegister)
	authRoute.Get("/line", auth.LoginLine)
	authRoute.Get("/line/callback", auth.LineCallback)
	authRoute.Get("/line/jwt", auth.LineGetJwt)
}
