package repository

import "github.com/google/uuid"

type LikedImpl struct{}

type Liked struct {
	PetID  int       `json:"pet_id" db:"pet_id"`
	UserID uuid.UUID `json:"user_id" db:"user_id"`
	Pet
}

func (i *LikedImpl) ListByUserID(userID uuid.UUID) (liked []Liked, err error) {
	stm := `SELECT * FROM "liked" l 
			INNER JOIN "pet" p 
			ON l.pet_id = p.id 
			WHERE l.user_id = $1`
	err = DB.Select(&liked, stm, userID)
	return
}

func (i *LikedImpl) Add(petID int, userID uuid.UUID) error {
	stm := `INSERT INTO "liked" (pet_id, user_id) VALUES ($1, $2)`
	_, err := DB.Exec(stm, petID, userID)
	return err
}

func (i *LikedImpl) Remove(petID int, userID uuid.UUID) error {
	stm := `DELETE FROM "liked" WHERE pet_id = $1 AND user_id = $2`
	_, err := DB.Exec(stm, petID, userID)
	return err
}
