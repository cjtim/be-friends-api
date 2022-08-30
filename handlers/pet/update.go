package pet

import (
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
)

func PetUpdate(c *fiber.Ctx) error {
	body := repository.Pet{}
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}
	err = repository.PetRepo.Update(body)
	if err != nil {
		return err
	}
	return c.JSON(body)
}
