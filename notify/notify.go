package notify

import model "forum.com/model"

type Message struct {
	title string
	body  string
	user  model.User
	post  interface{}
}

type Notify interface {
	Send(*Message) error
	Format(interface{}) (string, error)
}
