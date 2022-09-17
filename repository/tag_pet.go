package repository

type TagPetImpl struct{}

type TagPet struct {
	PetID int `json:"pet_id" db:"pet_id"`
	TagID int `json:"tag_id" db:"tag_id"`
}

func (i *TagPetImpl) Add(petID int, tagID int) error {
	stm := `INSERT INTO "tag_pet" (pet_id, tag_id) VALUES ($1, $2)`
	_, err := DB.Exec(stm, petID, tagID)
	return err
}
