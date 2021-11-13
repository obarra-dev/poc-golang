package main

import (
	"gin-service/controller"
	"gin-service/middleware"
	"gin-service/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    = service.New()
	videoController = controller.New(videoService)
)

func main() {
	r := gin.Default()
	r.Use(middleware.BasicAuth())
	r.Use(middleware.Logger())
	r.Use(gindump.Dump())

	r.GET("/videos", func(c *gin.Context) {
		c.JSON(200, videoController.FindAll())
	})

	r.POST("/videos", func(c *gin.Context) {
		c.JSON(200, videoController.Save(c))
	})

	r.GET("/test", func(c *gin.Context) {
		c.JSONP(200, gin.H{
			"message": "ok",
		})
	})


	r.Run(":8080")
}
