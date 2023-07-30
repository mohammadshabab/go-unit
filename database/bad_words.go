package database

import "database/sql"

type BadWords interface {
	FindAll() ([]string, error)
}
type BadWordsRepository struct {
	db *sql.DB
}

func (br *BadWordsRepository) NewBadWordsRepository(db *sql.DB) *BadWordsRepository {
	return &BadWordsRepository{
		db: db,
	}
}

func (br *BadWordsRepository) FindAll() (badWordList []string, err error) {
	sql := "SELECT name FROM bad_word"
	rows, err := br.db.Query(sql)
	if err != nil {
		return badWordList, err
	}
	var badWord string
	for rows.Next() {
		err := rows.Scan(&badWord)
		if err != nil {
			return badWordList, err
		}
		badWordList = append(badWordList, badWord)
	}
	return badWordList, nil
}
