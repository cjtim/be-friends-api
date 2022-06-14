package tag

import (
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
)

func TagCreate(c *fiber.Ctx) error {
	body := repository.Tag{}
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}

	tag, err := repository.TagRepo.Create(body.Name, body.Description, body.IsInternal)
	if err != nil {
		return err
	}
	return c.JSON(tag)
}
