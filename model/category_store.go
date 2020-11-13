package model

import (
	"database/sql"
)

type categoryStore struct {
	conn *sql.DB
}

var GlobalCategoryStore CategoryStore

const selectCategory = `select * from category `

func InitGlobalCategoryStore() {
	GlobalCategoryStore = DB.categoryStore
}

func (s *categoryStore) GetCategoryList() ([]*Category, error) {
	rows, err := s.conn.Query(selectCategory + "order by name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cats []*Category
	for rows.Next() {
		user, err := s.scanCategory(rows)
		if err != nil {
			return nil, err
		}
		cats = append(cats, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return cats, nil
}

func (s *categoryStore) scanCategory(scanner scanner) (*Category, error) {
	cat := new(Category)

	err := scanner.Scan(
		&cat.ID,
		&cat.Name,
		&cat.NameCode,
		&cat.Description,
	)
	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return cat, nil
}
