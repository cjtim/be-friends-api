package repository

import "time"

type TagImpl struct{}

type Tag struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description *string   `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func (t *TagImpl) List() (tags []Tag, err error) {
	stm := `SELECT id, name, description, created_at, updated_at FROM "tag"`
	err = DB.Select(&tags, stm)
	return
}

func (t *TagImpl) Create(name string, description *string) (tag Tag, err error) {
	stm := `
		INSERT INTO "tag" (name, description) 
		VALUES ($1, $2)
		RETURNING *
	`
	row := DB.QueryRowx(stm, name, description)
	err = row.StructScan(&tag)
	return
}
