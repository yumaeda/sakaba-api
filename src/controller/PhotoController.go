package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"sakaba.link/api/src/repository"
	"sakaba.link/api/src/service"
)

// PhotoController is a controller for Photo API.
type PhotoController struct{}

// GetAllPhotos returns all the photos.
func (c *PhotoController) GetAllPhotos(ctx *gin.Context) {
	photoRepository := repository.PhotoRepository{}
	allPhotos := photoRepository.GetAllPhotos()

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       allPhotos,
	})
}

// AddPhoto uploads the specified photo to the specified restaurant.
func (c *PhotoController) AddPhoto(ctx *gin.Context) {
	restaurantID := ctx.Param("restaurant_id")
	file, header, fileErr := ctx.Request.FormFile("file_content")
	if fileErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"statusCode": 400,
			"fileName":   header.Filename,
			"error":      fileErr.Error(),
		})
		return
	}

	s3Service := service.S3Service{}
	fileName := uuid.New().String()
	up, uploadErr := s3Service.Upload(restaurantID, fileName, file)
	if uploadErr == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"statusCode": 500,
			"error":      "Failed to upload file",
			"uploader":   up,
		})
		return
	}

	photoRepository := repository.PhotoRepository{}
	result := photoRepository.AddPhoto(restaurantID, fileName)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"statusCode": 500,
			"error":      "Failed to create a new photo meta data for the file",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       "New photo is uploaded",
	})
}
