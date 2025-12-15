package topic

import (
	"chalas.com/forum_project/API/internal/config"
	"chalas.com/forum_project/API/repository/topic"
)

type TopicService interface {
}

type topicService struct {
	cfg        *config.Config
	topicrepos topic.TopicRepository
}

func NewTopicService(cfg *config.Config, topicrepos topic.TopicRepository) TopicService {
	return &topicService{
		cfg:        cfg,
		topicrepos: topicrepos,
	}
}
