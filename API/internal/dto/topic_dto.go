package dto

type (
	CreateTopicRequest struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
	}

	CreateTopicResponse struct {
		TopicID int64 `json:"topic_id"`
	}

	UpdateTopicRequest struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
	}

	UpdateTopicResponse struct {
		TopicID int64 `json:"topic_id"`
	}

	DeleteTopicRequest struct {
		TopicID int64 `json:"topic_id" validate:"required"`
	}

	DeleteTopicResponse struct {
		TopicID int64 `json:"topic_id"`
	}
)
