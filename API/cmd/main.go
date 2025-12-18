package main

import (
	"fmt"
	"log"
	"net/http"

	"chalas.com/forum_project/API/internal/config"
	postHandler "chalas.com/forum_project/API/internal/handler/post"
	postRepos "chalas.com/forum_project/API/internal/repository/post"
	postService "chalas.com/forum_project/API/internal/service/post"

	topicHandler "chalas.com/forum_project/API/internal/handler/topic"
	topicRepos "chalas.com/forum_project/API/internal/repository/topic"
	topicService "chalas.com/forum_project/API/internal/service/topic"
	"chalas.com/forum_project/API/pkg/internalsql"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	chalasRouter := gin.Default()
	chalasValidate := validator.New()

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

	//Repos
	chalasTopicRepos := topicRepos.NewTopicRepository(chalasDB)
	chalasPostRepos := postRepos.NewPostRepository(chalasDB)

	//Services
	chalasTopicService := topicService.NewTopicService(chalasConfig, chalasTopicRepos)
	chalasPostService := postService.NewPostService(chalasConfig, chalasPostRepos)

	//Handlers
	chalasTopicHandler := topicHandler.NewTopicHandler(chalasRouter, chalasValidate, chalasTopicService)
	chalasPostHandler := postHandler.NewPostHandler(chalasRouter, chalasValidate, chalasPostService)

	chalasTopicHandler.RouteList()
	chalasPostHandler.RouteList()
	_ = chalasRouter.Run(fmt.Sprintf("%v:%s", chalasConfig.Chalas_Forum_Host, chalasConfig.WebAPP_Port))
}
