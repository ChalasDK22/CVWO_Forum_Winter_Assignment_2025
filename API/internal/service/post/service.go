package post

import (
	"context"
	"errors"
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
	DeletePost(ctx context.Context, postID int64, userID int64) (int, error)
	GetPostsByTopicID(ctx context.Context, topicID int64, search string, page int, limit int) ([]dto.GetPostWithAuthorResponse, int64, error)
	GetPosts(ctx context.Context, page int, size int) ([]dto.GetPostWithAuthorResponse, int, error)
	GetPostsByPostID(ctx context.Context, postID int64) ([]dto.GetPostWithAuthorResponse, error)
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

var (
	ErrNoPermission = errors.New("Unauthorized permission")
	ErrPostNotFound = errors.New("Post not found")
)

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
	ownerID, err := service.postrepos.GetPostOwnerID(ctx, postID)
	if err != nil {
		return http.StatusInternalServerError, ErrPostNotFound
	}

	if ownerID != userID {
		return http.StatusForbidden, ErrNoPermission
	}

	updateTime := time.Now()
	err = service.postrepos.UpdatePost(ctx, postID, &models.PostModel{
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

func (service *postService) DeletePost(ctx context.Context, postID int64, userID int64) (int, error) {
	ownerID, err := service.postrepos.GetPostOwnerID(ctx, postID)
	if err != nil {
		return http.StatusInternalServerError, ErrPostNotFound
	}

	if ownerID != userID {
		return http.StatusForbidden, ErrNoPermission
	}

	err = service.postrepos.DeletePost(ctx, postID, userID)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (service *postService) GetPostsByTopicID(ctx context.Context, topicID int64, search string, page int, limit int) ([]dto.GetPostWithAuthorResponse, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	totalpost, err := service.postrepos.CountPostsByTopicID(ctx, topicID, search)
	if err != nil {
		return nil, 0, err
	}

	posts, err := service.postrepos.GetPostByTopicID(ctx, topicID, search, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return posts, totalpost, nil
}

func (service *postService) GetPosts(ctx context.Context, page int, size int) ([]dto.GetPostWithAuthorResponse, int, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 50 {
		size = 10
	}

	offset := (page - 1) * size

	total, err := service.postrepos.CountPosts(ctx)
	if err != nil {
		return nil, 0, err
	}

	posts, err := service.postrepos.GetPosts(ctx, size, offset)
	if err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

func (service *postService) GetPostsByPostID(ctx context.Context, postID int64) ([]dto.GetPostWithAuthorResponse, error) {
	posts, err := service.postrepos.GetPostByPostID(ctx, postID)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
