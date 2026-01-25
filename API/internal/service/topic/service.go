package topic

import (
	"context"
	"errors"
	"net/http"

	"chalas.com/forum_project/API/internal/config"
	"chalas.com/forum_project/API/internal/dto"
	"chalas.com/forum_project/API/internal/models"

	"chalas.com/forum_project/API/internal/repository/topic"
)

type TopicService interface {
	CreateTopic(ctx context.Context, dtoReq *dto.CreateTopicRequest, userID int64) (int64, int, error)
	UpdateTopic(ctx context.Context, dtoReq *dto.UpdateTopicRequest, topicID int64, userID int64) (int, error)
	DeleteTopic(ctx context.Context, topicID int64, userID int64) (int, error)
	GetTopics(ctx context.Context, page int, size int) ([]models.TopicModel, int, error)
	GetTopicByTopicID(ctx context.Context, topicID int64) ([]dto.GetTopicByTopicIDResponse, error)
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
var (
	ErrNoPermission  = errors.New("Unauthorized permission")
	ErrTopicNotFound = errors.New("Topic not found")
)

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

func (service *topicService) UpdateTopic(ctx context.Context, dtoReq *dto.UpdateTopicRequest, topicID int64, userID int64) (int, error) {
	ownerID, err := service.topicrepos.GetOwnerID(ctx, topicID)
	if err != nil {
		return http.StatusInternalServerError, ErrTopicNotFound
	}

	if ownerID != userID {
		return http.StatusForbidden, ErrNoPermission
	}
	err = service.topicrepos.UpdateTopic(ctx, topicID, &models.TopicModel{
		UserID:     userID,
		Name:       dtoReq.Name,
		Desription: dtoReq.Description,
	})
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// Delete Topic

func (service *topicService) DeleteTopic(ctx context.Context, topicID int64, userID int64) (int, error) {
	ownerID, err := service.topicrepos.GetOwnerID(ctx, topicID)
	if err != nil {
		return http.StatusInternalServerError, ErrTopicNotFound
	}

	if ownerID != userID {
		return http.StatusForbidden, ErrNoPermission
	}
	rowsAffected, err := service.topicrepos.DeleteTopic(ctx, topicID, userID)

	if err != nil {
		return http.StatusInternalServerError, err
	}
	if rowsAffected == 0 {
		return http.StatusBadRequest, ErrNoPermission
	}

	return http.StatusOK, nil
}

func (service *topicService) GetTopics(ctx context.Context, page int, size int) ([]models.TopicModel, int, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 50 {
		size = 10
	}

	offset := (page - 1) * size

	total, err := service.topicrepos.CountTopics(ctx)
	if err != nil {
		return nil, 0, err
	}

	topics, err := service.topicrepos.GetTopics(ctx, size, offset)
	if err != nil {
		return nil, 0, err
	}

	return topics, total, nil
}

func (service *topicService) GetTopicByTopicID(ctx context.Context, topicID int64) ([]dto.GetTopicByTopicIDResponse, error) {
	topics, err := service.topicrepos.GetTopicByTopicID(ctx, topicID)
	if err != nil {
		return nil, err
	}

	return topics, nil
}
