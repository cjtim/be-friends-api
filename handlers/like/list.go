package like

import (
	"github.com/cjtim/be-friends-api/internal/auth"
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
)

func List(c *fiber.Ctx) error {
	claims, err := auth.GetUserExtendedFromFiberCtx(c)
	if err != nil {
		return err
	}
	liked, err := repository.LikedRepo.ListByUserID(claims.ID)
	if err != nil {
		return err
	}
	return c.JSON(liked)
}
