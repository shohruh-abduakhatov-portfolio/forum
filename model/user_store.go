package model

import (
	"database/sql"
	"strings"
)

type userStore struct {
	conn *sql.DB
}

var GlobalUserStore UserStore

func InitGlobalUserStore() {
	GlobalUserStore = DB.userStore
}

const selectFromUsers = `select * from user `

// New creates a new user.
func (s *userStore) New(user User) (string, error) {
	stmt, err := s.conn.Prepare("insert into user(id, username, email, password, role_id, permission_id, photo_id) values(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return "", err
	}

	_, err = stmt.Exec(user.ID, user.Username, user.Email, user.Password, user.RoleID, user.PermissionID, user.PhotoID)
	if err != nil {
		return "", err
	}

	return user.ID, err
}

func (s *userStore) Update(user User) error {
	stmt, err := s.conn.Prepare("update user set username=?, email=?, password=?, where id=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Username, user.Email, user.Password, user.ID)
	if err != nil {
		return err
	}
	return nil
}

// Find finds a user by ID.
func (s *userStore) Find(userID string) (*User, error) {
	row := s.conn.QueryRow(selectFromUsers+` where id=$1`, userID)
	return s.scanUser(row)
}

func (s *userStore) FindByUsername(username string) (*User, error) {
	if username == "" {
		return nil, nil
	}
	row := s.conn.QueryRow(selectFromUsers+` where username=$1`, strings.ToLower(username))
	return s.scanUser(row)
}

func (s *userStore) FindByEmail(email string) (*User, error) {
	if email == "" {
		return nil, nil
	}
	row := s.conn.QueryRow(selectFromUsers+` where email=$1`, strings.ToLower(email))
	return s.scanUser(row)
}

// GetMany finds users by IDs.
func (s *userStore) GetMany(ids []string) (map[string]*User, error) {
	users := make(map[string]*User)
	if len(ids) == 0 {
		return users, nil
	}

	for _, id := range ids {
		users[id] = nil
	}

	var params []interface{}
	for id := range users {
		params = append(params, id)
	}

	rows, err := s.conn.Query(
		selectFromUsers+` where id in (`+placeholders(1, len(params))+`)`,
		params...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user, err := s.scanUser(rows)
		if err != nil {
			return nil, err
		}
		users[user.ID] = user
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	for _, user := range users {
		if user == nil {
			return nil, ErrNotFound
		}
	}
	return users, nil
}

// GetAdmins finds all the admin users.
func (s *userStore) GetAll() ([]*User, error) {
	var users []*User

	rows, err := s.conn.Query(selectFromUsers + ` where role!=0`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user, err := s.scanUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// GetAdmins finds all the admin users.
func (s *userStore) GetAdmins() ([]*User, error) {
	var users []*User

	rows, err := s.conn.Query(selectFromUsers + ` where role==0`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user, err := s.scanUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// SetName updates user.Name value. It returns ErrConflict if the given name is already taken.
func (s *userStore) SetName(id string, name string) error {
	stmt, err := s.conn.Prepare("update user set username=? where id=?")
	if err != nil {
		// todo if isUniqueConstraintError(err) {
		// 	return model.ErrConflict
		// }
		return err
	}

	_, err = stmt.Exec(name, id)
	if err != nil {
		return err
	}
	return nil
}

// SetAvatar updates user.Avatar value.
func (s *userStore) SetAvatar(id int64, avatar string) error {
	stmt, err := s.conn.Prepare("update user set avatar=? where id=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(avatar, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *userStore) scanUser(scanner scanner) (*User, error) {
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
