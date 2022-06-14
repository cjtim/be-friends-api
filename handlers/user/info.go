package user

import (
	"database/sql"

	"github.com/cjtim/be-friends-api/internal/user"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UserInfo(c *fiber.Ctx) error {
	strId := c.Query("id")
	if strId == "" {
		return fiber.ErrBadRequest
	}
	userId, err := uuid.Parse(strId)
	if err != nil {
		return fiber.ErrBadRequest
	}
	tags, err := user.GetUserInfoById(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return fiber.ErrBadRequest
		}
		return err
	}
	return c.JSON(tags)
}
