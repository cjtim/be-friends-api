package pet

import (
	"github.com/cjtim/be-friends-api/internal/auth"
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
)

// PetCreate - Create pet
// @Summary		 Create pet
// @Description  Create pet
// @Tags         pet
// @Produce      json
// @accept 		 json
// @Security 	Bearer
// @Param		 body 	body	 	pet.CreateBody			true	"Pet details"
// @Success      200  	{object}  	repository.Pet
// @Failure      400  	{string}  	string
// @Failure      500  	{string}  	string
// @Router       /api/v1/pet [post]
func PetCreate(c *fiber.Ctx) error {
	body := repository.Pet{}
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	claims, err := auth.GetUserExtendedFromFiberCtx(c)
	if err != nil {
		return err
	}
	pet, err := repository.PetRepo.Create(repository.Pet{
		Name:        body.Name,
		Description: body.Description,
		Lat:         float64(body.Lat),
		Lng:         float64(body.Lng),
		Status:      repository.NEW,
		UserID:      claims.ID,
	})

	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(&pet)
}
