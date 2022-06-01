package repository

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type UserImpl struct{}

type Users struct {
	ID         uuid.UUID `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Email      *string   `json:"email" db:"email"`
	Password   *string   `json:"password" db:"password"`
	LineUid    *string   `json:"line_uid" db:"line_uid"`
	PictureURL *string   `json:"picture_url" db:"picture_url"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

func (u *UserImpl) List() ([]Users, error) {
	users := []Users{}
	rows := DB.QueryRow(`SELECT * FROM Users`)
	err := rows.Scan(&users)
	if err != nil {
		return []Users{}, err
	}
	return users, err
}

func (u *UserImpl) GetById(id uuid.UUID) (Users, error) {
	stm := `SELECT * FROM "Users" WHERE id = $1`
	row, err := DB.Queryx(stm, id)
	if err != nil {
		return Users{}, nil
	}
	result := Users{}
	err = row.StructScan(&result)
	return result, err
}

func (u *UserImpl) UpsertLine(user Users) (result Users, err error) {
	stmUpdate := `
		UPDATE "Users"
		SET picture_url = :picture_url, updated_at = NOW()
		WHERE line_uid = :line_uid
		RETURNING *
	`
	row, err := DB.NamedQuery(stmUpdate, user)
	if err != nil {
		return
	}
	noUpdatedRow := !row.Next()
	if noUpdatedRow {
		stmInsert := `
			INSERT INTO "Users" (name, line_uid, picture_url)
			VALUES (:name, :line_uid, :picture_url)
			RETURNING *
		`
		row, err = DB.NamedQuery(stmInsert, user)
		if err != nil {
			return
		}
		if row.Next() {
			row.StructScan(&result)
			zap.L().Info("NEW USER LOGIN", zap.String("line_uid", *result.LineUid))
			return result, err
		}
		err = errors.New("cannot get inserted result")
		zap.L().Error("NEW USER LOGIN", zap.String("line_uid", *result.LineUid), zap.Error(err))
		return Users{}, err
	}
	err = row.StructScan(&result)
	zap.L().Info("OLD USER LOGIN", zap.String("line_uid", *result.LineUid))
	return result, err
}
