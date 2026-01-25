package dto

import "time"

type (
	Author struct {
		UserID   int64  `json:"user_id"`
		Username string `json:"username"`
	}

	GetPostWithAuthorResponse struct {
		PostID   int64     `json:"post_id"`
		Title    string    `json:"title"`
		Content  string    `json:"content"`
		TopicID  int64     `json:"topic_id"`
		PostDate time.Time `json:"post_date"`
		Author   Author    `json:"author"`
	}

	GetCommentWithAuthorResponse struct {
		PostID    int64     `json:"post_id"`
		CommentID int64     `json:"comment_id"`
		Content   string    `json:"content"`
		PostDate  time.Time `json:"post_date"`
		Author    Author    `json:"author"`
	}

	GetTopicByTopicIDResponse struct {
		TopicID     int64  `json:"topic_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		UserID      int64  `json:"user_id"`
	}
)
