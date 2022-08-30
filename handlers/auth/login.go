package auth

import (
	"net/http"

	"github.com/cjtim/be-friends-api/internal/auth"
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type loginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @Summary		 Register with email password
// @Description  Register with email password
// @Tags         auth
// @Produce      json
// @accept 		 json
// @Param        body	body	repository.User	true	"Register body"
// @Success      200  {string}  string "JWT token"
// @Failure      400  {string}  string
// @Failure      500  {string}  string
// @Router       /api/v1/auth/register [post]
func AuthRegister(c *fiber.Ctx) error {
	register := repository.User{}
	err := c.BodyParser(&register)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Cannot parse body")
	}

	// Save to DB
	newUser, err := auth.CreateOrgEmailPassword(register)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"CS_user__email\"" {
			return c.Status(http.StatusBadRequest).SendString("User already registed")
		}
		return c.SendStatus(http.StatusInternalServerError)
	}

	// New JWT token
	j, token, err := auth.GetNewToken(newUser.ID)
	if err != nil {
		zap.L().Error("cannot generate jwt token", zap.Error(err))
		return c.SendStatus(http.StatusInternalServerError)
	}
	auth.SetCookie(c, token, j.Claims)
	return c.Status(http.StatusOK).SendString(token)
}

// @Summary		 Login email password
// @Description  Login email password
// @Tags         auth
// @Produce      json
// @accept 		 json
// @Param        body	body	loginBody	true	"Login body"
// @Success      200  {string}  string "JWT token"
// @Failure      400  {string}  string
// @Failure      500  {string}  string
// @Router       /api/v1/auth/login [post]
func AuthLogin(c *fiber.Ctx) error {
	credential := loginBody{}
	err := c.BodyParser(&credential)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Cannot parse body")
	}
	u, err := auth.Login(credential.Email, credential.Password)
	if err != nil {
		zap.L().Error("cannot login", zap.Error(err))
		return c.SendStatus(http.StatusBadRequest)
	}

	// New JWT token
	j, token, err := auth.GetNewToken(u.ID)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	auth.SetCookie(c, token, j.Claims)
	return c.Status(http.StatusOK).SendString(token)
}
