package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/controller"
)

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello world"})
}

func main() {
	videoController := controller.VideoController{}
	router := gin.New()
	router.GET("/", index)
	router.GET("/videos/", videoController.GetAllVideos)
	router.GET("/videos/:id", videoController.GetVideosByRestaurantId)
	router.Run(":8080")
}
