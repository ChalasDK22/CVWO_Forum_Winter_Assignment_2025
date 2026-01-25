package user

import (
	"fmt"
	"net/http"

	"chalas.com/forum_project/API/internal/dto"
	"chalas.com/forum_project/API/internal/middleware"
	"chalas.com/forum_project/API/internal/service/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	api          *gin.Engine
	validate     *validator.Validate
	jwtSecretKey string
	userService  user.UserService
}

// New Handler
func NewUserHandler(api *gin.Engine, validate *validator.Validate, userService user.UserService, jwtKey string) *Handler {
	return &Handler{
		api:          api,
		validate:     validate,
		jwtSecretKey: jwtKey,
		userService:  userService,
	}
}

// Create User
func (h *Handler) RegisterUser(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req dto.RegisterRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when binding JSON": fmt.Sprintf("binding JSON user from %v encountered %v", req, err.Error())})
		return
	}

	err = h.validate.Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when validating": fmt.Sprintf("validating user from %v encountered %v", req, err.Error())})
		return
	}

	userID, statusCode, err := h.userService.Register(ctx, &req)
	if err != nil {
		c.JSON(statusCode, gin.H{"error when create": fmt.Sprintf("creating user encountered %v", err.Error())})
		return
	}

	c.JSON(statusCode, dto.RegisterResponse{
		UserID: userID,
	})
}

func (h *Handler) Login(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req dto.LoginRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when binding JSON": fmt.Sprintf("binding JSON user from %v encountered %v", req, err.Error())})
		return
	}

	err = h.validate.Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error when validating": fmt.Sprintf("validating user from %v encountered %v", req, err.Error())})
		return
	}

	token, statusCode, err := h.userService.Login(ctx, &req)
	if err != nil {
		c.JSON(statusCode, gin.H{"error when logging in": fmt.Sprintf("log in user encountered %v", err.Error())})
		return
	}

	userInfo, err := h.userService.GetUserByUsername(ctx, req.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error when login": fmt.Sprintf("getting user encountered %v", err.Error())})
		return
	}

	c.JSON(statusCode, dto.LoginResponse{
		Token:    token,
		UserID:   userInfo.UserID,
		Username: userInfo.Username,
	})
}

func (h *Handler) Profile(c *gin.Context) {
	ctx := c.Request.Context()

	userIDanyType, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error when getting user": "unauthorized"})
		return
	}

	userID := userIDanyType.(int64)
	user, err := h.userService.GetUserByID(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error when getting user": "Cannot get user by ID"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Route
func (h *Handler) RouteList() {
	authRoute := h.api.Group("/auth")
	authRoute.POST("/register", h.RegisterUser)
	authRoute.POST("/login", h.Login)

	userRoute := h.api.Group("/user")
	userRoute.Use(middleware.AuthMiddleware(h.jwtSecretKey))
	userRoute.GET("/profile", h.Profile)
}
