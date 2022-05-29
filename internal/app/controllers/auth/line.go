package auth

import (
	"net/http"
	"net/url"

	"github.com/cjtim/be-friends-api/configs"
	"github.com/cjtim/be-friends-api/internal/pkg/line"
	"github.com/cjtim/be-friends-api/internal/pkg/users"
	"github.com/gofiber/fiber/v2"
)

// LoginLine - GET line login url
func LoginLine(c *fiber.Ctx) error {
	url := http.Request{
		URL: &url.URL{
			Scheme: "https",
			Host:   "access.line.me",
			Path:   "/oauth2/v2.1/authorize",
		},
	}
	q := url.URL.Query()
	q.Add("state", "12345abcde")
	q.Add("scope", "profile openid")
	q.Add("response_type", "code")
	q.Add("redirect_uri", configs.Config.LineLoginCallback)
	q.Add("client_id", configs.Config.LineClientID)
	url.URL.RawQuery = q.Encode()
	return c.Status(http.StatusOK).SendString(url.URL.String())
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

	// 2. Create JWT
	userInfo := users.User{
		ID:            profile.LineUid,
		Name:          profile.Name,
		ProfilePic:    profile.Picture,
		LoginMethodID: 2,
	}
	_, _, cookie, err := userInfo.GetNewToken()
	if err != nil {
		return err
	}
	c.Cookie(cookie)

	// 3. Update database
	err = profile.CreateLineUser()
	if err != nil {
		return err
	}

	return c.Redirect(configs.Config.LoginSuccessURL)
}
