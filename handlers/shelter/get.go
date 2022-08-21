package shelter

import (
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetShelters - list all shelters from database
// @Summary		 list all shelters from database
// @Description  list all shelters from database
// @Tags         shelter
// @Produce      json
// @Success      200  {object}  []repository.User
// @Failure      400  {string}  string
// @Failure      500  {string}  string
// @Router       /api/v1/shelter [get]
func GetShelters(c *fiber.Ctx) error {
	orgs, err := repository.UserRepo.GetOrganizations()
	if err != nil {
		return err
	}
	return c.JSON(orgs)
}

// GetShelterById - list shelter by id from database
// @Summary		 list shelter by id from database
// @Description  list shelter by id from database
// @Tags         shelter
// @Produce      json
// @Param		 id		path		string 			true	"shelter id"
// @Success      200  	{object}  	repository.User
// @Failure      400  	{string}  	string
// @Failure      500  	{string}  	string
// @Router       /api/v1/shelter/{id} [get]
func GetShelterById(c *fiber.Ctx) error {
	idStr := c.Params("id", "")
	if idStr == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		return err
	}

	org, err := repository.UserRepo.GetOrganizationById(id)
	if err != nil {
		return err
	}
	return c.JSON(org)
}
