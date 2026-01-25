package post

import (
	"context"
	"database/sql"

	"chalas.com/forum_project/API/internal/dto"
	"chalas.com/forum_project/API/internal/models"
)

type PostRepository interface {
	InsertPost(ctx context.Context, model *models.PostModel) (int64, error)
	UpdatePost(ctx context.Context, postID int64, model *models.PostModel) error
	DeletePost(ctx context.Context, postID int64, userID int64) error
	GetPostOwnerID(ctx context.Context, postID int64) (int64, error)
	GetPostByTopicID(ctx context.Context, topicID int64, search string, limit int, offset int) ([]dto.GetPostWithAuthorResponse, error)
	CountPostsByTopicID(ctx context.Context, topicID int64, search string) (int64, error)
	GetPosts(ctx context.Context, limit int, offset int) ([]dto.GetPostWithAuthorResponse, error)
	CountPosts(ctx context.Context) (int, error)
	GetPostByPostID(ctx context.Context, postID int64) ([]dto.GetPostWithAuthorResponse, error)
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
	query := "INSERT INTO `POSTS` (`topicID`, `user_id`, `post_date`, `title`, `content`) VALUES (?, ?, ?, ?, ?)"
	insertedRow, err := repos.db.ExecContext(ctx, query, &model.TopicID, &model.UserID, &model.PostDate, &model.Title, &model.Content)

	if err != nil {
		return 0, err
	}

	id, _ := insertedRow.LastInsertId()
	return id, nil
}

func (repos *postRepository) UpdatePost(ctx context.Context, postID int64, model *models.PostModel) error {
	query := "UPDATE `POSTS` SET `topicID` = ?, `user_id` = ?, `post_date` = ?, `title` = ?, `content` = ? WHERE `post_id` = ?"
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

func (repos *postRepository) DeletePost(ctx context.Context, postID int64, userID int64) error {
	query := "DELETE FROM `POSTS` WHERE `post_id` = ? and user_id = ?"
	deletedRow, err := repos.db.ExecContext(ctx, query, postID, userID)
	if err != nil {
		return err
	}

	rowAffected, _ := deletedRow.RowsAffected()

	if rowAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (repo *postRepository) GetPostOwnerID(ctx context.Context, postID int64) (int64, error) {
	query := "SELECT `user_id` FROM `POSTS` WHERE `post_id` = ?"
	result := repo.db.QueryRowContext(ctx, query, postID)
	var userID int64
	err := result.Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (repo *postRepository) GetPostByTopicID(ctx context.Context, topicID int64, search string, limit int, offset int) ([]dto.GetPostWithAuthorResponse, error) {
	query := "SELECT p.post_id, p.topicID, u.user_id, p.post_date, p.title, p.content, u.username FROM `POSTS` p JOIN `USERS` u on p.user_id = u.user_id WHERE `topicID` = ? AND ( ? = '' OR title LIKE CONCAT('%', ? , '%') OR content LIKE CONCAT('%', ? , '%'))ORDER BY `post_date` DESC LIMIT ? OFFSET ?"
	rowsAffected, err := repo.db.QueryContext(ctx, query, topicID, search, search, search, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rowsAffected.Close()

	posts := make([]dto.GetPostWithAuthorResponse, 0)
	for rowsAffected.Next() {
		var post dto.GetPostWithAuthorResponse
		err = rowsAffected.Scan(
			&post.PostID,
			&post.TopicID,
			&post.Author.UserID,
			&post.PostDate,
			&post.Title,
			&post.Content,
			&post.Author.Username,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, rowsAffected.Err()
}

func (repo *postRepository) CountPostsByTopicID(ctx context.Context, topicID int64, search string) (int64, error) {
	query := "SELECT COUNT(*) FROM `POSTS` WHERE topicID = ? AND (? = '' OR title LIKE CONCAT('%', ? , '%') OR content LIKE CONCAT('%', ? , '%'))"

	var totalnum int64
	err := repo.db.QueryRowContext(ctx, query, topicID, search, search, search).Scan(&totalnum)

	return totalnum, err
}

func (repo *postRepository) GetPosts(ctx context.Context, limit int, offset int) ([]dto.GetPostWithAuthorResponse, error) {
	query := "SELECT p.post_id, p.topicID, u.user_id, p.post_date, p.title, p.content, u.username FROM `POSTS` p JOIN `USERS` u on p.user_id = u.user_id ORDER BY post_date DESC LIMIT ? OFFSET ?"
	rows, err := repo.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	posts := make([]dto.GetPostWithAuthorResponse, 0)
	for rows.Next() {
		var post dto.GetPostWithAuthorResponse
		err = rows.Scan(&post.PostID, &post.TopicID, &post.Author.UserID, &post.PostDate, &post.Title, &post.Content, &post.Author.Username)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (repo *postRepository) CountPosts(ctx context.Context) (int, error) {
	query := "SELECT COUNT(*) FROM `POSTS`"
	var count int
	err := repo.db.QueryRowContext(ctx, query).Scan(&count)
	return count, err
}

func (repo *postRepository) GetPostByPostID(ctx context.Context, postID int64) ([]dto.GetPostWithAuthorResponse, error) {
	query := "SELECT p.post_id, p.topicID, u.user_id, p.post_date, p.title, p.content, u.username FROM `POSTS` p JOIN `USERS` u on p.user_id = u.user_id WHERE `post_id` = ?"
	rowsAffected, err := repo.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rowsAffected.Close()

	posts := make([]dto.GetPostWithAuthorResponse, 0)
	for rowsAffected.Next() {
		var post dto.GetPostWithAuthorResponse
		err = rowsAffected.Scan(
			&post.PostID,
			&post.TopicID,
			&post.Author.UserID,
			&post.PostDate,
			&post.Title,
			&post.Content,
			&post.Author.Username,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, rowsAffected.Err()
}
