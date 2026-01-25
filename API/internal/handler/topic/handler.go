package topic

import (
	"fmt"
	"net/http"
	"strconv"

	"chalas.com/forum_project/API/internal/dto"
	"chalas.com/forum_project/API/internal/middleware"
	"chalas.com/forum_project/API/internal/service/post"
	"chalas.com/forum_project/API/internal/service/topic"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	api          *gin.Engine
	validate     *validator.Validate
	topicService topic.TopicService
	postService  post.PostService
}

// New Handler
func NewTopicHandler(api *gin.Engine, validate *validator.Validate, topicService topic.TopicService, postService post.PostService) *Handler {
	return &Handler{
		api:          api,
		validate:     validate,
		topicService: topicService,
		postService:  postService,
	}
}

// Create Topic
func (h *Handler) CreateTopic(c *gin.Context) {
	var (
		ctx    = c.Request.Context()
		req    dto.CreateTopicRequest
		userID = c.GetInt64("user_id")
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when binding JSON": fmt.Sprintf("binding JSON topic from %v encountered %v", req, err.Error())})
		return
	}

	err = h.validate.Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when validating": fmt.Sprintf("validating topic from %v encountered %v", req, err.Error())})
		return
	}

	topicID, statusCode, err := h.topicService.CreateTopic(ctx, &req, userID)
	if err != nil {
		c.JSON(statusCode, gin.H{"error when create": fmt.Sprintf("creating topic encountered %v", err.Error())})
		return
	}
	c.JSON(statusCode, dto.CreateTopicResponse{TopicID: topicID})
	return
}

// Update Topic
func (h *Handler) UpdateTopic(c *gin.Context) {
	var (
		ctx     = c.Request.Context()
		req     dto.UpdateTopicRequest
		userID  = c.GetInt64("user_id")
		topicID = c.Param("topic_id")
	)
	newTopicID, err := strconv.ParseInt(topicID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when convert topic ID": fmt.Sprintf("updated topic ID from %v encountered %v", topicID, err.Error())})
	}

	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when binding JSON": fmt.Sprintf("binding JSON topic from %v encountered %v", req, err.Error())})
		return
	}

	err = h.validate.Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when validating": fmt.Sprintf("validating topic from %v encountered %v", req, err.Error())})
		return
	}

	statusCode, err := h.topicService.UpdateTopic(ctx, &req, newTopicID, userID)
	if err != nil {
		switch err {
		case topic.ErrTopicNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error when update": fmt.Sprintf("topic does not exist")})
			return
		case topic.ErrNoPermission:
			c.JSON(http.StatusForbidden, gin.H{"error when update": fmt.Sprintf("permission denied")})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error when update": fmt.Sprintf("updating topic encountered %v", err.Error())})
			return
		}
	}
	c.JSON(statusCode, dto.UpdateTopicResponse{TopicID: newTopicID})
}

// Delete Topic
func (h *Handler) DeleteTopic(c *gin.Context) {
	var (
		ctx     = c.Request.Context()
		topicID = c.Param("topic_id")
		userID  = c.GetInt64("user_id")
	)

	deletedTopicID, err := strconv.ParseInt(topicID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when convert topic ID": fmt.Sprintf("deleted topic ID from %v encountered %v", topicID, err.Error())})
		return
	}

	statusCode, err := h.topicService.DeleteTopic(ctx, deletedTopicID, userID)
	if err != nil {
		switch err {
		case topic.ErrTopicNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error when delete": fmt.Sprintf("topic does not exist")})
			return
		case topic.ErrNoPermission:
			c.JSON(http.StatusForbidden, gin.H{"error when delete": fmt.Sprintf("permission denied")})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error when ": fmt.Sprintf("deleting topic encountered %v", err.Error())})
			return
		}
	}
	c.JSON(statusCode, dto.DeleteTopicResponse{TopicID: deletedTopicID})
}

func (h *Handler) GetPostByTopicID(c *gin.Context) {
	var (
		ctx     = c.Request.Context()
		topicID = c.Param("topic_id")
	)

	newTopicID, err := strconv.ParseInt(topicID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when convert topic ID": fmt.Sprintf("convert topicID from %v encountered %v", topicID, err.Error())})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.DefaultQuery("search", "")

	posts, total, err := h.postService.GetPostsByTopicID(ctx, newTopicID, search, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"topic_id": newTopicID,
		"page":     page,
		"limit":    limit,
		"total":    total,
		"posts":    posts,
	})
}

func (h *Handler) GetTopics(c *gin.Context) {
	var (
		ctx = c.Request.Context()
	)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	topics, total, err := h.topicService.GetTopics(ctx, page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalPages := (total + size - 1) / size
	c.JSON(http.StatusOK, gin.H{
		"topics":     topics,
		"page":       page,
		"size":       size,
		"total":      total,
		"totalPages": totalPages,
	})
}

func (h *Handler) GetTopicByTopicID(c *gin.Context) {
	var (
		ctx     = c.Request.Context()
		topicID = c.Param("topic_id")
	)

	newTopicID, err := strconv.ParseInt(topicID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when convert topic ID": fmt.Sprintf("convert topicID from %v encountered %v", topicID, err.Error())})
		return
	}

	topic, err := h.topicService.GetTopicByTopicID(ctx, newTopicID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"topic_id": newTopicID,
		"topic":    topic,
	})
}

// Route
func (h *Handler) RouteList(secretKey string) {
	routeAuth := h.api.Group("/topics")
	routeAuth.Use(middleware.AuthMiddleware(secretKey))
	routeAuth.POST("/", h.CreateTopic)
	routeAuth.PUT("/:topic_id/update", h.UpdateTopic)
	routeAuth.DELETE("/:topic_id/delete", h.DeleteTopic)
	routeAuth.GET("/:topic_id/posts", h.GetPostByTopicID)
	routeAuth.GET("/api/topics", h.GetTopics)
	routeAuth.GET("/:topic_id", h.GetTopicByTopicID)
}
