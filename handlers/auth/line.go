package auth

import (
	"net/http"

	users "github.com/cjtim/be-friends-api/internal/jwt"
	"github.com/cjtim/be-friends-api/internal/line"
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
)

// LoginLine - GET line login url
func LoginLine(c *fiber.Ctx) error {
	url := line.GetLoginURL()
	return c.Status(http.StatusOK).SendString(url)
}

// LineCallback - Users being redirect here to register with us
func LineCallback(c *fiber.Ctx) error {
	code := c.Query("code", "")
	if code == "" {
		return c.SendStatus(http.StatusBadRequest)
	}
	state := c.Query("state", "")
	if state == "" {
		return c.SendStatus(http.StatusBadRequest)
	}

	// 1. Get profile from LINE
	token, err := line.GetJWT(code, state)
	if err != nil {
		return err
	}

	profile, err := line.GetProfile(token)
	if err != nil {
		return err
	}

	// 2. Update database
	userDB, err := profile.CreateLineUser()
	if err != nil {
		return err
	}

	// 3. Create JWT
	userInfo := repository.User{
		ID:         userDB.ID,
		Name:       profile.Name,
		Email:      userDB.Email,
		LineUid:    userDB.LineUid,
		PictureURL: userDB.PictureURL,
	}
	_, jwtToken, err := users.GetNewToken(&userInfo)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).SendString(jwtToken)
}
