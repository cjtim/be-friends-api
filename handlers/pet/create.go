package pet

import (
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
)

type createBody struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Lat         int     `json:"lat"`
	Lng         int     `json:"lng"`
	Images      []byte  `json:"images"`
}

func PetCreate(c *fiber.Ctx) error {
	body := createBody{}
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}
	err = repository.PetRepo.Create(repository.Pet{
		Name:        body.Name,
		Description: body.Description,
		Lat:         float64(body.Lat),
		Lng:         float64(body.Lng),
	})

	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusCreated)
}
