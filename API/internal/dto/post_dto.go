package dto

type (
	CreatePostRequest struct {
		TopicID int64  `json:"topic_id" validate:"required"`
		Title   string `json:"title" validate:"required"`
		Content string `json:"content" validate:"required"`
	}

	CreatePostResponse struct {
		PostID int64 `json:"post_id"`
	}

	UpdatePostRequest struct {
		TopicID int64  `json:"topic_id" validate:"required"`
		Title   string `json:"title" validate:"required"`
		Content string `json:"content" validate:"required"`
	}

	UpdatePostResponse struct {
		PostID int64 `json:"post_id"`
	}

	DeletePostRequest struct {
		PostID int64 `json:"post_id" validate:"required"`
	}

	DeletePostResponse struct {
		PostID int64 `json:"post_id"`
	}
)
