package repository

import "github.com/google/uuid"

type InterestedImpl struct{}

type Interested struct {
	PetID  int       `json:"pet_id" db:"pet_id"`
	UserID uuid.UUID `json:"user_id" db:"user_id"`
	Pet
}

func (i *InterestedImpl) ListByUserID(userID uuid.UUID) (output []Interested, err error) {
	stm := `SELECT * FROM "interested" l 
			INNER JOIN "pet" p 
			ON l.pet_id = p.id 
			WHERE l.user_id = $1`
	err = DB.Select(&output, stm, userID)
	return
}

func (i *InterestedImpl) ListByPetID(petId int) (output []User, err error) {
	stm := `SELECT u.* FROM "interested" l 
			INNER JOIN "user" u 
			ON u.id = l.user_id 
			WHERE l.pet_id = $1`
	err = DB.Select(&output, stm, petId)
	return
}

func (i *InterestedImpl) Add(petID int, userID uuid.UUID) error {
	stm := `INSERT INTO "interested" (pet_id, user_id) VALUES ($1, $2)`
	_, err := DB.Exec(stm, petID, userID)
	return err
}

func (i *InterestedImpl) Remove(petID int, userID uuid.UUID) error {
	stm := `DELETE FROM "interested" WHERE pet_id = $1 AND user_id = $2`
	_, err := DB.Exec(stm, petID, userID)
	return err
}
