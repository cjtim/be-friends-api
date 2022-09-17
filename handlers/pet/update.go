package pet

import (
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
)

func PetUpdate(c *fiber.Ctx) error {
	body := CreateBody{}
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}
	err = repository.PetRepo.Update(body.Pet)
	if err != nil {
		return err
	}
	// Add tag
	err = repository.TagPetRepo.DeleteAll(body.Pet.ID)
	if err != nil {
		return err
	}
	for _, v := range body.TagIds {
		err = repository.TagPetRepo.Add(body.Pet.ID, v)
		if err != nil {
			return err
		}
	}
	// repository.PetPicRepo.DeleteAll()
	return c.JSON(body)
}
