package repositories

import (
	"messenger-module/db"
)

type DBRepository struct {
	database db.Database
}

func NewDBRepository(database db.Database) *DBRepository {
	return &DBRepository{database: database}
}
