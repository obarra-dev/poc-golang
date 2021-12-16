package controller

import (
	"gin-service/entity"
	"gin-service/service"
	"gin-service/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type VideoController interface {
	FindAll(ctx *gin.Context) []entity.Video
	FindOne(ctx *gin.Context) entity.Video
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type videoController struct {
	service service.VideoService
}

var validate *validator.Validate

func NewVideoController(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &videoController{
		service: service,
	}
}

func (c *videoController) FindOne(ctx *gin.Context) entity.Video {
	return c.service.FindOne(ctx.Param("GUID"))
}

func (c *videoController) FindAll(ctx *gin.Context) []entity.Video {
	return c.service.FindWithFilter(ctx.Query("title"))
}

func (c *videoController) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.BindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}

func (c *videoController) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
