package like

import (
	"strconv"

	"github.com/cjtim/be-friends-api/internal/auth"
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
)

func Add(c *fiber.Ctx) error {
	idStr := c.Params("pet_id", "")
	if idStr == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	claims, err := auth.GetUserExtendedFromFiberCtx(c)
	if err != nil {
		return err
	}

	// TODO: List database
	err = repository.LikedRepo.Add(id, claims.ID)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).SendString(string(rune(id)))
}
