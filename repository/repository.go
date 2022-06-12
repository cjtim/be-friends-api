package repository

import (
	"github.com/cjtim/be-friends-api/configs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	DB       *sqlx.DB
	UserRepo *UserImpl
	TagRepo  *TagImpl
)

func Connect() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", configs.Config.DATABASE_URL)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	DB = db
	return db, err
}
