package tag

import (
	"strconv"

	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
)

func TagDelete(c *fiber.Ctx) error {
	strID := c.Params("id")
	if strID == "" {
		return fiber.ErrBadRequest
	}
	id, err := strconv.Atoi(strID)
	if err != nil {
		return fiber.ErrBadRequest
	}
	err = repository.TagRepo.Delete(id)
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusOK)
}
