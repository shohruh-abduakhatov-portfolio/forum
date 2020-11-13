package model

// Topic is a discussion topic.
type PostCategory struct {
	Post     *Post     `json:"post"`
	Category *Category `json:"category"`
}
