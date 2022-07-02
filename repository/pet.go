package repository

import "time"

type PetImpl struct{}

type Pet struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description *string   `json:"description" db:"description"`
	Lat         float64   `json:"lat" db:"lat"`
	Lng         float64   `json:"lng" db:"lng"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type PetWithPic struct {
	Pet
	PictureURLs string `json:"picture_urls" db:"picture_urls"`
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
