package model

import (
	"database/sql"
)

func (s *Store) scanRole(scanner scanner) (*User, error) {
	u := new(User)
	err := scanner.Scan(
		&u.ID,
		&u.Username,
		&u.Email,
		&u.Password,
		&u.DateCreated,
		&u.RoleID,
		&u.PermissionID,
		&u.PhotoID,
	)
	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return u, nil
}
