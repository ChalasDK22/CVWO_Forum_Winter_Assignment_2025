package main

import (
	"log"
	"net/http"
	"time"

	"chalas.com/forum_project/API/internal/config"
	commentHandler "chalas.com/forum_project/API/internal/handler/comment"
	commentRepos "chalas.com/forum_project/API/internal/repository/comment"
	commentService "chalas.com/forum_project/API/internal/service/comment"

	userHandler "chalas.com/forum_project/API/internal/handler/user"
	userRepos "chalas.com/forum_project/API/internal/repository/user"
	userService "chalas.com/forum_project/API/internal/service/user"

	postHandler "chalas.com/forum_project/API/internal/handler/post"
	postRepos "chalas.com/forum_project/API/internal/repository/post"
	postService "chalas.com/forum_project/API/internal/service/post"

	topicHandler "chalas.com/forum_project/API/internal/handler/topic"
	topicRepos "chalas.com/forum_project/API/internal/repository/topic"
	topicService "chalas.com/forum_project/API/internal/service/topic"
	"chalas.com/forum_project/API/pkg/internalsql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	chalasRouter := gin.Default()
	chalasValidate := validator.New()

	//Allow connections from the frontend
	chalasRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://chalasdk-cvwo-forum.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	//Config
	chalasConfig, err := config.ConfigLoad()
	if err != nil {

		log.Fatal(err)
	}

	chalasDB, err := internalsql.ConnectAPI_MYSQL(chalasConfig)
	if err != nil {
		log.Fatal(err)
	}

	chalasRouter.Use(gin.Logger())
	chalasRouter.Use(gin.Recovery())

	chalasRouter.GET("/checking-chalas-health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	chalasJWTKey := chalasConfig.Chalas_JWT

	//Repos
	chalasTopicRepos := topicRepos.NewTopicRepository(chalasDB)
	chalasPostRepos := postRepos.NewPostRepository(chalasDB)
	chalasUserRepos := userRepos.NewUserRepository(chalasDB)
	chalasCommentRepos := commentRepos.NewCommentRepository(chalasDB)

	//Services
	chalasTopicService := topicService.NewTopicService(chalasConfig, chalasTopicRepos)
	chalasPostService := postService.NewPostService(chalasConfig, chalasPostRepos)
	chalasUserService := userService.NewUserService(chalasConfig, chalasUserRepos)
	chalasCommentService := commentService.NewCommentService(chalasConfig, chalasCommentRepos)

	//Handlers
	chalasTopicHandler := topicHandler.NewTopicHandler(chalasRouter, chalasValidate, chalasTopicService, chalasPostService)
	chalasPostHandler := postHandler.NewPostHandler(chalasRouter, chalasValidate, chalasPostService, chalasCommentService)
	chalasUserHandler := userHandler.NewUserHandler(chalasRouter, chalasValidate, chalasUserService, chalasJWTKey)
	chalasCommentHandler := commentHandler.NewCommentHandler(chalasRouter, chalasValidate, chalasCommentService)

	chalasTopicHandler.RouteList(chalasJWTKey)
	chalasPostHandler.RouteList(chalasJWTKey)
	chalasUserHandler.RouteList()
	chalasCommentHandler.RouteList(chalasJWTKey)
	port := chalasConfig.WebAPP_Port
	if port == "" {
		port = "8080"
	}

	if err := chalasRouter.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
