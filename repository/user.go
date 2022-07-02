package repository

import (
	"encoding/json"
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

type UserExtended struct {
	User
	// Custome fields
	Tags    *json.RawMessage `json:"tags" db:"tags"`
	IsAdmin *bool            `json:"is_admin" db:"is_admin"`
}

// func (t *UserImpl) GetUser(userId uuid.UUID) (user User, err error) {
// 	stm := `SELECT * FROM user u WHERE u.id = $1`
// 	err = DB.Get(&user, stm)
// 	return user, err
// }

func (t *UserImpl) GetUserExtended(userID uuid.UUID) (user UserExtended, err error) {
	stm := `
	SELECT
		u.id, u.name, u.email, u.line_uid, u.picture_url, u.created_at, u.updated_at,
		(
			SELECT COALESCE(json_agg(tag), '[]')
			FROM (
				SELECT t.id as id, t.name as name
				FROM "tag_user" tu
				INNER JOIN "tag" t on t.id = tu.tag_id 
				WHERE tu.user_id = u.id AND t.is_internal = FALSE
			) tag
		) AS tags,
		(
			SELECT EXISTS (
				SELECT 1 
				FROM "tag_user" tu
				INNER JOIN "tag" t ON t.id = tu.tag_id 
				WHERE t.name = 'Admin' AND tu.user_id = u.id 
			)
		) AS is_admin
	FROM "user" u
	WHERE u.id = $1
	`
	err = DB.Get(&user, stm, userID)
	return
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
	return result, err
}
