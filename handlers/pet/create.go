package pet

import "github.com/gofiber/fiber/v2"

type createBody struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	lat         int     `json:"lat"`
	lng         int     `json:"lng"`
	images      []byte  `json:"images"`
}

func PetCreate(c *fiber.Ctx) error {

}
