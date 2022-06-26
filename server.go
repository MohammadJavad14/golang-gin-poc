package main

import (
	"github.com/MohammadJavad14/golang-gin-poc.git/controllers"
	"github.com/MohammadJavad14/golang-gin-poc.git/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService        = service.New()
	videoController controllers.VideoController = controllers.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}
