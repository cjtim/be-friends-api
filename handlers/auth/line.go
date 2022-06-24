package auth

import (
	"fmt"
	"net/http"

	"github.com/cjtim/be-friends-api/configs"
	"github.com/cjtim/be-friends-api/internal/auth"
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// LoginLine	 GET line login url
// @Summary		 Get LINE login url
// @Description  Get LINE login url and register user's host to redirect back
// @Tags         auth
// @Produce      plain
// @Param        host	query	string	true	"localhost:3000"
// @Success      200  {string}  string	"https://access.line.me/oauth2/v2.1/authorize"
// @Failure      500  {string}  string
// @Router       /api/v1/auth/line [get]
func LoginLine(c *fiber.Ctx) error {
	host := c.Query("host")
	url := auth.GetLoginURL(host)
	if url == "" {
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.Status(http.StatusOK).SendString(url)
}

// @Summary		 Receive callback from LINE and redirect user back to the website they're coming from
// @Description  Redirect user back to the website they're coming from
// @Tags         auth
// @Produce      plain
// @Param        state	query	string	true	"123456abcdef"
// @Success      200  {string}  string	"OK"
// @Header       200  {string}  Location  "https://localhost:3000/user"
// @Failure      400  {string}  string
// @Router       /api/v1/auth/line/callback [get]
func LineCallback(c *fiber.Ctx) error {
	state := c.Query("state", "")
	if state == "" {
		return c.SendStatus(http.StatusBadRequest)
	}

	clientHost, err := repository.RedisCallback.GetCallback(state)
	if err != nil || clientHost == "" {
		zap.L().Error(
			"error redis - cannot get callback by state",
			zap.String("state", state),
			zap.Error(err),
		)
	}
	redirectURL := fmt.Sprintf("http://%s%s?%s", clientHost, configs.Config.LINE_WEB_CALLBACK_PATH, string(c.Request().URI().QueryString()))
	return c.Redirect(redirectURL)
}

// LineGetJwt - After user back to website.
// @Summary		 Exchange code from line to jwt
// @Description  Exchange code from line to jwt
// @Tags         auth
// @Produce      plain
// @Param        state	query	string	true	"123456abcdef"
// @Param        code	query	string	true	"123456abcdef"
// @Success      200  {string}  string	"JWT TOKEN...."
// @Failure      400  {string}  string
// @Failure      500  {string}  string
// @Router       /api/v1/auth/line/jwt [get]
func LineGetJwt(c *fiber.Ctx) error {
	code := c.Query("code", "")
	if code == "" {
		return c.SendStatus(http.StatusBadRequest)
	}
	state := c.Query("state", "")
	if state == "" {
		return c.SendStatus(http.StatusBadRequest)
	}

	j, jwtToken, err := auth.Callback(code)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	auth.SetCookie(c, jwtToken, j.Claims)
	return c.SendString(jwtToken)
}
