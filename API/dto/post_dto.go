package dto

type (
	Create_Topic_Request struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
	}

	Create_Topic_Response struct {
		TopicID int64 `json:"topic_id"`
	}
)
