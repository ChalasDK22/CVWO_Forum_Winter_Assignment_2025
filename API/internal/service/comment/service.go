package comment

import (
	"context"
	"errors"
	"net/http"
	"time"

	"chalas.com/forum_project/API/internal/config"
	"chalas.com/forum_project/API/internal/dto"
	"chalas.com/forum_project/API/internal/models"
	"chalas.com/forum_project/API/internal/repository/comment"
)

type CommentService interface {
	CreateComment(ctx context.Context, dtoReq *dto.CreateCommentRequest, userID int64) (int64, int, error)
	UpdateComment(ctx context.Context, dtoReq *dto.UpdateCommentRequest, userID int64, commentID int64) (int, error)
	DeleteComment(ctx context.Context, commentID int64, userID int64) (int, error)
	GetCommentByPostID(ctx context.Context, postID int64, page int, limit int) ([]dto.GetCommentWithAuthorResponse, int64, error)
}

type commentService struct {
	cfg          *config.Config
	commentrepos comment.CommentRepository
}

func NewCommentService(cfg *config.Config, commentrepos comment.CommentRepository) CommentService {
	return &commentService{
		cfg:          cfg,
		commentrepos: commentrepos,
	}
}

var (
	ErrNoPermission    = errors.New("Unauthorized permission")
	ErrCommentNotFound = errors.New("Comment not found")
)

func (service *commentService) CreateComment(ctx context.Context, dtoReq *dto.CreateCommentRequest, userID int64) (int64, int, error) {
	postDate := time.Now()
	commentInsertedID, err := service.commentrepos.CreateComment(ctx, &models.CommentModel{
		UserID:   userID,
		Content:  dtoReq.Content,
		PostDate: postDate,
		PostID:   dtoReq.PostID,
	})
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	return commentInsertedID, http.StatusOK, nil
}

func (service *commentService) UpdateComment(ctx context.Context, dtoReq *dto.UpdateCommentRequest, userID int64, commentID int64) (int, error) {
	ownerID, err := service.commentrepos.GetCommentOwnerID(ctx, commentID)
	if err != nil {
		return http.StatusInternalServerError, ErrCommentNotFound
	}

	if ownerID != userID {
		return http.StatusForbidden, ErrNoPermission
	}

	updateTime := time.Now()
	err = service.commentrepos.UpdateComment(ctx, commentID, &models.CommentModel{
		UserID:   userID,
		PostID:   dtoReq.PostID,
		Content:  dtoReq.Content,
		PostDate: updateTime,
	})
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (service *commentService) DeleteComment(ctx context.Context, commentID int64, userID int64) (int, error) {
	ownerID, err := service.commentrepos.GetCommentOwnerID(ctx, commentID)
	if err != nil {
		return http.StatusInternalServerError, ErrCommentNotFound
	}

	if ownerID != userID {
		return http.StatusForbidden, ErrNoPermission
	}

	err = service.commentrepos.DeleteComment(ctx, commentID, userID)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (service *commentService) GetCommentByPostID(ctx context.Context, postID int64, page int, limit int) ([]dto.GetCommentWithAuthorResponse, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	totalComment, err := service.commentrepos.CountComment(ctx, postID)
	if err != nil {
		return nil, 0, err
	}

	comments, err := service.commentrepos.GetCommentByPostID(ctx, postID, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return comments, totalComment, nil
}
