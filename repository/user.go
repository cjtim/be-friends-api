package repository

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type UserImpl struct{}

type User struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Email       *string   `json:"email" db:"email"`
	Password    *string   `json:"password" db:"password"`
	LineUid     *string   `json:"line_uid" db:"line_uid"`
	Description *string   `json:"description" db:"description"`
	PictureURL  *string   `json:"picture_url" db:"picture_url"`
	Phone       *string   `json:"phone" db:"phone"`
	IsOrg       bool      `json:"is_org" db:"is_org"`
	IsAdmin     bool      `json:"is_admin" db:"is_admin"`

	Lat *float64 `json:"lat" db:"lat"`
	Lng *float64 `json:"lng" db:"lng"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (t *UserImpl) GetUser(userId uuid.UUID) (user User, err error) {
	stm := `SELECT * FROM "user" u WHERE u.id = $1`
	err = DB.Get(&user, stm, userId)
	return user, err
}

func (t *UserImpl) GetOrganizations(includeUser string) (users []User, err error) {
	baseStm := `SELECT u.id, u.name, u.email, u.line_uid, u.description, u.picture_url, u.phone, u.is_admin, u.is_org, u.lat, u.lng, u.created_at, u.updated_at FROM "user" u`
	if includeUser != "" {
		err = DB.Select(&users, baseStm)
		return users, err
	}
	err = DB.Select(&users, baseStm+" WHERE u.is_org = TRUE")
	return users, err
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
			INSERT INTO "user" (name, line_uid, picture_url, is_org)
			VALUES (:name, :line_uid, :picture_url, :is_org)
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

func (u *UserImpl) Register(user User) (result User, err error) {
	stmInsert := `
		INSERT INTO "user" (name, email, password, description, phone, picture_url, is_org, lat, lng)
		VALUES (:name, :email, :password, :description, :phone, :picture_url, :is_org, :lat, :lng)
		RETURNING *
	`
	rows, err := DB.NamedQuery(stmInsert, user)
	if err != nil {
		zap.L().Error("error register user", zap.String("email", *user.Email), zap.String("name", user.Name), zap.Error(err))
		return
	}
	if rows.Next() {
		rows.StructScan(&result)
		zap.L().Info("NEW USER register", zap.Any("id", result.ID), zap.Error(err))
		return result, err
	}
	return User{}, errors.New("error register user - cannot parse inserted row")
}

func (u *UserImpl) GetOrgByEmailWithPassword(email string) (User, error) {
	result := User{}
	stm := `SELECT * FROM "user" WHERE email = $1 AND is_org = TRUE`
	err := DB.Get(&result, stm, email)
	return result, err
}

func (u *UserImpl) UpdateUser(user User) error {
	stmUpdate := `
		UPDATE "user"
		SET 
			name = :name,
			description = :description,
			picture_url = :picture_url,
			phone = :phone,
			lat = :lat,
			lng = :lng,
			updated_at = NOW()
		WHERE id = :id
		RETURNING *
	`
	_, err := DB.NamedQuery(stmUpdate, user)
	return err
}
