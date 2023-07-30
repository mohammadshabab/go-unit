package database

import (
	"database/sql"

	"github.com/mohammadshabab/go-unit/entity"
)

type AttemptHistory interface {
	IncrementFailure(user entity.User) error
	CountFailures(user entity.User) (int, error)
}
type AttemptHistoryRepository struct {
	db *sql.DB
}

func NewAttemptHistoryRepository(db *sql.DB) *AttemptHistoryRepository {
	return &AttemptHistoryRepository{
		db: db,
	}
}

func (a *AttemptHistoryRepository) IncrementFailure(user entity.User) error {
	sql := "INSERT INTO attempt_hisotry (user_id) VALUES (?)"
	_, err := a.db.Exec(sql, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (a *AttemptHistoryRepository) CountFailures(user entity.User) (count int, err error) {
	sql := "SELECT count(user_id) FROM attempt_history WHERE user_id = ?"
	row := a.db.QueryRow(sql)
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
