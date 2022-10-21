package interest

import (
	"strconv"
	"strings"

	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

	userId, err := uuid.Parse(c.Query("user_id"))
	if err != nil {
		return err
	}
	step := strings.ToUpper(c.Query("step", "REVIEWING"))

	err = repository.InterestedRepo.UpdateStep(id, userId, step)
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusOK)
}
