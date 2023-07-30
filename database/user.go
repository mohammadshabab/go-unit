package database

import (
	"database/sql"

	"github.com/mohammadshabab/go-unit/entity"
)

type User interface {
	Add(user entity.User) error
}
type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) Add(user entity.User) error {
	sql := "INSERT INTO user (name, email, description) VALUES (?, ?, ?)"
	_, err := u.db.Exec(sql, user.Name, user.Email, user.Description)
	if err != nil {
		return err
	}
	return nil
}
