package main

import (
	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/controller"
)

func CORS(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "https://sakaba.link")
	ctx.Header("Access-Control-Allow-Methods", "GET")
	ctx.Header("Access-Control-Allow-Headers", "*")
	ctx.Header("Content-Type", "application/json")
	/*
		if ctx.Request.Method != "OPTIONS" {
			ctx.Next()
		} else {
			ctx.AbortWithStatus(http.StatusOK)
		}
	*/
}

func main() {
	homeController := controller.HomeController{}
	videoController := controller.VideoController{}

	router := gin.Default()
	router.Use(CORS)
	router.GET("/", homeController.Index)
	router.GET("/videos/", videoController.GetAllVideos)
	router.GET("/videos/:id", videoController.GetVideosByRestaurantId)
	router.Run(":8080")
}
