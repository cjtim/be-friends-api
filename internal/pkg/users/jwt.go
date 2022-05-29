package users

import (
	"time"

	"github.com/cjtim/be-friends-api/configs"
	"github.com/golang-jwt/jwt/v4"
)

func (u *User) GetNewToken() (*jwt.Token, string, error) {
	exp := time.Now().Add(time.Hour * 72)
	// Create the Claims
	claims := jwt.MapClaims{
		"id":         u.ID,
		"name":       u.Name,
		"profilePic": u.ProfilePic,
		"exp":        exp.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(configs.Config.JWTSecret))
	if err != nil {
		return token, "", err
	}
	return token, t, err
}
