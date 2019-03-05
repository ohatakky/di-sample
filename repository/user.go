package repository

import (
	"database/sql"
	"di-sample/model"
)

type UserRepository interface {
	Save(u model.User) error
}

func NewUserRepository(*sql.DB) UserRepository {

	return nil
}
