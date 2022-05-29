package auth

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cjtim/be-friends-api/configs"
	"github.com/cjtim/be-friends-api/internal/pkg/line"
	"github.com/cjtim/be-friends-api/internal/pkg/users"
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
)

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

	// BUG new user always create at login
	createdUser, err := repository.DB.Users.CreateOne(
		repository.Users.Name.Set(profile.Name),
		repository.Users.LoginMethodID.Set(2),
	).Exec(context.Background())

	if err != nil {
		return err
	}

	createdLineUser, err := repository.DB.LineUsers.UpsertOne(
		repository.LineUsers.LineUId.Equals(profile.LineUid),
	).Create(repository.LineUsers.LineUId.Set(profile.LineUid),
		repository.LineUsers.Name.Set(profile.Name),
		repository.LineUsers.User.Link(
			repository.Users.ID.Equals(createdUser.ID),
		),
		repository.LineUsers.ProfilePic.Set(profile.Picture),
	).Update(
		repository.LineUsers.ProfilePic.Set(profile.Picture),
	).Exec(context.Background())

	if err != nil {
		return err
	}

	fmt.Println(createdLineUser)

	return c.Redirect(configs.Config.LoginSuccessURL)
}
