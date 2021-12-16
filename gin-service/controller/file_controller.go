package controller

import (
	"fmt"
	"gin-service/helper"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type fileController struct {
}

func NewFileController() *fileController {
	return &fileController{
	}
}


func (s *fileController) DownloadFile(ctx *gin.Context) {
	path, err := filepath.Abs("file/give")
	if err != nil {
		log.Fatal(err)
	}
	fileName := "sample.pdf"

	ctx.FileAttachment(path+"/"+fileName, fileName)
}


func (s *fileController) UploadManyFile(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["file"]
	path, err := filepath.Abs("file/save")
	if err != nil {
		helper.DisplayError(ctx, "Error in Uploading Multiple file , keep the key = 'file' to avoid this error ", err)
	}
	for _, file := range files {
		log.Println(file.Filename)
		err := ctx.SaveUploadedFile(file, path+"/"+file.Filename)
		if err != nil {
			helper.DisplayError(ctx, "Error in Uploading Multiple file ,as the storage dir. not found ", err)
		}
	}

	fmt.Println(ctx.PostForm("key"))
	ctx.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

func (s *fileController) UploadSingleFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		helper.DisplayError(ctx, "Error in Uploading file, keep the key = 'file' to avoid this error ", err)
	}
	log.Println(file.Filename)

	path, err := filepath.Abs("file/save")
	if err != nil {
		helper.DisplayError(ctx, "Error in Uploading file, as the storage dir. not found ", err)
	}

	err = ctx.SaveUploadedFile(file, path+"/"+file.Filename)
	if err != nil {
		helper.DisplayError(ctx, "Error in Uploading Multiple file, as the storage dir. not found ", err)
	}
	ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func (s *fileController) UploadBinary(ctx *gin.Context) {
	byteBody, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		helper.DisplayError(ctx, "Error in Uploading binary file, cannot convert to byte ", err)
	}

	err = ioutil.WriteFile("file/save/myFilebinary.pdf", byteBody, 0644)
	if err != nil {
		helper.DisplayError(ctx, "Error in Uploading binary file, as the storage dir. not found ", err)
	}
	ctx.String(http.StatusOK, "got it")
}