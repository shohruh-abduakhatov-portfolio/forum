package model

import (
	"database/sql"
)

type commentStore struct {
	conn *sql.DB
}

var GlobalCommentStore CommentStore

const selectFromComment = ` select c.*, u.username, u.email, u.photo_id  
							from comment as c
							left join user as u 
								on c.user_id = u.id `

func InitGlobalCommentStore() {
	GlobalCommentStore = DB.commentStore
}

func (s *commentStore) New(c *Comment) (int64, error) {
	stmt, err := s.conn.Prepare(
		"insert into comment(post_id, user_id,  comment_dt, comment) values(?, ?, ?, ?)")

	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(c.PostID, c.User.ID, c.CreatedAt, c.Comment)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	if id == 0 {
		return 0, ErrNotInserted
	}

	return id, err
}

func (s *commentStore) Get(id int64) (*Comment, error) {
	return nil, nil
}

func (s *commentStore) GetByTopic(postID int64) ([]*Comment, error) {
	rows, err := s.conn.Query(selectFromComment+` where post_id=?`, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*Comment
	for rows.Next() {
		user, err := s.scanComment(rows)
		if err != nil {
			return nil, err
		}
		comments = append(comments, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}

func (s *commentStore) scanComment(scanner scanner) (*Comment, error) {
	u := new(Comment)
	user := new(User)

	err := scanner.Scan(
		&u.ID,
		&u.PostID,
		&user.ID,
		&u.CreatedAt,
		&u.Comment,
		&user.Username,
		&user.Email,
		&user.PhotoID,
	)

	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	u.User = user

	return u, nil
}
