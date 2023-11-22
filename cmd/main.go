package main

import (
	"prac/controller"
	"prac/db_client"

	"github.com/gin-gonic/gin"
)

func main() {
	dbclient.InitialiseDBConnection()

	router := gin.Default()

	router.POST("/", controller.CreatePost)
	router.GET("/:id", controller.GetPosts)
	router.GET("/my-id", controller.GetPost)
	router.DELETE("/:id", controller.DelPost)
	router.PUT("/")

	if err := router.Run(":5000"); err != nil {
		panic(err.Error())
	}

}
