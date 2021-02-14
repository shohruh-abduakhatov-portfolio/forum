package model

import (
	"database/sql"
	"fmt"
	"time"
)

type sessionStore struct {
	conn *sql.DB
}

var GlobalSessionStore SessionStore

const selectFromSession = `select * from session `

func InitGlobalSessionStore() {
	GlobalSessionStore = DB.sessionStore
}

func (s *sessionStore) Find(id string) (*Session, error) {
	row := s.conn.QueryRow(selectFromSession+` where id=$1`, id)
	return s.scanSession(row)
}

func (s *sessionStore) Save(session *Session) error {
	stmt, err := s.conn.Prepare("insert into session(id, userId, expiry) values(?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(session.ID, session.UserID, session.Expiry)
	if err != nil {
		return err
	}

	s.GetAllSessions()

	return nil
}

func (s *sessionStore) Delete(session *Session) error {
	stmt, err := s.conn.Prepare("delete from session where id=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(session.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *sessionStore) GetAllSessions() error {
	query := selectFromSession
	rows, err := s.conn.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		user, err := s.scanSession(rows)
		if err != nil {
			return err
		}
		fmt.Println(">> user:", user.ID, user.UserID, user.Expired)
	}
	if err = rows.Err(); err != nil {
		return err
	}
	return nil
}

func (s *sessionStore) scanSession(scanner scanner) (*Session, error) {
	u := new(Session)
	var expiry string
	err := scanner.Scan(
		&u.ID,
		&u.UserID,
		&expiry,
	)
	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	parsed, err := time.Parse("2006-01-02T15:04:05.999999999-07:00", expiry)
	if err == nil {
		u.Expiry = parsed
	}
	return u, nil
}
