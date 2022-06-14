package tag

import (
	"strconv"

	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
)

func TagUpdate(c *fiber.Ctx) error {
	strID := c.Query("id")
	name := c.Query("name")
	desc := c.Query("description")
	strIsInternal := c.Query("is_internal")

	if strID == "" || strIsInternal == "" {
		return fiber.ErrBadRequest
	}
	id, err := strconv.Atoi(strID)
	if err != nil {
		return fiber.ErrBadRequest
	}
	isInternal, err := strconv.ParseBool(strIsInternal)
	if err != nil {
		return fiber.ErrBadRequest
	}

	tag, err := repository.TagRepo.Update(id, name, &desc, &isInternal)
	if err != nil {
		return err
	}
	return c.JSON(tag)
}
