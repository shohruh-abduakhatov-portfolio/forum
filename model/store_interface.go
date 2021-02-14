package model

// IStore is a forum data store interface.
type IStore interface {
	Users() UserStore
	Topics() PostStore
	Comments() CommentStore
	Categories() CategoryStore
}

// UserStore is a forum user data store interface.
type UserStore interface {
	New(user User) (string, error)
	Update(user User) error
	Find(userID string) (*User, error)
	FindByUsername(username string) (*User, error)
	FindByEmail(email string) (*User, error)
	GetMany(ids []string) (map[string]*User, error)
	GetAll() ([]*User, error)
	GetAdmins() ([]*User, error)
	SetName(id string, name string) error
	SetAvatar(id int64, avatar string) error
}

// SessionStore is a user session store interface
type SessionStore interface {
	Find(string) (*Session, error)
	Save(*Session) error
	Delete(*Session) error
	GetAllSessions() error
}

// PostStore is a forum topic data store interface.
type PostStore interface {
	New(*Post) (int64, error)
	Modify(*Post) error
	NewPostCategory(post *Post, ids []int) error
	NewUserPost(post *Post) error
	DeletePostCategories(ids []string) error
	Get(id int64) (*Post, []*Category, *User, error)
	GetLatest(offset, limit int) ([]*Post, error)
	GetByCategory(offset, limit, category int) ([]*Post, error)
	GetByUser(offset, limit int, userID string) ([]*Post, error)
	GetLiked(offset, limit int, user string) ([]*Post, error)
	// SetTitle(id int64, title string) error
	Delete(id int64) error
	IncrementReaction(id int64, reaction string) error
	DecrementReaction(id int64, reaction string) error
	IncrementCommentCount(id int64) error
	HasReacted(int, string, int) (bool, error)
	NewUserReaction(int, string, int) error
}

// CommentStore is a forum comment data store interface.
type CommentStore interface {
	New(comment *Comment) (int64, error)
	Get(id int64) (*Comment, error)
	GetByTopic(postID int64) ([]*Comment, error)
}

type CategoryStore interface {
	GetCategoryList() ([]*Category, error)
}
