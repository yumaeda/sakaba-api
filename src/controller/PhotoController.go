package controller

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yumaeda/sakaba-api/src/model"
	"github.com/yumaeda/sakaba-api/src/repository"
	"github.com/yumaeda/sakaba-api/src/service"
)

// PhotoController is a controller for Photo API.
type PhotoController struct {
	Repository repository.PhotoRepository
}

// GetPhotosByRestaurantID returns all the photos for the specified restaurant.
func (c PhotoController) GetPhotosByRestaurantID(ctx *gin.Context) {
	photos := c.Repository.GetPhotosByRestaurantID(ctx.Param("id"))

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       photos,
	})
}

// GetLatestPhotos returns latest photos.
func (c PhotoController) GetLatestPhotos(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"body": c.Repository.GetLatestPhotos(),
	})
}

// AddPhoto uploads the specified photo to the specified restaurant.
func (c PhotoController) AddPhoto(ctx *gin.Context) {
	var errorMessage string
	var json model.PhotoRequest
	if err := ctx.ShouldBindJSON(&json); err == nil {
		base64Data := json.FileContent[strings.IndexByte(json.FileContent, ',')+1:]
		file, fileErr := base64.StdEncoding.DecodeString(base64Data)
		if fileErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fileErr.Error()})
			return
		}

		fileName := uuid.New().String()
		s3Service := service.S3Service{}
		uploadErr := s3Service.Upload(json.RestaurantID, fileName, file)
		if uploadErr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": uploadErr.Error(),
			})
			return
		}

		dbError := c.Repository.AddPhoto(json.RestaurantID, fileName)
		if dbError == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"statusCode": 200,
				"body":       "New photo is added.",
			})
			return
		} else {
			errorMessage = dbError.Error()
		}
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"statusCode": 400,
		"error":      "Photo insertion failed [" + errorMessage + "].",
	})
}
