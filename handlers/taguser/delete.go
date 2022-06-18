package taguser

import (
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type deleteBody struct {
	UserID uuid.UUID `json:"user_id"`
	TagID  int       `json:"tag_id"`
}

func Delete(c *fiber.Ctx) error {
	body := deleteBody{}
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}
	err = repository.TagUserRepo.DeleteById(body.UserID, body.TagID)
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusOK)
}
