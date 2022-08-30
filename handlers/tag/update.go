package tag

import (
	"strconv"

	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
)

func TagUpdate(c *fiber.Ctx) error {
	strID := c.Params("id")
	name := c.Query("name")
	desc := c.Query("description")

	if strID == "" {
		return fiber.ErrBadRequest
	}
	id, err := strconv.Atoi(strID)
	if err != nil {
		return fiber.ErrBadRequest
	}

	tag, err := repository.TagRepo.Update(id, name, &desc)
	if err != nil {
		return err
	}
	return c.JSON(tag)
}
