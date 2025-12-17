package testutil

import "chalas.com/forum_project/API/internal/models"

func NewTopicTestingModel() *models.TopicModel {
	return &models.TopicModel{
		UserID:     1,
		Name:       "test",
		Desription: "test model",
	}
}

func NewUpdatedTopicTestModel() *models.TopicModel {
	return &models.TopicModel{
		UserID:     2,
		Name:       "updated test",
		Desription: "updated test model",
	}
}
