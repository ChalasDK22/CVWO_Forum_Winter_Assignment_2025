package post

import (
	"context"
	"database/sql"

	"chalas.com/forum_project/API/internal/models"
)

type PostRepository interface {
	InsertPost(ctx context.Context, model *models.PostModel) (int64, error)
	UpdatePost(ctx context.Context, postID int64, model *models.PostModel) error
	DeletePost(ctx context.Context, postID int64) error
}

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{
		db: db,
	}
}

func (repos *postRepository) InsertPost(ctx context.Context, model *models.PostModel) (int64, error) {
	query := "INSERT INTO `POSTS` (`topic_id`, `user_id`, `post_date`, `title`, `content`) VALUES (?, ?, ?, ?, ?)"
	insertedRow, err := repos.db.ExecContext(ctx, query, &model.TopicID, &model.UserID, &model.PostDate, &model.Title, &model.Content)

	if err != nil {
		return 0, err
	}

	id, _ := insertedRow.LastInsertId()
	return id, nil
}

func (repos *postRepository) UpdatePost(ctx context.Context, postID int64, model *models.PostModel) error {
	query := "UPDATE `POSTS` SET `topic_id` = ?, `user_id` = ?, `post_date` = ?, `title` = ?, `content` = ? WHERE `post_id` = ?"
	updatedRow, err := repos.db.ExecContext(ctx, query, &model.TopicID, &model.UserID, &model.PostDate, &model.Title, &model.Content, postID)
	if err != nil {
		return err
	}

	rowAffected, err := updatedRow.RowsAffected()
	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (repos *postRepository) DeletePost(ctx context.Context, postID int64) error {
	query := "DELETE FROM `POSTS` WHERE `post_id` = ?"
	deletedRow, err := repos.db.ExecContext(ctx, query, postID)
	if err != nil {
		return err
	}

	rowAffected, _ := deletedRow.RowsAffected()

	if rowAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
