package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/repository"
)

type VideoController struct{}

func (c *VideoController) GetAllVideos(ctx *gin.Context) {
	videoRepository := repository.VideoRepository{}
	allVideos := videoRepository.GetAllVideos()

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       allVideos,
	})
}

func (c *VideoController) GetVideosByRestaurantId(ctx *gin.Context) {
	id := ctx.Param("id")
	videoRepository := repository.VideoRepository{}
	videos := videoRepository.GetVideosByRestaurantId(id)

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       videos,
	})
}
