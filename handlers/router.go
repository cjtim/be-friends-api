package handlers

import (
	"os"

	"github.com/cjtim/be-friends-api/handlers/auth"
	"github.com/cjtim/be-friends-api/handlers/middlewares"
	"github.com/cjtim/be-friends-api/handlers/pet"
	"github.com/cjtim/be-friends-api/handlers/tag"
	"github.com/cjtim/be-friends-api/handlers/taguser"
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
	v1.Get("/version", func(c *fiber.Ctx) error {
		return c.SendString(os.Getenv("VERSION"))
	})

	authRoute := v1.Group("/auth")
	{
		authRoute.Get("/me", middlewares.JWTMiddleware, auth.Me)
		authRoute.Get("/logout", auth.Logout)
		authRoute.Post("/login", auth.AuthLogin)
		authRoute.Post("/register", auth.AuthRegister)
		authRoute.Get("/line", auth.LoginLine)
		authRoute.Get("/line/callback", auth.LineCallback)
		authRoute.Get("/line/jwt", auth.LineGetJwt)
	}

	tagRoute := v1.Group("/tag")
	{
		tagRoute.Get("", tag.TagList)
		tagRoute.Post("", tag.TagCreate)
		tagRoute.Put("", tag.TagUpdate)
		tagRoute.Delete("/:id", tag.TagDelete)
	}

	tagUserRoute := v1.Group("/tag_user")
	{
		tagUserRoute.Get("", middlewares.JWTMiddleware, taguser.Get)
		tagUserRoute.Post("", taguser.Upsert)
		tagUserRoute.Delete("", taguser.Delete)
	}

	petRoute := v1.Group("/pet")
	{
		petRoute.Get("", pet.PetList)
	}
}
