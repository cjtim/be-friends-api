package interest

import (
	"strconv"

	"github.com/cjtim/be-friends-api/internal/auth"
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
)

func List(c *fiber.Ctx) error {
	claims, err := auth.GetUserExtendedFromFiberCtx(c)
	if err != nil {
		return err
	}
	petIdStr := c.Query("pet_id")
	if petIdStr == "" {
		// List intested by user_id
		interested, err := repository.InterestedRepo.ListByUserID(claims.ID)
		if err != nil {
			return err
		}
		return c.JSON(interested)
	}
	// If pet_id exist
	// Fetch users that interested
	petId, err := strconv.Atoi(petIdStr)
	if err != nil {
		return err
	}
	users, err := repository.InterestedRepo.ListByPetID(petId)
	if err != nil {
		return err
	}
	return c.JSON(users)
}
