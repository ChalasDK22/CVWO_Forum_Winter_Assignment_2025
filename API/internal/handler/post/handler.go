package post

import (
	"fmt"
	"net/http"
	"strconv"

	"chalas.com/forum_project/API/internal/dto"
	"chalas.com/forum_project/API/internal/middleware"
	"chalas.com/forum_project/API/internal/service/comment"
	"chalas.com/forum_project/API/internal/service/post"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	api            *gin.Engine
	validate       *validator.Validate
	postService    post.PostService
	commentService comment.CommentService
}

// New Handler
func NewPostHandler(api *gin.Engine, validate *validator.Validate, postService post.PostService, commentService comment.CommentService) *Handler {
	return &Handler{
		api:            api,
		postService:    postService,
		commentService: commentService,
		validate:       validate,
	}
}

func (h *Handler) CreatePost(c *gin.Context) {
	var (
		ctx    = c.Request.Context()
		req    dto.CreatePostRequest
		userID = c.GetInt64("user_id")
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when binding JSON": fmt.Sprintf("binding JSON post from %v encountered %v", req, err.Error())})
		return
	}

	err = h.validate.Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when validating": fmt.Sprintf("validatting post from %v encountered %v", req, err.Error())})
		return
	}

	postID, statusCode, err := h.postService.CreatePost(ctx, &req, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error when create": fmt.Sprintf("creating post encountered %v", err.Error())})
		return
	}
	c.JSON(statusCode, dto.CreatePostResponse{PostID: postID})
}

func (h *Handler) UpdatePost(c *gin.Context) {
	var (
		ctx    = c.Request.Context()
		req    dto.UpdatePostRequest
		userID = c.GetInt64("user_id")
		postID = c.Param("post_id")
	)

	updatePostID, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when convert post ID": fmt.Sprintf("updated post ID from %v encountered %v", postID, err.Error())})
		return
	}

	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when binding JSON": fmt.Sprintf("binding JSON post from %v encountered %v", req, err.Error())})
		return
	}

	err = h.validate.Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when validating": fmt.Sprintf("validating post from %v encountered %v", req, err.Error())})
		return
	}

	statusCode, err := h.postService.UpdatePost(ctx, &req, userID, updatePostID)
	if err != nil {
		switch err {
		case post.ErrPostNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error when update": fmt.Sprintf("post does not exist")})
			return
		case post.ErrNoPermission:
			c.JSON(http.StatusForbidden, gin.H{"error when update": fmt.Sprintf("permission denied")})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error when update": fmt.Sprintf("updating post encountered %v", err.Error())})
			return
		}
	}

	c.JSON(statusCode, dto.UpdatePostResponse{PostID: updatePostID})
}

func (h *Handler) DeletePost(c *gin.Context) {
	var (
		ctx    = c.Request.Context()
		postID = c.Param("post_id")
		userID = c.GetInt64("user_id")
	)

	deletedPostID, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when convert post ID": fmt.Sprintf("deleted post ID from %v encountered %v", postID, err.Error())})
		return
	}

	statusCode, err := h.postService.DeletePost(ctx, deletedPostID, userID)
	if err != nil {
		switch err {
		case post.ErrPostNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error when update": fmt.Sprintf("post does not exist")})
			return
		case post.ErrNoPermission:
			c.JSON(http.StatusForbidden, gin.H{"error when update": fmt.Sprintf("permission denied")})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error when update": fmt.Sprintf("deleting post encountered %v", err.Error())})
			return
		}
	}

	c.JSON(statusCode, dto.DeletePostResponse{PostID: deletedPostID})
}

func (h *Handler) GetPosts(c *gin.Context) {
	var (
		ctx = c.Request.Context()
	)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	posts, total, err := h.postService.GetPosts(ctx, page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalPages := (total + size - 1) / size
	c.JSON(http.StatusOK, gin.H{
		"posts":      posts,
		"page":       page,
		"size":       size,
		"total":      total,
		"totalPages": totalPages,
	})
}

func (h *Handler) GetCommentByPostID(c *gin.Context) {
	var (
		ctx    = c.Request.Context()
		postID = c.Param("post_id")
	)

	newPostID, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when convert post ID": fmt.Sprintf("convert post ID from %v encountered %v", postID, err.Error())})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	comments, total, err := h.commentService.GetCommentByPostID(ctx, newPostID, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"post_id":  newPostID,
		"page":     page,
		"limit":    limit,
		"total":    total,
		"comments": comments,
	})
}

func (h *Handler) GetPostByPostID(c *gin.Context) {
	var (
		ctx    = c.Request.Context()
		postID = c.Param("post_id")
	)

	newPostID, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when convert post ID": fmt.Sprintf("convert postID from %v encountered %v", postID, err.Error())})
		return
	}

	post, err := h.postService.GetPostsByPostID(ctx, newPostID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"topic_id": newPostID,
		"post":     post,
	})
}

func (h *Handler) RouteList(secretKey string) {
	routeAuth := h.api.Group("/posts")
	routeAuth.Use(middleware.AuthMiddleware(secretKey))
	routeAuth.POST("/", h.CreatePost)
	routeAuth.GET("/:post_id", h.GetPostByPostID)
	routeAuth.PUT("/:post_id/update", h.UpdatePost)
	routeAuth.DELETE("/:post_id/delete", h.DeletePost)
	routeAuth.GET("/api/posts", h.GetPosts)
	routeAuth.GET("/:post_id/comments", h.GetCommentByPostID)
}
