package user

import "database/sql"

type UserRepository interface {
}

type userRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}