package comment

import (
	"fmt"
	"net/http"
	"strconv"

	"chalas.com/forum_project/API/internal/dto"
	"chalas.com/forum_project/API/internal/middleware"
	"chalas.com/forum_project/API/internal/service/comment"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	api            *gin.Engine
	validate       *validator.Validate
	commentService comment.CommentService
}

// New Handler
func NewCommentHandler(api *gin.Engine, validate *validator.Validate, commentService comment.CommentService) *Handler {
	return &Handler{
		api:            api,
		validate:       validate,
		commentService: commentService,
	}
}

func (h *Handler) CreateComment(c *gin.Context) {
	var (
		ctx    = c.Request.Context()
		req    dto.CreateCommentRequest
		userID = c.GetInt64("user_id")
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when binding JSON": fmt.Sprintf("binding JSON comment from %v encountered %v", req, err.Error())})
		return
	}

	err = h.validate.Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when validating": fmt.Sprintf("validatting comment from %v encountered %v", req, err.Error())})
		return
	}

	commentID, statusCode, err := h.commentService.CreateComment(ctx, &req, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error when create": fmt.Sprintf("creating comment encountered %v", err.Error())})
		return
	}
	c.JSON(statusCode, dto.CreateCommentResponse{CommentID: commentID})
}

func (h *Handler) UpdateComment(c *gin.Context) {
	var (
		ctx       = c.Request.Context()
		req       dto.UpdateCommentRequest
		userID    = c.GetInt64("user_id")
		commentID = c.Param("comment_id")
	)

	updateCommentID, err := strconv.ParseInt(commentID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error when convert comment ID": fmt.Sprintf("updated comment ID from %v encountered %v", commentID, err.Error())})
		return
	}

	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when binding JSON": fmt.Sprintf("binding JSON comment from %v encountered %v", req, err.Error())})
		return
	}

	err = h.validate.Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when validating": fmt.Sprintf("validatting comment from %v encountered %v", req, err.Error())})
		return
	}

	statusCode, err := h.commentService.UpdateComment(ctx, &req, userID, updateCommentID)
	if err != nil {
		switch err {
		case comment.ErrCommentNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error when update": fmt.Sprintf("comment does not exist")})
			return
		case comment.ErrNoPermission:
			c.JSON(http.StatusForbidden, gin.H{"error when update": fmt.Sprintf("permission denied")})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error when update": fmt.Sprintf("updating comment encountered %v", err.Error())})
			return
		}
	}

	c.JSON(statusCode, dto.UpdateCommentResponse{CommentID: updateCommentID})
}

func (h *Handler) DeleteComment(c *gin.Context) {
	var (
		ctx       = c.Request.Context()
		commentID = c.Param("comment_id")
		userID    = c.GetInt64("user_id")
	)

	deletedCommentID, err := strconv.ParseInt(commentID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error when convert comment ID": fmt.Sprintf("deleted comment ID from %v encountered %v", commentID, err.Error())})
		return
	}

	statusCode, err := h.commentService.DeleteComment(ctx, deletedCommentID, userID)
	if err != nil {
		switch err {
		case comment.ErrCommentNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error when update": fmt.Sprintf("comment does not exist")})
			return
		case comment.ErrNoPermission:
			c.JSON(http.StatusForbidden, gin.H{"error when update": fmt.Sprintf("permission denied")})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error when update": fmt.Sprintf("deleting comment encountered %v", err.Error())})
			return
		}
	}

	c.JSON(statusCode, dto.DeleteCommentResponse{CommentID: deletedCommentID})
}

func (h *Handler) RouteList(secretKey string) {
	routeAuth := h.api.Group("/comments")
	routeAuth.Use(middleware.AuthMiddleware(secretKey))
	routeAuth.POST("/", h.CreateComment)
	routeAuth.PUT("/:comment_id/update", h.UpdateComment)
	routeAuth.DELETE("/:comment_id/delete", h.DeleteComment)
}
