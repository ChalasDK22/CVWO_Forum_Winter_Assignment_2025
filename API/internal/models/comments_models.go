package models

import "time"

type CommentModel struct {
	UserID    int64
	PostID    int64
	CommentID int64
	Content   string
	PostDate  time.Time
}
