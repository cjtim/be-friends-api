package pet

import (
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
)

type CreateBody struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Lat         float32 `json:"lat"`
	Lng         float32 `json:"lng"`
}

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
	body := CreateBody{}
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}
	pet, err := repository.PetRepo.Create(repository.Pet{
		Name:        body.Name,
		Description: body.Description,
		Lat:         float64(body.Lat),
		Lng:         float64(body.Lng),
	})

	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(&pet)
}
