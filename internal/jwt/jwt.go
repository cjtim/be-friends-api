package jwt

import (
	"time"

	"github.com/cjtim/be-friends-api/configs"
	"github.com/cjtim/be-friends-api/repository"
	"github.com/golang-jwt/jwt/v4"
)

var (
	TOKEN_EXPIRE = time.Hour * 72
)

func GetNewToken(u *repository.User) (*jwt.Token, string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"id":          u.ID,
		"name":        u.Name,
		"email":       u.Email,
		"line_uid":    u.LineUid,
		"picture_url": u.PictureURL,
		"exp":         time.Now().Add(TOKEN_EXPIRE).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(configs.Config.JWTSecret))
	if err != nil {
		return token, "", err
	}
	err = repository.RedisRepo.AddJwt(t, TOKEN_EXPIRE)
	return token, t, err
}
