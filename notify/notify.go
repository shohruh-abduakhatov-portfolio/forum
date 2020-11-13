package model

type Message struct {
	title string
	body  string
	user  User
	post  interface{}
}

type Notify interface {
	Send(*Message) error
	Format(interface{}) (string, error)
}
