package handlers

import (
	"github.com/cjtim/be-friends-api/configs"
	"github.com/cjtim/be-friends-api/handlers/auth"
	"github.com/cjtim/be-friends-api/handlers/like"
	interest "github.com/cjtim/be-friends-api/handlers/like"
	"github.com/cjtim/be-friends-api/handlers/middlewares"
	"github.com/cjtim/be-friends-api/handlers/pet"
	pet_img "github.com/cjtim/be-friends-api/handlers/pet/img"
	"github.com/cjtim/be-friends-api/handlers/shelter"
	"github.com/cjtim/be-friends-api/handlers/tag"
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
		return c.SendString(configs.Config.VERSION)
	})

	authRoute := v1.Group("/auth")
	{
		authRoute.Get("/me", middlewares.JWTMiddleware, auth.Me)
		authRoute.Post("/me", middlewares.JWTMiddleware, auth.UpdateUser)

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
		tagRoute.Put("/:id", tag.TagUpdate)
		tagRoute.Delete("/:id", tag.TagDelete)
	}

	petRoute := v1.Group("/pet")
	{
		petRoute.Get("", pet.PetList)                                           // list pet
		petRoute.Get("/:pet_id", pet.PetDetails)                                // get pet by :pet_id
		petRoute.Post("", middlewares.JWTMiddleware, pet.PetCreate)             // create pet
		petRoute.Put("", middlewares.JWTMiddleware, pet.PetCreate)              // create pet
		petRoute.Post("/img", middlewares.JWTMiddleware, pet_img.PetFileUpload) // upload pet image
	}

	shelterRoute := v1.Group("/shelter")
	{
		shelterRoute.Get("", shelter.GetShelters)        // shelter list
		shelterRoute.Get("/:id", shelter.GetShelterById) // details
	}

	// TODO: imprement all
	likeRoute := v1.Group("/like")
	{
		likeRoute.Get("", like.List)
		likeRoute.Post("/:pet_id", like.Add)
		likeRoute.Delete("/:pet_id", like.Remove)
	}

	interestRoute := v1.Group("/interest")
	{
		interestRoute.Get("", interest.List)
		interestRoute.Post("/:pet_id", interest.Add)
		interestRoute.Delete("/:pet_id", interest.Remove)
	}
}
