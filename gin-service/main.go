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
	r.Use(middleware.Logger(), gindump.Dump())

	r.GET("/ping", func(c *gin.Context) {
		c.JSONP(200, gin.H{
			"message": "ok",
		})
	})

	api := r.Group("/api")
	api.Use(middleware.BasicAuth())
	{
		api.GET("/videos", func(c *gin.Context) {
			c.JSON(200, videoController.FindAll())
		})

		api.POST("/videos", func(c *gin.Context) {
			c.JSON(200, videoController.Save(c))
		})
	}

	r.Run(":8080")
}
