// Package postgresql provides a PostgreSQL implementation of the bebop data store interface.
package model

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	Conn          *sql.DB
	userStore     *userStore
	sessionStore  *sessionStore
	postStore     *postStore
	commentStore  *commentStore
	categoryStore *categoryStore
}

var (
	DB     *Store
	dbName = ""
)

var errConfiErro = errors.New("Config error")

func Init() {
	dbName = os.Getenv("DBName")
	db, err := Connect()
	if err != nil {
		panic("Cannot init db")
	}
	DB = db
	InitGlobalUserStore()
	InitGlobalSessionStore()
	InitGlobalPostStore()
	InitGlobalCommentStore()
	InitGlobalCategoryStore()
}

// Connect connects to a store.
func Connect() (*Store, error) {
	if dbName == "" {
		return nil, errConfiErro
	}
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(20)

	if err := db.Ping(); err != nil {
		return nil, err
	}
	s := &Store{
		Conn:          db,
		userStore:     &userStore{db},
		sessionStore:  &sessionStore{db},
		postStore:     &postStore{db},
		commentStore:  &commentStore{db},
		categoryStore: &categoryStore{db},
	}
	// err = s.Migrate()
	// if err != nil {
	// 	panic("Cannot connect to DB")
	// }
	// err = s.CreateUsers()
	// if err != nil {
	// 	panic("Cannot create role/permission:", err)
	// }

	return s, nil
}

type scanner interface {
	Scan(v ...interface{}) error
}

// Migrate migrates the store database.
func (s *Store) Migrate() error {
	for _, q := range migrate {
		_, err := s.Conn.Exec(q)
		if err != nil {
			return fmt.Errorf("sql exec error: %s; query: %q", err, q)
		}
	}
	return nil
}

func (s *Store) CreateUsers() error {
	for _, q := range createUsers {
		_, err := s.Conn.Exec(q)
		if err != nil {
			if IsSqliteError(err) && IsUniqueConstraintError(err) {
				continue
			} else {
				return fmt.Errorf("sql exec error: %s; query: %q", err, q)
			}
		}
	}
	return nil
}

func placeholders(start, count int) string {
	buf := new(bytes.Buffer)
	for i := start; i < start+count; i++ {
		buf.WriteByte('$')
		buf.WriteString(strconv.Itoa(i))
		if i < start+count-1 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}
