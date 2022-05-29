package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// Me - JWT claim infomation
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
