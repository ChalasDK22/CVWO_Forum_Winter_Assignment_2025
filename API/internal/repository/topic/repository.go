package topic

import (
	"context"
	"database/sql"

	"chalas.com/forum_project/API/internal/models"
)

type TopicRepository interface {
	InsertTopic(ctx context.Context, model *models.TopicModel) (int64, error)
	GetTopicbyID(ctx context.Context, topic_id int64) (*models.TopicModel, error)
	UpdateTopic(ctx context.Context, topic_id int64, model *models.TopicModel) error
	DeleteTopic(ctx context.Context, topicID int64) error
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
	query := "UPDATE `TOPICS` SET `user_id` = ?, `name` = ?, description = ? WHERE topic_id = ?"

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
func (repo *topicRepository) DeleteTopic(ctx context.Context, topicID int64) error {
	query := "DELETE FROM `TOPICS` WHERE topic_id = ?"

	deletedRow, err := repo.db.ExecContext(ctx, query, topicID)

	if err != nil {
		return err
	}

	rowsAffected, _ := deletedRow.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
