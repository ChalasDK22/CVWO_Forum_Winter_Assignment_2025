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
	chalas_router := gin.Default()
	chalas_validate := validator.New()

	//Config
	chalas_config, err := config.ConfigLoad()
	if err != nil {

		log.Fatal(err)
	}

	chalas_db, err := internalsql.ConnectAPI_MYSQL(chalas_config)
	if err != nil {
		log.Fatal(err)
	}

	chalas_router.Use(gin.Logger())
	chalas_router.Use(gin.Recovery())

	chalas_router.GET("/checking-chalas-health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	//Repos
	chalasTopicRepos := topicRepos.NewTopicRepository(chalas_db)
	chalasPostRepos := postRepos.NewPostRepository(chalas_db)

	//Services
	chalasTopicService := topicService.NewTopicService(chalas_config, chalasTopicRepos)
	chalasPostService := postService.NewPostService(chalas_config, chalasPostRepos)

	//Handlers
	chalasTopicHandler := topicHandler.NewTopicHandler(chalas_router, chalas_validate, chalasTopicService)
	chalasPostHandler := postHandler.NewPostHandler(chalas_router, chalas_validate, chalasPostService)

	chalasTopicHandler.RouteList()
	chalasPostHandler.RouteList()
	chalas_router.Run(fmt.Sprintf("%v:%s", chalas_config.Chalas_Forum_Host, chalas_config.WebAPP_Port))
	fmt.Println((chalasTopicHandler))
}
