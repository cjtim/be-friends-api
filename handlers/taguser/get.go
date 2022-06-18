package taguser

import (
	"github.com/cjtim/be-friends-api/internal/auth"
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
)

func Get(c *fiber.Ctx) error {
	user, err := auth.GetUserExtendedFromFiberCtx(c)
	if err != nil {
		return err
	}
	tags, err := repository.TagUserRepo.GetTagsByUserID(user.ID)
	if err != nil {
		return err
	}
	return c.JSON(tags)
}
