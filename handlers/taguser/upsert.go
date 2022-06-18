package taguser

import (
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type upsertBody struct {
	UserID uuid.UUID `json:"user_id"`
	TagIDs []int     `json:"tag_ids"`
}

func Upsert(c *fiber.Ctx) error {
	body := upsertBody{}
	err := c.BodyParser(&body)
	if err != nil {
		return err
	}
	err = repository.TagUserRepo.Upsert(body.UserID, body.TagIDs)
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusOK)
}
