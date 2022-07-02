package auth

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/cjtim/be-friends-api/configs"
	"github.com/cjtim/be-friends-api/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var (
	TOKEN_EXPIRE  = time.Hour * 72
	FiberLocalKey = "user"
)

type CustomClaims struct {
	jwt.RegisteredClaims

	ID         uuid.UUID `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Email      *string   `json:"email" db:"email"`
	Password   *string   `json:"password" db:"password"`
	LineUid    *string   `json:"line_uid" db:"line_uid"`
	PictureURL *string   `json:"picture_url" db:"picture_url"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`

	// Custome fields
	Tags    json.RawMessage `json:"tags" db:"tags"` // not in used
	IsAdmin *bool           `json:"is_admin" db:"is_admin"`
}

func GetUserExtendedFromFiberCtx(c *fiber.Ctx) (*CustomClaims, error) {
	user, ok := c.Locals(FiberLocalKey).(*jwt.Token)
	if ok {
		claims, ok := user.Claims.(*CustomClaims)
		if ok {
			return claims, nil
		}
	}
	return nil, errors.New("cannot get user")
}

func RemoveCookie(c *fiber.Ctx) {
	c.Cookie(&fiber.Cookie{
		Name:    configs.Config.JWTCookies,
		Value:   "",
		Path:    "/",
		Expires: time.Now(),
	})
}

func SetCookie(c *fiber.Ctx, token string, claim jwt.Claims) {
	cliams := claim.(CustomClaims)
	c.Cookie(&fiber.Cookie{
		Name:    configs.Config.JWTCookies,
		Value:   token,
		Path:    "/",
		Expires: cliams.ExpiresAt.Local(),
	})
}

func GetNewToken(userID uuid.UUID) (*jwt.Token, string, error) {
	u, err := repository.UserRepo.GetUserExtended(userID)
	if err != nil {
		return nil, "", err
	}
	// Create the Claims
	claims := CustomClaims{
		ID:         u.ID,
		Name:       u.Name,
		Email:      u.Email,
		LineUid:    u.LineUid,
		PictureURL: u.PictureURL,
		Tags:       *u.Tags,
		UpdatedAt:  u.UpdatedAt,
		CreatedAt:  u.CreatedAt,
		IsAdmin:    u.IsAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TOKEN_EXPIRE)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(configs.Config.JWTSecret))
	if err != nil {
		return token, "", err
	}
	err = repository.RedisJwt.AddJwt(t, TOKEN_EXPIRE)
	return token, t, err
}
