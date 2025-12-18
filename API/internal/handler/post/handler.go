package post

import (
	"fmt"
	"net/http"
	"strconv"

	"chalas.com/forum_project/API/internal/dto"
	"chalas.com/forum_project/API/internal/service/post"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	api         *gin.Engine
	validate    *validator.Validate
	postService post.PostService
}

// New Handler
func NewPostHandler(api *gin.Engine, validate *validator.Validate, postService post.PostService) *Handler {
	return &Handler{
		api:         api,
		postService: postService,
		validate:    validate,
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
		c.JSON(http.StatusInternalServerError, gin.H{"error when convert post ID": fmt.Sprintf("updated post ID from %v encountered %v", postID, err.Error())})
		return
	}

	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when binding JSON": fmt.Sprintf("binding JSON post from %v encountered %v", req, err.Error())})
		return
	}

	err = h.validate.Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when validating": fmt.Sprintf("validatting post from %v encountered %v", req, err.Error())})
		return
	}

	statusCode, err := h.postService.UpdatePost(ctx, &req, userID, updatePostID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error when update": fmt.Sprintf("updating post encountered %v", err.Error())})
		return
	}

	c.JSON(statusCode, dto.UpdatePostResponse{PostID: updatePostID})
}

func (h *Handler) DeletePost(c *gin.Context) {
	var (
		ctx    = c.Request.Context()
		postID = c.Param("post_id")
	)

	deletedPostID, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error when convert post ID": fmt.Sprintf("deleted post ID from %v encountered %v", postID, err.Error())})
		return
	}

	statusCode, err := h.postService.DeletePost(ctx, deletedPostID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error when delete": fmt.Sprintf("deleting post encountered %v", err.Error())})
		return
	}

	c.JSON(statusCode, dto.DeletePostResponse{PostID: deletedPostID})
}

func (h *Handler) RouteList() {
	h.api.POST("/post", h.CreatePost)
	h.api.PUT("/post/:post_id/update", h.UpdatePost)
	h.api.DELETE("/post/:post_id/delete", h.DeletePost)
}
