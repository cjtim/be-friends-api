package auth

import (
	"net/http"

	"github.com/cjtim/be-friends-api/internal/auth"
	"github.com/gofiber/fiber/v2"
)

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AuthRegister(c *fiber.Ctx) error {
	register := RegisterBody{}
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
	_, token, err := auth.GetNewToken(&newUser)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).SendString(token)
}

func AuthLogin(c *fiber.Ctx) error {
	credential := LoginBody{}
	c.BodyParser(&credential)

	u, err := auth.Login(credential.Email, credential.Password)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	// New JWT token
	_, token, err := auth.GetNewToken(&u)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).SendString(token)
}
