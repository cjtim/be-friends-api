package auth

import (
	"net/http"

	"github.com/cjtim/be-friends-api/internal/auth"
	"github.com/gofiber/fiber/v2"
)

type loginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AuthRegister(c *fiber.Ctx) error {
	register := registerBody{}
	err := c.BodyParser(&register)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Cannot parse body")
	}

	// Save to DB
	newUser, err := auth.CreateUserEmailPassword(register.Name, register.Email, register.Password)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	// New JWT token
	_, token, err := auth.GetNewToken(newUser.ID)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).SendString(token)
}

func AuthLogin(c *fiber.Ctx) error {
	credential := loginBody{}
	err := c.BodyParser(&credential)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Cannot parse body")
	}
	u, err := auth.Login(credential.Email, credential.Password)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	// New JWT token
	_, token, err := auth.GetNewToken(u.ID)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).SendString(token)
}
