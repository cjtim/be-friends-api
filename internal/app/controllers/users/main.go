package users

import (
	"net/http"
	"net/url"

	"github.com/cjtim/be-friends-api/configs"
	"github.com/cjtim/be-friends-api/internal/pkg/line"
	"github.com/cjtim/be-friends-api/internal/pkg/users"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Me(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return c.JSON(claims)
}

// func Login(c *fiber.Ctx) error {
// 	user := c.FormValue("user")
// 	pass := c.FormValue("pass")

// 	// Throws Unauthorized error
// 	if user != "john" || pass != "doe" {
// 		return c.SendStatus(fiber.StatusUnauthorized)
// 	}

// 	userInfo := users.User{
// 		ID:            "1",
// 		Name:          "john",
// 		ProfilePic:    "",
// 		LoginMethodID: 1,
// 	}
// 	_, t, err := userInfo.GetNewToken()
// 	if err != nil {
// 		return c.SendStatus(http.StatusInternalServerError)
// 	}
// 	return c.JSON(fiber.Map{"token": t})
// }

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

func LineCallback(c *fiber.Ctx) error {
	code := c.Query("code", "")
	if code == "" {
		return c.SendStatus(http.StatusBadRequest)
	}
	token, err := line.GetJWT(code)
	if err != nil {
		return err
	}

	profile, err := line.GetProfile(token)
	if err != nil {
		return err
	}

	userInfo := users.User{
		ID:            profile.LineUid,
		Name:          profile.Name,
		ProfilePic:    profile.Picture,
		LoginMethodID: 2,
	}

	_, t, err := userInfo.GetNewToken()
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	c.Cookie(&fiber.Cookie{
		Name:  configs.Config.JWTCookies,
		Value: t,
	})

	return c.Redirect(configs.Config.LoginSuccessURL)
}
