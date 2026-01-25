package dto

type (
	CreateCommentRequest struct {
		Content string `json:"content" validate:"required"`
		PostID  int64  `json:"post_id" validate:"required"`
	}

	CreateCommentResponse struct {
		CommentID int64 `json:"comment_id"`
	}

	UpdateCommentRequest struct {
		Content string `json:"content" validate:"required"`
		PostID  int64  `json:"post_id" validate:"required"`
	}

	UpdateCommentResponse struct {
		CommentID int64 `json:"comment_id"`
	}

	DeleteCommentRequest struct {
		CommentID int64 `json:"comment_id" validate:"required"`
	}

	DeleteCommentResponse struct {
		CommentID int64 `json:"comment_id"`
	}
)
