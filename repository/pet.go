package repository

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type PetImpl struct{}

type Pet struct {
	ID          int     `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
	Lat         float64 `json:"lat" db:"lat"`
	Lng         float64 `json:"lng" db:"lng"`

	UserID uuid.UUID `json:"user_id" db:"user_id"`

	Status Status `json:"status" db:"status"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type PetWithPic struct {
	Pet
	PictureURLs json.RawMessage `json:"picture_urls" db:"picture_urls"`
}

func (p *PetImpl) List() (pets []PetWithPic, err error) {
	stm := `
	SELECT 
		p.*,
		(
			SELECT COALESCE(json_agg(pic), '[]')
			FROM (
				SELECT picture_url
				FROM "pic_pet" pp
				WHERE pp.pet_id = p.id
			) pic
		) AS picture_urls
	FROM pet p
	`
	err = DB.Select(&pets, stm)
	return pets, err
}

func (p *PetImpl) Create(pet Pet) (Pet, error) {
	stm := `
	INSERT INTO "pet" (name, description, lat, lng)
	VALUES ($1, $2, $3, $4)
	RETURNING *
	`
	err := DB.Get(&pet, stm, pet.Name, pet.Description, pet.Lat, pet.Lng)
	return pet, err
}

func (p *PetImpl) Update(pet Pet) error {
	stm := `
		UPDATE "tag" 
		SET 
			name = :name, 
			description = :description,
			lat = :lat,
			lng = :lng,
			status = :status
		WHERE id = :id
		RETURNING *
	`
	_, err := DB.NamedQuery(stm, pet)
	return err
}
