package post

import (
	"context"
	"net/http"
	"time"

	"chalas.com/forum_project/API/internal/config"
	"chalas.com/forum_project/API/internal/dto"
	"chalas.com/forum_project/API/internal/models"
	"chalas.com/forum_project/API/internal/repository/post"
)

type PostService interface {
	CreatePost(ctx context.Context, dtoReq *dto.CreatePostRequest, userID int64) (int64, int, error)
	UpdatePost(ctx context.Context, dtoReq *dto.UpdatePostRequest, userID int64, postID int64) (int, error)
	DeletePost(ctx context.Context, postID int64) (int, error)
}

type postService struct {
	cfg       *config.Config
	postrepos post.PostRepository
}

// New Post
func NewPostService(cfg *config.Config, postrepos post.PostRepository) PostService {
	return &postService{
		cfg:       cfg,
		postrepos: postrepos,
	}
}

func (service *postService) CreatePost(ctx context.Context, dtoReq *dto.CreatePostRequest, userID int64) (int64, int, error) {
	postTime := time.Now()
	postInsertedID, err := service.postrepos.InsertPost(ctx, &models.PostModel{
		UserID:   userID,
		TopicID:  dtoReq.TopicID,
		Content:  dtoReq.Content,
		Title:    dtoReq.Title,
		PostDate: postTime,
	})
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	return postInsertedID, http.StatusOK, nil
}

func (service *postService) UpdatePost(ctx context.Context, dtoReq *dto.UpdatePostRequest, userID int64, postID int64) (int, error) {
	updateTime := time.Now()
	err := service.postrepos.UpdatePost(ctx, postID, &models.PostModel{
		UserID:   userID,
		TopicID:  dtoReq.TopicID,
		Content:  dtoReq.Content,
		PostDate: updateTime,
		Title:    dtoReq.Title,
	})
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (service *postService) DeletePost(ctx context.Context, postID int64) (int, error) {
	err := service.postrepos.DeletePost(ctx, postID)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
