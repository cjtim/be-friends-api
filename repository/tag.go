package repository

import "time"

type TagImpl struct{}

type Tag struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description *string   `json:"description" db:"description"`
	IsInternal  *bool     `json:"is_internal" db:"is_internal"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func (t *TagImpl) List() (tags []Tag, err error) {
	stm := `SELECT id, name, description, is_internal, created_at, updated_at FROM "tag"`
	err = DB.Select(&tags, stm)
	return
}

func (t *TagImpl) Create(name string, description *string, isInternal *bool) (tag Tag, err error) {
	stm := `
		INSERT INTO "tag" (name, description, is_internal) 
		VALUES ($1, $2, $3)
		RETURNING *
	`
	err = DB.Get(&tag, stm, name, description, isInternal)
	return
}
