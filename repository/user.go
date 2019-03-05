package repository

import (
	"database/sql"
	"di-sample/model"
)

type UserRepository interface {
	Save(u model.User) error
}

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) userRepository {

	return userRepository{db}
}

func (u userRepository) Save(model.User) error {
	return nil
}
