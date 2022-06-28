package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yumaeda/sakaba-api/src/repository"
)

// VideoController is a controller for Video API.
type VideoController struct {
	Repository repository.VideoRepository
}

// GetAllVideos returns all the videos.
func (c VideoController) GetAllVideos(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       c.Repository.GetAllVideos(),
	})
}

// GetVideosByRestaurantID returns videos for the restaurant.
func (c VideoController) GetVideosByRestaurantID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       c.Repository.GetVideosByRestaurantID(ctx.Param("id")),
	})
}
