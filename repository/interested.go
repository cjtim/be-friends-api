package repository

import "github.com/google/uuid"

type InterestedImpl struct{}

type Interested struct {
	PetID  int       `json:"pet_id" db:"pet_id"`
	UserID uuid.UUID `json:"user_id" db:"user_id"`
}

func (i *InterestedImpl) ListByUserID(userID uuid.UUID) (output []Interested, err error) {
	stm := `SELECT * FROM "interested" WHERE user_id = $1`
	err = DB.Select(&output, stm, userID)
	return
}

func (i *InterestedImpl) Add(petID int, userID uuid.UUID) error {
	stm := `INSERT INTO "interested" (pet_id, user_id) VALUES ($1, $2)`
	_, err := DB.Exec(stm, petID, userID)
	return err
}

func (i *InterestedImpl) Remove(petID int, userID uuid.UUID) error {
	stm := `DELETE "interested" WHERE pet_id = $1, user_id = $2`
	_, err := DB.Exec(stm, petID, userID)
	return err
}