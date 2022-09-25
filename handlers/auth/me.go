package auth

import (
	"github.com/cjtim/be-friends-api/internal/auth"
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
)

// Me - JWT claim infomation
// @Summary		 JWT claim infomation
// @Description  JWT claim infomation
// @Tags         auth
// @Produce      json
// @Security 	Bearer
// @Success      200  {object}  repository.User
// @Failure      400  {string}  string
// @Failure      500  {string}  string
// @Router       /api/v1/auth/me [get]
func Me(c *fiber.Ctx) error {
	claims, err := auth.GetUserExtendedFromFiberCtx(c)
	if err != nil {
		return err
	}
	user, err := repository.UserRepo.GetUser(claims.ID)
	if err != nil {
		return err
	}
	return c.JSON(user)
}
