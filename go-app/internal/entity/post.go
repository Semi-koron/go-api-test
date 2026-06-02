package entity

type Post struct {
	ID		  int
	UserID      int
	ReplyPostID *int
	GroupID     *int
	Content     string
	CreatedAt   string
	UpdatedAt   string
}