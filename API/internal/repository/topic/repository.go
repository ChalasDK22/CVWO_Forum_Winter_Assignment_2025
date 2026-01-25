package topic

import (
	"context"
	"database/sql"

	"chalas.com/forum_project/API/internal/dto"
	"chalas.com/forum_project/API/internal/models"
)

type TopicRepository interface {
	InsertTopic(ctx context.Context, model *models.TopicModel) (int64, error)
	GetTopicbyID(ctx context.Context, topicId int64) (*models.TopicModel, error)
	UpdateTopic(ctx context.Context, topicId int64, model *models.TopicModel) error
	DeleteTopic(ctx context.Context, topicID int64, userID int64) (int64, error)
	GetTopics(ctx context.Context, limit int, offset int) ([]models.TopicModel, error)
	GetOwnerID(ctx context.Context, topicID int64) (int64, error)
	CountTopics(ctx context.Context) (int, error)
	GetTopicByTopicID(ctx context.Context, topicID int64) ([]dto.GetTopicByTopicIDResponse, error)
}

type topicRepository struct {
	db *sql.DB
}

func NewTopicRepository(db *sql.DB) TopicRepository {
	return &topicRepository{
		db: db,
	}
}

// Repo Method to create topic
func (repo *topicRepository) InsertTopic(ctx context.Context, model *models.TopicModel) (int64, error) {
	query := "INSERT INTO `TOPICS` (`user_id`, `name`, `description`) VALUES ( ?, ?, ?)"

	inserted_row, err := repo.db.ExecContext(ctx, query, model.UserID, model.Name, model.Desription)
	if err != nil {
		return 0, err
	}

	id, _ := inserted_row.LastInsertId()
	return id, nil
}

func (repo *topicRepository) GetTopicbyID(ctx context.Context, topicID int64) (*models.TopicModel, error) {
	query := "SELECT topicID, user_id, name, description FROM `TOPICS` WHERE topicID = ?"

	selected_row := repo.db.QueryRowContext(ctx, query, topicID)
	var resultedTopic models.TopicModel
	err := selected_row.Scan(&resultedTopic.TopicID, &resultedTopic.UserID, &resultedTopic.Name, resultedTopic.Desription)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &resultedTopic, nil
}

// Repo Method to update topic
func (repo *topicRepository) UpdateTopic(ctx context.Context, topicID int64, model *models.TopicModel) error {
	query := "UPDATE `TOPICS` SET `user_id` = ?, `name` = ?, description = ? WHERE topicID = ?"

	updatedRow, err := repo.db.ExecContext(ctx, query, model.UserID, model.Name, model.Desription, topicID)

	if err != nil {
		return err
	}

	rowsAffected, _ := updatedRow.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// Repo Method to delete topic
func (repo *topicRepository) DeleteTopic(ctx context.Context, topicID int64, userID int64) (int64, error) {
	query := "DELETE FROM `TOPICS` WHERE topicID = ? and user_id = ?"

	deletedRow, err := repo.db.ExecContext(ctx, query, topicID, userID)

	if err != nil {
		return 0, err
	}

	rowsAffected, _ := deletedRow.RowsAffected()
	if rowsAffected == 0 {
		return 0, sql.ErrNoRows
	}

	return rowsAffected, nil
}

func (repo *topicRepository) GetOwnerID(ctx context.Context, topicID int64) (int64, error) {
	query := "SELECT user_id FROM `TOPICS` WHERE topicID = ?"
	result := repo.db.QueryRowContext(ctx, query, topicID)
	var userID int64
	err := result.Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (repo *topicRepository) GetTopics(ctx context.Context, limit int, offset int) ([]models.TopicModel, error) {
	query := "SELECT topicID, user_id, name, description FROM `TOPICS` LIMIT ? OFFSET ?"
	rows, err := repo.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	topics := []models.TopicModel{}
	for rows.Next() {
		var topic models.TopicModel
		err = rows.Scan(&topic.TopicID, &topic.UserID, &topic.Name, &topic.Desription)
		if err != nil {
			return nil, err
		}
		topics = append(topics, topic)
	}
	return topics, nil
}

func (repo *topicRepository) CountTopics(ctx context.Context) (int, error) {
	query := "SELECT COUNT(*) FROM `TOPICS`"
	var count int
	err := repo.db.QueryRowContext(ctx, query).Scan(&count)
	return count, err
}

func (repo *topicRepository) GetTopicByTopicID(ctx context.Context, topicID int64) ([]dto.GetTopicByTopicIDResponse, error) {
	query := "SELECT topicID, name, description, user_id FROM `TOPICS` WHERE `topicID` = ? "
	rowsAffected, err := repo.db.QueryContext(ctx, query, topicID)
	if err != nil {
		return nil, err
	}
	defer rowsAffected.Close()

	topics := make([]dto.GetTopicByTopicIDResponse, 0)
	for rowsAffected.Next() {
		var topic dto.GetTopicByTopicIDResponse
		err = rowsAffected.Scan(
			&topic.TopicID,
			&topic.Name,
			&topic.Description,
			&topic.UserID,
		)
		if err != nil {
			return nil, err
		}
		topics = append(topics, topic)
	}
	return topics, rowsAffected.Err()
}
