package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"sakaba.link/api/src/model"
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
	var json model.PhotoRequest
	if err := ctx.ShouldBindJSON(&json); err == nil {
		header, paramErr := ctx.FormFile("file_content")
		if paramErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": paramErr.Error()})
			return
		}

		file, fileErr := header.Open()
		if fileErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fileErr.Error()})
			return
		}

		fileName := uuid.New().String()
		s3Service := service.S3Service{}
		up, uploadErr := s3Service.Upload(json.RestaurantID, fileName, file)
		if uploadErr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":    "Failed to upload file",
				"uploader": up,
			})
			return
		}

		photoRepository := repository.PhotoRepository{}
		result := photoRepository.AddPhoto(json.RestaurantID, fileName)
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
}
