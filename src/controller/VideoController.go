package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/repository"
)

// VideoController is a controller for Video API.
type VideoController struct{}

// GetAllVideos returns all the videos.
func (c *VideoController) GetAllVideos(ctx *gin.Context) {
	videoRepository := repository.VideoRepository{}
	allVideos := videoRepository.GetAllVideos()

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       allVideos,
	})
}

// GetVideosByRestaurantID returns videos for the restaurant.
func (c *VideoController) GetVideosByRestaurantID(ctx *gin.Context) {
	id := ctx.Param("id")
	videoRepository := repository.VideoRepository{}
	videos := videoRepository.GetVideosByRestaurantID(id)

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       videos,
	})
}
