package auth

import (
	"net/http"

	"github.com/cjtim/be-friends-api/configs"
	"github.com/cjtim/be-friends-api/internal/pkg/line"
	"github.com/cjtim/be-friends-api/internal/pkg/users"
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

	// 1. Get profile from LINE
	token, err := line.GetJWT(code)
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
	userInfo := repository.Users{
		ID:         userDB.ID,
		Name:       profile.Name,
		Email:      userDB.Email,
		LineUid:    userDB.LineUid,
		PictureURL: userDB.PictureURL,
	}
	_, _, cookie, err := users.GetNewToken(&userInfo)
	if err != nil {
		return err
	}
	c.Cookie(cookie)

	return c.Redirect(configs.Config.LoginSuccessURL)
}
