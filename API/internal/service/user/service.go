package user

import (
	"context"
	"errors"
	"net/http"
	"time"

	"chalas.com/forum_project/API/internal/config"
	"chalas.com/forum_project/API/internal/dto"
	"chalas.com/forum_project/API/internal/models"
	"chalas.com/forum_project/API/internal/repository/user"
	"chalas.com/forum_project/API/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (int64, int, error)
	Login(ctx context.Context, req *dto.LoginRequest) (string, int, error)
	GetUserByID(ctx context.Context, userID int64) (*models.UserModel, error)
	GetUserByUsername(ctx context.Context, username string) (*models.UserModel, error)
}

type userService struct {
	cfg       *config.Config
	userrepos user.UserRepository
}

func NewUserService(cfg *config.Config, userrepos user.UserRepository) UserService {
	return &userService{
		cfg:       cfg,
		userrepos: userrepos}
}

func (service *userService) Register(ctx context.Context, req *dto.RegisterRequest) (int64, int, error) {
	userCheckExist, err := service.userrepos.GetUser(ctx, req.Username)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	if userCheckExist != nil {
		return 0, http.StatusBadRequest, errors.New("User already exists. Please log in!")
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	now := time.Now()
	userTypedInfoModel := &models.UserModel{
		Username:         req.Username,
		Password:         string(hasedPassword),
		RegistrationDate: now,
	}

	userID, err := service.userrepos.CreateUser(ctx, userTypedInfoModel)
	if err != nil {
		return 0, http.StatusInternalServerError, errors.New("Failed to create user")
	}

	return userID, http.StatusCreated, nil
}

func (service *userService) Login(ctx context.Context, req *dto.LoginRequest) (string, int, error) {
	userCheckExist, err := service.userrepos.GetUser(ctx, req.Username)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	if userCheckExist == nil {
		return "", http.StatusNotFound, errors.New("User does not exist. Please try again!")
	}
	err = bcrypt.CompareHashAndPassword([]byte(userCheckExist.Password), []byte(req.Password))
	if err != nil {
		return "", http.StatusUnauthorized, errors.New("Password is incorrect! Please try again!")
	}

	token, err := jwt.CreateJWTToken(userCheckExist.UserID, userCheckExist.Username, service.cfg.Chalas_JWT)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return token, http.StatusOK, nil
}

func (service *userService) GetUserByID(ctx context.Context, userID int64) (*models.UserModel, error) {
	return service.userrepos.GetUserByID(ctx, userID)
}

func (service *userService) GetUserByUsername(ctx context.Context, username string) (*models.UserModel, error) {
	return service.userrepos.GetUser(ctx, username)
}
