package topic

import (
	"context"
	"net/http"

	"chalas.com/forum_project/API/internal/config"
	"chalas.com/forum_project/API/internal/dto"
	"chalas.com/forum_project/API/internal/models"

	"chalas.com/forum_project/API/internal/repository/topic"
)

type TopicService interface {
	CreateTopic(ctx context.Context, dtoReq *dto.CreateTopicRequest, userID int64) (int64, int, error)
	UpdateTopic(ctx context.Context, dtoReq *dto.Update_Topic_Request, topicID int64, userID int64) (int, error)
}

type topicService struct {
	cfg        *config.Config
	topicrepos topic.TopicRepository
}

// NewTopic
func NewTopicService(cfg *config.Config, topicrepos topic.TopicRepository) TopicService {
	return &topicService{
		cfg:        cfg,
		topicrepos: topicrepos,
	}
}

// Create Topic
func (service *topicService) CreateTopic(ctx context.Context, dtoReq *dto.CreateTopicRequest, userID int64) (int64, int, error) {
	topicInsertedID, err := service.topicrepos.InsertTopic(ctx, &models.TopicModel{
		UserID:     userID,
		Name:       dtoReq.Name,
		Desription: dtoReq.Description,
	})
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	return topicInsertedID, http.StatusOK, nil
}

//Update Topic

func (service *topicService) UpdateTopic(ctx context.Context, dtoReq *dto.Update_Topic_Request, topicID int64, userID int64) (int, error) {
	err := service.topicrepos.UpdateTopic(ctx, topicID, &models.TopicModel{
		UserID:     userID,
		Name:       dtoReq.Name,
		Desription: dtoReq.Description,
	})
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
