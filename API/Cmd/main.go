package main

import (
	"fmt"
	"log"

	"chalas.com/forum_project/API/Config"
	"github.com/gin-gonic/gin"
)

func main() {
	chalas_router := gin.Default()
	chalas_config, err := Config.ConfigLoad()
	if err != nil {
		log.Fatal(err)
	}

	//chalas_router.Use(gin.Logger())
	//chalas_router.Use(gin.Recovery())
	chalas_router.Run(fmt.Sprintf("%v:%s", chalas_config.Chalas_Forum_Host, chalas_config.WebAPP_Port))
}
