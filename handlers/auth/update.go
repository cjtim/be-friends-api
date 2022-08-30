package auth

import (
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
)

func UpdateUser(c *fiber.Ctx) error {
	user := repository.User{}
	err := c.BodyParser(&user)
	if err != nil {
		return err
	}
	err = repository.UserRepo.UpdateUser(user)
	if err != nil {
		return err
	}
	return c.JSON(user)
}
