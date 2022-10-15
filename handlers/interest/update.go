package interest

import (
	"strconv"
	"strings"

	"github.com/cjtim/be-friends-api/internal/auth"
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
)

func UpdateStep(c *fiber.Ctx) error {
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
	step := strings.ToUpper(c.Query("step", "REVIEWING"))

	err = repository.InterestedRepo.UpdateStep(id, claims.ID, step)
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusOK)
}
