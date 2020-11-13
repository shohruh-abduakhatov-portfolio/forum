package model

import (
	"time"
	"unicode/utf8"
)

// Comment is a single comment on a topic.
type Comment struct {
	ID        int64     `json:"id"`
	PostID    int64     `json:"post_id"`
	User      *User     `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	Comment   string    `json:"comment"`
}

const (
	commentContentMinLen = 1
	commentContentMaxLen = 10000
)

func NewComment(user *User, postID int64, comment string) (*Comment, error) {
	c := &Comment{
		PostID:    postID,
		User:      user,
		CreatedAt: time.Now(),
		Comment:   comment,
	}
	// Validate min len
	if err := c.ValidCommentContent(); err != nil {
		return c, err
	}

	// save comment
	id, err := GlobalCommentStore.New(c)
	if err != nil && !IsItemError(ErrNotInserted) {
		return nil, err
	}
	if id == 0 {
		return nil, ErrNotInserted
	}

	// increment post commentsCount
	err = GlobalPostStore.IncrementCommentCount(c.PostID)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// ValidCommentContent checks if comment content is valid.
func (c *Comment) ValidCommentContent() error {
	if !utf8.ValidString(c.Comment) {
		return errInvalidCommentString
	}

	length := utf8.RuneCountInString(c.Comment)
	if !(commentContentMinLen <= length && length <= commentContentMaxLen) {
		return errCommentContentMinLen
	}

	return nil
}
