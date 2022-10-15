package repository

import "github.com/google/uuid"

type InterestedImpl struct{}

type Interested struct {
	PetID  int       `json:"pet_id" db:"pet_id"`
	UserID uuid.UUID `json:"user_id" db:"user_id"`
	Step   string    `json:"step" db:"step"`
}

type InterestedPet struct {
	Interested
	Pet
}

type InterestedUser struct {
	Interested
	User
}

func (i *InterestedImpl) ListByUserID(userID uuid.UUID) (output []InterestedPet, err error) {
	stm := `SELECT * FROM "interested" l 
			INNER JOIN "pet" p 
			ON l.pet_id = p.id 
			WHERE l.user_id = $1`
	err = DB.Select(&output, stm, userID)
	return
}

func (i *InterestedImpl) ListByPetID(petId int) (output []InterestedUser, err error) {
	stm := `SELECT * FROM "interested" l 
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

func (i *InterestedImpl) UpdateStep(petID int, userID uuid.UUID, step string) error {
	stm := `UPDATE "interested"
			SET
				step = $3
			WHERE pet_id = $1 AND user_id = $2`
	_, err := DB.Exec(stm, petID, userID, step)
	return err
}
