package topic

import "database/sql"

type TopicRepository interface {
}

type topicRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *topicRepository {
	return &topicRepository{
		db: db,
	}
}
