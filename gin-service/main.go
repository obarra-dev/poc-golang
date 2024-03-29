package main

import (
	"gin-service/controller"
	"gin-service/middleware"
	"gin-service/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"net/http"
)

var (
	videoService    = service.New()
	videoController = controller.NewVideoController(videoService)
	fileService = controller.NewFileController()
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
		api.GET("/videos/:GUID", func(c *gin.Context) {
			c.JSON(http.StatusOK, videoController.FindOne(c))
		})
		api.GET("/videos", func(c *gin.Context) {
			c.JSON(http.StatusOK, videoController.FindAll(c))
		})

		api.POST("/videos", func(c *gin.Context) {
			if err := videoController.Save(c); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusCreated, gin.H{})
			}
		})
	}

	fileAPI := r.Group("/files")
	{
		fileAPI.GET("/download", fileService.DownloadFile)
		fileAPI.POST("/uploadmany", fileService.UploadManyFile)
		fileAPI.POST("/uploadone", fileService.UploadSingleFile)
		fileAPI.POST("/uploadbinary", fileService.UploadBinary)

	}

	r.Static("/css", "./templates/css")
	r.LoadHTMLGlob("templates/*.html")
	view := r.Group("/view")
	{
		view.GET("/videos", videoController.ShowAll)
	}

	r.Run(":8080")
}
