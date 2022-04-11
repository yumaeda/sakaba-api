package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/repository"
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
