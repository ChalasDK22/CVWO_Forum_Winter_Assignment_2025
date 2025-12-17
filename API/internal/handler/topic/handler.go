package topic

import (
	"net/http"
	"strconv"

	"chalas.com/forum_project/API/internal/dto"
	"chalas.com/forum_project/API/internal/service/topic"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type topicHandler struct {
	api          *gin.Engine
	validate     *validator.Validate
	topicService topic.TopicService
}

// New Handler
func NewTopicHandler(api *gin.Engine, validate *validator.Validate, topicService topic.TopicService) *topicHandler {
	return &topicHandler{
		api:          api,
		validate:     validate,
		topicService: topicService,
	}
}

// Create Topic
func (h *topicHandler) CreateTopic(c *gin.Context) {
	var (
		ctx    = c.Request.Context()
		req    dto.CreateTopicRequest
		userID = c.GetInt64("user_id")
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.validate.Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	topicID, statusCode, err := h.topicService.CreateTopic(ctx, &req, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(statusCode, dto.CreateTopicResponse{TopicID: topicID})
}

// Update Topic
func (h *topicHandler) UpdateTopic(c *gin.Context) {
	var (
		ctx     = c.Request.Context()
		req     dto.UpdateTopicRequest
		userID  = c.GetInt64("user_id")
		topicID = c.Param("topic_id")
	)
	newTopicID, err := strconv.ParseInt(topicID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.validate.Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	statusCode, err := h.topicService.UpdateTopic(ctx, &req, newTopicID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(statusCode, dto.UpdateTopicResponse{TopicID: newTopicID})
}

func (h *topicHandler) RouteList() {
	h.api.POST("/topics", h.CreateTopic)
	h.api.PUT("/topics/:topic_id", h.UpdateTopic)
}
