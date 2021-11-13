package main

import (
	"gin-service/controller"
	"gin-service/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	r := gin.Default()

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
