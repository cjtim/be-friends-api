package repository

type TagPetImpl struct{}

type TagPet struct {
	PetID int `json:"pet_id" db:"pet_id"`
	TagID int `json:"tag_id" db:"tag_id"`
}
