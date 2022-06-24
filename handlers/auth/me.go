package auth

import (
	"github.com/cjtim/be-friends-api/internal/auth"
	"github.com/gofiber/fiber/v2"
)

// Me - JWT claim infomation
// @Summary		 JWT claim infomation
// @Description  JWT claim infomation
// @Tags         auth
// @Produce      json
// @Security 	Bearer
// @Success      200  {object}  auth.CustomClaims
// @Failure      400  {string}  string
// @Failure      500  {string}  string
// @Router       /api/v1/auth/me [get]
func Me(c *fiber.Ctx) error {
	claims, err := auth.GetUserExtendedFromFiberCtx(c)
	if err != nil {
		return err
	}
	return c.JSON(claims)
}
