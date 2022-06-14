package tag

import (
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
)

func TagList(c *fiber.Ctx) error {
	// if is_admin return including internal tags
	tags, err := repository.TagRepo.List()
	if err != nil {
		return err
	}
	return c.JSON(tags)
}
