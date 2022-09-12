package pet

import (
	"strconv"

	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	userId := c.Query("user_id")
	if userId != "" {
		uid, err := uuid.Parse(userId)
		if err != nil {
			return err
		}
		pets, err := repository.PetRepo.ListByUserId(uid)
		if err != nil {
			return err
		}
		return c.JSON(pets)
	}
	pets, err := repository.PetRepo.List()
	if err != nil {
		return err
	}
	return c.JSON(pets)
}

func PetDetails(c *fiber.Ctx) error {
	idStr := c.Params("pet_id", "")
	if idStr == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	pet, err := repository.PetRepo.GetById(id)
	if err != nil {
		return err
	}

	return c.JSON(pet)
}
