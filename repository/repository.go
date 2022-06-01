package repository

import (
	"github.com/cjtim/be-friends-api/configs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	DB       *sqlx.DB
	UserRepo *UserImpl
)

func Connect() (*sqlx.DB, error) {
	return sqlx.Open("postgres", configs.Config.DATABASE_URL)
}
