package model

import (
	"reflect"
	"strconv"
	"time"
	"unicode/utf8"
)

// Topic is a discussion topic.
type Post struct {
	ID           int64       `json:"id"`
	UserID       string      `json:"userId"`
	User         *User       `json:"user"`
	Title        string      `json:"title"`
	Text         string      `json:"text"`
	CreatedAt    time.Time   `json:"createdAt"`
	LikeCount    int64       `json:"likeCount"`
	DislikeCount int64       `json:"dislikeCount"`
	CommentCount int64       `json:"commentCount"`
	PhotoID      string      `json:"photoId"`
	Categories   []*Category `json:"categories"`
}

const (
	topicTitleMinLen = 1
	topicTitleMaxLen = 100
)

func NewPost(user *User, title, text, photoPath string) (*Post, error) {
	post := &Post{
		User:      user,
		Title:     title,
		Text:      text,
		CreatedAt: time.Now(),
		PhotoID:   photoPath,
	}
	// Validate min len
	if err := post.ValidTextLen(); err != nil {
		return post, err
	}

	return post, nil
}

func ParseCategoryArr(categoryIdList []string) ([]int, error) {
	if len(categoryIdList) == 0 {
		return nil, errNoCategory
	}

	idList := make([]int, len(categoryIdList))
	for i, val := range categoryIdList {
		id, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		idList[i] = id
	}
	return idList, nil
}

func ValidCategory(categories interface{}) error {
	if !(reflect.ValueOf(categories).Len() == 0) {
		return errNoCategory
	}
	return nil
}

// ValidTopicTitle checks if topic title is valid.
func (p *Post) ValidTextLen() error {
	if !utf8.ValidString(p.Title) {
		return errInvalidTopicString
	}

	length := utf8.RuneCountInString(p.Title)
	if !(topicTitleMinLen <= length && length <= topicTitleMaxLen) {
		return errTopicContentMinLen
	}

	length = utf8.RuneCountInString(p.Text)
	if !(topicTitleMinLen <= length) {
		return errTextContentMinLen
	}
	return nil
}
