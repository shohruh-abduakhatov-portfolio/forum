package model

// Reaction is a single reaction on a topic.
type Reaction struct {
	PostID   int64 `json:"postId"`
	UserID   int64 `json:"userId"`
	Reaction int   `json:"reaction"`
}

const (
	LIKE = iota
	DISLIKE
)

const (
	LIKE_COUNT    = "like_count"
	DISLIKE_COUNT = "dislike_count"
)
