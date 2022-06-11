package repository

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type UserImpl struct{}

type User struct {
	ID         uuid.UUID `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Email      *string   `json:"email" db:"email"`
	Password   *string   `json:"password" db:"password"`
	LineUid    *string   `json:"line_uid" db:"line_uid"`
	PictureURL *string   `json:"picture_url" db:"picture_url"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

func (u *UserImpl) List() ([]User, error) {
	users := []User{}
	rows := DB.QueryRow(`SELECT * FROM "user"`)
	err := rows.Scan(&users)
	if err != nil {
		return []User{}, err
	}
	return users, err
}

func (u *UserImpl) GetById(id uuid.UUID) (User, error) {
	stm := `SELECT * FROM "user" WHERE id = $1`
	row, err := DB.Queryx(stm, id)
	if err != nil {
		return User{}, err
	}
	result := User{}
	err = row.StructScan(&result)
	return result, err
}

func (u *UserImpl) UpsertLine(user User) (User, error) {
	result := User{}
	stmUpdate := `
		UPDATE "user"
		SET picture_url = :picture_url, updated_at = NOW()
		WHERE line_uid = :line_uid
		RETURNING *
	`
	row, err := DB.NamedQuery(stmUpdate, user)
	if err != nil {
		return result, err
	}
	noUpdatedRow := !row.Next()
	if noUpdatedRow {
		stmInsert := `
			INSERT INTO "user" (name, line_uid, picture_url)
			VALUES (:name, :line_uid, :picture_url)
			RETURNING *
		`
		row, err = DB.NamedQuery(stmInsert, user)
		if err != nil {
			return result, err
		}
		if row.Next() {
			err = row.StructScan(&result)
			zap.L().Info("NEW USER LOGIN", zap.String("line_uid", *result.LineUid))
			return result, err
		}
		err = errors.New("cannot get inserted result")
		zap.L().Error("NEW USER LOGIN", zap.String("line_uid", *result.LineUid), zap.Error(err))
		return result, err
	}
	err = row.StructScan(&result)
	zap.L().Info("OLD USER LOGIN", zap.String("line_uid", *result.LineUid), zap.Error(err))
	return result, err
}

func (u *UserImpl) RegisterUser(user User) (result User, err error) {
	stmInsert := `
		INSERT INTO "user" (name, email, password)
		VALUES (:name, :email, :password)
		RETURNING *
	`
	rows, err := DB.NamedQuery(stmInsert, user)
	if err != nil {
		zap.L().Error("error register user", zap.String("email", *user.Email), zap.String("name", user.Name))
		return
	}
	if rows.Next() {
		rows.StructScan(&result)
		zap.L().Info("NEW USER register", zap.Any("id", result.ID))
		return result, err
	}
	return User{}, errors.New("error register user - cannot parse inserted row")
}

func (u *UserImpl) GetUserByEmailWithPassword(email string) (User, error) {
	result := User{}
	stm := `SELECT * FROM "user" WHERE email = $1`
	err := DB.Get(&result, stm, email)
	if err != nil {
		return User{}, err
	}
	return result, err
}
