package pet

import (
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
)

// PetList - list all pets from database
// @Summary		 list all pets from database
// @Description  list all pets from database
// @Tags         pet
// @Produce      json
// @Success      200  {object}  repository.Pet
// @Failure      400  {string}  string
// @Failure      500  {string}  string
// @Router       /api/v1/pet [get]
func PetList(c *fiber.Ctx) error {
	pets, err := repository.PetRepo.List()
	if err != nil {
		return err
	}
	return c.JSON(pets)
}
