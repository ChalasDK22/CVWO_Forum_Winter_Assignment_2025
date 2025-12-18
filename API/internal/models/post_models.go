package models

import "time"

type PostModel struct {
	PostID   int64
	TopicID  int64
	UserID   int64
	Title    string
	Content  string
	PostDate time.Time
}
