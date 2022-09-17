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
	Liked       json.RawMessage `json:"liked" db:"liked"`
	Interested  json.RawMessage `json:"interested" db:"interested"`
	Tags        json.RawMessage `json:"tags" db:"tags"`
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
		) AS picture_urls,
		(
			SELECT COALESCE(json_agg(user_id), '[]')
			FROM (
				SELECT user_id
				FROM "liked" lk
				WHERE lk.pet_id = p.id
			) liked
		) AS liked,
		(
			SELECT COALESCE(json_agg(user_id), '[]')
			FROM (
				SELECT user_id
				FROM "interested" itrt
				WHERE itrt.pet_id = p.id
			) itrt
		) AS interested,
		(
			SELECT COALESCE(json_agg(tags), '[]')
			FROM (
				SELECT t.*
				FROM "tag_pet" tp
				JOIN "tag" t
				ON t.id = tp.tag_id
				WHERE tp.pet_id = p.id
			) tags
		) AS tags
	FROM pet p
	`
	err = DB.Select(&pets, stm)
	return pets, err
}

func (p *PetImpl) ListByUserId(userId uuid.UUID) (pets []PetWithPic, err error) {
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
		) AS picture_urls,
		(
			SELECT COALESCE(json_agg(user_id), '[]')
			FROM (
				SELECT user_id
				FROM "liked" lk
				WHERE lk.pet_id = p.id
			) liked
		) AS liked,
		(
			SELECT COALESCE(json_agg(user_id), '[]')
			FROM (
				SELECT user_id
				FROM "interested" itrt
				WHERE itrt.pet_id = p.id
			) itrt
		) AS interested,
		(
			SELECT COALESCE(json_agg(tags), '[]')
			FROM (
				SELECT t.*
				FROM "tag_pet" tp
				JOIN "tag" t
				ON t.id = tp.tag_id
				WHERE tp.pet_id = p.id
			) tags
		) AS tags
	FROM pet p
	WHERE p.user_id = $1
	`
	err = DB.Select(&pets, stm, userId)
	return pets, err
}

func (p *PetImpl) GetById(id int) (pet PetWithPic, err error) {
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
		) AS picture_urls,
		(
			SELECT COALESCE(json_agg(user_id), '[]')
			FROM (
				SELECT user_id
				FROM "liked" lk
				WHERE lk.pet_id = p.id
			) liked
		) AS liked,
		(
			SELECT COALESCE(json_agg(user_id), '[]')
			FROM (
				SELECT user_id
				FROM "interested" itrt
				WHERE itrt.pet_id = p.id
			) itrt
		) AS interested,
		(
			SELECT COALESCE(json_agg(tags), '[]')
			FROM (
				SELECT t.*
				FROM "tag_pet" tp
				JOIN "tag" t
				ON t.id = tp.tag_id
				WHERE tp.pet_id = p.id
			) tags
		) AS tags
	FROM pet p
	WHERE id = $1
	`
	err = DB.Get(&pet, stm, id)
	return pet, err
}

func (p *PetImpl) Create(pet Pet) (Pet, error) {
	stm := `
	INSERT INTO "pet" (name, description, lat, lng, user_id, status)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING *
	`
	err := DB.Get(&pet, stm, pet.Name, pet.Description, pet.Lat, pet.Lng, pet.UserID, pet.Status)
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
