package comment

import (
	"context"
	"database/sql"

	"chalas.com/forum_project/API/internal/dto"
	"chalas.com/forum_project/API/internal/models"
)

type CommentRepository interface {
	CreateComment(ctx context.Context, model *models.CommentModel) (int64, error)
	UpdateComment(ctx context.Context, commentID int64, model *models.CommentModel) error
	GetCommentOwnerID(ctx context.Context, commentID int64) (int64, error)
	DeleteComment(ctx context.Context, commentID int64, userID int64) error
	GetCommentByPostID(ctx context.Context, postID int64, limit int, offset int) ([]dto.GetCommentWithAuthorResponse, error)
	CountComment(ctx context.Context, postID int64) (int64, error)
}

type commentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (repos *commentRepository) CreateComment(ctx context.Context, model *models.CommentModel) (int64, error) {
	query := "INSERT INTO `COMMENTS` (`user_id`, `post_date`, `post_id`, `content`) VALUES (?, ?, ?, ?)"
	insertedRow, err := repos.db.ExecContext(ctx, query, &model.UserID, &model.PostDate, &model.PostID, &model.Content)

	if err != nil {
		return 0, err
	}

	id, _ := insertedRow.LastInsertId()
	return id, nil
}

func (repos *commentRepository) UpdateComment(ctx context.Context, commentID int64, model *models.CommentModel) error {
	query := "UPDATE `COMMENTS` SET `user_id` = ?, `post_date` = ?, `post_id` = ?, `content` = ? WHERE `comment_id` = ?"
	updatedRow, err := repos.db.ExecContext(ctx, query, &model.UserID, &model.PostDate, &model.PostID, &model.Content, commentID)
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

func (repo *commentRepository) GetCommentOwnerID(ctx context.Context, commentID int64) (int64, error) {
	query := "SELECT `user_id` FROM `COMMENTS` WHERE `comment_id` = ?"
	result := repo.db.QueryRowContext(ctx, query, commentID)
	var userID int64
	err := result.Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (repos *commentRepository) DeleteComment(ctx context.Context, commentID int64, userID int64) error {
	query := "DELETE FROM `COMMENTS` WHERE `comment_id` = ? and user_id = ?"
	deletedRow, err := repos.db.ExecContext(ctx, query, commentID, userID)
	if err != nil {
		return err
	}

	rowAffected, _ := deletedRow.RowsAffected()

	if rowAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (repos *commentRepository) GetCommentByPostID(ctx context.Context, postID int64, limit int, offset int) ([]dto.GetCommentWithAuthorResponse, error) {
	query := "SELECT u.user_id, c.post_date, c.post_id, c.content, c.comment_id, u.username FROM `COMMENTS` c JOIN `USERS` u on c.user_id = u.user_id WHERE `post_id` = ? ORDER BY post_date DESC LIMIT ? OFFSET ?"
	rows, err := repos.db.QueryContext(ctx, query, postID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	comments := make([]dto.GetCommentWithAuthorResponse, 0)
	for rows.Next() {
		var comment dto.GetCommentWithAuthorResponse
		err = rows.Scan(
			&comment.Author.UserID,
			&comment.PostDate,
			&comment.PostID,
			&comment.Content,
			&comment.CommentID,
			&comment.Author.Username)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (repos *commentRepository) CountComment(ctx context.Context, postID int64) (int64, error) {
	var count int64
	query := "SELECT COUNT(*) FROM `COMMENTS` WHERE `comment_id` = ?"
	err := repos.db.QueryRowContext(ctx, query, postID).Scan(&count)
	return count, err
}
