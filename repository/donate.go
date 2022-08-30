package repository

import (
	"time"

	"github.com/google/uuid"
)

type DonateImpl struct{}

type Donate struct {
	ID int `json:"id" db:"id"`

	PictureURL  string    `json:"picture_url" db:"picture_url"`
	Description *string   `json:"description" db:"description"`
	Amount      float64   `json:"amount" db:"amount"`
	UserID      uuid.UUID `json:"user_id" db:"user_id"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
